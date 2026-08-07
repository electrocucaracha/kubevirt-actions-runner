<!-- Markdownlint-disable MD024 -->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [18.9.1] - 2026-08-07

### Fixed

- Stabilized the fmt target by ensuring that formatting changes from golangci-lint are preserved in the final output through a consistent order of code formatting tool execution. [91b75ef2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/91b75ef21f20b9d2e6c93f133452f801bf924273)

## [18.9.0] - 2026-08-07

### Added

- Enabled enforcement of consistent terminology usage in documentation and prose by defaulting the terminology rule on in textlint configuration. [41b9df58](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/41b9df58e41171b78d4e7b2829f101783305fd41)

## [18.8.1] - 2026-08-07

### Changed

- Clarified terminology throughout the changelog to consistently reference tools and technologies such as GitHub Actions, README, Go linter, and spell checking, enhancing readability and professionalism without introducing any breaking behavior or migration requirements. [2f49d0c7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2f49d0c7a0d0d5c7bb23ba596ed940780b3d8e42)

## [18.8.0] - 2026-08-07

### Added

- Enabled more accurate spelling and linting tools by expanding the custom wordlist to include new terms, abbreviations, and project-specific identifiers that reduce false positives on domain-specific vocabulary. [643c0806](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/643c08061cf82ee17b14a1a837b851238cecf265)

## [18.7.0] - 2026-08-07

### Added

- Enabled clear and consistent release notes for users and maintainers through the introduction of a comprehensive changelog following the Keep a Changelog format and Semantic Versioning, which includes versioned entries with links to commits, summaries of changes, and context for each release. [224740a1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/224740a1bede9df8c51b7b063079c99934e65395)

## [18.6.0] - 2026-08-07

### Added

- Enabled more flexible Markdown writing by allowing second H1 headings and longer lines in documents, while maintaining Jekyll front-matter as the primary source for main H1 headings. [ddb2d6cb](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ddb2d6cb659fe3cdc492d95cd5bab4a4d0593868)

## [18.5.1] - 2026-08-07

### Changed

- Optimized the markdownlint hook to utilize the faster and more maintainable markdownlint-cli linter resulting in improved developer experience through reduced setup friction and a more efficient linting process. [b34b19ff](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b34b19ffe41f26fbded7c9b3a0bbccf487af670a)

## [18.5.0] - 2026-08-07

### Added

- Enabled automatic Go linting and style enforcement during formatting, ensuring code quality and consistency across the project by installing golangci-lint if missing and running it with autofix. [5bd8a1c0](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5bd8a1c0d8e68b74aaec0db8745e0ef2e62f015c)

## [18.4.0] - 2026-08-07

### Added

- Enhanced test case clarity by introducing empty lines after os.Args assignments to improve readability and separate setup steps without affecting API or CLI contracts. [ce71f230](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ce71f230ddffff728aa4cee724850626fa571e4b)

## [18.3.0] - 2026-08-07

### Added

- Enabled automatic enforcement of code quality during formatting by installing and running golangci-lint with autofix capabilities. [1dea3151](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1dea31513fd2a0a71dad18fdaa5bb09eb8a4090e)

## [18.2.0] - 2026-08-07

### Added

- Improved readability of test cases by introducing empty lines after os.Args assignments to enhance visual separation of setup steps without affecting API or CLI contracts and security implications. [8066fc30](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8066fc30e5de25547ab071825c777144e952eda5)

## [18.1.2] - 2026-08-07

### Changed

- Optimized the workflow frequency by restricting execution to Fridays that fall between the 15th and 21st of the month, effectively targeting the third Friday. [dc77fdfe](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/dc77fdfe2173331ffd57fcad07fd5ded122fbd17)

## [18.1.1] - 2026-08-07

### Changed

- Updated several GitHub Actions and Go dependencies to their latest patch versions for compatibility and security reasons, requiring downstream consumers to verify integration due to potential subtle API or behavior changes in updated dependencies. [b6d8fe1f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b6d8fe1f542823876cf4b0407414e8896a91245a)

## [18.1.0] - 2026-08-07

### Added

- Improved test coverage for cmd/kar main entrypoint by exercising previously untested paths, increasing statement coverage from 58.9% to 94.5% and mutation testing coverage from 90.91% to 95.45%, without modifying production code or changing existing behavior. [959b4bc7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/959b4bc759d910a75e1368d9eed0832a4bf9c9d0)

## [18.0.0] - 2026-08-07

### Removed

- Simplified the .golangci.yml configuration by removing unused dependencies github.com/golang/mock/gomock and github.com/pkg/errors that are no longer imported in the codebase. [ef5d590c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ef5d590c72fc705d697aa51e13e74105345dba78)

## [17.1.8] - 2026-07-31

### Changed

- Simplified error handling in tests by introducing shared variables for simulated failures, such as watch and transient get errors, reducing duplication and improving maintainability without changing functional behavior. [6413c56c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6413c56c73575ff51117df92718042727af391d7)

## [17.1.7] - 2026-07-31

### Changed

- Updated dependencies to their latest versions, enabling improved compatibility and bugfixes for Go modules including kubevirt.io/API, kubevirt.io/client-go, and github.com/k8snetworkplumbingwg/network-attachment-definition-client, as well as GitHub Actions workflows. [60d1a24b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/60d1a24bc49bb980e4e133c305b4d3e9ae818434)

## [17.1.6] - 2026-07-31

### Changed

- Simplified test logic by extracting assertion helpers for error and shutdown checks, resulting in improved test readability and reduced maintenance effort due to the extracted helper functions. [c45de95a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/c45de95a41bf322ac7ec431cf8efdeb46c6ae8dc)

## [17.1.5] - 2026-07-31

### Changed

- Stabilized code coverage for cmd/kar and internal runner packages by introducing additional tests that now pass, resulting in a 5.2% increase in overall coverage from 84.9% to 90.1%. [e6fd7bf4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e6fd7bf450d5c856d286bf8f0d0b96f18f1c3605)

## [17.1.4] - 2026-07-31

### Changed

- Simplified the update logic for golangci-lint and gremlins versions by consolidating it into a single generic helper function, making no changes to observable behavior. [6e039c41](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6e039c41a7371c81dd1a7b38169b4c98b9570082)

## [17.1.3] - 2026-07-24

### Changed

- Simplified VCS build setting keys and runMainApp function signature to reduce cognitive overhead and clarify the application's contract without introducing any functional changes. [34c1e7c5](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/34c1e7c556556f559954d572ee1a1e0cf93da9c7)

## [17.1.2] - 2026-07-24

### Changed

- Updated GitHub Actions and Go dependencies across all workflows to ensure compatibility with the latest features and security patches, including upgrading actions/setup-go, docker/login-action, and several Go module dependencies, as well as bumping the codespell pre-commit hook for improved typo detection. [28d84a26](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/28d84a261199fb70a10668695712ce97f79efe5a)

## [17.1.1] - 2026-07-24

### Fixed

- Resolved issues with artifact upload reliability in CI runs by correcting exported logs directory ownership to prevent permission errors and improving debuggability of future log collection problems. [a1fbe002](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a1fbe0024a8b7416d5642a7e877648860fc182ba)

## [17.1.0] - 2026-07-24

### Added

- Increased code coverage for the cmd/kar/main.go file from 0% to 60.6%, enabling more reliable testing and diagnostics of critical functionality. [3169678e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3169678effbd5b48e4696a7ddafa92302ef569e7)

## [17.0.0] - 2026-07-24

### Removed

- Simplified the Printf method in LoggerImpl to delegate directly to Infof, eliminating redundant logic without altering the API contract or observable behavior for users and maintainers. [e61ae8f8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e61ae8f8d50d3fb1b2b217d770e390a1e3b408fc)

## [16.2.1] - 2026-07-17

### Changed

- Enabled Copilot CLI to perform comprehensive verification tasks by relaxing default restrictions on tool and path usage while maintaining security for critical operations through explicit denial of dangerous shell commands. [667cbe2c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/667cbe2c43d278b508f2eafb33e3e5df2039fc5d)

## [16.2.0] - 2026-07-15

### Added

- Enabled dual-mode linter failure reporting, providing human-readable summaries and structured machine-report payloads that improve automation, integration, and communication for maintainers. [e62e5a8d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e62e5a8dced88c923c54b8c6bdffbe8ba0e10040)

## [16.1.5] - 2026-07-11

### Fixed

- Resolved permission errors when accessing Kubernetes cluster logs in the CI environment by running the kind export logs command with elevated permissions. [3a620304](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3a620304c02f10c794c07a19b872abe8d7093ba0)

## [16.1.4] - 2026-07-11

### Changed

- Simplified logic for resolving the latest semantic version tag from remote repositories by extracting it into a reusable function that replaces duplicated inlined code across three update functions with no functional changes expected as a result. [df0a1397](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/df0a13974305307575f5659b7e25e664c3b7ec12)

## [16.1.3] - 2026-07-11

### Changed

