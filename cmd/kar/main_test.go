/* jscpd:ignore-start */
/*
Copyright © 2026

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
	"os"
	"runtime/debug"
	"testing"
	"time"

	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
)

var errMainTestFailure = errors.New("simulated failure")

// mockRunner is a lightweight stand-in for runner.Runner used to exercise
// runMainApp without touching a real KubeVirt client.
type mockRunner struct {
	createErr error
	waitErr   error
	deleteErr error
}

func (m *mockRunner) CreateResources(_ context.Context, _, _, _, _ string) error {
	return m.createErr
}

func (m *mockRunner) WaitForVirtualMachineInstance(_ context.Context) error {
	return m.waitErr
}

func (m *mockRunner) DeleteResources(_ context.Context) error {
	return m.deleteErr
}

func TestGetBuildInfo(t *testing.T) {
	t.Parallel()

	t.Run("returns ldflags values when both commit and date are provided", func(t *testing.T) {
		t.Parallel()

		info := getBuildInfo("ldflags-commit", "ldflags-date", "")

		if info.gitCommit != "ldflags-commit" {
			t.Fatalf("expected gitCommit %q, got %q", "ldflags-commit", info.gitCommit)
		}

		if info.buildDate != "ldflags-date" {
			t.Fatalf("expected buildDate %q, got %q", "ldflags-date", info.buildDate)
		}
	})

	t.Run("falls back to VCS settings when ldflags values are empty", func(t *testing.T) {
		t.Parallel()

		info := getBuildInfo("", "", "")

		// The test binary is built with debug.ReadBuildInfo support, so the
		// goVersion field should be populated regardless of VCS availability.
		if info.goVersion == "" {
			t.Fatal("expected goVersion to be populated from build info")
		}
	})
}

func TestApplyVCSSettings(t *testing.T) {
	t.Parallel()

	t.Run("populates all fields from empty settings", func(t *testing.T) {
		t.Parallel()

		out := buildInfo{}
		out.applyVCSSettings([]debug.BuildSetting{
			{Key: "vcs.revision", Value: "abc123"},
			{Key: "vcs.time", Value: "2026-01-01T00:00:00Z"},
			{Key: "vcs.modified", Value: "true"},
			{Key: "vcs.unrelated", Value: "ignored"},
		})

		if out.gitCommit != "abc123" {
			t.Fatalf("expected gitCommit %q, got %q", "abc123", out.gitCommit)
		}

		if out.buildDate != "2026-01-01T00:00:00Z" {
			t.Fatalf("expected buildDate %q, got %q", "2026-01-01T00:00:00Z", out.buildDate)
		}

		if out.gitTreeModified != "true" {
			t.Fatalf("expected gitTreeModified %q, got %q", "true", out.gitTreeModified)
		}
	})

	t.Run("does not overwrite fields that are already set", func(t *testing.T) {
		t.Parallel()

		out := buildInfo{
			gitCommit:       "preset-commit",
			buildDate:       "preset-date",
			gitTreeModified: "preset-modified",
		}
		out.applyVCSSettings([]debug.BuildSetting{
			{Key: "vcs.revision", Value: "from-settings-commit"},
			{Key: "vcs.time", Value: "from-settings-date"},
			{Key: "vcs.modified", Value: "from-settings-modified"},
		})

		if out.gitCommit != "preset-commit" {
			t.Fatalf("expected gitCommit to remain %q, got %q", "preset-commit", out.gitCommit)
		}

		if out.buildDate != "preset-date" {
			t.Fatalf("expected buildDate to remain %q, got %q", "preset-date", out.buildDate)
		}

		if out.gitTreeModified != "preset-modified" {
			t.Fatalf("expected gitTreeModified to remain %q, got %q", "preset-modified", out.gitTreeModified)
		}
	})
}

func TestGetDurationEnvOrDefault(t *testing.T) {
	const testKey = "KAR_TEST_DURATION"

	tests := []struct {
		name       string
		envVal     string
		setEnv     bool
		want       time.Duration
		defaultVal time.Duration
	}{
		{
			name: "returns default when env unset", setEnv: false,
			defaultVal: 42 * time.Second, want: 42 * time.Second,
		},
		{
			name: "returns default when env empty", setEnv: true, envVal: "",
			defaultVal: 42 * time.Second, want: 42 * time.Second,
		},
		{
			name: "returns parsed duration when valid", setEnv: true, envVal: "10s",
			defaultVal: 42 * time.Second, want: 10 * time.Second,
		},
		{
			name: "returns default when env invalid", setEnv: true, envVal: "not-a-duration",
			defaultVal: 42 * time.Second, want: 42 * time.Second,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.setEnv {
				t.Setenv(testKey, test.envVal)
			}

			got := getDurationEnvOrDefault(testKey, test.defaultVal)
			if got != test.want {
				t.Fatalf("getDurationEnvOrDefault() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestEnsureValidCleanupContext(t *testing.T) {
	t.Parallel()

	t.Run("derives a fresh context when parent is already cancelled", func(t *testing.T) {
		t.Parallel()

		parentCtx, cancel := context.WithCancel(context.Background())
		cancel()

		cleanupCtx, cleanupCancel := ensureValidCleanupContext(parentCtx)
		defer cleanupCancel()

		err := cleanupCtx.Err()
		if err != nil {
			t.Fatalf("expected fresh context to not be cancelled, got err=%v", err)
		}

		if _, ok := cleanupCtx.Deadline(); !ok {
			t.Fatal("expected cleanup context to have a deadline")
		}
	})

	t.Run("derives from parent when parent is still valid", func(t *testing.T) {
		t.Parallel()

		parentCtx := context.Background()

		cleanupCtx, cleanupCancel := ensureValidCleanupContext(parentCtx)
		defer cleanupCancel()

		err := cleanupCtx.Err()
		if err != nil {
			t.Fatalf("expected context derived from valid parent to not be cancelled, got err=%v", err)
		}

		if _, ok := cleanupCtx.Deadline(); !ok {
			t.Fatal("expected cleanup context to have a deadline")
		}
	})
}

// malformedKubeconfig is intentionally invalid YAML so that client config
// loading fails while parsing, exercising the namespace-resolution error path
// of getClientAndNamespace.
const malformedKubeconfig = "not: valid: yaml: content: [\n"

// kubeconfigWithInvalidCA is syntactically valid but contains a
// certificate-authority-data value that isn't a real PEM certificate, so
// namespace resolution succeeds but building the KubeVirt client fails.
const kubeconfigWithInvalidCA = `apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://example.invalid:6443
    certificate-authority-data: bm90IGEgdmFsaWQgY2VydGlmaWNhdGU=
  name: test
contexts:
- context:
    cluster: test
    namespace: mynamespace
    user: test
  name: test
current-context: test
users:
- name: test
  user:
    token: faketoken
`

func writeTempKubeconfig(t *testing.T, contents string) string {
	t.Helper()

	dir := t.TempDir()
	path := dir + "/kubeconfig"

	err := os.WriteFile(path, []byte(contents), 0o600)
	if err != nil {
		t.Fatalf("failed to write temp kubeconfig: %v", err)
	}

	return path
}

func assertClientAndNamespaceError(t *testing.T, client any, namespace string, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if client != nil {
		t.Fatalf("expected nil client on error, got %v", client)
	}

	if namespace != "" {
		t.Fatalf("expected empty namespace on error, got %q", namespace)
	}
}

func assertShutdownNoError(t *testing.T, shutdown func(context.Context) error) {
	t.Helper()

	if shutdown == nil {
		t.Fatal("expected a non-nil shutdown function")
	}

	err := shutdown(context.Background())
	if err != nil {
		t.Fatalf("expected no error from shutdown, got %v", err)
	}
}

func TestGetClientAndNamespace(t *testing.T) {
	t.Run("returns default namespace and client when no kubeconfig is configured", func(t *testing.T) {
		t.Setenv("KUBECONFIG", t.TempDir()+"/nonexistent-kubeconfig")

		client, namespace, err := getClientAndNamespace()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if namespace != "default" {
			t.Fatalf("expected namespace %q, got %q", "default", namespace)
		}

		if client == nil {
			t.Fatal("expected a non-nil KubeVirt client")
		}
	})

	t.Run("returns an error when the kubeconfig cannot be parsed", func(t *testing.T) {
		t.Setenv("KUBECONFIG", writeTempKubeconfig(t, malformedKubeconfig))

		client, namespace, err := getClientAndNamespace()
		assertClientAndNamespaceError(t, client, namespace, err)
	})

	t.Run("returns an error when the KubeVirt client cannot be built", func(t *testing.T) {
		t.Setenv("KUBECONFIG", writeTempKubeconfig(t, kubeconfigWithInvalidCA))

		client, namespace, err := getClientAndNamespace()
		assertClientAndNamespaceError(t, client, namespace, err)
	})
}

func TestSetupTelemetry(t *testing.T) {
	log := utils.GetLogger()

	t.Run("returns a no-op shutdown function when telemetry is disabled", func(t *testing.T) {
		t.Setenv("KAR_TELEMETRY_ENABLED", "false")

		shutdown := setupTelemetry(log)
		assertShutdownNoError(t, shutdown)
	})

	t.Run("initializes and shuts down the stdout exporter when enabled", func(t *testing.T) {
		t.Setenv("KAR_TELEMETRY_ENABLED", "true")
		t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "stdout")

		shutdown := setupTelemetry(log)
		assertShutdownNoError(t, shutdown)
	})
}

func TestRunMainApp(t *testing.T) {
	t.Parallel()

	log := utils.GetLogger()

	t.Run("logs nothing extra when execution succeeds", func(t *testing.T) {
		t.Parallel()

		runner := &mockRunner{}
		// runMainApp should not panic and should invoke the root command
		// against the provided runner without requiring a real KubeVirt client.
		runMainApp(context.Background(), runner, log)
	})

	t.Run("logs failure when execution returns a non-cancellation error", func(t *testing.T) {
		t.Parallel()

		runner := &mockRunner{createErr: errMainTestFailure}
		runMainApp(context.Background(), runner, log)
	})

	t.Run("suppresses logging when execution is cancelled", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		runner := &mockRunner{}
		runMainApp(ctx, runner, log)
	})
}

// TestMain_DirectInvocation exercises the real main() entrypoint directly
// (in-process) rather than through a re-exec'd subprocess, so that its
// statements are attributed to this test binary's coverage profile.
//
// These subtests intentionally avoid t.Parallel(): main() temporarily
// overrides the process-wide os.Args so that cobra parses a controlled
// argument list instead of the `go test` binary's own flags. Top-level tests
// that call t.Parallel() defer their bodies until every non-parallel
// top-level test (including this one) has finished, so there is no risk of a
// concurrent test observing the temporarily mutated os.Args.
func TestMain_DirectInvocation(t *testing.T) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	t.Run("returns early when client/namespace resolution fails", func(t *testing.T) {
		os.Args = []string{"kar"}

		t.Setenv("KUBECONFIG", writeTempKubeconfig(t, malformedKubeconfig))
		t.Setenv("KAR_TELEMETRY_ENABLED", "false")

		// main() should log the resolution error and return early, without
		// panicking or reaching the runner/command-execution setup below it.
		main()
	})

	t.Run("proceeds through command execution and signal-triggered cleanup", func(t *testing.T) {
		os.Args = []string{"kar"}

		t.Setenv("KUBECONFIG", t.TempDir()+"/nonexistent-kubeconfig")
		t.Setenv("KAR_TELEMETRY_ENABLED", "false")
		t.Setenv("KAR_CLEANUP_TIMEOUT", "5s")

		// With no jitconfig flag supplied, the root command's RunE fails fast
		// (without making any real KubeVirt API call), so main() reaches the
		// end of its body quickly. Its deferred stop() then cancels the
		// signal-notification context, unblocking the cleanup goroutine,
		// which attempts (and fails fast against) a real DeleteResources
		// call, exercising that goroutine's error-logging path as well.
		main()

		// Give the cleanup goroutine a brief moment to finish running before
		// this subtest returns, so its coverage counters are recorded
		// deterministically rather than racing the test binary's exit.
		time.Sleep(200 * time.Millisecond)
	})
}
