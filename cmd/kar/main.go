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
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/electrocucaracha/kubevirt-actions-runner/cmd/kar/app"
	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"kubevirt.io/client-go/kubecli"
)

const (
	defaultCleanupTimeout = 5 * time.Minute
	shutdownTimeout       = 5 * time.Second
)

type buildInfo struct {
	gitCommit       string
	gitTreeModified string
	buildDate       string
	goVersion       string
}

func getBuildInfo() buildInfo {
	out := buildInfo{}

	if info, ok := debug.ReadBuildInfo(); ok {
		out.goVersion = info.GoVersion

		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				out.gitCommit = setting.Value
			case "vcs.time":
				out.buildDate = setting.Value
			case "vcs.modified":
				out.gitTreeModified = setting.Value
			}
		}
	}

	return out
}

func getCleanupTimeout() time.Duration {
	log := utils.GetLogger()

	if val := os.Getenv("KAR_CLEANUP_TIMEOUT"); val != "" {
		d, err := time.ParseDuration(val)
		if err == nil {
			return d
		}

		log.Printf("Invalid KAR_CLEANUP_TIMEOUT value: %q, using default %s", val, defaultCleanupTimeout)
	}

	return defaultCleanupTimeout
}

func ensureValidCleanupContext(parent context.Context) (context.Context, context.CancelFunc) {
	if parent.Err() != nil {
		return context.WithTimeout(context.TODO(), getCleanupTimeout())
	}

	return context.WithTimeout(parent, getCleanupTimeout())
}

func setupTelemetry(log any) func(context.Context) error {
	telemetryCfg := runner.GetTelemetryConfig()

	shutdownTelemetry, err := runner.InitializeTelemetry(context.Background(), telemetryCfg)
	if err != nil {
		// Cast log to the correct type
		if logVal, ok := log.(interface{ Warnf(s string, args ...any) }); ok {
			logVal.Warnf("failed to initialize telemetry: %v", err)
		}
	}

	if shutdownTelemetry == nil {
		return func(_ context.Context) error { return nil }
	}

	return shutdownTelemetry
}

func getClientAndNamespace() (kubecli.KubevirtClient, string, error) {
	clientConfig := kubecli.DefaultClientConfig(&pflag.FlagSet{})

	namespace, _, err := clientConfig.Namespace()
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to get namespace")
	}

	virtClient, err := kubecli.GetKubevirtClientFromClientConfig(clientConfig)
	if err != nil {
		return nil, "", errors.Wrap(err, "cannot obtain KubeVirt client")
	}

	return virtClient, namespace, nil
}

func runMainApp(ctx context.Context, opts app.Opts, kr runner.Runner, log any) {
	rootCmd := app.NewRootCommand(ctx, kr, opts)

	execErr := rootCmd.Execute()
	if execErr != nil && !errors.Is(errors.Cause(execErr), context.Canceled) {
		if logVal, ok := log.(interface{ Println(args ...any) }); ok {
			logVal.Println("execute command failed:", execErr)
		}
	}
}

func main() {
	var opts app.Opts

	log := utils.GetLogger()
	buildInfo := getBuildInfo()
	log.Printf("starting kubevirt action runner\ncommit: %v\tmodified: %v\tdate: %v\tgo: %v\n",
		buildInfo.gitCommit, buildInfo.gitTreeModified, buildInfo.buildDate, buildInfo.goVersion)

	// Initialize telemetry
	shutdownTelemetry := setupTelemetry(log)

	defer func() {
		if shutdownTelemetry != nil {
			shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()

			err := shutdownTelemetry(shutdownCtx)
			if err != nil {
				log.Warnf("failed to shutdown telemetry: %v", err)
			}
		}
	}()

	virtClient, namespace, err := getClientAndNamespace()
	if err != nil {
		log.Warnf("error getting client or namespace: %v\n", err)

		return
	}

	kubevirtRunner := runner.NewRunner(namespace, virtClient)

	log.Printf("cleanup timeout is set to: %s", getCleanupTimeout())

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