- Enabled proper coverage reporting by granting write permission to contents for analysis, allowing the gwatts/go-coverage-action to annotate commits during Git notes usage without affecting other workflow steps. [43682a91](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/43682a9162e583d2635cff01804b2156196f512a)

## [16.1.2] - 2026-07-11

### Changed

- Upgraded several Go module dependencies to their latest versions to ensure compatibility and security, including github.com/onsi/ginkgo/v2, github.com/onsi/gomega, and google.golang.org/gRPC. [91b83fad](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/91b83fad0cba07cee037b275e41b245f311807f7)

## [16.1.1] - 2026-07-11

### Changed

- Updated GitHub Actions to their latest patch versions for improved security and future-proofing, reducing maintenance overhead and improving workflow reliability. [27df61ea](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/27df61ea20eff9faddab7c92f329b8e89bab222b)

## [16.1.0] - 2026-07-11

### Added

- Automatically resolves and updates golangci-lint, gremlins, and rtk to their latest available versions in GitHub Actions workflows, reducing manual intervention and minimizing the risk of using outdated tools. [893b6ed1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/893b6ed1b5c69b74545ea52cc50c4e53d2a4b0e6)

## [16.0.4] - 2026-07-11

### Changed

- Hardened copilot-cli configuration to restrict tool usage and access to sensitive resources while maintaining necessary workflow functionality. [d76f5ab9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d76f5ab9fc31059aa71f675b7b86f146fd1aaff4)

## [16.0.3] - 2026-07-07

### Changed

- Optimized the CI workflow to conditionally run Go-related jobs only when Go files have changed resulting in reduced resource usage and build times while maintaining publish job functionality. [dc1af83b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/dc1af83b5ab8fac1f73b64e6c7d88bccdb49c8fe)

## [16.0.2] - 2026-07-07

### Changed

- Standardized GitHub Actions workflows for clarity and maintainability by renaming them to accurately reflect their purpose and triggers, clarifying event triggers and path filters, pinning tool versions, and improving job and step names for workflow logs readability. [40bc7fa6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/40bc7fa60bd108109c2c8261b0949247eaf8a31b)

## [16.0.1] - 2026-07-02

### Changed

- Optimized test coverage by introducing additional deterministic tests that cover previously uncovered code paths resulting in an overall increase from 90.6% to 94.7%. [98dc50b8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/98dc50b837319eeb179f641b73ea38bdfea93727)

## [16.0.0] - 2026-07-02

### Removed

- Simplified demo scripts by eliminating ~20 lines of duplicated shell logic and dead code through the consolidation of shared status-reporting functionality into a single `get_status()` function. [81cb5b24](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/81cb5b2487adb8b645273a5bfbd8cc65f57d0379)

## [15.2.2] - 2026-07-02

### Fixed

- Resolves issues where creating a Virtual Machine Instance that already exists results in unnecessary failures by returning the existing VMI instead of halting execution and logging the "already exists" error separately from other failure cases. [bdf23444](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bdf234447b8738eb55eabb738281dfe2a9b41098)

## [15.2.1] - 2026-07-01

### Changed

- Simplified the creation of DataVolumes by extracting conditional creation into a new helper method, createOptionalDataVolume, which is now used by the createResources function to simplify its logic without introducing any functional changes or breaking API contracts. [28078917](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/28078917388d6a2defac219d242e073857f11445)

## [15.2.0] - 2026-07-01

### Added

- Enabled accurate spell checking in Kubernetes and testing contexts by introducing commonly used terms such as kube, testvm, kubectl, and Kubernetes into the wordlist without affecting application logic. [1a1fb5d4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1a1fb5d4a2a5380c9b667d6393a29f8e12d6816b)

## [15.1.0] - 2026-07-01

### Added

- Enabled centralized template management by allowing users to specify the namespace from which to fetch the VM template independently of the runner's namespace through a new flag and corresponding environment variable, reducing template duplication and simplifying lifecycle management while maintaining backward compatibility. [531fb694](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/531fb6943bae02604b7d0207e5187d142e9e8049)

## [15.0.0] - 2026-06-25

### Removed

- Streamlined project documentation by eliminating redundant generic code review instructions template that is now maintained elsewhere. [860862be](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/860862be222c55bce2e14ffdf6bd2f044e96d5f6)

## [14.0.0] - 2026-06-25

### Removed

- Eliminated outdated documentation about the review-and-refactor skill to keep the repository up to date and reduce confusion about current practices. [c2817b39](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/c2817b39cdf53dbcb84beb37881167a99afad06e)

## [13.0.0] - 2026-06-25

### Removed

- Eliminated redundant documentation to prevent contributor confusion by removing outdated GNU Make Makefile authoring instructions from the instructions directory. [43f34dd2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/43f34dd23d74f03cc6d4cf6161714275d424dd20)

## [12.0.0] - 2026-06-25

### Removed

- Simplified the workflow by eliminating redundant configuration through consolidation of responsibilities under a unified engineering guidance agent, resulting in no impact on existing automation or engineering processes. [d5aad069](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d5aad069bd60629594ba8b978042ac4bc8957005)

## [11.1.0] - 2026-06-25

### Added

- Enabled global Copilot token management and usage tracking by integrating the rtk token saver into both fixer and verifier workflows, improving observability of token usage and supporting future automation or reporting needs with no impact on existing job logic. [b7eeca91](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b7eeca912fccbebcfe10df8715a4057e05413460)

## [11.0.1] - 2026-06-23

### Changed

- Optimized license headers in Go source files to prevent jscpd from flagging them as duplicate code thereby reducing noise in duplication reports without affecting application logic or tests and with minimal security impact. [1de8a0f4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1de8a0f49d86acb70717027ed4f2aaacd04a42b4)

## [11.0.0] - 2026-06-23

### Removed

- Upgraded GitHub Actions workflows to utilize the latest Copilot CLI features and security updates by switching from v2.0 to v3.2 without introducing any breaking behavior or API changes. [32bd4257](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/32bd4257fc536557fbc962d9af3f87de4041aea4)

## [10.0.2] - 2026-06-23

### Changed

- Enhanced GitHub Actions to reliably update action hashes for SemVer tags in workflow files, addressing edge cases and improving overall automation robustness. [8be2c280](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8be2c280b656ebd0f36a96cb1f863e156fa7ca94)

## [10.0.1] - 2026-06-23

### Changed

- Normalized GitHub Actions workflows now consistently reference the latest tagged versions of actions, reducing ambiguity and simplifying maintenance. [bafb6e71](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bafb6e71131fdcf835c29df97c26340aabe25463)

## [10.0.0] - 2026-06-19

### Removed

- Eliminated unnecessary whitespace in the data volume test to improve code readability and consistency without affecting API or CLI contracts, config schema, security, or migration requirements. [7a59bc19](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7a59bc19b457f0f67615f7a38a3aa6de9a73dce0)

## [9.7.4] - 2026-06-19

### Changed

- Upgraded several Go module dependencies to their latest versions, ensuring continued support and reducing the risk of vulnerabilities from outdated packages, without modifying any application code. [a6d30861](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a6d30861343aca0ca05da146dcfa0433d15c0796)

## [9.7.3] - 2026-06-19

### Changed

- Simplified telemetry and runner tests to improve maintainability by introducing reusable helper functions and closures that preserve existing test coverage and behavior without altering it. [924b7377](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/924b737736920b7091abd93dff24f1a2d806cca6)

## [9.7.2] - 2026-06-19

### Changed

- Updated the ai-prepare-commit-msg hook to leverage upstream improvements and bugfixes for commit message generation, requiring no configuration changes but monitoring future breaking changes in hook behavior. [3f0464d1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3f0464d1e50712a766a6445de1a8df5331fd007a)

## [9.7.1] - 2026-06-18

### Changed

- Simplified the runner code to reduce technical debt without introducing any observable changes in behavior, API contract, security, or config schema. [2e96c41f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2e96c41f91268fa464c44a39fcd3e97f27e11131)

## [9.7.0] - 2026-06-18

### Added

- Enabled comprehensive testing of previously untested code paths resulting in improved overall statement coverage from 70.5% to 92.2%. [68d46dd9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/68d46dd9eb8525ee3028681cfc4265dda7904207)

## [9.6.2] - 2026-06-18

### Changed

- Stabilized documentation clarity by updating QA subagent code blocks to use "text" language and correcting the heading from "antipatterns" to "Antipatterns". [5f182399](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5f18239916873c8a42c56c23eacd5c541ab721e8)

## [9.6.1] - 2026-06-18

### Changed

- Enhanced documentation clarity by introducing embedded diagrams that explicitly illustrate architecture, quickstart, telemetry pipeline, and timeout behavior for improved onboarding and troubleshooting. [bda81f17](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bda81f17671365389e2cc58f6c9f6a4d7d4527f5)

## [9.6.0] - 2026-06-18

### Added

- Automated code coverage improvement is now enabled through a dedicated QA subagent that analyzes gaps in unit tests, generates targeted tests, and submits focused pull requests to improve coverage, formalizing and automating test planning, bug reporting, and quality standards while reducing manual oversight and regression risk. [b6922b18](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b6922b18bad6fed9eebd840db5c79c55ee1480e5)

