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
	"github.com/spf13/viper"
)

const testCommandUse = "test"

// TestBindFlags_ErrorWhenSetFails verifies that bindFlags returns an error
// when a value sourced from the environment cannot be applied to the flag,
// e.g. a non-numeric string bound to an integer flag.
func TestBindFlags_ErrorWhenSetFails(t *testing.T) {
	t.Setenv("TESTINT", "not-an-int")

	cmd := &cobra.Command{Use: testCommandUse}
	cmd.Flags().Int("testint", 0, "test int flag")

	viperInstance := viper.New()
	viperInstance.AutomaticEnv()

	err := viperInstance.BindEnv("testint")
	if err != nil {
		t.Fatalf("failed to bind env var: %v", err)
	}

	err = bindFlags(cmd, viperInstance)
	if err == nil {
		t.Fatal("expected an error when the flag value cannot be applied")
	}
}

// TestBindFlags_SkipsChangedFlags verifies that bindFlags does not override
// a flag value that was already explicitly set on the command line.
func TestBindFlags_SkipsChangedFlags(t *testing.T) {
	t.Setenv("TESTSTR", "from-env")

	cmd := &cobra.Command{Use: testCommandUse}
	cmd.Flags().String("teststr", "default", "test string flag")

	err := cmd.Flags().Set("teststr", "from-cli")
	if err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	viperInstance := viper.New()
	viperInstance.AutomaticEnv()

	err = viperInstance.BindEnv("teststr")
	if err != nil {
		t.Fatalf("failed to bind env var: %v", err)
	}

	err = bindFlags(cmd, viperInstance)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := cmd.Flags().Lookup("teststr").Value.String(); got != "from-cli" {
		t.Fatalf("expected flag to remain %q, got %q", "from-cli", got)
	}
}

// TestBindFlags_SkipsUnsetEnvVars verifies that bindFlags leaves the flag at
// its default value when no corresponding environment variable is set.
func TestBindFlags_SkipsUnsetEnvVars(t *testing.T) {
	t.Parallel()

	cmd := &cobra.Command{Use: testCommandUse}
	cmd.Flags().String("testunset", "default", "test string flag")

	viperInstance := viper.New()
	viperInstance.AutomaticEnv()

	err := bindFlags(cmd, viperInstance)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := cmd.Flags().Lookup("testunset").Value.String(); got != "default" {
		t.Fatalf("expected flag to remain default, got %q", got)
	}
}
