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

package utils_test

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
)

// TestFatalExitsProcess verifies that Fatal terminates the process with a
// non-zero exit code. Since zap's Fatal calls os.Exit(1), the assertion runs
// in a subprocess to avoid killing the test binary itself.
func TestFatalExitsProcess(t *testing.T) {
	t.Parallel()

	if os.Getenv("KAR_TEST_INVOKE_FATAL") == "1" {
		utils.GetLogger().Fatal("simulated fatal error")

		return
	}

	//nolint:gosec
	cmd := exec.CommandContext(context.Background(), os.Args[0], "-test.run=TestFatalExitsProcess")

	cmd.Env = append(os.Environ(), "KAR_TEST_INVOKE_FATAL=1")

	err := cmd.Run()

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		t.Fatalf("expected process to exit with a non-zero status, got err=%v", err)
	}

	if exitErr.ExitCode() != 1 {
		t.Fatalf("expected exit code 1, got %d", exitErr.ExitCode())
	}
}

func TestLoggerMethods(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	tests := []struct {
		name string
		run  func()
	}{
		{name: "Printf", run: func() { logger.Printf("test %s", "message") }},
		{name: "Println", run: func() { logger.Println("test message") }},
		{name: "Infof", run: func() { logger.Infof("test %s", "info") }},
		{name: "Warnf", run: func() { logger.Warnf("test %s", "warning") }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			test.run()
		})
	}
}

func verifyLoggerImpl(t *testing.T) {
	t.Helper()
	utils.ResetLoggerForTesting()

	logger := utils.GetLogger()
	if logger == nil {
		t.Fatal("GetLogger returned nil")
	}
}

func TestGetLoggerLevels(t *testing.T) {
	tests := []struct {
		name  string
		level string
	}{
		{name: "DefaultLevel", level: ""},
		{name: "DebugLevel", level: "debug"},
		{name: "InfoLevel", level: "info"},
		{name: "WarnLevel", level: "warn"},
		{name: "WarningLevel", level: "warning"},
		{name: "ErrorLevel", level: "error"},
		{name: "FatalLevel", level: "fatal"},
		{name: "UnknownLevel", level: "unknown"},
		{name: "UppercaseLevel", level: "DEBUG"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("KAR_LOG_LEVEL", test.level)
			verifyLoggerImpl(t)
		})
	}
}

func TestGetLoggerSingleton(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "info")

	utils.ResetLoggerForTesting()

	logger1 := utils.GetLogger()
	logger2 := utils.GetLogger()

	if logger1 != logger2 {
		t.Fatal("GetLogger should return the same instance")
	}
}
