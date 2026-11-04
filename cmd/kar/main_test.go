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
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
)

const testCommandName = "kar"

const cliHelperModeEnv = "KAR_CLI_HELPER_MODE"

const cleanupWaitHelperModeEnv = "KAR_CLEANUP_TEST_MODE"

const cleanupFinishedMarker = "cleanup request finished"

const testTrueString = "true"

var errMainTestFailure = errors.New("simulated failure")

var errMainTestWaitTimeout = errors.New("timeout while waiting for the virtual machine instance")

// mockRunner is a lightweight stand-in for runner.Runner used to exercise
// runMainApp without touching a real KubeVirt client.
type mockRunner struct {
	createErr  error
	waitErr    error
	deleteErr  error
	deleteFunc func(context.Context) error
}

func (m *mockRunner) CreateResources(_ context.Context, _, _, _, _ string) error {
	return m.createErr
}

func (m *mockRunner) WaitForVirtualMachineInstance(_ context.Context) error {
	return m.waitErr
}

func (m *mockRunner) DeleteResources(ctx context.Context) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx)
	}

	return m.deleteErr
}

// TestDefaultTimeoutConstants locks in the intended duration values of the
// package-level timeout constants, guarding against accidental arithmetic
// changes (e.g. `5 * time.Minute` becoming `5 + time.Minute`) that would
// otherwise silently alter runtime behavior without failing any other test.
func TestDefaultTimeoutConstants(t *testing.T) {
	t.Parallel()

	if defaultCleanupTimeout != 5*time.Minute {
		t.Fatalf("expected defaultCleanupTimeout to be %v, got %v", 5*time.Minute, defaultCleanupTimeout)
	}

	if defaultWaitTimeout != 1*time.Hour {
		t.Fatalf("expected defaultWaitTimeout to be %v, got %v", 1*time.Hour, defaultWaitTimeout)
	}

	if shutdownTimeout != 5*time.Second {
		t.Fatalf("expected shutdownTimeout to be %v, got %v", 5*time.Second, shutdownTimeout)
	}
}

// TestNewDefaultTimeoutFunctions exercises the constructor functions backing
// the package-level timeout variables directly, ensuring their arithmetic
// expressions are covered independently of package initialization order and
// giving mutation testing a direct, unambiguous target to kill.
func TestNewDefaultTimeoutFunctions(t *testing.T) {
	t.Parallel()

	if got := newDefaultCleanupTimeout(); got != 5*time.Minute {
		t.Fatalf("newDefaultCleanupTimeout() = %v, want %v", got, 5*time.Minute)
	}

	if got := newDefaultWaitTimeout(); got != 1*time.Hour {
		t.Fatalf("newDefaultWaitTimeout() = %v, want %v", got, 1*time.Hour)
	}

	if got := newShutdownTimeout(); got != 5*time.Second {
		t.Fatalf("newShutdownTimeout() = %v, want %v", got, 5*time.Second)
	}
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

	t.Run("returns ldflags values unmodified when no build info is available", func(t *testing.T) {
		t.Parallel()

		// Temporarily swap the readBuildInfo seam to simulate
		// debug.ReadBuildInfo() reporting ok=false, which happens when a
		// binary is built without module information embedded.
		origReadBuildInfo := readBuildInfo
		readBuildInfo = func() (*debug.BuildInfo, bool) { return nil, false }

		defer func() { readBuildInfo = origReadBuildInfo }()

		info := getBuildInfo("ldflags-commit", "ldflags-date", "ldflags-modified")

		if info.gitCommit != "ldflags-commit" {
			t.Fatalf("expected gitCommit %q, got %q", "ldflags-commit", info.gitCommit)
		}

		if info.buildDate != "ldflags-date" {
			t.Fatalf("expected buildDate %q, got %q", "ldflags-date", info.buildDate)
		}

		if info.gitTreeModified != "ldflags-modified" {
			t.Fatalf("expected gitTreeModified %q, got %q", "ldflags-modified", info.gitTreeModified)
		}

		if info.goVersion != "" {
			t.Fatalf("expected goVersion to remain empty, got %q", info.goVersion)
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
			{Key: "vcs.modified", Value: testTrueString},
			{Key: "vcs.unrelated", Value: "ignored"},
		})

		if out.gitCommit != "abc123" {
			t.Fatalf("expected gitCommit %q, got %q", "abc123", out.gitCommit)
		}

		if out.buildDate != "2026-01-01T00:00:00Z" {
			t.Fatalf("expected buildDate %q, got %q", "2026-01-01T00:00:00Z", out.buildDate)
		}

		if out.gitTreeModified != testTrueString {
			t.Fatalf("expected gitTreeModified %q, got %q", testTrueString, out.gitTreeModified)
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

		shutdown := setupTelemetry(context.Background(), log)
		assertShutdownNoError(t, shutdown)
	})

	t.Run("initializes and shuts down the stdout exporter when enabled", func(t *testing.T) {
		t.Setenv("KAR_TELEMETRY_ENABLED", "true")
		t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "stdout")

		shutdown := setupTelemetry(context.Background(), log)
		assertShutdownNoError(t, shutdown)
	})

	t.Run("logs a warning and still returns a usable shutdown function on init failure", func(t *testing.T) {
		t.Setenv("KAR_TELEMETRY_ENABLED", "true")
		t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "otlp")
		t.Setenv("KAR_TELEMETRY_OTLP_ENDPOINT", "http://localhost:19999")

		// An already-cancelled context makes OTLP exporter creation fail
		// synchronously, exercising setupTelemetry's log.Warnf branch.
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		shutdown := setupTelemetry(ctx, log)
		if shutdown == nil {
			t.Fatal("expected a non-nil shutdown function even when initialization fails")
		}

		assertShutdownNoError(t, shutdown)
	})
}