## [9.5.0] - 2026-06-17

### Added

- Standardized terminology across the codebase by introducing Ctrl, jitconfig, ParseDuration, and SIGTERM to linter-enforced wordlist usage. [5156f568](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5156f56833fd52f5c898f18b756e8590e561d286)

## [9.4.3] - 2026-06-17

### Fixed

- Stabilized compatibility for all environments building the application by updating the Dockerfile to utilize the latest version of Git, 2.54.0-r0, which includes recent bugfixes and security improvements over the previously used outdated version 2.52.0-r0. [d6535c99](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d6535c99de063288b7b2a656a90ad47e47507649)

## [9.4.2] - 2026-06-17

### Changed

- Organized documentation into distinct sections for tutorials, how-to guides, reference, and explanations to improve discoverability and align with established best practices. [f95c9db9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f95c9db9c5798e098bafd2cd64324ff2461d5de2)

## [9.4.1] - 2026-06-17

### Fixed

- Stabilized the build environment by updating the base image to alpine 3.24 which includes the latest security patches and bugfixes without introducing any breaking changes to the build process or application functionality. [480ebb71](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/480ebb7159f791b4ab087ac660f52ba1d60fea6a)

## [9.4.0] - 2026-06-17

### Added

- Enabled Gremlins mutation testing in the build workflow to improve code resilience and identify potential weaknesses in the test suite, which runs after unit tests only when Go files have changed and pushes the Docker image only after all tests pass. [fcff167a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/fcff167a0b8c40263b5198d3b260e441af961630)

## [9.3.0] - 2026-06-17

### Added

- Enabled automatic commit message suggestions based on changes in the staging area to improve consistency and quality across the project. [a4c56584](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a4c565847b26098f44715876a74b083d684900d1)

## [9.2.15] - 2026-06-12

### Changed

- Simplified the codebase to reduce technical debt by eliminating unnecessary checks and consolidating defaults, ensuring backwards compatibility without introducing any breaking behavior or API changes. [817b68c5](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/817b68c5ab5096417d1f47d51be93a7505ce76e0)

## [9.2.14] - 2026-06-05

### Changed

- Optimized version updates in continuous integration workflows to preserve partial changes and improve reliability by resolving dependency mismatches and refining failure reporting without breaking existing APIs. [1ebce834](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1ebce8341782aee6a6a15d84bddeb173163b73cb)

## [9.2.13] - 2026-06-05

### Changed

- Modernized project dependencies to align with updated Go versions and latest OpenTelemetry, Kubernetes, and gRPC releases without introducing breaking behavior or requiring migration efforts from API clients, Kubernetes APIs, or other external libraries. [8def3a3f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8def3a3f1a3d69ab45213135216d55093a46302d)

## [9.2.12] - 2026-06-04

### Changed

- Optimized GitHub Actions workflows by pinning actions/checkout to v6.0.3 ensuring consistent behavior across all reliant workflows without introducing any breaking changes or modifying the API contract. [5eca7d3c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5eca7d3cc3197be3ad58f5a4e8613b55bfbc4c82)

## [9.2.11] - 2026-06-04

### Changed

- Simplified error handling and reduced technical debt by consolidating variable declarations and removing unnecessary imports without introducing any breaking behavior or API changes. [f337f2f8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f337f2f89629dbf521d318111c4bdcbae19b4d54)

## [9.2.10] - 2026-06-03

### Changed

- Updated dependencies to the latest available versions, including actions/ai-inference 2.1.1, docker/login-action 4.2.0, docker/metadata-action 6.1.0, and markdownlint v0.16.0, ensuring compatibility with other changes in the project. [1fec048d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1fec048d5f576bde09581e19291863e54d4d51e9)

## [9.2.9] - 2026-06-01

### Changed

- Simplified and optimized telemetry shutdown logic in cmd/kar/main.go to eliminate redundant checks and improve performance without introducing any breaking changes. [7b823ed2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7b823ed204c424b3f38a750808bf8a809f244175)

## [9.2.8] - 2026-05-28

### Fixed

- Enabled GoCI configurations to utilize the newer and more secure gomodguard_v2 linter instead of its deprecated predecessor, requiring users who rely on this tool for code analysis to update their configurations accordingly. [bd8401ff](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bd8401ff8a66037232b9dcf7f2cb7477d252772a)

## [9.2.7] - 2026-05-28

### Changed

- Stabilized the VMI wait/status flow in the runner to improve robustness and protect API server integrity by reconnecting the VMI watch on stream closure with context-aware backoff before reconnection and tracking a `readyReported` flag to emit the 'Running and Ready' log only once per phase transition. [dd550f89](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/dd550f89292c510ccbcdd29ec659c8475ec26d4c)

## [9.2.6] - 2026-05-28

### Fixed

- Improved job resilience by automatically reconnecting VMI watches on stream closure during long-running jobs without requiring user configuration changes. [e0024a27](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e0024a27c97d139f119b0b0d3266c2cef6de33d3)

## [9.2.5] - 2026-05-22

### Changed

- Enabled package-level sentinel errors in internal/runner.go to prevent err113 issues and moved error declarations outside if statements in cmd/kar/app/root.go to satisfy noinlineerr requirements. [7beb720d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7beb720dc91d41e542832da607717312cd1db04f)

## [9.2.4] - 2026-05-22

### Changed

- Updated versions files to latest available versions, enabling seamless build and linting processes for Docker images and GitHub Actions workflows without requiring any migration steps or breaking behavior adjustments. [a4432b61](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a4432b61b25397ae4c2d11307bb22a003014f948)

## [9.2.3] - 2026-05-15

### Changed

- The logging functionality has been optimized to improve maintainability and reduce complexity in the codebase. [7f56e593](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7f56e593c6a6d6ac21e63ca197ac8303aa847fdd)

## [9.2.2] - 2026-05-15

### Changed

- Updated versions of the `actions/ai-inference` action to 2.1.0 in three GitHub workflows, potentially impacting build and linter analysis results due to API contract changes that may affect accuracy or output of these analyses. [f01c4513](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f01c451346c9306198837a01d96ea1839457ab87)

## [9.2.1] - 2026-05-14

### Changed

- Optimized the default wait timeout for the KAR runner to 1 hour from its previous value of 10 minutes, affecting users who configure timeouts via environment variables such as KAR_WAIT_TIMEOUT. [1aaf5577](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1aaf5577dda0daa1181a10c4ff0a407818352fa3)

## [9.2.0] - 2026-05-14

### Added

- Enabled configurable KAR_WAIT_TIMEOUT and VMI Running+Ready provisioning milestone logging, allowing users to customize timeouts and interpret provisioning milestones more effectively in their workflows. [6a853fc3](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6a853fc3c99c9a644998aa3e3c73576e5246df32)

## [9.1.6] - 2026-05-08

### Changed

- Simplified logging and build info creation to improve maintainability and adherence to best practices without introducing breaking changes or security vulnerabilities. [3c1ca8ce](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3c1ca8ce936ce86798f9fa171d71698fd7fcd89e)

## [9.1.5] - 2026-05-01

### Fixed

- Resolved an error in actions/setup-copilot by downgrading austenstone/copilot-cli from v3.0 to v2.0 and disabling the Copilot CLI auto-update via a pinned_actions list in update_versions.sh. [9ac1d458](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/9ac1d45853edb80722230f39a18f49aa6b944379)

## [9.1.4] - 2026-04-20

### Changed

- Updated GitHub Actions workflows to utilize newer versions of actions, specifically GoTestTools/gotestfmt-action was updated from 2.2.0 to 2.3.0 and actions/upload-artifact was updated from 7.0.0 to 7.0.1 without introducing breaking behavior or requiring migration steps. [6a84162a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6a84162a986a6a31b4639086575e495c0a5380f7)

## [9.1.3] - 2026-04-11

### Fixed

- Resolved documentation duplicity by removing redundant content from the Diátaxis Documentation Expert skill. [fea2e201](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/fea2e201652786d5a11b89717a3109e292d4a34c)

## [9.1.2] - 2026-04-07

### Fixed

- Resolved NATURAL_LANGUAGE and GITLEAKS super-linter failures, addressed Go linting issues, applied code review feedback, and added default branch handling for unrecognized VMI phases without introducing any breaking behavior or migration requirements. [9b026b66](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/9b026b66a888705b442148cc652ab73b8cc49801)

## [9.1.1] - 2026-04-05

### Fixed

- Stabilized the WaitForVirtualMachineInstance method by introducing a 5-minute timeout to prevent indefinite blocking and improve test reliability, clarity, and security. [d5babba4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d5babba40cd8ef22322232cc69378726d09dab93)

## [9.1.0] - 2026-04-05

### Added

- Enabled custom GitHub instructions to ensure code quality and documentation standards in the repository by emphasizing testing, validation, documentation updates, and clean code practices for contributors, reviewers, and maintainers. [0e9d3162](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/0e9d316275320e0da4058a45077554296706d7f6)

## [9.0.5] - 2026-04-03

### Changed

