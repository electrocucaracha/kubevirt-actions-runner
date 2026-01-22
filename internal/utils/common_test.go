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
	"reflect"
	"testing"

	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
)

func TestLoggerImplPrintf(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Printf("test %s", "message")
}

func TestLoggerImplPrintln(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Println("test message")
}

func TestLoggerImplDebugf(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Debugf("test %s", "debug")
}

func TestLoggerImplDebug(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Debug("test debug")
}

func TestLoggerImplInfof(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Infof("test %s", "info")
}

func TestLoggerImplInfo(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Info("test info")
}

func TestLoggerImplInfow(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Infow("test message", "key1", "value1", "key2", "value2")
}

func TestLoggerImplWarnf(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Warnf("test %s", "warning")
}

func TestLoggerImplWarn(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Warn("test warning")
}

func TestLoggerImplWarnw(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Warnw("test message", "key1", "value1")
}

func TestLoggerImplErrorf(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Errorf("test %s", "error")
}

func TestLoggerImplError(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Error("test error")
}

func TestLoggerImplErrorw(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()
	// Should not panic
	logger.Errorw("test message", "key1", "value1")
}

func verifyLoggerImpl(t *testing.T) {
	t.Helper()
	// Need to reset the singleton for testing
	updateLoggerForTesting()

	logger := utils.GetLogger()
	if logger == nil {
		t.Fatal("GetLogger returned nil")
	}

	if reflect.TypeFor[*utils.LoggerImpl]().String() != utils.LoggerImplTypeString {
		t.Errorf("Expected *utils.LoggerImpl, got %s", reflect.TypeFor[*utils.LoggerImpl]().String())
	}
}

func TestGetLoggerDefaultLevel(t *testing.T) {
	// Clear the environment variable and reset the singleton
	t.Setenv("KAR_LOG_LEVEL", "")
	verifyLoggerImpl(t)
}

func TestGetLoggerDebugLevel(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "debug")
	verifyLoggerImpl(t)
}

func TestGetLoggerInfoLevel(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "info")
	verifyLoggerImpl(t)
}

func TestGetLoggerWarnLevel(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "warn")
	verifyLoggerImpl(t)
}

func TestGetLoggerWarningLevel(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "warning")
	verifyLoggerImpl(t)
}

func TestGetLoggerErrorLevel(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "error")
	verifyLoggerImpl(t)
}

func TestGetLoggerFatalLevel(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "fatal")
	verifyLoggerImpl(t)
}

func TestGetLoggerUnknownLevel(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "unknown")
	verifyLoggerImpl(t)
}

func TestGetLoggerUppercaseLevel(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "DEBUG")
	verifyLoggerImpl(t)
}

func TestGetLoggerSingleton(t *testing.T) {
	t.Setenv("KAR_LOG_LEVEL", "info")

	// Need to reset the singleton for testing
	updateLoggerForTesting()

	logger1 := utils.GetLogger()
	logger2 := utils.GetLogger()

	if logger1 != logger2 {
		t.Fatal("GetLogger should return the same instance")
	}
}

func TestLoggerInterface(t *testing.T) {
	t.Parallel()

	logger := utils.GetLogger()

	// Verify logger implements Logger interface
	var _ utils.Logger = logger
}

func updateLoggerForTesting() {
	utils.ResetLoggerForTesting()
}
