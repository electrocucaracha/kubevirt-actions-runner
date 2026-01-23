/*
Copyright © 2025

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

package runner

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

// TelemetryConfig holds telemetry configuration.
type TelemetryConfig struct {
	// Enabled enables or disables telemetry.
	Enabled bool
	// ExportType specifies the export type: "otlp", "stdout", or "" for disabled.
	ExportType string
	// OTLPEndpoint is the OTLP collector endpoint (e.g., http://localhost:4318).
	OTLPEndpoint string
	// ServiceName is the service name for telemetry.
	ServiceName string
	// ServiceVersion is the service version for telemetry.
	ServiceVersion string
}

// GetTelemetryConfig returns the telemetry configuration from environment variables.
func GetTelemetryConfig() TelemetryConfig {
	return TelemetryConfig{
		Enabled:        os.Getenv("KAR_TELEMETRY_ENABLED") == "true",
		ExportType:     os.Getenv("KAR_TELEMETRY_EXPORT_TYPE"),   // "otlp" or "stdout"
		OTLPEndpoint:   os.Getenv("KAR_TELEMETRY_OTLP_ENDPOINT"), // e.g., "http://localhost:4318"
		ServiceName:    getEnvOrDefault("KAR_TELEMETRY_SERVICE_NAME", "kubevirt-actions-runner"),
		ServiceVersion: getEnvOrDefault("KAR_TELEMETRY_SERVICE_VERSION", "unknown"),
	}
}

func getEnvOrDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}

// InitializeTelemetry sets up OpenTelemetry tracing based on the configuration.
// Returns a shutdown function that should be called before the application exits.
func InitializeTelemetry(ctx context.Context, cfg TelemetryConfig) (func(context.Context) error, error) {
	log := GetLogger()

	if !cfg.Enabled {
		log.Infof("Telemetry is disabled")

		return func(_ context.Context) error { return nil }, nil
	}

	log.Infof("Initializing telemetry with export type: %s", cfg.ExportType)

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.ServiceName),
			semconv.ServiceVersionKey.String(cfg.ServiceVersion),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	exporter, err := createExporter(ctx, cfg)
	if err != nil {
		return nil, err
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)

	otel.SetTracerProvider(tracerProvider)

	log.Infof("Telemetry initialized successfully")

	return func(shutdownCtx context.Context) error {
		return tracerProvider.Shutdown(shutdownCtx)
	}, nil
}

func createExporter(ctx context.Context, cfg TelemetryConfig) (trace.SpanExporter, error) {
	log := GetLogger()

	switch cfg.ExportType {
	case "otlp":
		return createOTLPExporter(ctx, cfg)
	case "stdout":
		return createStdoutExporter()
	default:
		if cfg.ExportType != "" {
			log.Warnf("Unknown export type: %s, using stdout", cfg.ExportType)
		}

		return createStdoutExporter()
	}
}

func createOTLPExporter(ctx context.Context, cfg TelemetryConfig) (trace.SpanExporter, error) {
	log := GetLogger()

	endpoint := cfg.OTLPEndpoint
	if endpoint == "" {
		endpoint = "http://localhost:4318"
	}

	log.Infof("Using OTLP exporter with endpoint: %s", endpoint)

	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(endpoint),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP exporter: %w", err)
	}

	return exporter, nil
}

func createStdoutExporter() (trace.SpanExporter, error) {
	log := GetLogger()

	log.Infof("Using stdout exporter")

	exporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout exporter: %w", err)
	}

	return exporter, nil
}