func TestShutdownTelemetryAndLog(t *testing.T) {
	t.Parallel()

	log := utils.GetLogger()

	t.Run("does not log when shutdown succeeds", func(t *testing.T) {
		t.Parallel()

		called := false

		shutdownTelemetryAndLog(context.Background(), func(context.Context) error {
			called = true

			return nil
		}, log)

		if !called {
			t.Fatal("expected shutdown function to be invoked")
		}
	})

	t.Run("logs a warning when shutdown returns an error", func(t *testing.T) {
		t.Parallel()

		// shutdownTelemetryAndLog should not panic and should log the
		// returned error via log.Warnf rather than propagating it.
		shutdownTelemetryAndLog(context.Background(), func(context.Context) error {
			return errMainTestFailure
		}, log)
	})
}

func TestRunCleanup(t *testing.T) {
	t.Parallel()

	log := utils.GetLogger()

	t.Run("does not log when DeleteResources succeeds", func(t *testing.T) {
		t.Parallel()

		runner := &mockRunner{}
		runCleanup(context.Background(), runner, log)
	})

	t.Run("logs cleanup failure when DeleteResources returns an error", func(t *testing.T) {
		t.Parallel()

		runner := &mockRunner{deleteErr: errMainTestFailure}
		runCleanup(context.Background(), runner, log)
	})
}

func TestRunMainApp(t *testing.T) {
	t.Parallel()

	log := utils.GetLogger()

	t.Run("returns success when execution succeeds", func(t *testing.T) {
		t.Parallel()

		runner := &mockRunner{}
		if got := runMainApp(context.Background(), runner, log, []string{}); got != 0 {
			t.Fatalf("runMainApp() = %d, want 0", got)
		}
	})

	t.Run("returns failure when runner execution fails", func(t *testing.T) {
		t.Parallel()

		runner := &mockRunner{createErr: errMainTestFailure}
		if got := runMainApp(context.Background(), runner, log, []string{}); got != 1 {
			t.Fatalf("runMainApp() = %d, want 1", got)
		}
	})

	t.Run("returns failure when execution is cancelled", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		runner := &mockRunner{waitErr: context.Canceled}
		if got := runMainApp(ctx, runner, log, []string{}); got != 1 {
			t.Fatalf("runMainApp() = %d, want 1", got)
		}
	})
}

