# Contributing to Kubevirt Actions Runner

Thank you for your interest in contributing to the Kubevirt Actions Runner project. This document provides clear guidelines and instructions for developers to ensure a smooth contribution process.

## Development Setup

This project includes **Dev Container** support, providing a fully configured development environment. This ensures consistency across setups and allows contributors to start working without manual configuration.

### Option 1: GitHub Codespaces

Launch a cloud-based development environment instantly:

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://github.com/codespaces/new?repo=electrocucaracha/kubevirt-actions-runner)

### Option 2: Local Dev Container (Visual Studio Code)

1. Install [Visual Studio Code](https://code.visualstudio.com/).
2. Install the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).
3. Ensure you have a container runtime installed and running:
   - [Docker Desktop](https://www.docker.com/products/docker-desktop/) (Windows/macOS)
   - [Docker Engine](https://docs.docker.com/engine/install/) or [Podman](https://podman.io/) (Linux)
4. Open this repository in Visual Studio Code.
5. When prompted, reopen the project in a Dev Container.

## Code Standards

### Style Guidelines

Follow the coding standards described in the [Go instructions file](.github/instructions/go.instructions.md). These include:

- Writing simple, clear, idiomatic Go code per Effective Go and Google style.
- Favoring clarity, early returns, and useful zero values.
- Using proper naming conventions (e.g., lowercase packages, MixedCaps, small interfaces).
- Organizing modules effectively and testing with table-driven tests.
- Managing concurrency safely and validating input.
- Documenting exported APIs clearly.

### Running Code Quality Checks

Run linting validation:

```bash
make lint
```

### Running Tests

Execute all tests:

```bash
make test
```

#### Test Requirements

- Define clear test scopes: unit, integration, and (when needed) end-to-end tests, aligned with package boundaries and public APIs.
- Write table-driven tests with descriptive names, covering both happy paths and error scenarios.
- Use subtests (`t.Run`) for structure, and isolate dependencies with small interfaces and test doubles.
- Validate behavior, not implementation details; assert on observable outputs and side effects.
- Ensure determinism: avoid flaky tests by controlling time, randomness, and concurrency.
- Run `go test`, linters, and the race detector (`-race`) in CI; keep tests fast, reliable, and close to the code they verify.

## Commits and Pull Requests

### Commit Messages

Use the Conventional Commits specification for commit messages:

- **Format**: `<type>(<scope>): <description>`
- **Types**: feat, fix, docs, style, refactor, test, chore
- **Examples**:
  - `feat(notion): add support for custom database queries`
  - `fix(youtube): handle unavailable transcripts gracefully`
  - `docs: update installation instructions`

### Pull Request Process

1. Create a feature branch from `master`.
2. Make your changes and test thoroughly.
3. Ensure all tests pass and there are no linting issues.
4. Push to your fork and submit a pull request.
5. In the PR description:
   - Explain the changes clearly.
   - Reference related issues.
   - Describe any configuration changes needed.

### Review Process

- All PRs require code review before merging.
- Address feedback promptly.
- Ensure automated checks pass before merging.

## Documentation

- Update `README.md` if adding features or changing behavior.
- Add new documents to the `docs/` folder.
- Follow the [documentation-writer skill](.github/skills/documentation-writer/SKILL.md) for creating high-quality documentation.

## Reporting Issues

When reporting bugs:

- Describe the problem clearly.
- Include steps to reproduce.
- Provide environment details (e.g., Go version, Kubernetes version, KubeVirt, GitHub Runner, OS).
- Share error messages and logs.

## License

By contributing to this project, you agree that your contributions will be licensed under the Apache License 2.0.
