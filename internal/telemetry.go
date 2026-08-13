/* jscpd:ignore-start */
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
/* jscpd:ignore-end */

package runner

import (
	"context"
	"fmt"
	"os"

	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

func getEnvOrDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}

// InitializeTelemetry sets up OpenTelemetry tracing from environment variables.
// Returns a shutdown function that should be called before the application exits.
func InitializeTelemetry(ctx context.Context) (func(context.Context) error, error) {
	log := utils.GetLogger()

	enabled := os.Getenv("KAR_TELEMETRY_ENABLED") == "true"
	if !enabled {
		log.Infof("Telemetry is disabled")

		return func(_ context.Context) error { return nil }, nil
	}

	exportType := os.Getenv("KAR_TELEMETRY_EXPORT_TYPE")
	log.Infof("Initializing telemetry with export type: %s", exportType)

	serviceName := getEnvOrDefault("KAR_TELEMETRY_SERVICE_NAME", "kubevirt-actions-runner")
	serviceVersion := getEnvOrDefault("KAR_TELEMETRY_SERVICE_VERSION", "unknown")

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(serviceVersion),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	exporter, err := createExporter(ctx, exportType)
	if err != nil {
		return func(_ context.Context) error { return nil }, err
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

func createExporter(ctx context.Context, exportType string) (trace.SpanExporter, error) {
	log := utils.GetLogger()

	switch exportType {
	case "otlp":
		endpoint := getEnvOrDefault("KAR_TELEMETRY_OTLP_ENDPOINT", "http://localhost:4318")
		log.Infof("Using OTLP exporter with endpoint: %s", endpoint)

		exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(endpoint))
		if err != nil {
			return nil, fmt.Errorf("failed to create OTLP exporter: %w", err)
		}

		return exporter, nil
	default:
		if exportType != "" {
			log.Warnf("Unknown export type: %s, using stdout", exportType)
		}

		log.Infof("Using stdout exporter")

		exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			return nil, fmt.Errorf("failed to create stdout exporter: %w", err)
		}

		return exporter, nil
	}
}
