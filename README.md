# KubeVirt Actions Runner

<!-- markdown-link-check-disable-next-line -->

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GitHub Super-Linter](https://github.com/electrocucaracha/kubevirt-actions-runner/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/marketplace/actions/super-linter)

<!-- markdown-link-check-disable-next-line -->

![visitors](https://visitor-badge.laobi.icu/badge?page_id=electrocucaracha.kubevirt-actions-runner)
[![Scc Code Badge](https://sloc.xyz/github/electrocucaracha/kubevirt-actions-runner?category=code)](https://github.com/boyter/scc/)
[![Scc COCOMO Badge](https://sloc.xyz/github/electrocucaracha/kubevirt-actions-runner?category=cocomo)](https://github.com/boyter/scc/)

## Overview

The `kubevirt-actions-runner` project provides a robust solution for running GitHub Actions
workflows in isolated, ephemeral virtual machines using [KubeVirt](https://kubevirt.io/).
By integrating with [Actions Runner Controller (ARC)](https://github.com/actions/actions-runner-controller),
this project enables you to execute CI/CD jobs in highly customizable VM environments with
enhanced security and flexibility.

This project acts as a bridge between Kubernetes and GitHub Actions, automatically provisioning dedicated VMs for each workflow job and cleaning them up after completion.

![Diagram](docs/assets/diagram.png)

## Reasons to Use KubeVirt Actions Runner

While GitHub-hosted runners work well for standard workflows, they have limitations. This project addresses these limitations by offering:

- **Custom Environments**: Run jobs requiring specific kernel modules, system services, or custom OS configurations.
- **Enhanced Isolation**: Execute untrusted code or security-sensitive workflows in fully isolated VMs.
- **OS Flexibility**: Support for Windows VMs and other operating systems beyond standard Linux containers.
- **Ephemeral Instances**: Fresh, clean VM for every job run ensures reproducibility.
- **System-Level Control**: Full control over VM resources, storage, and network configuration.

## Key Features

- **Ephemeral VM Creation**: Automatically provisions and destroys VMs for each job.
- **Kubernetes-Native**: Seamless integration with Kubernetes clusters and ARC.
- **Customizable**: Tailor VM specifications, resources, and configurations per workflow.
- **Lifecycle Management**: Automatic VM cleanup and resource management.
- **Cloud-Native**: Built with Go and follows cloud-native best practices.

## Getting Started

### Prerequisites

- Kubernetes cluster (v1.24+)
- KubeVirt installed and configured
- Actions Runner Controller (ARC) deployed
- GitHub Personal Access Token (PAT) with appropriate permissions

## Documentation

Full documentation is available at the [official site](https://electrocucaracha.github.io/kubevirt-actions-runner/).

## Learn More

For a detailed walkthrough of the project, check out the following resources:

- **[KCD Guadalajara 2025 Presentation](https://www.slideshare.net/slideshow/migrating-github-actions-with-nested-virtualization-to-cloud-native-ecosystem-pptx/277448656)**: Presentation materials
- **[Video Recording](https://www.youtube.com/watch?v=ccb8y_Ij30k)**: Watch the full presentation

## Contributing

We welcome contributions from the community, including:

- Bug reports and fixes
- New features
- Documentation improvements
- Ideas and suggestions

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to get involved.

## License

This project is licensed under the Apache License 2.0. See the [license](https://opensource.org/licenses/Apache-2.0) for details.
