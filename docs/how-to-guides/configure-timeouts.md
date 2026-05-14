# How to configure runner timeouts

## Goal

This guide explains how to configure the VMI wait timeout and the cleanup timeout
for the `kubevirt-actions-runner` application.
Use these settings to tune behavior for long-running VM-backed GitHub Actions jobs
or to adjust the grace period given to resource cleanup.

## Prerequisites

- `kubevirt-actions-runner` is installed and functional.
- Basic knowledge of environment variables and shell commands.

## Configuration

Timeouts are configured via environment variables.
Below is a summary of the available options:

| Environment Variable  | Default | Description                                                                                                              |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------------------------------ |
| `KAR_WAIT_TIMEOUT`    | `10m0s` | Maximum time to wait for the VMI to reach `Running` phase with `Ready=True` condition before treating the job as failed. |
| `KAR_CLEANUP_TIMEOUT` | `5m0s`  | Maximum time allowed for resource cleanup (VMI and DataVolume deletion) after job completion.                            |

Both variables accept any valid Go duration string,
for example `30m`, `1h`, or `90s`.
Invalid values are logged and the default is used instead.

## VMI provisioning-success semantics

The runner considers a VM-backed job **successfully provisioned** as soon as
the VMI reaches the `Running` phase **and** the KubeVirt `Ready` condition
becomes `True`.
At that point the runner exits its wait loop and proceeds to cleanup,
without waiting for the VMI to reach the terminal `Succeeded` phase.

This change allows long-running jobs (jobs that keep the VM alive for many minutes
or hours) to be treated as successful once the VM is usable,
rather than requiring the entire VM lifecycle to complete within the wait timeout.

The runner still treats a VMI that enters the `Failed` phase as an error,
and it still recognizes `Succeeded` as a terminal success for cases where
the VMI completes normally before the `Running + Ready` signal is observed.

## Steps to configure `KAR_WAIT_TIMEOUT`

### 1. Set the environment variable

Increase the wait timeout to 30 minutes:

```bash
export KAR_WAIT_TIMEOUT=30m
```

Or reduce it for faster failure detection in environments where VMs boot quickly:

```bash
export KAR_WAIT_TIMEOUT=3m
```

### 2. Run the application

```bash
./kar -t vm-template -r my-runner -c '<jit-config>'
```

At startup the runner logs the effective timeout:

```text
wait timeout is set to: 30m0s
```

If an invalid value is provided, the runner logs a warning and falls back to
the default:

```text
Invalid KAR_WAIT_TIMEOUT value: "bad-value", using default 10m0s
```

## Steps to configure `KAR_CLEANUP_TIMEOUT`

### 1. Set `KAR_CLEANUP_TIMEOUT`

```bash
export KAR_CLEANUP_TIMEOUT=10m
```

### 2. Restart the runner

```bash
./kar -t vm-template -r my-runner -c '<jit-config>'
```

At startup the runner logs the effective timeout:

```text
cleanup timeout is set to: 10m0s
```

## Next Steps

- Enable telemetry to observe provisioning latency.
  See [How to enable Telemetry](enable-telemetry.md).