- Modernized versions files and tool dependencies to resolve linting issues and ensure compatibility with the latest CI tools. [bbf36644](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bbf366448cf0046eec577ec096da65399a2ff96e)

## [9.0.4] - 2026-03-30

### Fixed

- Resolved the pinned Git version in the Dockerfile to 2.52.0-r0 from 2.49.1-r0, requiring users who were relying on the previous version to rebuild their containers. [5e82fb61](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5e82fb61c9e855c2749736a7f7adc9975556a102)

## [9.0.3] - 2026-03-27

### Changed

- Updated dependencies to align with latest available releases impacting GitHub Pages settings and deployment actions potentially requiring migration steps for users relying on these dependencies in their workflows. [6a06c04d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6a06c04d758eec28da63f99a870028de0ebb9f4b)

## [9.0.2] - 2026-03-20

### Fixed

- Resolved demo execution reliability by pre-building kar binary and waiting for virt-handler readiness before running the demo. [34f22cd4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/34f22cd4fc530d5e6f6a16e8f36acbba0cd8b939)

## [9.0.1] - 2026-03-20

### Fixed

- Resolved scheduled version update workflow failures by robustly extracting major.minor Go versions and updating the Dockerfile base image to match. [d04e41af](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d04e41af9b9a11662843cb19151e0953daaac718)

## [9.0.0] - 2026-03-13

### Removed

- Prevented fatal crashes in DeleteResources by making AppContext initialization optional and introducing the HasAppContext function to check for initialization before proceeding with deletion. [cdaaa3dc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/cdaaa3dc544bab4371eee86157aa8e5a51a4fbb5)

## [8.0.2] - 2026-03-13

### Changed

- Updated GitHub Actions workflows to include detailed Docker image build process and tool dependency information for enhanced debugging and troubleshooting capabilities across multiple development tasks. [d194d7ac](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d194d7aceeefca883a0f8eb6178e3612311ed77a)

## [8.0.1] - 2026-03-13

### Fixed

- The Docker build process was stabilized to allow users submitting changes directly to the master branch without requiring pull requests, maintaining the existing API and CLI contract with no migration steps required. [26910142](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2691014241233ba9f001385849b6a68f9bf62100)

## [8.0.0] - 2026-03-13

### Removed

- Streamlined access to relevant information by eliminating unnecessary index.md content that previously cluttered the testbed's documentation sections on architecture and design principles. [1966ee0f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1966ee0f1eed4782a8a5aff6d9771b29f966d189)

## [7.20.7] - 2026-03-13

### Fixed

- resolved golangci-lint exitAfterDefer and paralleltest failures in appcontext ensuring consistent behavior during testing without introducing any breaking changes. [a28a6002](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a28a6002f71957b9f2ff10f62a66697cf7614f2c)

## [7.20.6] - 2026-03-13

### Fixed

- Stabilized user understanding of KubeVirt Actions Runner's functionality by adding a diagram image to the README.md file that addresses limitations of GitHub-hosted runners in standard workflows. [d9424fda](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d9424fdacb4f11459f3a4fd7bc51c29221d1fbe2)

## [7.20.5] - 2026-03-13

### Fixed

- resolved GitHub Pages deployment issues by stripping Runme-specific code fence attributes from documentation before Jekyll rendering. [b93ae3f6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b93ae3f6ed9dfd33bdb9fa3da0daabe5eeb0693b)

## [7.20.4] - 2026-03-13

### Fixed

- The Docker build process was resolved by correcting an issue that prevented it from importing the project's Git repository correctly during the build process. [3f62f463](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3f62f463bb53e081e90330245803ccf69e0c8cca)

## [7.20.3] - 2026-03-13

### Fixed

- Resolved super-linter failures in NATURAL_LANGUAGE, Markdown, and JSCPD checks by updating configuration settings and modifying README.md content to conform to textlint and markdownlint rules. [5e56b630](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5e56b63013afb82a1518cad7dcdd293f62f81095)

## [7.20.2] - 2026-03-13

### Changed

- Simplified documentation in the testbed guide to eliminate duplicate content and redundant flowcharts ensuring consistency for users. [732ec425](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/732ec425c0523f7c0dd4fc3ea16168acb10ec6ec)

## [7.20.1] - 2026-03-13

### Fixed

- Optimized GitHub Actions workflow triggers to improve job filtering and reduce technical debt by introducing more selective remediation label matching and eliminating pull request review triggering in the build process. [d4b3c217](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d4b3c21772631da7f28f35d5cfc0724202c1b800)

## [7.20.0] - 2026-03-13

### Added

- Enabled GitHub workflows documentation for users to understand and utilize existing CI validation, maintenance, and automation workflows without introducing breaking behavior or API changes. [ecf4b03a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ecf4b03a5eeae8507bec3220a0bc034a1ee23c5b)

## [7.19.0] - 2026-03-13

### Added

- Enhanced documentation for KubeVirt Actions Runner's README.md file now clearly outlines key features including ephemeral VM creation, increased isolation, customizable system-level configuration support, and seamless integration with Kubernetes-native tooling. [b496c5a4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b496c5a461b9b0c1db705c21a33e064c5e2f6b5e)

## [7.18.0] - 2026-03-13

### Added

- Improved code readability and maintainability through enhanced documentation and metadata variables for build-time information. [bfd76a05](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bfd76a053fa91589f4fcb601a6b0a6562aa94ff2)

## [7.17.0] - 2026-03-13

### Added

- Enabled easier maintenance and updates of the architecture diagram by relocating it to its own separate document in the `docs/explanations` directory without introducing any breaking behavior, API changes, or security impact. [821604c8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/821604c85d05c125b1ffcece32be1e06bc0d8626)

## [7.16.0] - 2026-03-13

### Added

- Enabled the Cayman theme for project documentation, resulting in an updated visual appearance online without introducing any breaking behavior or API changes. [56f2592a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/56f2592a7a714a1c4dad751bbfd5d5a123933736)

## [7.15.0] - 2026-03-13

### Added

- Improved documentation navigation and organization have been introduced through an enhanced GitHub Pages setup that includes restructured project sections, updated tutorials, and a corrected telemetry guide for Jaeger UI traces verification. [84d2e4f0](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/84d2e4f08fbb2e08f480459d3b028adf89a93739)

## [7.14.0] - 2026-03-13

### Added

- The quickstart documentation is now more accessible and easier to maintain due to its relocation from the readme into a separate file. [6c472840](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6c472840477f076c6bfdaded92049d142e55cb21)

## [7.13.2] - 2026-02-28

### Fixed

- Resolved the Markdown linting issues in the docs/how-to-enable-telemetry.md file to ensure consistency across example sections without introducing any breaking behavior or requiring migration steps. [a10c04b7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a10c04b7ff8d10567019a939ed9b3b3da328a421)

## [7.13.1] - 2026-02-28

### Fixed

- Corrected linting issues to ensure adherence to critical rules and improved code readability by removing duplicated package declarations, adding empty lines between logical code groups, and clarifying documentation requirements in readme files. [974b827d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/974b827d41932beae430585a8d2f228b19a63eda)

## [7.13.0] - 2026-02-28

### Added

- Enabled clear guidelines for developers contributing to the Kubevirt Actions Runner project by introducing a contributing file outlining setup instructions and code standards. [3258012c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3258012ca9022f7783fcabc15a12968fdb85de7b)

## [7.12.0] - 2026-02-28

### Added

- Enabled better observability and debugging capabilities for users through distributed tracing by introducing a new guide on how to enable telemetry for the `kubevirt-actions-runner` application, including configuration options and exporter types, in favor of previous telemetry documentation. [5f9be7fa](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5f9be7fa0120ec66b7420de0928eaa826977f111)

## [7.11.2] - 2026-02-27

### Fixed

- The documentation for the Diátaxis Documentation Expert skill was updated to correct Markdown and natural language issues, improving clarity and consistency in project tone, style, and terminology. [02c8b8d5](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/02c8b8d5a2eb56cadcf4805725dd6fa5722a808b)

## [7.11.1] - 2026-02-27

### Fixed

- The wordlist used in GitHub Actions has been updated to include several new terms without introducing any breaking behavior or migration requirements. [18f94788](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/18f9478899af0ba2e4a2b13c96fd97642d3b1e8b)

## [7.11.0] - 2026-02-27

### Added

- Enabled setup of testbed environments for kubevirt-actions-runner through comprehensive documentation that outlines prerequisites, architecture, deployment flow, automated install script, and demo usage. [f7a19689](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f7a1968957b8dd59dbd92673e249c43200684aa5)

## [7.10.0] - 2026-02-27

### Added

- Enabled expert technical writers to create high-quality software documentation by introducing guidelines based on the Diátaxis Framework's principles and structure that outline four document types, workflow, and contextual awareness for clarity, accuracy, user-centricity, and consistency. [1b8c9e14](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1b8c9e140fd239e15da42d220df140d02335d489)

## [7.9.3] - 2026-02-14

### Fixed

