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

func TestGetTelemetryConfig_Defaults(t *testing.T) {
	t.Parallel()

	cfg := runner.GetTelemetryConfig()

	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}

	if cfg.OTLPEndpoint != "http://localhost:4318" {
		t.Errorf("expected default OTLP endpoint, got %q", cfg.OTLPEndpoint)
	}

	if cfg.ServiceName != "kubevirt-actions-runner" {
		t.Errorf("expected default service name, got %q", cfg.ServiceName)
	}

	if cfg.ServiceVersion != "unknown" {
		t.Errorf("expected default service version, got %q", cfg.ServiceVersion)
	}

	if cfg.ExportType != "" {
		t.Errorf("expected empty ExportType by default, got %q", cfg.ExportType)
	}
}

func TestGetTelemetryConfig_WithEnvVars(t *testing.T) {
	t.Setenv("KAR_TELEMETRY_ENABLED", "true")
	t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "stdout")
	t.Setenv("KAR_TELEMETRY_OTLP_ENDPOINT", "http://custom-collector:4318")
	t.Setenv("KAR_TELEMETRY_SERVICE_NAME", "my-custom-service")
	t.Setenv("KAR_TELEMETRY_SERVICE_VERSION", "2.0.0")

	cfg := runner.GetTelemetryConfig()

	if !cfg.Enabled {
		t.Error("expected Enabled=true")
	}

	if cfg.ExportType != "stdout" {
		t.Errorf("expected ExportType=stdout, got %q", cfg.ExportType)
	}

	if cfg.OTLPEndpoint != "http://custom-collector:4318" {
		t.Errorf("expected custom OTLP endpoint, got %q", cfg.OTLPEndpoint)
	}

	if cfg.ServiceName != "my-custom-service" {
		t.Errorf("expected custom service name, got %q", cfg.ServiceName)
	}

	if cfg.ServiceVersion != "2.0.0" {
		t.Errorf("expected custom service version, got %q", cfg.ServiceVersion)
	}
}

func TestGetTelemetryConfig_EnabledFalseWhenNotTrue(t *testing.T) {
	t.Setenv("KAR_TELEMETRY_ENABLED", "false")

	cfg := runner.GetTelemetryConfig()

	if cfg.Enabled {
		t.Error("expected Enabled=false when env var is 'false'")
	}
}

func TestInitializeTelemetry_Disabled(t *testing.T) {
	t.Parallel()

	cfg := runner.TelemetryConfig{Enabled: false}

	shutdown, err := runner.InitializeTelemetry(context.Background(), cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_StdoutExporter(t *testing.T) {
	t.Parallel()

	cfg := runner.TelemetryConfig{
		Enabled:        true,
		ExportType:     "stdout",
		ServiceName:    testServiceName,
		ServiceVersion: testServiceVersion,
	}

	shutdown, err := runner.InitializeTelemetry(context.Background(), cfg)
	if err != nil {
		t.Fatalf("unexpected error initializing telemetry: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_UnknownExportType(t *testing.T) {
	t.Parallel()

	cfg := runner.TelemetryConfig{
		Enabled:        true,
		ExportType:     "unknown-exporter",
		ServiceName:    testServiceName,
		ServiceVersion: testServiceVersion,
	}

	// Unknown exporter falls back to stdout and should not error.
	shutdown, err := runner.InitializeTelemetry(context.Background(), cfg)
	if err != nil {
		t.Fatalf("unexpected error for unknown exporter: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_EmptyExportType(t *testing.T) {
	t.Parallel()

	cfg := runner.TelemetryConfig{
		Enabled:        true,
		ExportType:     "",
		ServiceName:    testServiceName,
		ServiceVersion: testServiceVersion,
	}

	// Empty export type falls back to stdout without a warning log.
	shutdown, err := runner.InitializeTelemetry(context.Background(), cfg)
	if err != nil {
		t.Fatalf("unexpected error for empty exporter type: %v", err)
	}

	requireShutdownNoError(t, shutdown)
}

func TestInitializeTelemetry_OTLPExporter(t *testing.T) {
	t.Parallel()

	// The OTLP exporter is created with a fake endpoint. The SDK connects lazily,
	// so creation succeeds; the shutdown may fail to flush but that is tolerated.
	cfg := runner.TelemetryConfig{
		Enabled:        true,
		ExportType:     "otlp",
		OTLPEndpoint:   "http://localhost:19999",
		ServiceName:    testServiceName,
		ServiceVersion: testServiceVersion,
	}

	shutdown, err := runner.InitializeTelemetry(context.Background(), cfg)
	if err != nil {
		t.Fatalf("unexpected error initializing OTLP telemetry: %v", err)
	}

	// Ignore shutdown error; the fake endpoint will reject the flush.
	_ = shutdown(context.Background())
}