func TestRunAndCleanupWaitsForCleanup(t *testing.T) {
	t.Setenv("KAR_CLEANUP_TIMEOUT", "1s")

	log := utils.GetLogger()
	cleanupStarted := make(chan struct{})
	releaseCleanup := make(chan struct{}, 1)

	runner := &mockRunner{
		waitErr: errMainTestWaitTimeout,
		deleteFunc: func(ctx context.Context) error {
			close(cleanupStarted)

			select {
			case <-releaseCleanup:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		},
	}

	ctx, stop := context.WithCancel(context.Background())
	t.Cleanup(stop)

	exitCode := make(chan int, 1)
	go func() {
		exitCode <- runAndCleanup(ctx, stop, runner, log, []string{})
	}()

	select {
	case <-cleanupStarted:
	case <-time.After(time.Second):
		t.Fatal("cleanup did not start")
	}

	select {
	case got := <-exitCode:
		releaseCleanup <- struct{}{}

		t.Fatalf("runAndCleanup() returned %d before cleanup completed", got)
	case <-time.After(20 * time.Millisecond):
	}

	releaseCleanup <- struct{}{}

	if got := <-exitCode; got != 1 {
		t.Fatalf("runAndCleanup() = %d, want 1", got)
	}
}

func TestRunAndCleanupHonorsCleanupTimeout(t *testing.T) {
	t.Setenv("KAR_CLEANUP_TIMEOUT", "20ms")

	cleanupTimedOut := make(chan struct{})
	runner := &mockRunner{
		waitErr: errMainTestWaitTimeout,
		deleteFunc: func(ctx context.Context) error {
			<-ctx.Done()
			close(cleanupTimedOut)

			return ctx.Err()
		},
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	if got := runAndCleanup(ctx, stop, runner, utils.GetLogger(), []string{}); got != 1 {
		t.Fatalf("runAndCleanup() = %d, want 1", got)
	}

	select {
	case <-cleanupTimedOut:
	default:
		t.Fatal("cleanup did not finish through its configured timeout")
	}
}

// TestRun_DirectInvocation exercises run() in-process so its statements are
// attributed to this test binary's coverage profile.
//
// These subtests intentionally avoid t.Parallel(): run() temporarily
// overrides the process-wide os.Args so that cobra parses a controlled
// argument list instead of the `go test` binary's own flags. Top-level tests
// that call t.Parallel() defer their bodies until every non-parallel
// top-level test (including this one) has finished, so there is no risk of a
// concurrent test observing the temporarily mutated os.Args.
func TestRun_DirectInvocation(t *testing.T) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	t.Run("returns early when client/namespace resolution fails", func(t *testing.T) {
		os.Args = []string{testCommandName}

		t.Setenv("KUBECONFIG", writeTempKubeconfig(t, malformedKubeconfig))
		t.Setenv("KAR_TELEMETRY_ENABLED", "false")

		// run() should report the resolution error and return early, without
		// panicking or reaching the runner/command-execution setup below it.
		if got := run(); got != 1 {
			t.Fatalf("run() = %d, want 1", got)
		}
	})

	t.Run("proceeds through command execution and signal-triggered cleanup", func(t *testing.T) {
		os.Args = []string{testCommandName}

		t.Setenv("KUBECONFIG", t.TempDir()+"/nonexistent-kubeconfig")
		t.Setenv("KAR_TELEMETRY_ENABLED", "false")
		t.Setenv("KAR_CLEANUP_TIMEOUT", "5s")

		// With no jitconfig flag supplied, the root command's RunE fails fast
		// without making any real KubeVirt API call. runAndCleanup then calls
		// DeleteResources synchronously, exercising its error-logging path.
		if got := run(); got != 1 {
			t.Fatalf("run() = %d, want 1", got)
		}
	})

	t.Run("returns failure for invalid flags", func(t *testing.T) {
		os.Args = []string{testCommandName, "--invalid-flag"}

		t.Setenv("KUBECONFIG", t.TempDir()+"/nonexistent-kubeconfig")
		t.Setenv("KAR_TELEMETRY_ENABLED", "false")

		if got := run(); got != 1 {
			t.Fatalf("run() = %d, want 1", got)
		}
	})

	t.Run("returns success for help", func(t *testing.T) {
		os.Args = []string{testCommandName, "--help"}

		t.Setenv("KUBECONFIG", t.TempDir()+"/nonexistent-kubeconfig")
		t.Setenv("KAR_TELEMETRY_ENABLED", "false")

		if got := run(); got != 0 {
			t.Fatalf("run() = %d, want 0", got)
		}
	})
}

func TestCLIProcessExitStatus(t *testing.T) {
	tests := []struct {
		name         string
		mode         string
		wantExitCode int
		wantOutput   string
	}{
		{name: "invalid flag exits with failure", mode: "invalid-flag", wantExitCode: 1},
		{name: "help exits successfully", mode: "help", wantOutput: "Usage:"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv(cliHelperModeEnv, test.mode)
			t.Setenv("KUBECONFIG", t.TempDir()+"/nonexistent-kubeconfig")
			t.Setenv("KAR_TELEMETRY_ENABLED", "false")

			//nolint:gosec // The test invokes its own binary with a fixed test selector.
			cmd := exec.CommandContext(t.Context(), os.Args[0], "-test.run=^TestCLIHelperProcess$")
			output, err := cmd.CombinedOutput()

			if test.wantExitCode == 0 {
				if err != nil {
					t.Fatalf("CLI process failed: %v\n%s", err, output)
				}
			} else {
				var exitErr *exec.ExitError
				if !errors.As(err, &exitErr) {
					t.Fatalf("expected CLI process exit code %d, got err=%v\n%s", test.wantExitCode, err, output)
				}

				if got := exitErr.ExitCode(); got != test.wantExitCode {
					t.Fatalf("CLI process exit code = %d, want %d\n%s", got, test.wantExitCode, output)
				}
			}

			if test.wantOutput != "" && !strings.Contains(string(output), test.wantOutput) {
				t.Fatalf("CLI output does not contain %q:\n%s", test.wantOutput, output)
			}
		})
	}
}