- The GitHub Actions workflow for on-demand CI jobs now triggers for all relevant pull requests, including opened, synchronized, and reopened ones, in addition to submitted reviews, potentially increasing the number of tests run without introducing any breaking behavior or migration requirements. [4797133c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4797133cdfae46eb148249f5b1bbe5585bdb4404)

## [7.9.2] - 2026-02-14

### Fixed

- Resolved inconsistencies in janitor agent documentation by correcting overengineering terminology to maintain proper consistency and accuracy without introducing any breaking behavior or migration requirements. [a1de3c32](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a1de3c325daa17d12a7c6579413925d7cdd57ea9)

## [7.9.1] - 2026-02-14

### Fixed

- Resolved GitHub workflows to more effectively guide users in submitting Pull Requests that reduce technical debt by clarifying the steps for creating a new branch and submitting a request with improvements. [b1689a76](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b1689a7624afd7f4a803923bfc68c2441cb39738)

## [7.9.0] - 2026-02-14

### Added

- Enabled automated cleanup of unused code elements and tech debt remediation through integration into GitHub workflows for Docker build fixes and Go linting issues. [40df54db](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/40df54db7f6d7b0d3a1c1b6ec5f7472910c554f2)

## [7.8.0] - 2026-02-14

### Added

- Introduced automated analysis of Docker build failures and direct fixes through a new workflow, reducing manual effort and improving developer productivity. [13ad46df](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/13ad46df4a09af1ce2083dc8fc97d1d67c2b2f8f)

## [7.7.1] - 2026-02-12

### Fixed

- Resolved all 14 golangci-lint wsl_v5 whitespace violations in the codebase, ensuring adherence to linter rules and maintaining code readability without introducing any breaking behavior or API/CLI contract changes. [6467558c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6467558c0b962aaf8a056fba99db516ddefe8524)

## [7.7.0] - 2026-02-12

### Added

- Enabled automatic linting and issue resolution through integration with golangci-lint agent in GitHub Actions workflows without introducing breaking behavior or API changes but requiring migration of existing workflows to leverage the new functionality. [4d9d33dd](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4d9d33dd6d7df1b194fed5a457dadc311cb2cb9f)

## [7.6.1] - 2026-02-02

### Fixed

- The Docker build process now includes Git information in the image, enabling developers to leverage this data for debugging and development purposes without requiring any migration steps. [82d918e2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/82d918e241ce995c5bae41b0153c603d9af8cd09)

## [7.6.0] - 2026-02-02

### Added

- Enabled improved testing coverage for projects that use Docker images by introducing a new smoke test for image building and testing in addition to Go code. [150f7617](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/150f7617d7955198635303e82ad3ca3616e6f0fa)

## [7.5.6] - 2026-02-02

### Fixed

- Resolved false positives in code spell checking by excluding go.* files from the `.codespellrc` configuration and disabling code spell checking during linting in GitHub Actions workflows and Makefiles. [35cd82aa](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/35cd82aa4d019ac55c42f9319356be21b1562087)

## [7.5.5] - 2026-02-02

### Changed

- Updated GitHub Actions versions in workflows to 3.7.0 for docker/login-action, 3.2.0 for actions/attest-build-provenance, and 8.4.0 for super-linter/super-linter. [14fb8e04](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/14fb8e046a7367ea3a9289b1e3581437b42c12ad)

## [7.5.4] - 2026-02-02

### Fixed

- Resolved the issue with go mod update by downgrading grpc-gateway from v2.27.5 to v2.27.3 requiring a potential migration step for users relying on this version. [bae3ef04](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bae3ef04f4e7fd05577d59bf4f48ee5e50520be0)

## [7.5.3] - 2026-01-23

### Fixed

- GitHub Actions permissions were updated to remove write access for artifact-metadata, which may break workflows relying on this permission, and users should review their configurations to ensure compatibility. [8123ca44](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8123ca44be22ce96f940cf7bf0676b057f1af681)

## [7.5.2] - 2026-01-23

### Fixed

- Resolved an issue in the janitor prompt's natural language linting by correcting the spelling of "overengineering" to its original form "overengineering", maintaining consistency with the rest of the text without introducing any breaking behavior. [d62d911b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d62d911bbc02efd4c1c0ff381744b095e6cdd197)

## [7.5.1] - 2026-01-23

### Fixed

- Artifact metadata permissions have been updated to allow write access, enabling users to modify the metadata associated with artifacts in workflows without requiring any migration steps for existing configurations. [b5de370b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b5de370b4ebd76f86518f7136de15246eee70ae7)

## [7.5.0] - 2026-01-23

### Added

- Enabled cleanup of codebases through the introduction of the janitor prompt, which facilitates elimination of technical debt by deleting redundant code, simplifying complex structures, and managing dependencies, resulting in improved code quality and reduced maintenance burden. [4c5b3963](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4c5b396317f8da5003bfe9b3c6c7ce7d1421982f)

## [7.4.9] - 2026-01-23

### Fixed

- The application now accurately displays build information including the Git commit hash, build date, and Git tree modification status within its interface. [804b7118](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/804b7118393feb177f0cea0872205530bcb795ec)

## [7.4.8] - 2026-01-23

### Fixed

- The Docker image build process was stabilized by modifying the SHELL command to ensure compatibility with the build environment without introducing any breaking behavior and requiring migration steps. [3f1db7fc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3f1db7fc5b517b9ba2556f5f36ccf1cee82a2397)

## [7.4.7] - 2026-01-23

### Changed

- Upgraded the Ginkgo testing framework to v2.27.5 and the related Gomega library to v1.38.2. [8d238714](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8d2387141e7c7dd68406f0557390d0bf203d01f7)

## [7.4.6] - 2026-01-23

### Fixed

- Resolved issues related to the Docker build process by correcting the shell used from /bin/Bash to /bin/sh in affected Dockerfiles. [ea0f5947](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ea0f5947f5b909fd03226f6bd88d9bf8b0e1d216)

## [7.4.5] - 2026-01-23

### Changed

- Updated containerized-data-importer-api dependency to version v1.64.0 for compatibility and routine maintenance. [67437cf8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/67437cf8744775a6d0ff0e7be66136963073cae4)

## [7.4.4] - 2026-01-23

### Fixed

- Expanded build conditions for GitHub Actions now include Go module files, allowing users to more comprehensively test their Go projects and adjust expectations for code coverage results accordingly. [2eb75b26](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2eb75b26b676a7f3cf7fac5c690c62dd07bc7d9e)

## [7.4.3] - 2026-01-23

### Fixed

- Git credentials are now persisted by the GitHub Actions workflow for building the project, enabling users who rely on this feature to authenticate with their repositories without requiring adjustments to their workflows. [61518c3e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/61518c3e1556debc47e6c8535f771c9526ee3a46)

## [7.4.2] - 2026-01-23

### Changed

- Updated versions files to ensure maintainers and users have access to the latest dependency information without requiring any migration steps. [9a0096df](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/9a0096df2706062d98706e2d4e0c927e1c631b52)

## [7.4.1] - 2026-01-23

### Fixed

- Resolved super-linter issues in GitHub workflows for Go linter and linter by adding configuration to ignore certain sections of code and modifying the failed build issue action. [beba3810](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/beba381000485ef35884aeeffb6b9257004c9a05)

## [7.4.0] - 2026-01-23

### Added

- Improved issue descriptions are now generated by GitHub Actions for Go linter and Super-linter, providing developers with more comprehensive and informative issues to implement fixes confidently. [239c941a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/239c941a2c71cdbeca8b7bfe17b3892a47046649)

## [7.3.2] - 2026-01-23

### Fixed

- Stabilized linter compliance by addressing issues detected by golangci, ensuring maintainability and consistency without introducing breaking behavior or migration requirements. [25373bbf](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/25373bbf7735800a7a68364fb4320e25b745b68f)

## [7.3.1] - 2026-01-23

### Fixed

- The Docker build process now stabilizes by enabling dynamic construction of metadata in the compiled binary through the ldflags settings when VCS information is unavailable. [b3865e21](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b3865e21d185cd1e2484c016fd00e7a07e9c21f0)

## [7.3.0] - 2026-01-23

### Added

- Enabled automated analysis of Go code for linting issues through integration of the GitHub Actions workflow with golangci-lint tool runs and AI-driven output analysis. [bb37fca7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bb37fca76b130407392b0b921420fa693ef8d8df)

## [7.2.0] - 2026-01-22

### Added

- Enabled the efficient and flexible updating of Go versions in CI workflows through a single command that automates updates across all relevant files without requiring manual modifications or migration steps. [2d8d5f87](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2d8d5f874cc9ef4e9a250e968c2a9bf66b2d02f4)

## [7.1.3] - 2026-01-22

### Fixed

- Resolved the setup-go cache poisoning vulnerability by updating the minimum required Go version to 1.25 and requiring manual adjustments for users relying on cached versions of Go modules due to the removal of cached dependencies during builds. [94803a9f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/94803a9f080821866abdb64d4d43ca4605ad5dcb)

## [7.1.2] - 2026-01-22

### Changed

- Updated the AI-inference action version to 2.0.5 in the GitHub workflow, requiring manual migration steps for users running older versions of this action. [33fab97a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/33fab97ad91f65559fbfa4ddbd93a37930fce789)

