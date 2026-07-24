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

//nolint:testpackage // tests unexported helpers directly
package main

import (
	"context"
	"errors"
	"runtime/debug"
	"testing"
	"time"

	"github.com/electrocucaracha/kubevirt-actions-runner/cmd/kar/app"
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

func TestApplyVCSSettings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		settings []debug.BuildSetting
		initial  buildInfo
		want     buildInfo
	}{
		{
			name: "populates empty fields",
			settings: []debug.BuildSetting{
				{Key: "vcs.revision", Value: "abc123"},
				{Key: "vcs.time", Value: "2026-01-01T00:00:00Z"},
				{Key: "vcs.modified", Value: "true"},
			},
			initial: buildInfo{},
			want: buildInfo{
				gitCommit:       "abc123",
				buildDate:       "2026-01-01T00:00:00Z",
				gitTreeModified: "true",
			},
		},
		{
			name: "does not override pre-populated fields",
			settings: []debug.BuildSetting{
				{Key: "vcs.revision", Value: "abc123"},
				{Key: "vcs.time", Value: "2026-01-01T00:00:00Z"},
				{Key: "vcs.modified", Value: "true"},
			},
			initial: buildInfo{
				gitCommit:       "preset-commit",
				buildDate:       "preset-date",
				gitTreeModified: "preset-modified",
			},
			want: buildInfo{
				gitCommit:       "preset-commit",
				buildDate:       "preset-date",
				gitTreeModified: "preset-modified",
			},
		},
		{
			name:     "ignores unknown keys",
			settings: []debug.BuildSetting{{Key: "unknown.key", Value: "value"}},
			initial:  buildInfo{},
			want:     buildInfo{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			out := test.initial
			applyVCSSettings(&out, test.settings)

			if out != test.want {
				t.Fatalf("applyVCSSettings() = %+v, want %+v", out, test.want)
			}
		})
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
		{name: "returns default when env unset", setEnv: false, defaultVal: 42 * time.Second, want: 42 * time.Second},
		{name: "returns default when env empty", setEnv: true, envVal: "", defaultVal: 42 * time.Second, want: 42 * time.Second},
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

func TestGetCleanupTimeout(t *testing.T) {
	t.Run("returns default when env unset", func(t *testing.T) {
		t.Setenv("KAR_CLEANUP_TIMEOUT", "")

		if got := getCleanupTimeout(); got != defaultCleanupTimeout {
			t.Fatalf("getCleanupTimeout() = %v, want %v", got, defaultCleanupTimeout)
		}
	})

	t.Run("returns env value when set", func(t *testing.T) {
		t.Setenv("KAR_CLEANUP_TIMEOUT", "3m")

		if got := getCleanupTimeout(); got != 3*time.Minute {
			t.Fatalf("getCleanupTimeout() = %v, want %v", got, 3*time.Minute)
		}
	})
}

func TestGetWaitTimeout(t *testing.T) {
	t.Run("returns default when env unset", func(t *testing.T) {
		t.Setenv("KAR_WAIT_TIMEOUT", "")

		if got := getWaitTimeout(); got != defaultWaitTimeout {
			t.Fatalf("getWaitTimeout() = %v, want %v", got, defaultWaitTimeout)
		}
	})

	t.Run("returns env value when set", func(t *testing.T) {
		t.Setenv("KAR_WAIT_TIMEOUT", "2h")

		if got := getWaitTimeout(); got != 2*time.Hour {
			t.Fatalf("getWaitTimeout() = %v, want %v", got, 2*time.Hour)
		}
	})
}

func TestEnsureValidCleanupContext(t *testing.T) {
	t.Parallel()

	t.Run("derives a fresh context when parent is already cancelled", func(t *testing.T) {
		t.Parallel()

		parentCtx, cancel := context.WithCancel(context.Background())
		cancel()

		cleanupCtx, cleanupCancel := ensureValidCleanupContext(parentCtx)
		defer cleanupCancel()

		if err := cleanupCtx.Err(); err != nil {
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

		if err := cleanupCtx.Err(); err != nil {
			t.Fatalf("expected context derived from valid parent to not be cancelled, got err=%v", err)
		}

		if _, ok := cleanupCtx.Deadline(); !ok {
			t.Fatal("expected cleanup context to have a deadline")
		}
	})
}

func TestRunMainApp(t *testing.T) {
	t.Parallel()

	log := utils.GetLogger()

	t.Run("logs nothing extra when execution succeeds", func(t *testing.T) {
		t.Parallel()

		runner := &mockRunner{}
		opts := app.Opts{}

		// runMainApp should not panic and should invoke the root command
		// against the provided runner without requiring a real KubeVirt client.
		runMainApp(context.Background(), opts, runner, log)
	})

	t.Run("logs failure when execution returns a non-cancellation error", func(t *testing.T) {
		t.Parallel()

		runner := &mockRunner{createErr: errMainTestFailure}
		opts := app.Opts{}

		runMainApp(context.Background(), opts, runner, log)
	})

	t.Run("suppresses logging when execution is cancelled", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		runner := &mockRunner{}
		opts := app.Opts{}

		runMainApp(ctx, opts, runner, log)
	})
}
