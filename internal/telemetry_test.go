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

package runner_test

import (
	"context"
	"testing"

	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
)

const (
	testServiceName    = "test-service"
	testServiceVersion = "1.0.0"
)

func requireShutdownNoError(t *testing.T, shutdown func(context.Context) error) {
	t.Helper()

	err := shutdown(context.Background())
	if err != nil {
		t.Fatalf("shutdown returned unexpected error: %v", err)
	}
}

func TestInitializeTelemetry_DisabledByDefault(t *testing.T) {
	t.Parallel()

	shutdown, err := runner.InitializeTelemetry(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_WithEnvVars(t *testing.T) {
	t.Setenv("KAR_TELEMETRY_ENABLED", "true")
	t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "stdout")
	t.Setenv("KAR_TELEMETRY_SERVICE_NAME", "my-custom-service")
	t.Setenv("KAR_TELEMETRY_SERVICE_VERSION", "2.0.0")

	shutdown, err := runner.InitializeTelemetry(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_EnabledFalseByDefault(t *testing.T) {
	t.Setenv("KAR_TELEMETRY_ENABLED", "false")

	shutdown, err := runner.InitializeTelemetry(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_StdoutExporter(t *testing.T) {
	t.Setenv("KAR_TELEMETRY_ENABLED", "true")
	t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "stdout")
	t.Setenv("KAR_TELEMETRY_SERVICE_NAME", testServiceName)
	t.Setenv("KAR_TELEMETRY_SERVICE_VERSION", testServiceVersion)

	shutdown, err := runner.InitializeTelemetry(context.Background())
	if err != nil {
		t.Fatalf("unexpected error initializing telemetry: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_UnknownExportType(t *testing.T) {
	t.Setenv("KAR_TELEMETRY_ENABLED", "true")
	t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "unknown-exporter")
	t.Setenv("KAR_TELEMETRY_SERVICE_NAME", testServiceName)
	t.Setenv("KAR_TELEMETRY_SERVICE_VERSION", testServiceVersion)

	// Unknown exporter falls back to stdout and should not error.
	shutdown, err := runner.InitializeTelemetry(context.Background())
	if err != nil {
		t.Fatalf("unexpected error for unknown exporter: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_EmptyExportType(t *testing.T) {
	t.Setenv("KAR_TELEMETRY_ENABLED", "true")
	t.Setenv("KAR_TELEMETRY_SERVICE_NAME", testServiceName)
	t.Setenv("KAR_TELEMETRY_SERVICE_VERSION", testServiceVersion)

	// Empty export type falls back to stdout without a warning log.
	shutdown, err := runner.InitializeTelemetry(context.Background())
	if err != nil {
		t.Fatalf("unexpected error for empty exporter type: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_OTLPExporter(t *testing.T) {
	// The OTLP exporter is created with a fake endpoint. The SDK connects lazily,
	// so creation succeeds; the shutdown may fail to flush but that is tolerated.
	t.Setenv("KAR_TELEMETRY_ENABLED", "true")
	t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "otlp")
	t.Setenv("KAR_TELEMETRY_OTLP_ENDPOINT", "http://localhost:19999")
	t.Setenv("KAR_TELEMETRY_SERVICE_NAME", testServiceName)
	t.Setenv("KAR_TELEMETRY_SERVICE_VERSION", testServiceVersion)

	shutdown, err := runner.InitializeTelemetry(context.Background())
	if err != nil {
		t.Fatalf("unexpected error initializing OTLP telemetry: %v", err)
	}

	// Ignore shutdown error; the fake endpoint will reject the flush.
	_ = shutdown(context.Background())
}
