# OpenTelemetry Integration

This document describes how to enable and configure OpenTelemetry (OTEL) tracing for the KubeVirt Actions Runner application.

## Overview

OpenTelemetry telemetry is integrated into the application to provide distributed tracing capabilities. The application exports traces to either OTLP (OpenTelemetry Protocol) endpoints or to stdout for local testing.

## Features

- **Distributed Tracing**: Track the execution of key operations:
  - Resource creation (VirtualMachineInstance and DataVolume)
  - Virtual machine instance monitoring
  - Resource cleanup and deletion
- **Flexible Exporters**: Support for multiple export backends:
  - OTLP HTTP exporter (for production telemetry backends like Jaeger, DataDog, etc.)
  - Stdout exporter (for local development and debugging)

## Configuration

Telemetry is configured via environment variables:

| Environment Variable            | Default                   | Description                                                |
| ------------------------------- | ------------------------- | ---------------------------------------------------------- |
| `KAR_TELEMETRY_ENABLED`         | `false`                   | Enable or disable telemetry (`true` or `false`)            |
| `KAR_TELEMETRY_EXPORT_TYPE`     | ``                        | Export type: `otlp` or `stdout`                            |
| `KAR_TELEMETRY_OTLP_ENDPOINT`   | `http://localhost:4318`   | OTLP collector HTTP endpoint (only for `otlp` export type) |
| `KAR_TELEMETRY_SERVICE_NAME`    | `kubevirt-actions-runner` | Service name for telemetry                                 |
| `KAR_TELEMETRY_SERVICE_VERSION` | `unknown`                 | Service version for telemetry                              |

## Usage

### Enable Telemetry with OTLP Exporter

```bash
export KAR_TELEMETRY_ENABLED=true
export KAR_TELEMETRY_EXPORT_TYPE=otlp
export KAR_TELEMETRY_OTLP_ENDPOINT=http://jaeger-collector:4318
export KAR_TELEMETRY_SERVICE_NAME=my-runner-service
export KAR_TELEMETRY_SERVICE_VERSION=1.0.0
./kar -t vm-template -r my-runner -c '<jit-config>'
```

This will send traces to a Jaeger or other OTLP-compatible collector.

## Traces Captured

The following operations are traced:

### CreateResources

- **Span Name**: `CreateResources`
- **Attributes**: `vmTemplate`, `runnerName`, `namespace`
- **Child Spans**:
  - `CreateVMI`: Creating the VirtualMachineInstance
  - `CreateDataVolume`: Creating the DataVolume (if applicable)
- **Events**: Errors are recorded if any occur

### WaitForVirtualMachineInstance

- **Span Name**: `WaitForVirtualMachineInstance`
- **Attributes**: `vmiName`
- **Events**: Phase transitions are recorded (Pending, Scheduled, Running, Succeeded, Failed)

### DeleteResources

- **Span Name**: `DeleteResources`
- **Attributes**: `vmiName`
- **Child Spans**:
  - `DeleteDataVolume`: Deleting the DataVolume (if applicable)
- **Events**: Errors are recorded if any occur

## Integration with Observability Platforms

### Jaeger

To send traces to Jaeger:

```bash
# Start Jaeger locally
docker run -d --name jaeger \
  -p 4318:4318 \
  -p 16686:16686 \
  jaegertracing/all-in-one

# Run the application
export KAR_TELEMETRY_ENABLED=true
export KAR_TELEMETRY_EXPORT_TYPE=otlp
export KAR_TELEMETRY_OTLP_ENDPOINT=http://localhost:4318
./kar -t vm-template -r my-runner -c '<jit-config>'

# View traces at http://localhost:16686
```
