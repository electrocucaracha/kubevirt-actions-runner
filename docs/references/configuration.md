# Configuration reference

`kubevirt-actions-runner` reads runtime configuration from environment variables.
This page lists supported variables,
their defaults,
and accepted values.

## Timeout configuration

| Variable              | Default  | Description                                                         |
| --------------------- | -------- | ------------------------------------------------------------------- |
| `KAR_WAIT_TIMEOUT`    | `1h0m0s` | Maximum wait time for terminal VMI phases (`Succeeded` or `Failed`) |
| `KAR_CLEANUP_TIMEOUT` | `5m0s`   | Maximum time allotted to resource cleanup                           |

`KAR_WAIT_TIMEOUT` and `KAR_CLEANUP_TIMEOUT` accept
[Go duration](https://pkg.go.dev/time#ParseDuration)
format,
for example `90s`, `15m`, or `2h`.
If a value is invalid,
the default is used.

## Telemetry configuration

| Variable                        | Default                   | Description                                     |
| ------------------------------- | ------------------------- | ----------------------------------------------- |
| `KAR_TELEMETRY_ENABLED`         | `false`                   | Enables telemetry when set to `true`            |
| `KAR_TELEMETRY_EXPORT_TYPE`     | empty                     | Exporter type: `otlp` or `stdout`               |
| `KAR_TELEMETRY_OTLP_ENDPOINT`   | `http://localhost:4318`   | HTTP endpoint for OpenTelemetry Protocol export |
| `KAR_TELEMETRY_SERVICE_NAME`    | `kubevirt-actions-runner` | Telemetry service name                          |
| `KAR_TELEMETRY_SERVICE_VERSION` | `unknown`                 | Telemetry service version                       |

## Logging configuration

| Variable        | Default | Description                                   |
| --------------- | ------- | --------------------------------------------- |
| `KAR_LOG_LEVEL` | `info`  | Log verbosity level used by the runner logger |

## Related guides

- To configure telemetry in practice,
  see [Enable telemetry](../how-to-guides/enable-telemetry.md).
- To tune timeout behavior,
  see [Configure runner timeouts](../how-to-guides/configure-timeouts.md).
