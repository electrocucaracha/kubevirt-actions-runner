# Architecture Overview

This document provides an overview of the architecture used in the testbed for
`kubevirt-actions-runner`.

## Components

- **kind**: Creates a local Kubernetes cluster running in Docker.
- **KubeVirt**: Enables VM workloads inside Kubernetes.
- **VM Template**: Cloned dynamically by the runner.
- **Demo Workload**: Validates VM lifecycle management.

## Deployment Flow

The following flowchart describes the complete deployment and runtime sequence:

```mermaid
flowchart TD
  A[Start: Run automation] --> B[Install tools]
  B --> C[Configure cluster]
  C --> D[Deploy KubeVirt operator]
  D --> E[Apply VM template]
  E --> F[Configure RBAC]
  F --> G[Deploy runner scale set]
  G --> H[Runner pod mounts runner-info]
  H --> I[Runner creates VirtualMachineInstance]
  I --> J[VM boots and executes job]
  J --> K[Job completes]
  K --> L[Teardown: delete VMI / pod]
```