## [7.1.1] - 2026-01-22

### Changed

- Upgraded dependencies to include OpenTelemetry packages for tracing and metrics, potentially requiring configuration adjustments due to changes in performance data collection capabilities. [f3e71fd8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f3e71fd8eae14dfda78cdf7b5009c5fe60228a92)

## [7.1.0] - 2026-01-22

### Added

- Enabled Go version 1.25 or later for code coverage execution in CI/CD pipelines, ensuring compatibility with newer releases and maintaining uninterrupted workflows. [197eab85](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/197eab85ef612db8840b8b748c13978f06582c52)

## [7.0.2] - 2026-01-21

### Fixed

- The GitHub Actions workflow for building the project now fetches sufficient history to enable accurate and comprehensive code coverage analysis. [59a4d8b0](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/59a4d8b062161a03a94284315bf9850851389b74)

## [7.0.1] - 2026-01-21

### Changed

- Upgraded GitHub Actions workflows to utilize newer versions of actions and dependencies, ensuring compatibility without introducing breaking changes or security vulnerabilities. [4e9558bf](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4e9558bf1ccc0c6e92ccd2f66ab2a505a0b35a28)

## [7.0.0] - 2026-01-21

### Removed

- Simplified test code by eliminating duplicated logic and introducing a new helper function to reset the singleton for testing purposes without any breaking behavior or API changes. [98211f0b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/98211f0b9df4e3f36409962bdf28c45dcfd5f8d2)

## [6.6.2] - 2026-01-21

### Fixed

- Improved code coverage by enabling additional linter checks and test cases for logging functionality without introducing any breaking behavior or API changes. [b768e6af](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b768e6afbec6c7f7f04094097238250d2dc01dda)

## [6.6.1] - 2026-01-21

### Fixed

- Resolved spelling issues in the GitHub repository wordlist to ensure accurate code completion and linting results for users relying on this list. [1602de91](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1602de91178de2914accbf19d934325d60e654a0)

## [6.6.0] - 2026-01-21

### Added

- Automatic detection of misspelled words in code is now enabled through the addition of codespell to the project's pre-commit hooks and Makefile configuration. [d86c6ea9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d86c6ea9d3b0bb3bf7b7217b244bdd1d1ec91db9)

## [6.5.0] - 2026-01-21

### Added

- Enabled distributed tracing capabilities through OpenTelemetry integration with configurable exporters for multiple backends and environment variables to enable or disable telemetry settings. [637b9ddb](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/637b9ddb95d5b8b87b95ad060876207a37eacff0)

## [6.4.0] - 2026-01-21

### Added

- Enabled more flexible and customizable logging configurations through environment variables by switching to the Uber's Zap logging library without introducing any breaking behavior or migration requirements. [4177eb2a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4177eb2a7e16f4080f7429e48e15b6de1922b3ac)

## [6.3.6] - 2026-01-07

### Fixed

- Resolved an issue related to commit versions by updating the version of the `git-diff-action` used in GitHub workflows to 2.8.1, which may require migration steps from users relying on this action. [7db0a59c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7db0a59cbaca423663deec67694c4414ec12f084)

## [6.3.5] - 2026-01-07

### Fixed

- The GitHub Action workflow for linting has been stabilized to ensure compatibility with newer versions of the action by switching to a different action that outputs raw diffs instead of processed ones. [20da4f25](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/20da4f2528dc0752503b13a14dd2781c6aa12e1d)

## [6.3.4] - 2026-01-07

### Fixed

- The GitHub Actions workflow for super-linter validation now fetches the full Git history, which is required by the tool and may introduce performance implications and potentially reveal sensitive information previously hidden in the repository's commit history. [4be8477a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4be8477af65bc1e099a8ac281c2bc1153d2901d4)

## [6.3.3] - 2026-01-07

### Changed

- Upgraded versions files to utilize super-linter version 8.3.2 for linter validation in the GitHub Actions workflow. [870fffd8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/870fffd82a733f07bd56f9b9ea46275f0357e134)

## [6.3.2] - 2026-01-07

### Fixed

- Stabilized integration tests to accurately reflect cluster configuration and environment variables changes by updating job names, artifact upload, Kubernetes events retrieval, enabling the `useEmulation` flag for kubevirt daemonset, and collecting storage information with `df -h`, `lsblk`, and `lsmod`. [f45d1af6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f45d1af65f1aa555b1f74379ea24998075d5426c)

## [6.3.1] - 2026-01-06

### Fixed

- The demo script now properly sources system path configuration files avoiding errors due to missing environment variables when running Alpine demos. [3e2837c9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3e2837c977d4542052463aca3c8493cd039b4e25)

## [6.3.0] - 2026-01-06

### Added

- Enabled users to run ActionLint on their own infrastructure by introducing a new self-hosted runner configuration file that allows for customizable and flexible management of the tool's setup without any breaking behavior or migration requirements. [a33efc6f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a33efc6f2744920cf852851707e5a2bd11f50d4c)

## [6.2.1] - 2026-01-06

### Fixed

- Resolved a zizimor linting issue by updating the GitHub Actions workflow for on-demand CI to no longer persist credentials in the checkout step, potentially requiring users to reconfigure their workflows if they relied on persisted credentials. [3179385a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3179385a0dede46482b6320438e67cb952bcbd21)

## [6.2.0] - 2026-01-06

### Added

- Enabled on-demand system testing for the project through the introduction of new workflows and scripts that facilitate automated quality assurance and functionality verification. [abe0378d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/abe0378d55cde63f0250f8b99856745991b6c803)

## [6.1.1] - 2025-12-19

### Changed

- Upgraded several Kubernetes-related dependencies to their latest versions, including k8s.io/API v0.22.0 and others, which may require migration steps in dependent code due to potential breaking changes. [2eb806cf](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2eb806cf6c7e99cfb300a5801449f05111423be3)

## [6.1.0] - 2025-12-19

### Added

- Enabled GitHub Actions for reviewdog/action-misspell and actions/attest-build-provenance to update their commit hashes as expected by workflows relying on these actions, with no migration steps required. [d08c752d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d08c752d2217b116933008eb7e8069efd6476f23)

## [6.0.3] - 2025-12-19

### Changed

- Upgraded the pre-commit configuration to leverage improvements for commit message formatting in the latest version of the ai-prepare-commit-msg repository. [de736cf5](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/de736cf5c2aea260ddf5b38801bc61d231f8bd60)

## [6.0.2] - 2025-12-19

### Changed

- Upgraded dependencies in versions files to ensure continued support and compatibility for GitHub workflows, affecting build, linter, and pre-commit configurations. [68b98274](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/68b982748be6282b55e13f42d203a4663e125cf6)

## [6.0.1] - 2025-12-19

### Changed

- Upgraded GitHub workflows to utilize more recent versions of dependencies, specifically actions/checkout was updated from 6.0.0 to 6.0.1, without introducing any breaking behavior or requiring migration efforts. [41ca18b4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/41ca18b410287b74d509020524e34ecdd8910c5d)

## [6.0.0] - 2025-11-29

### Removed

- Optimized build times by disabling pre-commit linting checks that are already covered by other validation processes. [eb595e97](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/eb595e97a221743ba3a81192149e42fa6dd45bc5)

## [5.7.1] - 2025-11-29

### Changed

- Enabled successful linter validation on Ubuntu platforms by installing OpenSSL libraries prior to running the linter. [c4ba49c8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/c4ba49c8868583f346d7663e191387dda439a308)

## [5.7.0] - 2025-11-28

### Added

- Enabled charset checking by default for users who rely on this feature to enforce character encoding standards, and updated the config schema to allow disabling it via the "Charset" property under the "Disable" section. [e0e658f1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e0e658f19bbab60a5e6d9cfab4682f79e05c9b20)

## [5.6.1] - 2025-11-28

### Changed

- Updated version dependencies across workflows and configuration files to ensure compatibility and consistency, including the `actions/checkout` package from 5.x to 6.x, `actions/setup-go` from 6.0.0 to 6.1.0, and several other packages such as markdownlint, yamlfmt, and technote-space repositories, with a new repository and hook added to the `.pre-commit-config.yaml` file. [d97c90c4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d97c90c4fc3d1de3a025adc9932203eddef9cdd5)

## [5.6.0] - 2025-11-27

### Added

- Enabled pre-commit and prepare-commit-msg hooks for AI package maintenance, enforcing code quality checks before committing changes without introducing any breaking behavior or API contract changes; users must update their commit hooks according to the new configuration. [82c01dae](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/82c01dae6c3279ab45011594bb9aa1de830e86c3)

## [5.5.1] - 2025-10-26

### Changed

- The GitHub Actions workflow for linter checks now utilizes the Ministral-3B model instead of the previous medium-sized model which may affect analysis results and accuracy. [8268c245](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8268c245cbe59384567baac60c65830a263c48d8)

## [5.5.0] - 2025-10-26

### Added

