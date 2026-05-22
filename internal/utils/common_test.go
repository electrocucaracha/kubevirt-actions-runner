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

package utils_test

import (
	"testing"

	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
)

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
