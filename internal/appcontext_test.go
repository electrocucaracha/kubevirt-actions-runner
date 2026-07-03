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
	"errors"
	"os"
	"os/exec"
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

	//nolint:gosec
	cmd := exec.Command(os.Args[0], "-test.run=TestGetAppContextExitsWhenUninitialized")
	cmd.Env = append(os.Environ(), "KAR_TEST_INVOKE_GET_APP_CONTEXT=1")

	expectExitCode(t, cmd.Run(), 1)
}
