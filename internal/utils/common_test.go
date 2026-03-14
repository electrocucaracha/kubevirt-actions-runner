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
	logger := utils.GetLogger()

	tests := []struct {
		name string
		call func(*utils.LoggerImpl)
	}{
		{name: "Printf", call: func(log *utils.LoggerImpl) { log.Printf("test %s", "message") }},
		{name: "Println", call: func(log *utils.LoggerImpl) { log.Println("test message") }},
		{name: "Debugf", call: func(log *utils.LoggerImpl) { log.Debugf("test %s", "debug") }},
		{name: "Debug", call: func(log *utils.LoggerImpl) { log.Debug("test debug") }},
		{name: "Infof", call: func(log *utils.LoggerImpl) { log.Infof("test %s", "info") }},
		{name: "Info", call: func(log *utils.LoggerImpl) { log.Info("test info") }},
		{name: "Infow", call: func(log *utils.LoggerImpl) { log.Infow("test message", "key1", "value1", "key2", "value2") }},
		{name: "Warnf", call: func(log *utils.LoggerImpl) { log.Warnf("test %s", "warning") }},
		{name: "Warn", call: func(log *utils.LoggerImpl) { log.Warn("test warning") }},
		{name: "Warnw", call: func(log *utils.LoggerImpl) { log.Warnw("test message", "key1", "value1") }},
		{name: "Errorf", call: func(log *utils.LoggerImpl) { log.Errorf("test %s", "error") }},
		{name: "Error", call: func(log *utils.LoggerImpl) { log.Error("test error") }},
		{name: "Errorw", call: func(log *utils.LoggerImpl) { log.Errorw("test message", "key1", "value1") }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.call(logger)
		})
	}
}

func TestGetLoggerLevels(t *testing.T) {
	levels := []string{"", "debug", "info", "warn", "warning", "error", "fatal", "unknown", "DEBUG"}

	for _, level := range levels {
		t.Run(levelName(level), func(t *testing.T) {
			t.Setenv("KAR_LOG_LEVEL", level)
			utils.ResetLoggerForTesting()

			if utils.GetLogger() == nil {
				t.Fatal("GetLogger returned nil")
			}
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

func levelName(level string) string {
	if level == "" {
		return "default"
	}

	return level
}
