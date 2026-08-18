/* jscpd:ignore-start */
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
/* jscpd:ignore-end */

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
	vcsRevisionSetting = "vcs.revision"
	vcsTimeSetting     = "vcs.time"
	vcsModifiedSetting = "vcs.modified"
)

//nolint:gochecknoglobals
var (
	// Build-time variables set via ldflags during build.
	// These variables provide metadata about the build, such as the Git commit hash and build date.
	gitCommit       string
	buildDate       string
	gitTreeModified string

	// readBuildInfo is a seam over debug.ReadBuildInfo so tests can exercise
	// the "no build info available" branch of getBuildInfo deterministically.
	readBuildInfo = debug.ReadBuildInfo

	// defaultCleanupTimeout, defaultWaitTimeout, and shutdownTimeout are computed by
	// dedicated functions rather than declared as plain arithmetic constants. Go's
	// coverage instrumentation does not track top-level const declarations, so
	// mutation-testing tools (e.g. Gremlins) can never observe a test exercising
	// those arithmetic expressions and always flag them as "not covered". Wrapping
	// each value in a function makes the arithmetic part of an instrumented
	// statement that runs at package initialization and is verified by
	// TestDefaultTimeoutConstants.
	defaultCleanupTimeout = newDefaultCleanupTimeout()
	defaultWaitTimeout    = newDefaultWaitTimeout()
	shutdownTimeout       = newShutdownTimeout()
)

func newDefaultCleanupTimeout() time.Duration {
	//nolint:mnd // 5 minutes is the intended default cleanup window.
	return 5 * time.Minute
}

func newDefaultWaitTimeout() time.Duration {
	return 1 * time.Hour
}

func newShutdownTimeout() time.Duration {
	//nolint:mnd // 5 seconds is the intended default shutdown window.
	return 5 * time.Second
}

type buildInfo struct {
	gitCommit       string
	gitTreeModified string
	buildDate       string
	goVersion       string
}

func getBuildInfo(commit, date, modified string) buildInfo {
	out := buildInfo{
		gitCommit:       commit,
		buildDate:       date,
		gitTreeModified: modified,
	}

	info, ok := readBuildInfo()
	if !ok {
		return out
	}

	out.goVersion = info.GoVersion
	if commit != "" && date != "" {
		return out
	}

	out.applyVCSSettings(info.Settings)

	return out
}

func (out *buildInfo) applyVCSSettings(settings []debug.BuildSetting) {
	for _, setting := range settings {
		switch setting.Key {
		case vcsRevisionSetting:
			if out.gitCommit == "" {
				out.gitCommit = setting.Value
			}
		case vcsTimeSetting:
			if out.buildDate == "" {
				out.buildDate = setting.Value
			}
		case vcsModifiedSetting:
			if out.gitTreeModified == "" {
				out.gitTreeModified = setting.Value
			}
		}
	}
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

func ensureValidCleanupContext(parent context.Context) (context.Context, context.CancelFunc) {
	cleanupTimeout := getDurationEnvOrDefault("KAR_CLEANUP_TIMEOUT", defaultCleanupTimeout)
	if parent.Err() != nil {
		return context.WithTimeout(context.Background(), cleanupTimeout)
	}

	return context.WithTimeout(parent, cleanupTimeout)
}

func setupTelemetry(ctx context.Context, log *utils.LoggerImpl) func(context.Context) error {
	shutdownTelemetry, err := runner.InitializeTelemetry(ctx)
	if err != nil {
		log.Warnf("failed to initialize telemetry: %v", err)
	}

	return shutdownTelemetry
}

// shutdownTelemetryAndLog invokes the telemetry shutdown function and logs a
// warning if it fails, without terminating the process.
func shutdownTelemetryAndLog(
	ctx context.Context,
	shutdownTelemetry func(context.Context) error,
	log *utils.LoggerImpl,
) {
	err := shutdownTelemetry(ctx)
	if err != nil {
		log.Warnf("failed to shutdown telemetry: %v", err)
	}
}

// runCleanup deletes the runner's KubeVirt resources once the parent context
// is done, logging any failure returned by DeleteResources.
func runCleanup(ctx context.Context, kr runner.Runner, log *utils.LoggerImpl) {
	cleanupCtx, cancel := ensureValidCleanupContext(ctx)
	defer cancel()

	err := kr.DeleteResources(cleanupCtx)
	if err != nil {
		log.Println("cleanup failed:", err)
	}
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

func runMainApp(ctx context.Context, kr runner.Runner, log *utils.LoggerImpl) {
	rootCmd := app.NewRootCommand(ctx, kr, app.Opts{})

	execErr := rootCmd.Execute()
	if execErr != nil && !errors.Is(execErr, context.Canceled) {
		log.Println("execute command failed:", execErr)
	}
}

func main() {
	log := utils.GetLogger()
	buildInfo := getBuildInfo(gitCommit, buildDate, gitTreeModified)
	log.Printf("starting kubevirt action runner\ncommit: %v\tmodified: %v\tdate: %v\tgo: %v\n",
		buildInfo.gitCommit, buildInfo.gitTreeModified, buildInfo.buildDate, buildInfo.goVersion)

	shutdownTelemetry := setupTelemetry(context.Background(), log)

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		shutdownTelemetryAndLog(shutdownCtx, shutdownTelemetry, log)
	}()

	virtClient, namespace, err := getClientAndNamespace()
	if err != nil {
		log.Warnf("error getting client or namespace: %v\n", err)

		return
	}

	waitTimeout := getDurationEnvOrDefault("KAR_WAIT_TIMEOUT", defaultWaitTimeout)
	kubevirtRunner := runner.NewRunner(namespace, virtClient, waitTimeout)

	log.Printf("cleanup timeout is set to: %v", getDurationEnvOrDefault("KAR_CLEANUP_TIMEOUT", defaultCleanupTimeout))
	log.Printf("wait timeout is set to: %v", waitTimeout)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()

		runCleanup(ctx, kubevirtRunner, log)
	}()

	runMainApp(ctx, kubevirtRunner, log)
}
