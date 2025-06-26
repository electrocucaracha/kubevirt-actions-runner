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
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/electrocucaracha/kubevirt-actions-runner/cmd/kar/app"
	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"kubevirt.io/client-go/kubecli"
)

const defaultCleanupTimeout = 30 * time.Minute

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
	if val := os.Getenv("KAR_CLEANUP_TIMEOUT"); val != "" {
		if d, err := time.ParseDuration(val); err == nil {
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

func main() {
	var (
		opts app.Opts
		err  error
	)

	buildInfo := getBuildInfo()
	log.Printf("starting kubevirt action runner\ncommit: %v\tmodified: %v\tdate: %v\tgo: %v\n",
		buildInfo.gitCommit, buildInfo.gitTreeModified, buildInfo.buildDate, buildInfo.goVersion)

	clientConfig := kubecli.DefaultClientConfig(&pflag.FlagSet{})

	namespace, _, err := clientConfig.Namespace()
	if err != nil {
		log.Fatalf("error in namespace : %v\n", err)
	}

	virtClient, err := kubecli.GetKubevirtClientFromClientConfig(clientConfig)
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt client: %v\n", err)
	}

	runner := runner.NewRunner(namespace, virtClient)

	log.Printf("cleanup timeout is set to: %s", getCleanupTimeout())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()

		cleanupCtx, cancel := ensureValidCleanupContext(ctx)
		defer cancel()

		if err := runner.DeleteResources(cleanupCtx, runner.GetVMIName(), runner.GetDataVolumeName()); err != nil {
			log.Println("cleanup failed:", err)
		}
	}()

	rootCmd := app.NewRootCommand(ctx, runner, opts)

	if err := rootCmd.Execute(); err != nil && !errors.Is(errors.Cause(err), context.Canceled) {
		log.Println("execute command failed:", err)
	}
}
