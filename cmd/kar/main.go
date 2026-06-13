/*
Copyright © 2024

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/electrocucaracha/kubevirt-actions-runner/cmd/kar/app"
	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
	"github.com/spf13/pflag"
	"kubevirt.io/client-go/kubecli"
)

const (
	defaultCleanupTimeout = 5 * time.Minute
	defaultWaitTimeout    = 1 * time.Hour
	shutdownTimeout       = 5 * time.Second
)

//nolint:gochecknoglobals
var (
	// Build-time variables set via ldflags during build.
	// These variables provide metadata about the build, such as the Git commit hash and build date.
	gitCommit       string
	buildDate       string
	gitTreeModified string
)

type buildInfo struct {
	gitCommit       string
	gitTreeModified string
	buildDate       string
	goVersion       string
}

func applyVCSSettings(out *buildInfo, settings []debug.BuildSetting) {
	for _, setting := range settings {
		switch setting.Key {
		case "vcs.revision":
			if out.gitCommit == "" {
				out.gitCommit = setting.Value
			}
		case "vcs.time":
			if out.buildDate == "" {
				out.buildDate = setting.Value
			}
		case "vcs.modified":
			if out.gitTreeModified == "" {
				out.gitTreeModified = setting.Value
			}
		}
	}
}

func getBuildInfo(gitCommit, buildDate, gitTreeModified string) buildInfo {
	out := buildInfo{
		gitCommit:       gitCommit,
		buildDate:       buildDate,
		gitTreeModified: gitTreeModified,
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return out
	}

	out.goVersion = info.GoVersion
	if gitCommit != "" && buildDate != "" {
		return out
	}

	applyVCSSettings(&out, info.Settings)

	return out
}

func getDurationEnvOrDefault(key string, defaultValue time.Duration) time.Duration {
	if val := os.Getenv(key); val != "" {
		d, err := time.ParseDuration(val)
		if err == nil {
			return d
		}

		utils.GetLogger().Printf("Invalid %s value: %q, using default %s", key, val, defaultValue)
	}

	return defaultValue
}

func getCleanupTimeout() time.Duration {
	return getDurationEnvOrDefault("KAR_CLEANUP_TIMEOUT", defaultCleanupTimeout)
}

func getWaitTimeout() time.Duration {
	return getDurationEnvOrDefault("KAR_WAIT_TIMEOUT", defaultWaitTimeout)
}

func ensureValidCleanupContext(parent context.Context) (context.Context, context.CancelFunc) {
	if parent.Err() != nil {
		return context.WithTimeout(context.Background(), getCleanupTimeout())
	}

	return context.WithTimeout(parent, getCleanupTimeout())
}

func setupTelemetry(log *utils.LoggerImpl) func(context.Context) error {
	telemetryCfg := runner.GetTelemetryConfig()

	shutdownTelemetry, err := runner.InitializeTelemetry(context.Background(), telemetryCfg)
	if err != nil {
		log.Warnf("failed to initialize telemetry: %v", err)
	}

	return shutdownTelemetry
}

func getClientAndNamespace() (kubecli.KubevirtClient, string, error) {
	clientConfig := kubecli.DefaultClientConfig(&pflag.FlagSet{})

	namespace, _, err := clientConfig.Namespace()
	if err != nil {
		return nil, "", fmt.Errorf("failed to get namespace: %w", err)
	}

	virtClient, err := kubecli.GetKubevirtClientFromClientConfig(clientConfig)
	if err != nil {
		return nil, "", fmt.Errorf("cannot obtain KubeVirt client: %w", err)
	}

	return virtClient, namespace, nil
}

func runMainApp(ctx context.Context, opts app.Opts, kr runner.Runner, log *utils.LoggerImpl) {
	rootCmd := app.NewRootCommand(ctx, kr, opts)

	execErr := rootCmd.Execute()
	if execErr != nil && !errors.Is(execErr, context.Canceled) {
		log.Println("execute command failed:", execErr)
	}
}

func main() {
	var opts app.Opts

	log := utils.GetLogger()
	buildInfo := getBuildInfo(gitCommit, buildDate, gitTreeModified)
	log.Printf("starting kubevirt action runner\ncommit: %v\tmodified: %v\tdate: %v\tgo: %v\n",
		buildInfo.gitCommit, buildInfo.gitTreeModified, buildInfo.buildDate, buildInfo.goVersion)

	shutdownTelemetry := setupTelemetry(log)

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err := shutdownTelemetry(shutdownCtx)
		if err != nil {
			log.Warnf("failed to shutdown telemetry: %v", err)
		}
	}()

	virtClient, namespace, err := getClientAndNamespace()
	if err != nil {
		log.Warnf("error getting client or namespace: %v\n", err)

		return
	}

	waitTimeout := getWaitTimeout()
	kubevirtRunner := runner.NewRunner(namespace, virtClient, waitTimeout)

	log.Printf("cleanup timeout is set to: %v", getCleanupTimeout())
	log.Printf("wait timeout is set to: %v", waitTimeout)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()

		cleanupCtx, cancel := ensureValidCleanupContext(ctx)
		defer cancel()

		err := kubevirtRunner.DeleteResources(cleanupCtx)
		if err != nil {
			log.Println("cleanup failed:", err)
		}
	}()

	runMainApp(ctx, opts, kubevirtRunner, log)
}
