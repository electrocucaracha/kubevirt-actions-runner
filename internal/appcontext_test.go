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

package runner_test

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"sync"
	"testing"

	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
)

func expectExitCode(t *testing.T, err error, expected int) {
	t.Helper()

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		t.Fatalf("expected process to exit with a non-zero status, got err=%v", err)
	}

	if exitErr.ExitCode() != expected {
		t.Fatalf("expected exit code %d, got %d", expected, exitErr.ExitCode())
	}
}

func TestCancelAppContextResetsSingleton(t *testing.T) {
	t.Cleanup(runner.CancelAppContext)

	ctx := runner.NewAppContext("first-vmi", "first-dv")
	if got := ctx.GetVMIName(); got != "first-vmi" {
		t.Fatalf("expected first VMI name, got %q", got)
	}

	runner.CancelAppContext()

	ctx = runner.NewAppContext("second-vmi", "second-dv")
	if got := ctx.GetVMIName(); got != "second-vmi" {
		t.Fatalf("expected reset VMI name, got %q", got)
	}

	if got := ctx.GetDataVolumeName(); got != "second-dv" {
		t.Fatalf("expected reset data volume name, got %q", got)
	}
}

// TestHasAppContextLifecycle verifies that HasAppContext directly reflects
// the initialization state of the singleton across its full lifecycle:
// uninitialized, initialized, and reset via CancelAppContext.
func TestHasAppContextLifecycle(t *testing.T) {
	t.Parallel()
	t.Cleanup(runner.CancelAppContext)

	runner.CancelAppContext()

	if runner.HasAppContext() {
		t.Fatal("expected HasAppContext to be false before initialization")
	}

	runner.NewAppContext("lifecycle-vmi", "lifecycle-dv")

	if !runner.HasAppContext() {
		t.Fatal("expected HasAppContext to be true after initialization")
	}

	runner.CancelAppContext()

	if runner.HasAppContext() {
		t.Fatal("expected HasAppContext to be false after CancelAppContext")
	}
}

// TestNewAppContextIgnoresSubsequentValues verifies the documented behavior
// that once the AppContext singleton is created, further calls to
// NewAppContext return the existing instance and ignore the new arguments,
// until CancelAppContext resets it.
func TestNewAppContextIgnoresSubsequentValues(t *testing.T) {
	t.Parallel()
	t.Cleanup(runner.CancelAppContext)

	runner.CancelAppContext()

	first := runner.NewAppContext("original-vmi", "original-dv")
	second := runner.NewAppContext("ignored-vmi", "ignored-dv")

	if first != second {
		t.Fatal("expected NewAppContext to return the same singleton instance on subsequent calls")
	}

	if got := second.GetVMIName(); got != "original-vmi" {
		t.Fatalf("expected VMI name to remain %q, got %q", "original-vmi", got)
	}

	if got := second.GetDataVolumeName(); got != "original-dv" {
		t.Fatalf("expected data volume name to remain %q, got %q", "original-dv", got)
	}
}

// TestAppContextConcurrentAccess exercises NewAppContext, GetAppContext, and
// HasAppContext from many goroutines concurrently to verify the mutex-guarded
// singleton is safe under concurrent access (run with -race to detect data
// races) and that all callers observe a single, consistent instance.
func TestAppContextConcurrentAccess(t *testing.T) {
	t.Parallel()
	t.Cleanup(runner.CancelAppContext)

	runner.CancelAppContext()

	const goroutines = 50

	var (
		waitGroup sync.WaitGroup
		mutex     sync.Mutex
		instances = make([]*runner.AppContext, 0, goroutines)
	)

	waitGroup.Add(goroutines)

	for routine := range goroutines {
		go func(idx int) {
			defer waitGroup.Done()

			ctx := runner.NewAppContext("concurrent-vmi", "concurrent-dv")

			_ = runner.HasAppContext()

			if idx%2 == 0 {
				ctx = runner.GetAppContext()
			}

			mutex.Lock()

			instances = append(instances, ctx)
			mutex.Unlock()
		}(routine)
	}

	waitGroup.Wait()

	if len(instances) != goroutines {
		t.Fatalf("expected %d recorded instances, got %d", goroutines, len(instances))
	}

	for _, inst := range instances {
		if inst != instances[0] {
			t.Fatal("expected all goroutines to observe the same singleton instance")
		}
	}

	if got := instances[0].GetVMIName(); got != "concurrent-vmi" {
		t.Fatalf("expected VMI name %q, got %q", "concurrent-vmi", got)
	}
}

// TestGetAppContextExitsWhenUninitialized verifies that GetAppContext exits
// the process with status 1 when called before NewAppContext. The assertion
// runs in a subprocess since Fatal terminates the process.
func TestGetAppContextExitsWhenUninitialized(t *testing.T) {
	t.Parallel()

	if os.Getenv("KAR_TEST_INVOKE_GET_APP_CONTEXT") == "1" {
		runner.CancelAppContext()
		runner.GetAppContext()

		return
	}

	//nolint:gosec // re-executes the test binary itself (os.Args[0]); not user-controlled input.
	cmd := exec.CommandContext(context.Background(), os.Args[0], "-test.run=TestGetAppContextExitsWhenUninitialized")

	cmd.Env = append(os.Environ(), "KAR_TEST_INVOKE_GET_APP_CONTEXT=1")

	expectExitCode(t, cmd.Run(), 1)
}
