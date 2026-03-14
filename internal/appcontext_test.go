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

package runner_test

import (
	"testing"

	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
)

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
