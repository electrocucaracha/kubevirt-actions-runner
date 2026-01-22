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

const LoggerImplTypeString = "*utils.LoggerImpl"

// Logger is an interface for structured logging with multiple formatting options.
//
//nolint:interfacebloat
type Logger interface {
	Printf(format string, args ...any)
	Println(args ...any)
	Debugf(format string, args ...any)
	Debug(args ...any)
	Infof(format string, args ...any)
	Info(args ...any)
	Infow(msg string, keysAndValues ...any)
	Warnf(format string, args ...any)
	Warn(args ...any)
	Warnw(msg string, keysAndValues ...any)
	Errorf(format string, args ...any)
	Error(args ...any)
	Errorw(msg string, keysAndValues ...any)
	Fatalf(format string, args ...any)
	Fatal(args ...any)
}

type LoggerImpl struct {
	logger *zap.SugaredLogger
}

func (l *LoggerImpl) Printf(format string, args ...any) {
	l.logger.Infof(format, args...)
}

func (l *LoggerImpl) Println(args ...any) {
	l.logger.Info(args...)
}

func (l *LoggerImpl) Debugf(format string, args ...any) {
	l.logger.Debugf(format, args...)
}

func (l *LoggerImpl) Debug(args ...any) {
	l.logger.Debug(args...)
}

func (l *LoggerImpl) Infof(format string, args ...any) {
	l.logger.Infof(format, args...)
}

func (l *LoggerImpl) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *LoggerImpl) Infow(msg string, keysAndValues ...any) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *LoggerImpl) Warnf(format string, args ...any) {
	l.logger.Warnf(format, args...)
}

func (l *LoggerImpl) Warn(args ...any) {
	l.logger.Warn(args...)
}

func (l *LoggerImpl) Warnw(msg string, keysAndValues ...any) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *LoggerImpl) Errorf(format string, args ...any) {
	l.logger.Errorf(format, args...)
}

func (l *LoggerImpl) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *LoggerImpl) Errorw(msg string, keysAndValues ...any) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *LoggerImpl) Fatalf(format string, args ...any) {
	l.logger.Fatalf(format, args...)
}

func (l *LoggerImpl) Fatal(args ...any) {
	l.logger.Fatal(args...)
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

		logger, err := config.Build()
		if err != nil {
			devCfg := zap.NewDevelopmentConfig()
			logger, _ = devCfg.Build()
		}

		loggerInstance = &LoggerImpl{
			logger: logger.Sugar(),
		}
	})

	return loggerInstance
}

// ResetLoggerForTesting resets the logger singleton for testing purposes.
func ResetLoggerForTesting() {
	loggerInstance = nil
	loggerOnce = sync.Once{}
}
