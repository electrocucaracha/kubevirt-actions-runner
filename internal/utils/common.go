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

// Package utils provides utility functions and logging capabilities.
package utils

import (
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerImpl struct {
	*zap.SugaredLogger
}

func (l *LoggerImpl) Printf(format string, args ...any) {
	l.Infof(format, args...)
}

func (l *LoggerImpl) Println(args ...any) {
	l.Info(args...)
}

//nolint:gochecknoglobals
var (
	loggerInstance *LoggerImpl
	loggerOnce     sync.Once
)

// GetLogger returns a singleton LoggerImpl instance. The singleton is
// initialized lazily on first call using the `KAR_LOG_LEVEL` environment
// variable.
func GetLogger() *LoggerImpl {
	level := os.Getenv("KAR_LOG_LEVEL")

	loggerOnce.Do(func() {
		var lvl zapcore.Level

		switch strings.ToLower(level) {
		case "debug":
			lvl = zapcore.DebugLevel
		case "info":
			lvl = zapcore.InfoLevel
		case "warn", "warning":
			lvl = zapcore.WarnLevel
		case "error":
			lvl = zapcore.ErrorLevel
		case "fatal":
			lvl = zapcore.FatalLevel
		default:
			lvl = zapcore.InfoLevel
		}

		config := zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(lvl)

		baseLogger, err := config.Build()
		if err != nil {
			baseLogger, err = zap.NewDevelopmentConfig().Build()
			if err != nil {
				baseLogger = zap.NewNop()
			}
		}

		loggerInstance = &LoggerImpl{SugaredLogger: baseLogger.Sugar()}
	})

	return loggerInstance
}

// ResetLoggerForTesting resets the logger singleton for testing purposes.
func ResetLoggerForTesting() {
	loggerInstance = nil
	loggerOnce = sync.Once{}
}