- Go linting is now disabled by default in the project's GitHub Actions workflow and Makefile, impacting users who relied on these checks for code quality. [cc1e99cc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/cc1e99cc60590d846fe162935db35e76dca0c59f)

## [5.4.8] - 2025-10-26

### Changed

- Automated analysis of linter failures in CI workflows has been enhanced with AI-driven diagnosis and resolution capabilities. [be2469be](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/be2469be09612f45982376d93e47aac6de811ce8)

## [5.4.7] - 2025-10-23

### Changed

- Resolved the issue of the reviewdog/action-misspell action not being properly installed and used in the GitHub workflow by correctly referencing it in the configuration files. [6789b5ec](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6789b5ecbf6f0ddd26113a91d71ff96f889dfe5e)

## [5.4.6] - 2025-10-23

### Changed

- Stabilized GitHub workflows by updating actions for linting checks and the Makefile to reflect these changes. [bccb011e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bccb011e1e731341ff8b1808bc4e7da9de1d04a0)

## [5.4.5] - 2025-10-16

### Changed

- Updated dependency versions in version files to ensure compatibility and stability of the linter workflow. [6636e9d9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6636e9d9a4235635814e4d95e82cd30a9ec7c681)

## [5.4.4] - 2025-10-16

### Changed

- Upgraded versions files for docker/login-action and super-linter/super-linter to 3.6.0 and 8.2.0 respectively potentially impacting build and linter jobs due to changes in these dependencies. [6589a248](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6589a2489048d73f300f511749194b364c517f73)

## [5.4.3] - 2025-09-25

### Changed

- Updated the Docker container to use the latest Git version by changing the installed package from git=2.49.0-r0 to git=2.49.1-r0 without affecting API or CLI contracts and requiring no migration steps. [d539a614](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d539a614051d24f106aee9336d8dd2dcf152908c)

## [5.4.2] - 2025-09-25

### Changed

- Enabled proper access for code updates and Dockerfile modifications in Docker CI jobs by granting write permissions for pull requests and contents. [4373a749](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4373a7491b0c683ab9b104167ef74f6f65b7c066)

## [5.4.1] - 2025-09-25

### Changed

- Upgraded versions of actions in GitHub workflows to setup-go 6.0.0 and attest-build-provenance 3.0.0, requiring users relying on these specific action versions to take migration steps. [08d737b2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/08d737b265f8be9fa98d6e19e811281e782726c7)

## [5.4.0] - 2025-09-25

### Added

- Enabled notifications for failed jobs in GitHub Actions workflows by introducing a notify step that respects access controls and only sends issues to the repository when necessary. [7c9fc228](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7c9fc2283ce0ab1a2c00d3931e24f5f220278b17)

## [5.3.1] - 2025-09-25

### Changed

- Enabled PR creation for users who create pull requests with fine-grained PATs by persisting credentials securely without introducing breaking behavior or requiring migration. [08968a3a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/08968a3a13b13a7cda85e6f7fc43202f9785cf86)

## [5.3.0] - 2025-09-21

### Added

- Enabled contributors to easily set up development environments through cloud-based GitHub Codespaces and local Dev Containers via Visual Studio Code without manual configuration. [783eee16](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/783eee16aa533d5efb93a5e5108051f6631c0f76)

## [5.2.0] - 2025-09-21

### Added

- Automated pre-commit checks have been enabled for developers, running hooks for trailing whitespace, YAML syntax, shell script linting, Markdown linting, and YAML formatting during the commit process with configuration stored in a .pre-commit-config.yaml file. [d2d4df16](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d2d4df16d9eb3a6c6c79997dd285464b593ecf31)

## [5.1.6] - 2025-08-27

### Changed

- Updated Go dependencies to their latest versions, including changes to the build process that may require migration steps for users who have customized their build configurations. [99e90559](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/99e90559e12681fbb26b5a09308931908f4ab8a8)

## [5.1.5] - 2025-08-24

### Changed

- Upgraded GitHub Actions workflows to resolve linting issues by upgrading actions/checkout from version 4.2.2 to 5.0.0 and adding necessary permissions for contents: write and read in various jobs. [01b3f61b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/01b3f61b6fb11b48bf1ed64872ca9ec0d4d6e9c6)

## [5.1.4] - 2025-08-24

### Changed

- Optimized CI update process to maintain accurate and up-to-date build dependencies by correctly handling GitHub Actions commit hashes in the workflow file. [db5bff4d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/db5bff4d5bf5082d87c06a638e4ae50d1418b80f)

## [5.1.3] - 2025-08-18

### Changed

- The CI update job now correctly retrieves and uses the Go version from the official site for the project's dependencies and GitHub Actions. [4617acff](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4617acff424704f404746fcfe5913dbe7d3419fc)

## [5.1.2] - 2025-08-18

### Changed

- Aligned Bash scripts to consistently use spaces for indentation and set size to 4 characters by updating the `.editorconfig` file and modifying the `shfmt` command in the Makefile's `fmt` target with the `-i 4` option. [7d533f3c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7d533f3c22e69a71d33e3d49d89ca0c83184a243)

## [5.1.1] - 2025-08-07

### Changed

- Modernized versions of dependencies in GitHub workflows to ensure compatibility with latest tooling and dependencies. [05ae6e94](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/05ae6e94ae3cbeeb3eb9a17dc5cfe8c1522b24ee)

## [5.1.0] - 2025-08-07

### Added

- Enabled automated linting across different branches and repositories by default setting the super-linter's branch to follow the current GitHub head ref or repository name without requiring manual configuration changes. [6a520d1f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6a520d1f0cb34df425c4d0349ca6a39c4a5aac84)

## [5.0.4] - 2025-07-18

### Changed

- Upgraded dependencies for GitHub Actions workflows to introduce new linter rules and spell checking capabilities without requiring any migration steps or breaking changes to existing workflows. [5ee77a6d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5ee77a6d975ba7e66b436468975066e1add81c1e)

## [5.0.3] - 2025-07-18

### Changed

- Upgraded the GitHub workflow for linter validation to utilize super-linter version 8.0.0 from 7.4.0, requiring manual review of updated configuration but preserving API and CLI contracts, config schema, and security integrity. [59536498](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/59536498a1c90b81bb7fbd8aad03ef17ab92280e)

## [5.0.2] - 2025-07-04

### Changed

- Upgraded the spell checking tool versions used in GitHub Actions workflows to 1.0.5 from 1.0.4 without introducing any breaking changes or requiring migration efforts. [2739d1a6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2739d1a6c31d6a4302c6b6cf864ecedcb53ce4d3)

## [5.0.1] - 2025-06-27

### Changed

- The watch resources function now correctly handles various status phase changes in Virtual Machine Instances (VMIs) thanks to the addition of a new test case that simulates and verifies these scenarios under different conditions without introducing any breaking behavior. [08ef6fa6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/08ef6fa60abab64a76c01860727e6f258479aa61)

## [5.0.0] - 2025-06-27

### Removed

- The current status variable is no longer accessible to users who relied on it for monitoring or logging purposes, and they must update their code to access the current status through other means. [9e09a462](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/9e09a462c7ff0372478e0dacfffd6c6f2b101e74)

## [4.0.5] - 2025-06-27

### Changed

- Simplified access to Virtual Machine Instance and Data Volume names by providing an AppContext instance that can be accessed through the KubevirtRunner. [82bcdde1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/82bcdde13f4bd89127d718a0f38ffc7730ec6e83)

## [4.0.4] - 2025-06-27

### Changed

- Reduced the default cleanup timeout to 5 minutes, affecting users who rely on automatic cleanup tasks and requiring them to update their scripts and workflows accordingly. [b55b780e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b55b780ec51322bb9e310d2f658a344fc775b30d)

## [4.0.3] - 2025-06-27

### Changed

- Updated the Docker build process to use the latest Git version and updated the working directory to `/app`, which may necessitate adjustments in users' build scripts or configurations if they rely on specific behavior from the previous Dockerfile. [2b296ea6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2b296ea68358509d52850d7dc433d71604f61e56)

## [4.0.2] - 2025-06-27

### Changed

- Updated the kubevirt client-go dependencies to modernize the version of k8s.io/kube-openapi and related packages. [b5a272a8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b5a272a8ffa2e864317fe2525b4f3e3003c8bb45)

## [4.0.1] - 2025-06-27

### Changed

- Corrected the typo in log messages to accurately reflect resource creation status without introducing breaking behavior or requiring migration steps. [c289626f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/c289626ff1b073248d5fc1a1c3499c433a4f4aa3)

## [4.0.0] - 2025-06-27

### Removed

- Eliminated the reporting feature which will break behavior for users who relied on this functionality and requires them to review their workflows accordingly. [f9d61fdc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f9d61fdce8b904bafbfa50aea820ecfd19ea6adc)

## [3.4.4] - 2025-06-27

### Changed

- The workflow's ability to push code and create pull requests on Fridays at midnight has been optimized by updating Dockerfile GitHub Action permissions to allow write access for specific actions. [7a068441](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7a068441fd3ffaadd290d48b925d422506f3d561)

## [3.4.3] - 2025-06-26

### Changed

