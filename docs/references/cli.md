# CLI reference

Complete command-line reference for `kar`.

## Synopsis

```shell
kar [flags]
```

## Flags

| Flag                               | Short | Default       | Description                                                         |
| ---------------------------------- | ----- | ------------- | ------------------------------------------------------------------- |
| `--kubevirt-vm-template`           | `-t`  | `vm-template` | VirtualMachine template name used to create VirtualMachineInstances |
| `--kubevirt-vm-template-namespace` | `-n`  | `default`     | Namespace where the VirtualMachine template exists                  |
| `--runner-name`                    | `-r`  | `runner`      | Runner name used for generated resources                            |
| `--actions-runner-input-jitconfig` | `-c`  | empty         | Opaque just-in-time runner configuration payload                    |

## Environment variable mapping for flags

Flag values can be provided through environment variables.
The CLI maps flag names by replacing `-` with `_`
and using uppercase names.

Examples:

- `KUBEVIRT_VM_TEMPLATE` maps to `--kubevirt-vm-template`
- `KUBEVIRT_VM_TEMPLATE_NAMESPACE` maps to `--kubevirt-vm-template-namespace`
- `RUNNER_NAME` maps to `--runner-name`
- `ACTIONS_RUNNER_INPUT_JITCONFIG` maps to `--actions-runner-input-jitconfig`

If both a flag and an environment variable are provided,
the explicit flag value is used.

## Runtime behavior

Execution sequence:

1. Create VM-backed runner resources.
1. Wait for the target VirtualMachineInstance to complete.
1. Delete resources created by the runner.

When interrupted by `SIGTERM` or `Ctrl-C`,
the runner enters cleanup and attempts to remove created resources
within the configured cleanup timeout.

## Centralized template strategy

`--kubevirt-vm-template-namespace` lets you retrieve the VM template from a namespace
that is different from the runner execution namespace.
Use this to centralize templates as a single golden image source,
reduce template duplication,
and keep template lifecycle management in one place.