func TestCLIHelperProcess(t *testing.T) {
	t.Parallel()

	switch os.Getenv(cliHelperModeEnv) {
	case "invalid-flag":
		os.Args = []string{testCommandName, "--invalid-flag"}
	case "help":
		os.Args = []string{testCommandName, "--help"}
	default:
		return
	}

	main()
	t.Fatal("main() returned without exiting")
}

func TestCleanupWaitsBeforeProcessExit(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	//nolint:gosec // The test invokes its own binary with a fixed test selector.
	cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestCleanupWaitHelperProcess$")

	cmd.Env = append(os.Environ(), cleanupWaitHelperModeEnv+"="+testTrueString)
	output, err := cmd.CombinedOutput()

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 1 {
		t.Fatalf("helper process exit error = %v, want exit code 1\n%s", err, output)
	}

	if !strings.Contains(string(output), cleanupFinishedMarker) {
		t.Fatalf("process exited before cleanup finished:\n%s", output)
	}
}

func TestCleanupWaitHelperProcess(t *testing.T) {
	t.Parallel()

	if os.Getenv(cleanupWaitHelperModeEnv) != testTrueString {
		return
	}

	runner := &mockRunner{
		waitErr: errMainTestWaitTimeout,
		deleteFunc: func(context.Context) error {
			time.Sleep(50 * time.Millisecond)

			_, _ = fmt.Fprintln(os.Stdout, cleanupFinishedMarker)

			return nil
		},
	}

	ctx, stop := context.WithCancel(context.Background())
	os.Exit(runAndCleanup(ctx, stop, runner, utils.GetLogger(), []string{}))
}