- Enabled users to customize cleanup duration through environment variable-controlled timeouts and improved error handling for resource deletion operations. [8dae66f1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8dae66f17bd3c43e9caf343c61339a3b72abff44)

## [3.4.2] - 2025-06-26

### Changed

- Upgraded golangci-lint version, requiring users to review and adjust their project configurations for updated linter settings and configuration. [2383f2e7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2383f2e7701b0a8729377513652d1e3599d4f6d7)

## [3.4.1] - 2025-06-26

### Changed

- Improved error handling during cleanup is now enabled by logging errors at the debug level if they're not due to the resource already being deleted and ignoring them otherwise. [1a8e24b6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1a8e24b6f9cb211e2bc60a152eac6dbbdfa4290d)

## [3.4.0] - 2025-06-26

### Added

- Enabled logging of key Virtual Machine runner process stages, including resource creation, completion, and deletion, without introducing breaking behavior, altering the API or CLI contract, or requiring migration actions. [ecf9573c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ecf9573ce29fe666c9d9d18298fd473e93e1d178)

## [3.3.1] - 2025-06-24

### Changed

- Modernized documentation for the `kubevirt-actions-runner` to provide detailed information on its usage, architecture, and configuration, including key features, prerequisites, quick start guides, ephemeral VM creation, increased isolation, custom system-level configuration support, and integration with Actions Runner Controller and Kubernetes-native tooling. [73162ddb](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/73162ddb3679559f5318a20b9792d327529d57b7)

## [3.3.0] - 2025-06-24

### Added

- Enabled ignore-missing-schemas support for kubeconform by introducing the KUBERNETES_KUBECONFORM_OPTIONS flag in GitHub workflows. [3f923521](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3f9235215110b301105107945f9dacf3bb056f93)

## [3.2.1] - 2025-06-24

### Changed

- Updated the Docker image's base version to Alpine 3.22, which may require migration steps for users who have customized their images based on the previous version and now references golang:1.24-alpine3.22 instead of golang:1.24.0-alpine3.21 in the FROM instruction of the Dockerfile. [1c9bd97c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1c9bd97ca445825fe0780ff678ad2c3794682602)

## [3.2.0] - 2025-06-24

### Added

- Enabled Docker builds to bypass HEALTHCHECK checkov security checks by default, potentially requiring users to adjust their security settings accordingly. [899b9719](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/899b9719c809eec22c79aa1dfb4bf1d8c22c1177)

## [3.1.1] - 2025-06-24

### Changed

- Upgraded GitHub Action versions across multiple workflows to newer versions, primarily moving setup actions for Go, Docker login, metadata extraction, build and push, artifact attestation, linter validation, and spell checking from 4.x to 5.x or 6.x. [e5af489c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e5af489c754ce9150f9f143aa322ad9700a9ee85)

## [3.1.0] - 2025-04-09

### Added

- Enabled developers to continue development without strict adherence to go linter rules by disabling the project's go linter and introducing an option in the Makefile to disable go module validation. [8990975d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8990975d60f45959dd4fcfa10ae16362a586aa49)

## [3.0.4] - 2025-04-09

### Changed

- Updated to utilize the latest version of super-linter directly from its upstream repository, ensuring consistent and improved linter functionality without introducing breaking behavior or API changes requiring migration steps. [ac7cfeb6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ac7cfeb6501ede91e2621a6d6168a549936d5f69)

## [3.0.3] - 2025-02-14

### Changed

- Upgraded the project's Go version to 1.24.0, requiring corresponding updates to actions/setup-go and Docker images. [1cb578cc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1cb578cc1ea6026e6e9d3f9dea4926b6e9cbfafb)

## [3.0.2] - 2025-02-12

### Changed

- Updated linter configurations to reflect current best practices and package requirements, which may necessitate users reviewing their workflow files and updating project configurations accordingly. [45950e4a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/45950e4aff5278292376cd90142f1d87929b7a57)

## [3.0.1] - 2025-02-07

### Changed

- Runner methods now explicitly return errors on failure allowing for more precise error propagation and handling in the application. [d7477d74](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d7477d740dfe95d36f97c4101667ed6602eed4e1)

## [3.0.0] - 2025-02-04

### Removed

- Stabilized the delete resources use case by removing its unit tests which previously introduced breaking behavior and changed the API contract requiring two string arguments for the `DeleteResources` method of the `Runner` interface and its implementations. [cdb9f5b3](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/cdb9f5b3ed6d9be034c162fa5db7ec31707ef1dd)

## [2.4.1] - 2025-02-04

### Changed

- Enabled code coverage checks for Go projects in GitHub Actions workflow, requiring at least 40% code coverage to pass. [b6d9ca9f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b6d9ca9ffa129df18f92dbd08d648f200e587286)

## [2.4.0] - 2025-02-04

### Added

- Enabled secure access to repository contents and workflows without exposing sensitive credentials by authenticating actions with the WORKFLOW_TOKEN secret token instead of PAT usage. [40abec1b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/40abec1b5cefadc27fdf97c8476aa9967b645a5c)

## [2.3.3] - 2025-02-03

### Changed

- Improved validation for create resources use case by introducing error returns on empty inputs for virtual machine template, runner name, and Just-in-Time configuration. [e4c33b8d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e4c33b8d4b06d8d2f33da2882d91840ba171b5bf)

## [2.3.2] - 2025-01-31

### Changed

- Updated GitHub Actions versions now require workflows to be migrated if customized for specific action versions. [54039b3e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/54039b3e12294e55ea51036ba6f5ed67f25deab1)

## [2.3.1] - 2025-01-31

### Changed

- Improved unit tests for the kar command were enabled to cover various scenarios of runner creation processes including successful and failed outcomes. [01742283](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/0174228332f9223279ae9954c9073102e1a88978)

## [2.3.0] - 2025-01-23

### Added

- Enabled runtime display of build information including Git commit hash, tree modification status, build date, and Go version through new logging statements. [5adf666b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5adf666ba433c054e448e17270838cbf5a2d07dc)

## [2.2.2] - 2025-01-23

### Changed

- Enabled improved error handling for failed VMI instances by raising an error to provide a clear indication of the issue, introducing no breaking behavior and requiring no migration steps. [18208ce2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/18208ce291ef079420a4f8bb534958841851a8ee)

## [2.2.1] - 2025-01-23

### Changed

- Optimized log frequency to reduce the number of log messages sent every 5 minutes, potentially requiring adjustments in monitoring tools or scripts that depend on these logs. [8f5d519b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8f5d519b5bba5b6018a25a6a909d15f22f5f3e57)

## [2.2.0] - 2025-01-23

### Added

- Automated Dockerfile version updates are now triggered on pull requests to ensure consistent versions across all builds. [231e2a64](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/231e2a64dcdd209ebb1f9511435b58018ff4d11e)

## [2.1.2] - 2025-01-23

### Changed

- Updated the super-linter image reference in the Makefile to ghcr.io/super-linter/super-linter, which may break existing workflows relying on the old image and requires no migration steps or changes to the API or CLI contract. [edbf223c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/edbf223c93a3e06b68032cba69f8dc8881239489)

## [2.1.1] - 2025-01-08

### Changed

- Updated GH action versions enabled workflows for linter and spell checking to utilize the latest available functionality without requiring explicit migration steps. [484c6225](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/484c622500c0e48a9e113a9c3468b373536a850b)

## [2.1.0] - 2025-01-08

### Added

- The update version script now includes running `go mod tidy` to maintain the Go module's integrity, ensuring dependencies are correctly updated when changing versions. [6e24be6f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6e24be6f8ab5c5d9d4c40e75ee773de42fffbec2)

## [2.0.0] - 2025-01-08

### Removed

- Simplified the GoLangCI configuration file format to remove unnecessary lines and improve maintainability by reducing clutter in the configuration file. [4f0f484d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4f0f484d78b60f5ed26ca532a53736ce6002971c)

## [1.1.0] - 2025-01-08

### Added

- Automated creation of Docker containers is now enabled through a GitHub Actions workflow that triggers on push to the master branch and pull requests to the master branch, builds and publishes Docker images with versioned tags, and generates an artifact attestation. [6ac394bc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6ac394bc31f56208e573db40c4fc7db4a904db5b)

## [1.0.2] - 2024-12-19

### Changed

- DataVolumes now reference their owner through an OwnerReference field in the DataVolume's metadata, enabling users to manage DataVolumes based on their associated VirtualMachineInstances without requiring any modifications to their existing workflow. [a3df2c68](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a3df2c688cec490f8615ac5c882b25defe375ec4)

## [1.0.1] - 2024-12-19

### Changed

- Resources are now properly cleaned up after usage due to the replacement of context.WithCancel with NotifyContext method in several places. [350bb8ac](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/350bb8acd40d839229bf366e5cee03fd53e24b90)

## [1.0.0] - 2024-12-18

### Added

- Enabled initial project foundation for Kubevirt Actions Runner including configuration files workflows and dependencies that can be leveraged by developers to build upon this base. [f571e6da](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f571e6da5f69cb100d5dc85ef41672814bb51279)
