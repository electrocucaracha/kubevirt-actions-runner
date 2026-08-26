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

package app //nolint:testpackage // tests unexported bindFlags behavior directly

import (
	"testing"

	"github.com/spf13/cobra"
)

const testCommandUse = "test"

// TestInitializeConfig_SkipsChangedFlags verifies that initializeConfig does not override
// a flag value that was already explicitly set on the command line.
func TestInitializeConfig_SkipsChangedFlags(t *testing.T) {
	t.Setenv("TESTSTR", "from-env")

	cmd := new(cobra.Command)
	cmd.Use = testCommandUse
	cmd.Flags().String("teststr", "default", "test string flag")

	err := cmd.Flags().Set("teststr", "from-cli")
	if err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	err = initializeConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := cmd.Flags().Lookup("teststr").Value.String(); got != "from-cli" {
		t.Fatalf("expected flag to remain %q, got %q", "from-cli", got)
	}
}

// TestInitializeConfig_SkipsUnsetEnvVars verifies that initializeConfig leaves the flag at
// its default value when no corresponding environment variable is set.
func TestInitializeConfig_SkipsUnsetEnvVars(t *testing.T) {
	t.Parallel()

	cmd := new(cobra.Command)
	cmd.Use = testCommandUse
	cmd.Flags().String("testunset", "default", "test string flag")

	err := initializeConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := cmd.Flags().Lookup("testunset").Value.String(); got != "default" {
		t.Fatalf("expected flag to remain default, got %q", got)
	}
}

// TestInitializeConfig_AppliesUnsetEnvVars verifies that initializeConfig overrides a flag's
// default value from the matching environment variable when the flag was not explicitly set
// on the command line. This exercises the "found" branch of os.LookupEnv in initializeConfig,
// which previously only had incidental coverage in CI because GitHub Actions runners happen to
// export a RUNNER_NAME environment variable that collides with this project's "runner-name"
// flag. Relying on that ambient variable made coverage of this branch environment-dependent:
// it silently dropped to 0% whenever tests ran outside GitHub Actions (see flag.go).
func TestInitializeConfig_AppliesUnsetEnvVars(t *testing.T) {
	t.Setenv("TESTFROMENV", "value-from-env")

	cmd := new(cobra.Command)
	cmd.Use = testCommandUse
	cmd.Flags().String("testfromenv", "default", "test string flag")

	err := initializeConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := cmd.Flags().Lookup("testfromenv").Value.String(); got != "value-from-env" {
		t.Fatalf("expected flag to be set from environment variable, got %q", got)
	}
}

// TestInitializeConfig_InvalidEnvVarValueIsIgnored verifies that initializeConfig does not
// fail when the environment variable's value is rejected by the flag's Set method (e.g. an
// invalid value for a typed flag). The flag.Set error is intentionally discarded, so the flag
// must retain its previous value and initializeConfig must still return a nil error.
func TestInitializeConfig_InvalidEnvVarValueIsIgnored(t *testing.T) {
	t.Setenv("TESTBOOLFLAG", "not-a-bool")

	cmd := new(cobra.Command)
	cmd.Use = testCommandUse
	cmd.Flags().Bool("testboolflag", false, "test bool flag")

	err := initializeConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := cmd.Flags().Lookup("testboolflag").Value.String(); got != "false" {
		t.Fatalf("expected flag to remain at its previous value %q, got %q", "false", got)
	}
}
