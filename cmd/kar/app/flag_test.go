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

	cmd := &cobra.Command{Use: testCommandUse}
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

	cmd := &cobra.Command{Use: testCommandUse}
	cmd.Flags().String("testunset", "default", "test string flag")

	err := initializeConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := cmd.Flags().Lookup("testunset").Value.String(); got != "default" {
		t.Fatalf("expected flag to remain default, got %q", got)
	}
}
