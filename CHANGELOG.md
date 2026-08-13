<!-- Markdownlint-disable MD024 -->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [18.14.1] - 2026-08-07

### Changed

- Standardized casing and escaping for "readme" and "go.\*" in the changelog to improve clarity and consistency with project documentation standards. [f0ae54ff](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f0ae54ffcb45afaa9c3299e9e432c45326222547)

## [18.14.0] - 2026-08-07

### Added

- Enabled detailed changelog entries for recent releases, including new terms recognized by automated checks during linting and validation, documentation consistency improvements, and expanded historical notes, thereby enhancing traceability and clarity for users reviewing project history. [76140cb4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/76140cb438c1719ce1fc21d90abe3ca8bca4257f)

## [18.13.0] - 2026-08-07

### Added

- Enabled two missing terms, dfa and ebcee, in the wordlist to prevent false positives during linting and validation. [6eaf6714](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6eaf6714a837d04a6cf9b437d0c048b72ff9f17c)

## [18.12.1] - 2026-08-07

### Changed

- Clarified references to "readme" and "go.\*" in the changelog for improved consistency and alignment with project documentation standards without affecting functional behavior. [054746c2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/054746c222cb988486d6ca82d100f1723a939154)

## [18.12.0] - 2026-08-07

### Added

- Optimized project transparency by introducing detailed changelog entries for versions 18.10.0 to 18.11.1 that document improvements like wordlist optimization, reproducibility enhancements, and documentation alignment without introducing breaking changes or migration requirements. [a9ecba87](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a9ecba87e4b4869f6145c626c397b08f61bc1921)

## [18.11.1] - 2026-08-07

### Changed

- Standardized project documentation by renaming references to README.md to readme.md for consistency with project file naming conventions. [7fc97d7e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7fc97d7efd7e3f3e00d505ebcee41789296cf89b)

## [18.11.0] - 2026-08-07

### Added

- Enabled smoother CI runs and reduced noise for contributors by adding missing terms to the custom wordlist, specifically bugfixes and cecf, which were causing false positives in spelling checks. [cd9f928f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/cd9f928f5392769d7eb0a42d11c739c823692dfa)

## [18.10.1] - 2026-08-07

### Changed

- Stabilized linting and formatting results across environments by removing stale dependencies and virtual environments through an automated cleanup step prior to running these tasks. [b73686da](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b73686da69d98f2d8509a16d0577251b1a1ff73b)

## [18.10.0] - 2026-08-07

### Added

- Enabled detailed tracking of feature evolution and project progress through a comprehensive changelog documenting new features, enhancements, and fixes for versions 18.7.0 to 18.9.1 without introducing breaking behavior or requiring migration steps. [8046630c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8046630cb5e2f5f8e90796225313b82482da29a1)

## [18.9.1] - 2026-08-07

### Fixed

- Stabilized the fmt target by ensuring that all applied fixes are reflected in the final output through consistent ordering of code formatting tools. [91b75ef2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/91b75ef21f20b9d2e6c93f133452f801bf924273)

## [18.9.0] - 2026-08-07

### Added

- Enabled enforcement of consistent terminology usage in documentation and prose through textlint configuration. [41b9df58](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/41b9df58e41171b78d4e7b2829f101783305fd41)

## [18.8.1] - 2026-08-07

### Changed

- Standardized terminology and capitalization for tools, technologies, and workflow steps to improve readability and reduce confusion among contributors and users referencing the changelog. [2f49d0c7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2f49d0c7a0d0d5c7bb23ba596ed940780b3d8e42)

## [18.8.0] - 2026-08-07

### Added

- Enabled more accurate spelling and linting by expanding the custom wordlist to include new domain-specific terms, abbreviations, and project identifiers, reducing false positives on technology names, internal function names, and common acronyms observed in the codebase. [643c0806](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/643c08061cf82ee17b14a1a837b851238cecf265)

## [18.7.0] - 2026-08-07

### Added

- Introduced a comprehensive changelog following Keep a Changelog and Semantic Versioning to provide clear release notes for users and maintainers, including versioned entries with links to commits, summaries of changes, and context for each release. [224740a1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/224740a1bede9df8c51b7b063079c99934e65395)

## [18.6.0] - 2026-08-07

### Added

- Enabled more flexible Markdown content organization by allowing secondary headings in self-contained documentation while maintaining Jekyll front-matter as the primary title source. [ddb2d6cb](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ddb2d6cb659fe3cdc492d95cd5bab4a4d0593868)

## [18.5.1] - 2026-08-07

### Changed

- Optimized the markdownlint pre-commit hook to utilize the faster and more maintainable markdownlint-cli linter, thereby reducing setup friction in CI environments. [b34b19ff](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b34b19ffe41f26fbded7c9b3a0bbccf487af670a)

## [18.5.0] - 2026-08-07

### Added

- Enforces Go linting standards during formatting by installing and running golangci-lint to autofix style issues improving code quality and consistency across the project. [5bd8a1c0](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5bd8a1c0d8e68b74aaec0db8745e0ef2e62f015c)

## [18.4.0] - 2026-08-07

### Added

- Improved test case readability by introducing empty lines after os.Args assignments to visually separate setup steps without altering functionality. [ce71f230](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ce71f230ddffff728aa4cee724850626fa571e4b)

## [18.1.2] - 2026-08-07

### Changed

- Optimized workflow frequency by restricting execution to the third Friday of each month. [dc77fdfe](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/dc77fdfe2173331ffd57fcad07fd5ded122fbd17)

## [18.1.1] - 2026-08-07

### Changed

- Updated several GitHub Actions and Go module dependencies to their latest versions to ensure continued compatibility and security, incorporating upstream bugfixes and improvements from OpenTelemetry, kubevirt.io/containerized-data-importer-api, grpc-gateway, and github.com/go-openapi packages. [b6d8fe1f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b6d8fe1f542823876cf4b0407414e8896a91245a)

## [18.1.0] - 2026-08-07

### Added

- Stabilized code coverage for the cmd/kar package's main entrypoint by introducing a new test that exercises previously untested paths and improved statement coverage from 0% to 94.5%, also increasing mutation testing coverage to 95.45%. [959b4bc7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/959b4bc759d910a75e1368d9eed0832a4bf9c9d0)

## [18.0.0] - 2026-08-07

### Removed

- Simplified configuration by eliminating unnecessary dependencies in .golangci.yml, removing github.com/golang/mock/gomock and github.com/pkg/errors as they are no longer imported by the project. [ef5d590c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ef5d590c72fc705d697aa51e13e74105345dba78)

## [17.1.8] - 2026-07-31

### Changed

- Simplified error handling in tests by introducing shared variables for simulated failures. [6413c56c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6413c56c73575ff51117df92718042727af391d7)

## [17.1.7] - 2026-07-31

### Changed

- Updated several Go module dependencies and GitHub Actions workflows to their latest releases for improved compatibility, bugfixes, and security patches, introducing no breaking changes but requiring downstream consumers to verify integration with the updated dependencies. [60d1a24b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/60d1a24bc49bb980e4e133c305b4d3e9ae818434)

## [17.1.6] - 2026-07-31

### Changed

- Simplified error and shutdown assertions in tests to reduce duplication and improve readability, allowing developers to more easily verify these behaviors without introducing any breaking changes. [c45de95a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/c45de95a41bf322ac7ec431cf8efdeb46c6ae8dc)

## [17.1.5] - 2026-07-31

### Changed

- Optimized code coverage by adding tests for key functions in cmd/kar and internal runner packages resulting in increased test success rates from 84.9% to 90.1%. [e6fd7bf4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e6fd7bf450d5c856d286bf8f0d0b96f18f1c3605)

## [17.1.4] - 2026-07-31

### Changed

- Simplified the update logic for GitHub Actions by consolidating duplicated functions into a single generic helper function that can handle different actions without changing behavior. [6e039c41](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6e039c41a7371c81dd1a7b38169b4c98b9570082)

## [17.1.3] - 2026-07-24

### Changed

- Simplified VCS build setting keys and runMainApp function signature to reduce cognitive overhead and improve clarity in table-driven tests without introducing any functional changes to application behavior. [34c1e7c5](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/34c1e7c556556f559954d572ee1a1e0cf93da9c7)

## [17.1.2] - 2026-07-24

### Changed

- Upgraded GitHub Actions and Go dependencies to ensure compatibility with the latest features and security patches across all workflows. [28d84a26](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/28d84a261199fb70a10668695712ce97f79efe5a)

## [17.1.1] - 2026-07-24

### Fixed

- resolved issues with artifact uploads in CI runs by changing ownership of exported logs to the current user and listing directory contents for easier debugging. [a1fbe002](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a1fbe0024a8b7416d5642a7e877648860fc182ba)

## [17.1.0] - 2026-07-24

### Added

- Enabled increased code coverage for the cmd/kar/main.go file by introducing unit tests in main_test.go that now cover 60.6% of statements. [3169678e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3169678effbd5b48e4696a7ddafa92302ef569e7)

## [17.0.0] - 2026-07-24

### Removed

- Simplified logging behavior by eliminating redundant Printf method logic and delegating to Infof without introducing any breaking changes or migration requirements. [e61ae8f8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e61ae8f8d50d3fb1b2b217d770e390a1e3b408fc)

## [16.2.1] - 2026-07-17

### Changed

- Enabled comprehensive verification tasks for Copilot CLI by relaxing tool and path restrictions while maintaining security through explicit denial of dangerous shell commands. [667cbe2c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/667cbe2c43d278b508f2eafb33e3e5df2039fc5d)

## [16.2.0] - 2026-07-15

### Added

- Improved automation and easier integration with external tools for maintainers are now enabled through dual-mode linter failure reporting that includes human-readable summaries and structured machine-report JSON payloads. [e62e5a8d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e62e5a8dced88c923c54b8c6bdffbe8ba0e10040)

## [16.1.5] - 2026-07-11

### Fixed

- Resolved permission errors accessing Kubernetes cluster logs in CI environment by running kind export logs command with elevated permissions. [3a620304](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3a620304c02f10c794c07a19b872abe8d7093ba0)

## [16.1.4] - 2026-07-11

### Changed

- Simplified the resolution of semantic version tags from remote repositories by encapsulating the logic in a reusable function, reducing code duplication and making future maintenance easier with no functional changes expected. [df0a1397](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/df0a13974305307575f5659b7e25e664c3b7ec12)

## [16.1.3] - 2026-07-11

### Changed

- Enabled proper coverage reporting by allowing the gwatts/go-coverage-action to annotate commits during analysis. [43682a91](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/43682a9162e583d2635cff01804b2156196f512a)

## [16.1.2] - 2026-07-11

### Changed

- Upgraded several dependencies to their latest versions to address compatibility and security concerns, ensuring the project remains compatible with upstream libraries and benefits from recent bugfixes and performance improvements. [91b83fad](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/91b83fad0cba07cee037b275e41b245f311807f7)

## [16.1.1] - 2026-07-11

### Changed

- Updated GitHub Actions for the repository's workflows to their latest patch versions, ensuring security and bugfixes are applied, reducing maintenance overhead and improving reliability. [27df61ea](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/27df61ea20eff9faddab7c92f329b8e89bab222b)

## [16.1.0] - 2026-07-11

### Added

- Automatically resolves and updates golangci-lint, gremlins, and rtk versions in GitHub Actions workflows to ensure CI uses the latest dependencies with reduced manual intervention. [893b6ed1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/893b6ed1b5c69b74545ea52cc50c4e53d2a4b0e6)

## [16.0.4] - 2026-07-11

### Changed

- Hardened the copilot-cli configuration to restrict tool usage and limit path and URL access by explicitly denying dangerous shell commands and only allowing a curated set of trusted domains. [d76f5ab9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d76f5ab9fc31059aa71f675b7b86f146fd1aaff4)

## [16.0.3] - 2026-07-07

### Changed

- Optimized CI workflow efficiency by conditionally running Go-related jobs only when Go files have changed, reducing unnecessary resource usage and build times. [dc1af83b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/dc1af83b5ab8fac1f73b64e6c7d88bccdb49c8fe)

## [16.0.2] - 2026-07-07

### Changed

- Standardized GitHub Actions workflows to improve maintainability and onboarding by clarifying event triggers, adding concurrency groups, pinning tool versions, and updating documentation with a quick context section and an execution map. [40bc7fa6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/40bc7fa60bd108109c2c8261b0949247eaf8a31b)

## [16.0.1] - 2026-07-02

### Changed

- Optimized test coverage by adding deterministic and isolated tests for previously uncovered code paths via go tool cover analysis resulting in a 4.1% increase to 94.7%. [98dc50b8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/98dc50b837319eeb179f641b73ea38bdfea93727)

## [16.0.0] - 2026-07-02

### Removed

- Simplified demo scripts by eliminating duplicated status-reporting logic and dead code, resulting in ~20 lines of reduced shell complexity. [81cb5b24](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/81cb5b2487adb8b645273a5bfbd8cc65f57d0379)

## [15.2.2] - 2026-07-02

### Fixed

- Resolve attempts to create a Virtual Machine Instance that already exists by returning the existing VMI and logging an error message without introducing any breaking behavior or migration requirements. [bdf23444](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bdf234447b8738eb55eabb738281dfe2a9b41098)

## [15.2.1] - 2026-07-01

### Changed

- Simplified DataVolume creation logic to improve code organization and maintainability by extracting the optional creation into a new helper method createOptionalDataVolume without introducing any functional changes or breaking behavior. [28078917](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/28078917388d6a2defac219d242e073857f11445)

## [15.2.0] - 2026-07-01

### Added

- Enabled spell checks to accurately identify commonly used terms in Kubernetes and testing contexts by incorporating "kube", "testvm", and related variations into the wordlist without impacting application logic. [1a1fb5d4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1a1fb5d4a2a5380c9b667d6393a29f8e12d6816b)

## [15.1.0] - 2026-07-01

### Added

- Enabled centralized template management by allowing users to specify the namespace from which to fetch the VM template independently of the runner's namespace. [531fb694](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/531fb6943bae02604b7d0207e5187d142e9e8049)

## [15.0.0] - 2026-06-25

### Removed

- Simplified repository clutter by eliminating a generic code review instructions template that is no longer needed and is now maintained separately for each project, without affecting functional or behavioral code. [860862be](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/860862be222c55bce2e14ffdf6bd2f044e96d5f6)

## [14.0.0] - 2026-06-25

### Removed

- Simplified the repository by eliminating outdated review-and-refactor skill documentation that no longer reflects current project practices and workflows. [c2817b39](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/c2817b39cdf53dbcb84beb37881167a99afad06e)

## [13.0.0] - 2026-06-25

### Removed

- Simplified documentation by eliminating redundant Makefile authoring instructions that are no longer maintained and have been superseded by project-level documentation and inline comments within individual Makefiles. [43f34dd2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/43f34dd23d74f03cc6d4cf6161714275d424dd20)

## [12.0.0] - 2026-06-25

### Removed

- Simplified maintenance efforts for maintainers by eliminating redundant functionality through consolidation of the principal software engineer agent's responsibilities under a unified engineering guidance agent. [d5aad069](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d5aad069bd60629594ba8b978042ac4bc8957005)

## [11.1.0] - 2026-06-25

### Added

- Enabled global Copilot token management and usage tracking by integrating the rtk token saver into both fixer and verifier workflows without impacting existing job logic. [b7eeca91](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b7eeca912fccbebcfe10df8715a4057e05413460)

## [11.0.1] - 2026-06-23

### Changed

- Optimized license headers in Go source files to prevent jscpd from flagging them as duplicate code, reducing noise in duplication reports. [1de8a0f4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1de8a0f49d86acb70717027ed4f2aaacd04a42b4)

## [11.0.0] - 2026-06-23

### Removed

- The GitHub Actions workflows now utilize the latest Copilot CLI version at v3.2, ensuring compatibility with its features and security updates, while introducing a potential migration requirement for future updates. [32bd4257](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/32bd4257fc536557fbc962d9af3f87de4041aea4)

## [10.0.2] - 2026-06-23

### Changed

- Optimized automated updates of GitHub Action hashes to rely on SemVer tags only and handle annotated and lightweight tags more reliably. [8be2c280](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8be2c280b656ebd0f36a96cb1f863e156fa7ca94)

## [10.0.1] - 2026-06-23

### Changed

- Simplified GitHub Actions workflows to consistently reference their latest available tags, reducing the risk of using outdated action references and making future maintenance more straightforward. [bafb6e71](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bafb6e71131fdcf835c29df97c26340aabe25463)

## [10.0.0] - 2026-06-19

### Removed

- Simplified the data volume test in runner_test by removing unnecessary whitespace to improve code readability and consistency with surrounding tests without introducing breaking behavior or altering the existing functionality. [7a59bc19](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7a59bc19b457f0f67615f7a38a3aa6de9a73dce0)

## [9.7.4] - 2026-06-19

### Changed

- Upgraded several dependencies to their latest versions, ensuring continued support and reducing the risk of vulnerabilities from outdated packages, without changing any application code. [a6d30861](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a6d30861343aca0ca05da146dcfa0433d15c0796)

## [9.7.3] - 2026-06-19

### Changed

- Simplified test maintenance by extracting reusable logic and constants to reduce duplication and improve readability while preserving existing test coverage and behavior. [924b7377](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/924b737736920b7091abd93dff24f1a2d806cca6)

## [9.7.2] - 2026-06-19

### Changed

- Updated the ai-prepare-commit-msg hook to leverage upstream improvements and bugfixes for commit message generation without requiring configuration changes. [3f0464d1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3f0464d1e50712a766a6445de1a8df5331fd007a)

## [9.7.1] - 2026-06-18

### Changed

- Simplified the runner code to eliminate duplicated strings and checks, improving maintainability without introducing any breaking behavior or API changes. [2e96c41f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2e96c41f91268fa464c44a39fcd3e97f27e11131)

## [9.7.0] - 2026-06-18

### Added

- Enabled comprehensive test coverage for previously untested code paths, increasing overall statement coverage from 70.5% to 92.2%. [68d46dd9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/68d46dd9eb8525ee3028681cfc4265dda7904207)

## [9.6.2] - 2026-06-18

### Changed

- Clarified documentation content by updating code blocks to use "text" language and correcting the heading from "antipatterns" to "Antipatterns", maintaining alignment with standards without introducing breaking behavior or migration requirements. [5f182399](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5f18239916873c8a42c56c23eacd5c541ab721e8)

## [9.6.1] - 2026-06-18

### Changed

- Enhanced documentation clarity by introducing visual diagrams that explicitly illustrate key control flows and lifecycles for improved onboarding and troubleshooting. [bda81f17](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bda81f17671365389e2cc58f6c9f6a4d7d4527f5)

## [9.6.0] - 2026-06-18

### Added

- Enabled automated code coverage improvement through a dedicated QA subagent and GitHub Actions workflow that analyzes gaps, generates targeted tests, and submits focused pull requests to maintain high test quality. [b6922b18](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b6922b18bad6fed9eebd840db5c79c55ee1480e5)

## [9.5.0] - 2026-06-17

### Added

- Enabled consistent terminology across the project by adding Ctrl, jitconfig, ParseDuration, and SIGTERM to the wordlist used by linters for enforcing standardized usage throughout the codebase. [5156f568](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5156f56833fd52f5c898f18b756e8590e561d286)

## [9.4.3] - 2026-06-17

### Fixed

- Updated to the latest version of Git, 2.54.0-r0, ensuring compatibility with current build requirements and access to recent bugfixes and security improvements for all environments using this Dockerfile. [d6535c99](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d6535c99de063288b7b2a656a90ad47e47507649)

## [9.4.2] - 2026-06-17

### Changed

- Modernized documentation structure to clearly separate learning tutorials from task-oriented how-to guides and reference material, improving discoverability and navigation efficiency for users. [f95c9db9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f95c9db9c5798e098bafd2cd64324ff2461d5de2)

## [9.4.1] - 2026-06-17

### Fixed

- Resolved the base image to utilize the latest security patches and bugfixes available in alpine 3.24 without introducing any breaking changes to the build process or application functionality. [480ebb71](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/480ebb7159f791b4ab087ac660f52ba1d60fea6a)

## [9.4.0] - 2026-06-17

### Added

- Enabled Gremlins mutation testing in the build workflow to strengthen code resilience and fault detection by identifying potential weaknesses in the test suite. [fcff167a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/fcff167a0b8c40263b5198d3b260e441af961630)

## [9.3.0] - 2026-06-17

### Added

- Enabled automated commit message suggestions during the pre-commit process by re-enabling and updating the ai-prepare-commit-msg hook configuration to utilize the correct repository URL and revision. [a4c56584](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a4c565847b26098f44715876a74b083d684900d1)

## [9.2.15] - 2026-06-12

### Changed

- Simplified the codebase by eliminating unnecessary nil checks and improving consistency in environment variable usage, resulting in no observable impact on behavior for developers or operators. [817b68c5](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/817b68c5ab5096417d1f47d51be93a7505ce76e0)

## [9.2.14] - 2026-06-05

### Changed

- ci/update_versions.sh has been optimized to preserve partial updates and handle errors more robustly without breaking changes to the API contract or requiring migration efforts from developers or operators. [1ebce834](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1ebce8341782aee6a6a15d84bddeb173163b73cb)

## [9.2.13] - 2026-06-05

### Changed

- Updated the project's Go dependencies to their latest versions, including go.opentelemetry.io/otel and its related packages at version 1.44.0, k8s.io/API at 0.36.1, and kubevirt.io/API at 1.8.3. [8def3a3f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8def3a3f1a3d69ab45213135216d55093a46302d)

## [9.2.12] - 2026-06-04

### Changed

- Pinned actions/checkout to version 6.0.3 in various GitHub workflows, upgrading underlying dependencies without introducing breaking behavior. [5eca7d3c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5eca7d3cc3197be3ad58f5a4e8613b55bfbc4c82)

## [9.2.11] - 2026-06-04

### Changed

- Simplified error handling and reduced technical debt by consolidating variable declarations, removing redundant imports, and refactoring code to improve maintainability and readability without introducing any breaking changes. [f337f2f8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f337f2f89629dbf521d318111c4bdcbae19b4d54)

## [9.2.10] - 2026-06-03

### Changed

- Updated versions of GitHub Actions, Docker dependencies, and pre-commit configuration files to ensure consistency and reflect the latest available versions, potentially requiring manual updates from users who rely on these specific versions for their workflows. [1fec048d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1fec048d5f576bde09581e19291863e54d4d51e9)

## [9.2.9] - 2026-06-01

### Changed

- Simplified build information retrieval and telemetry shutdown handling to eliminate panics and golangci-lint violations without introducing breaking changes that affect developers or operators. [7b823ed2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7b823ed204c424b3f38a750808bf8a809f244175)

## [9.2.8] - 2026-05-28

### Fixed

- Enabled gomodguard_v2 by default in the GoCI configuration to replace and harden against the deprecated gomodguard linter. [bd8401ff](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bd8401ff8a66037232b9dcf7f2cb7477d252772a)

## [9.2.7] - 2026-05-28

### Changed

- Optimized VMI wait/status flow to reconnect watch on stream closure and protect API server against excessive requests during context cancellation, enabling more reliable phase transitions in WSL v5 environments. [dd550f89](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/dd550f89292c510ccbcdd29ec659c8475ec26d4c)

## [9.2.6] - 2026-05-28

### Fixed

- Stabilized the runner's behavior against temporary Kubernetes API issues by automatically reconnecting the VMI watch on stream closure without breaking any existing functionality. [e0024a27](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e0024a27c97d139f119b0b0d3266c2cef6de33d3)

## [9.2.5] - 2026-05-22

### Changed

- The package-level sentinel errors have been modernized in internal/runner.go to satisfy golangci-lint requirements and improve overall error handling. [7beb720d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7beb720dc91d41e542832da607717312cd1db04f)

## [9.2.4] - 2026-05-22

### Changed

- Updated versions files to ensure dependencies remain current with security patches and bugfixes, requiring no migration steps or breaking behavior changes. [a4432b61](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a4432b61b25397ae4c2d11307bb22a003014f948)

## [9.2.3] - 2026-05-15

### Changed

- Optimized test functions to eliminate redundant loop variable copies thereby improving performance and efficiency. [7f56e593](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7f56e593c6a6d6ac21e63ca197ac8303aa847fdd)

## [9.2.2] - 2026-05-15

### Changed

- Updated versions of the `actions/ai-inference` action in GitHub Actions workflows to 2.1.0 from 2.0.8 for analysis and linter checks. [f01c4513](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f01c451346c9306198837a01d96ea1839457ab87)

## [9.2.1] - 2026-05-14

### Changed

- Increased the default wait timeout from 10 minutes to 1 hour for Kubevirt Actions Runner, allowing more time for terminal VMI phases to complete and enabling users to configure the new timeout value using the KAR_WAIT_TIMEOUT environment variable. [1aaf5577](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1aaf5577dda0daa1181a10c4ff0a407818352fa3)

## [9.2.0] - 2026-05-14

### Added

- Enabled configurable wait timeout for terminal VMI phase, introducing Running+Ready milestone that ends wait without requiring Succeeded phase, with default timeout of 10 minutes adjustable via KAR_WAIT_TIMEOUT environment variable. [6a853fc3](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6a853fc3c99c9a644998aa3e3c73576e5246df32)

## [9.1.6] - 2026-05-08

### Changed

- Simplified logging and build information creation to improve maintainability without introducing breaking changes or security vulnerabilities. [3c1ca8ce](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3c1ca8ce936ce86798f9fa171d71698fd7fcd89e)

## [9.1.5] - 2026-05-01

### Fixed

- Resolved issues caused by actions/setup-copilot not being found by preventing austenstone/copilot-cli from auto-updating to v3.0 and instead locking it at version 2.0 in workflows. [9ac1d458](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/9ac1d45853edb80722230f39a18f49aa6b944379)

## [9.1.4] - 2026-04-20

### Changed

- Updated GitHub Actions workflows to utilize newer versions of several actions, ensuring continued correct function with the latest available tools. [6a84162a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6a84162a986a6a31b4639086575e495c0a5380f7)

## [9.1.3] - 2026-04-11

### Fixed

- Resolved documentation duplicity by removing redundant content from the Diátaxis Documentation Expert skill, allowing users and maintainers to review updated README.md for alternative resources and contributing information without impacting API or CLI contracts. [fea2e201](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/fea2e201652786d5a11b89717a3109e292d4a34c)

## [9.1.2] - 2026-04-07

### Fixed

- Resolved multiple linter failures and code review feedback issues, ensuring that Makefiles and shell scripts are now lint-free and adhere to best practices for maintainability and portability. [9b026b66](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/9b026b66a888705b442148cc652ab73b8cc49801)

## [9.1.1] - 2026-04-05

### Fixed

- Improved reliability by introducing a 5-minute timeout to the WaitForVirtualMachineInstance method and restricting GitHub tool access permissions for tighter security controls. [d5babba4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d5babba40cd8ef22322232cc69378726d09dab93)

## [9.1.0] - 2026-04-05

### Added

- Enabled clear and consistent documentation for various programming languages and tools by introducing custom GitHub instructions, agents, and skills to the repository. [0e9d3162](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/0e9d316275320e0da4058a45077554296706d7f6)

## [9.0.5] - 2026-04-03

### Changed

- Modernized versions files and GitHub Actions tool dependencies to ensure compatibility with the latest super-linter and other tools, resolving linting issues and simplifying configuration management for developers. [bbf36644](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bbf366448cf0046eec577ec096da65399a2ff96e)

## [9.0.4] - 2026-03-30

### Fixed

- Updated Docker images to utilize the latest Git version 2.52.0-r0, requiring users who relied on the previous version 2.49.1-r0 to rebuild their images. [5e82fb61](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5e82fb61c9e855c2749736a7f7adc9975556a102)

## [9.0.3] - 2026-03-27

### Changed

- Updated GitHub Actions workflows to utilize the latest available versions of Go and other dependencies without introducing breaking behavior or requiring migration efforts. [6a06c04d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6a06c04d758eec28da63f99a870028de0ebb9f4b)

## [9.0.2] - 2026-03-20

### Fixed

- Resolved demo execution reliability by ensuring the virt-handler is fully operational before proceeding. [34f22cd4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/34f22cd4fc530d5e6f6a16e8f36acbba0cd8b939)

## [9.0.1] - 2026-03-20

### Fixed

- Resolved scheduled version update workflow failures by robustly extracting major.minor Go versions and updating the Dockerfile FROM-line using the Docker Hub API. [d04e41af](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d04e41af9b9a11662843cb19151e0953daaac718)

## [9.0.0] - 2026-03-13

### Removed

- Stabilized DeleteResources functionality to prevent fatal crashes on uninitialized AppContext by ensuring the presence of an initialized context before proceeding. [cdaaa3dc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/cdaaa3dc544bab4371eee86157aa8e5a51a4fbb5)

## [8.0.2] - 2026-03-13

### Changed

- Updated the workflow configuration for GitHub Actions to include more detailed information in the build process and Docker image creation steps with no breaking changes required. [d194d7ac](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d194d7aceeefca883a0f8eb6178e3612311ed77a)

## [8.0.1] - 2026-03-13

### Fixed

- Resolved an issue in Docker build workflows for GitHub Actions, eliminating the requirement for pull requests targeting the master branch to trigger builds and allowing users to simplify their workflow configurations accordingly. [26910142](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2691014241233ba9f001385849b6a68f9bf62100)

## [8.0.0] - 2026-03-13

### Removed

- Simplified navigation through project sections by eliminating redundant content that was previously accessible via index.md information. [1966ee0f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1966ee0f1eed4782a8a5aff6d9771b29f966d189)

## [7.20.7] - 2026-03-13

### Fixed

- Resolved issues with golangci-lint and Docker build process in appcontext, ensuring successful linting and building for users relying on these tools. [a28a6002](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a28a6002f71957b9f2ff10f62a66697cf7614f2c)

## [7.20.6] - 2026-03-13

### Fixed

- Resolved the lack of visual context in the README.md file by adding a diagram image that enhances user understanding of the project's functionality and its relationship with Kubernetes and GitHub Actions. [d9424fda](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d9424fdacb4f11459f3a4fd7bc51c29221d1fbe2)

## [7.20.5] - 2026-03-13

### Fixed

- Resolved GitHub Pages deployment issues by enabling a Python script to remove Runme-specific code fence attributes from Markdown files in the documentation site. [b93ae3f6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b93ae3f6ed9dfd33bdb9fa3da0daabe5eeb0693b)

## [7.20.4] - 2026-03-13

### Fixed

- Resolved issues related to importing the Git repository in Docker builds from non-git directories by removing unnecessary instructions and ensuring successful project code import. [3f62f463](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3f62f463bb53e081e90330245803ccf69e0c8cca)

## [7.20.3] - 2026-03-13

### Fixed

- Resolved super-linter failures for NATURAL_LANGUAGE, Markdown, and JSCPD checks by addressing textlint and markdownlint issues in the README.md file. [5e56b630](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5e56b63013afb82a1518cad7dcdd293f62f81095)

## [7.20.2] - 2026-03-13

### Changed

- Simplified the testbed guide documentation to remove duplicate content and improve clarity for users by replacing a detailed architecture overview link in place of the removed flowchart without introducing any breaking behavior or migration requirements. [732ec425](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/732ec425c0523f7c0dd4fc3ea16168acb10ec6ec)

## [7.20.1] - 2026-03-13

### Fixed

- The GH action triggers for workflows have been optimized to only run the `fixers.yml` workflow when an opened issue already includes a matching remediation label. [d4b3c217](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d4b3c21772631da7f28f35d5cfc0724202c1b800)

## [7.20.0] - 2026-03-13

### Added

- Enabled GitHub workflows documentation for CI validation, maintenance, and automation, including purpose, trigger, and description for each workflow file without introducing breaking behavior or API changes. [ecf4b03a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ecf4b03a5eeae8507bec3220a0bc034a1ee23c5b)

## [7.19.0] - 2026-03-13

### Added

- Improved the README.md file documentation to better explain the KubeVirt Actions Runner project's features and benefits for running GitHub Actions workflows in isolated virtual machines using KubeVirt. [b496c5a4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b496c5a461b9b0c1db705c21a33e064c5e2f6b5e)

## [7.18.0] - 2026-03-13

### Added

- Code comments now provide metadata about the build, including the Git commit hash and build date, allowing developers to track version information. [bfd76a05](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bfd76a053fa91589f4fcb601a6b0a6562aa94ff2)

## [7.17.0] - 2026-03-13

### Added

- Introduced an architecture overview diagram in docs/explanations/architecture-overview.md and updated the README.md to include a link to it. [821604c8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/821604c85d05c125b1ffcece32be1e06bc0d8626)

## [7.16.0] - 2026-03-13

### Added

- Introduced Cayman theme for documentation, which may affect how documentation appears online and requires users to review their documentation settings for the desired appearance. [56f2592a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/56f2592a7a714a1c4dad751bbfd5d5a123933736)

## [7.15.0] - 2026-03-13

### Added

- Optimized documentation navigation by reorganizing GitHub Pages under consistent naming conventions and moving files to the `how-to-guides` and `tutorials` directories without introducing breaking behavior or migration requirements. [84d2e4f0](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/84d2e4f08fbb2e08f480459d3b028adf89a93739)

## [7.14.0] - 2026-03-13

### Added

- Enabled improved organization and clarity of quickstart instructions by moving documentation to a separate guide in docs/how-to-setup-quickstart.md. [6c472840](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6c472840477f076c6bfdaded92049d142e55cb21)

## [7.13.2] - 2026-02-28

### Fixed

- Resolved the Markdown linting issues in the docs without introducing any breaking behavior or requiring migration steps. [a10c04b7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a10c04b7ff8d10567019a939ed9b3b3da328a421)

## [7.13.1] - 2026-02-28

### Fixed

- Stabilized linting for Go code by removing duplicate package declarations and clarifying formatting rules for code organization and documentation to improve adherence to community standards. [974b827d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/974b827d41932beae430585a8d2f228b19a63eda)

## [7.13.0] - 2026-02-28

### Added

- Established clear guidelines for contributors to ensure consistency in code quality and adherence to best practices through the introduction of a contributing file outlining coding standards, testing, security, documentation, and development workflow. [3258012c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3258012ca9022f7783fcabc15a12968fdb85de7b)

## [7.12.0] - 2026-02-28

### Added

- Improved telemetry setup is now simplified for users through detailed configuration instructions and examples for OTLP exporters and stdout exporters, as well as verification steps for Jaeger and local debugging. [5f9be7fa](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5f9be7fa0120ec66b7420de0928eaa826977f111)

## [7.11.2] - 2026-02-27

### Fixed

- Resolved issues related to Markdown and natural language in the Diátaxis Documentation Expert skill documentation, ensuring consistency in tone and style without introducing any breaking behavior or migration requirements. [02c8b8d5](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/02c8b8d5a2eb56cadcf4805725dd6fa5722a808b)

## [7.11.1] - 2026-02-27

### Fixed

- Resolved missing words in dictionary used for GitHub Actions wordlist to ensure proper handling of keywords during automated tasks without affecting API or CLI contracts and introducing any security risks. [18f94788](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/18f9478899af0ba2e4a2b13c96fd97642d3b1e8b)

## [7.11.0] - 2026-02-27

### Added

- Introduced setup testbed documentation enabling users to configure and validate a local testing environment for `kubevirt-actions-runner`, including automated install scripts and customizable VM templates. [f7a19689](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f7a1968957b8dd59dbd92673e249c43200684aa5)

## [7.10.0] - 2026-02-27

### Added

- Enabled users to contribute high-quality software documentation by introducing a skill outlining four document types and guidelines for contextual awareness following the Diátaxis Framework's principles and structure. [1b8c9e14](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1b8c9e140fd239e15da42d220df140d02335d489)

## [7.9.3] - 2026-02-14

### Fixed

- Triggering of on-demand CI jobs is now optimized to include all pull requests, including opened, synchronized, and reopened ones, as well as submitted reviews, without introducing any breaking changes that require maintainers to review their setup. [4797133c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4797133cdfae46eb148249f5b1bbe5585bdb4404)

## [7.9.2] - 2026-02-14

### Fixed

- Resolved documentation inaccuracies in the Janitor agent's "Clean any codebase" section by correcting minor wording issues to ensure precise terminology is used. [a1de3c32](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a1de3c325daa17d12a7c6579413925d7cdd57ea9)

## [7.9.1] - 2026-02-14

### Fixed

- Simplified the janitor prompt task in the workflow to clearly guide users through reducing technical debt and maintaining clean codebases without breaking public APIs. [b1689a76](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b1689a7624afd7f4a803923bfc68c2441cb39738)

## [7.9.0] - 2026-02-14

### Added

- Enabled the janitor agent to perform cleanup tasks on codebases, including tech debt remediation, simplification, and removal of unnecessary abstractions, which impacts the Docker build process by removing a workflow that resolved build and runtime failures. [40df54db](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/40df54db7f6d7b0d3a1c1b6ec5f7472910c554f2)

## [7.8.0] - 2026-02-14

### Added

- Automated Docker image build and runtime failure analysis is now enabled through the addition of AI-powered container failure analysis and a failed-build-issue-action in GitHub workflows. [13ad46df](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/13ad46df4a09af1ce2083dc8fc97d1d67c2b2f8f)

## [7.7.1] - 2026-02-12

### Fixed

- Resolved all 14 golangci-lint wsl_v5 whitespace violations in test files to maintain code readability without introducing breaking behavior or API/CLI contract changes. [6467558c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6467558c0b962aaf8a056fba99db516ddefe8524)

## [7.7.0] - 2026-02-12

### Added

- Enabled automatic resolution of linter findings in the repository through the addition of a GitHub Actions workflow for golangci-lint issues that preserves existing behavior and tests while following idiomatic Go patterns. [4d9d33dd](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4d9d33dd6d7df1b194fed5a457dadc311cb2cb9f)

## [7.6.1] - 2026-02-02

### Fixed

- The build process now includes Git information for users relying on this data in their applications. [82d918e2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/82d918e241ce995c5bae41b0153c603d9af8cd09)

## [7.6.0] - 2026-02-02

### Added

- Enabled improved testing of image builds through the addition of a smoke test job to the GitHub Actions workflow that verifies the functionality of Docker images upon code changes. [150f7617](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/150f7617d7955198635303e82ad3ca3616e6f0fa)

## [7.5.6] - 2026-02-02

### Fixed

- Improved the accuracy of linting by resolving false positives from spell checks in go files and disabling them in GitHub Actions workflows. [35cd82aa](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/35cd82aa4d019ac55c42f9319356be21b1562087)

## [7.5.5] - 2026-02-02

### Changed

- Updated GitHub Actions versions in build and linter workflows to 3.7.0 for docker/login-action and 8.4.0 for super-linter/super-linter, potentially requiring manual intervention if hardcoded references exist elsewhere. [14fb8e04](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/14fb8e046a7367ea3a9289b1e3581437b42c12ad)

## [7.5.4] - 2026-02-02

### Fixed

- Resolved the go.mod file to reflect updated dependencies by downgrading certain packages to earlier versions such as v0.29.0 for golang.org/x/mod and v1.77.0 for google.golang.org/gRPC which may require manual migration steps in some cases. [bae3ef04](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bae3ef04f4e7fd05577d59bf4f48ee5e50520be0)

## [7.5.3] - 2026-01-23

### Fixed

- Resolved an issue where GitHub Actions workflows relying on write access to artifact metadata were failing due to missing permissions by updating the workflow's permissions accordingly and requiring users to review their existing workflows for necessary updates. [8123ca44](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8123ca44be22ce96f940cf7bf0676b057f1af681)

## [7.5.2] - 2026-01-23

### Fixed

- The natural language linting in the janitor prompt has been resolved to correct minor issues affecting how unnecessary abstractions are identified by using "overengineering" instead of "overengineering". [d62d911b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d62d911bbc02efd4c1c0ff381744b095e6cdd197)

## [7.5.1] - 2026-01-23

### Fixed

- Artifact metadata permissions have been updated to allow write access, adding 'artifact-metadata: write' to the GitHub Actions workflow configuration, which may require adjustments in workflows that rely on artifact metadata. [b5de370b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b5de370b4ebd76f86518f7136de15246eee70ae7)

## [7.5.0] - 2026-01-23

### Added

- Enabled users to perform cleanup tasks on codebases through the addition of a janitor prompt that includes tools for measuring usage, deleting safely, simplifying incrementally, validating continuously, and documenting nothing without introducing breaking behavior or API changes. [4c5b3963](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4c5b396317f8da5003bfe9b3c6c7ce7d1421982f)

## [7.4.9] - 2026-01-23

### Fixed

- Stabilized build information is now automatically generated and printed to the console at runtime, displaying the commit hash, build date, and Git tree modification status without affecting API or CLI contract. [804b7118](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/804b7118393feb177f0cea0872205530bcb795ec)

## [7.4.8] - 2026-01-23

### Fixed

- Stabilized container execution by modifying the Dockerfile to use /bin/ash instead of /bin/sh for SHELL commands, ensuring compatibility in certain environments and potentially requiring adjustments in scripts that rely on this behavior. [3f1db7fc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3f1db7fc5b517b9ba2556f5f36ccf1cee82a2397)

## [7.4.7] - 2026-01-23

### Changed

- Updated Ginkgo library to version v2.27.5 and corresponding Gomega library to version 1.38.2, requiring potential migration steps for users with customized test suites due to the version bump. [8d238714](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8d2387141e7c7dd68406f0557390d0bf203d01f7)

## [7.4.6] - 2026-01-23

### Fixed

- Resolved the Docker build issue by updating the default shell from "/bin/Bash" to "/bin/sh", which may impact users relying on Bash-specific features during container builds and runs. [ea0f5947](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ea0f5947f5b909fd03226f6bd88d9bf8b0e1d216)

## [7.4.5] - 2026-01-23

### Changed

- Updated the containerized-data-importer-api dependency to version 1.64.0 from 1.63.1 without introducing breaking changes or requiring migration steps in dependent projects. [67437cf8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/67437cf8744775a6d0ff0e7be66136963073cae4)

## [7.4.4] - 2026-01-23

### Fixed

- The minimum required code coverage threshold for Go projects has been hardened to 45% from its previous level of 40%. [2eb75b26](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2eb75b26b676a7f3cf7fac5c690c62dd07bc7d9e)

## [7.4.3] - 2026-01-23

### Fixed

- The GitHub Actions workflow for building the project now persists credentials by default, which may impact users who rely on this behavior and should review their workflows to ensure they handle credential persistence as intended. [61518c3e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/61518c3e1556debc47e6c8535f771c9526ee3a46)

## [7.4.2] - 2026-01-23

### Changed

- Updated versions files to reflect new dependencies, which may require users to review and update their dependency management scripts for integration or testing purposes. [9a0096df](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/9a0096df2706062d98706e2d4e0c927e1c631b52)

## [7.4.1] - 2026-01-23

### Fixed

- The GitHub Actions workflows for Go linter checks were stabilized to provide improved issue reporting when linter checks fail by generating detailed issue descriptions and ignoring certain parts of the analysis. [beba3810](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/beba381000485ef35884aeeffb6b9257004c9a05)

## [7.4.0] - 2026-01-23

### Added

- Enabled developers to generate more comprehensive issue descriptions through improved Markdown templates and analysis output processing without introducing breaking behavior or requiring migration efforts. [239c941a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/239c941a2c71cdbeca8b7bfe17b3892a47046649)

## [7.3.2] - 2026-01-23

### Fixed

- Resolved linting issues by refactoring build-time variables and removing ldflags from the codebase, resulting in no breaking behavior or migration requirements. [25373bbf](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/25373bbf7735800a7a68364fb4320e25b745b68f)

## [7.3.1] - 2026-01-23

### Fixed

- Resolved issues related to Docker build variables by enabling accurate tracking of Git commit information in the built binary through modifications to the `go build` command. [b3865e21](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b3865e21d185cd1e2484c016fd00e7a07e9c21f0)

## [7.3.0] - 2026-01-23

### Added

- Enabled automated analysis of code quality and generation of detailed issue descriptions for reported problems through the addition of a Go linter reporter to the GitHub Actions workflow. [bb37fca7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bb37fca76b130407392b0b921420fa693ef8d8df)

## [7.2.0] - 2026-01-22

### Added

- Enabled automated Go version updates in CI scripts, ensuring consistent and efficient maintenance of Go versions without requiring manual intervention. [2d8d5f87](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2d8d5f874cc9ef4e9a250e968c2a9bf66b2d02f4)

## [7.1.3] - 2026-01-22

### Fixed

- Resolved the setup-go cache poisoning vulnerability by requiring explicit Go versions in workflows instead of relying on the cache and removing the vulnerable cache from all relevant configurations. [94803a9f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/94803a9f080821866abdb64d4d43ca4605ad5dcb)

## [7.1.2] - 2026-01-22

### Changed

- Upgraded the ai-inference action to version 2.0.5 in GitHub workflow linter configuration, requiring no migration steps for users but potentially necessitating adjustments if previously customized for the older version. [33fab97a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/33fab97ad91f65559fbfa4ddbd93a37930fce789)

## [7.1.1] - 2026-01-22

### Changed

- Enabled support for metrics and tracing by incorporating OpenTelemetry packages into the Go dependencies without introducing any breaking behavior, thus requiring developers to adapt their usage of new features but leaving API and CLI contracts unchanged. [f3e71fd8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f3e71fd8eae14dfda78cdf7b5009c5fe60228a92)

## [7.1.0] - 2026-01-22

### Added

- Enabled consistent code coverage results across different environments by explicitly setting the go version to "^1.25" in the GitHub Actions workflow. [197eab85](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/197eab85ef612db8840b8b748c13978f06582c52)

## [7.0.2] - 2026-01-21

### Fixed

- Resolved the coverage action's history depth issue by enabling it to fetch a deeper history of commits up to 10, which affects users relying on this feature and does not introduce any breaking behavior or security concerns. [59a4d8b0](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/59a4d8b062161a03a94284315bf9850851389b74)

## [7.0.1] - 2026-01-21

### Changed

- Upgraded GitHub Actions workflows to rely on newer versions of actions/checkout and other dependencies potentially requiring updates to workflow files. [4e9558bf](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4e9558bf1ccc0c6e92ccd2f66ab2a505a0b35a28)

## [7.0.0] - 2026-01-21

### Removed

- Simplified test efficiency by eliminating duplicated code in test functions through consolidation into the new function `verifyLoggerImpl`. [98211f0b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/98211f0b9df4e3f36409962bdf28c45dcfd5f8d2)

## [6.6.2] - 2026-01-21

### Fixed

- Improved test coverage was resolved through enhanced linter checks and updated test cases that now effectively cover various logging scenarios with different log levels. [b768e6af](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b768e6afbec6c7f7f04094097238250d2dc01dda)

## [6.6.1] - 2026-01-21

### Fixed

- Resolved spelling issues in the GitHub repository wordlist to improve auto-completion functionality for users without introducing any breaking behavior or affecting API or CLI contracts. [1602de91](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1602de91178de2914accbf19d934325d60e654a0)

## [6.6.0] - 2026-01-21

### Added

- Enabled spell checking during the pre-commit phase to enforce better coding standards and reduce the likelihood of typos causing issues, with codespell configured to skip go.sum files to avoid false positives. [d86c6ea9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d86c6ea9d3b0bb3bf7b7217b244bdd1d1ec91db9)

## [6.5.0] - 2026-01-21

### Added

- Enabled distributed tracing capabilities through the addition of telemetry support, allowing developers to track key operations and customize configuration via environment variables for OTLP export type and endpoint. [637b9ddb](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/637b9ddb95d5b8b87b95ad060876207a37eacff0)

## [6.4.0] - 2026-01-21

### Added

- Enabled more flexible and customizable logging capabilities through the introduction of Uber's Zap library, which provides enhanced support for structured logging and formatting along with new configuration options via environment variables like `KAR_LOG_LEVEL`. [4177eb2a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4177eb2a7e16f4080f7429e48e15b6de1922b3ac)

## [6.3.6] - 2026-01-07

### Fixed

- Resolved an issue with commit versions in GitHub Actions workflows by updating the version of the `git-diff-action` to 2.8.1, which may require users who rely on specific commit hashes to update their scripts accordingly. [7db0a59c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7db0a59cbaca423663deec67694c4414ec12f084)

## [6.3.5] - 2026-01-07

### Fixed

- Resolved the GitHub Action workflow for linting to correctly handle diff outputs by switching from a processed diff to raw output. [20da4f25](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/20da4f2528dc0752503b13a14dd2781c6aa12e1d)

## [6.3.4] - 2026-01-07

### Fixed

- Resolved issues with super-linter validation by requiring full Git history in the GitHub Actions workflow to prevent potential breakage from shallow clones. [4be8477a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4be8477af65bc1e099a8ac281c2bc1153d2901d4)

## [6.3.3] - 2026-01-07

### Changed

- Upgraded versions files to use version 8.3.2 of the super-linter for linter validation in GitHub Actions workflows without introducing any breaking behavior and preserving an unchanged API contract. [870fffd8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/870fffd82a733f07bd56f9b9ea46275f0357e134)

## [6.3.2] - 2026-01-07

### Fixed

- Resolved issues with integration tests by updating them to reflect changes in cluster configuration and utilizing the newgrp command for log export. [f45d1af6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f45d1af65f1aa555b1f74379ea24998075d5426c)

## [6.3.1] - 2026-01-06

### Fixed

- The demo script now correctly loads system paths before running the Alpine demo instance without errors related to path resolution. [3e2837c9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3e2837c977d4542052463aca3c8493cd039b4e25)

## [6.3.0] - 2026-01-06

### Added

- Enabled users to run ActionLint on their own infrastructure by introducing a self-hosted runner configuration file that adheres to the Apache License 2.0 and includes labels for identification. [a33efc6f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a33efc6f2744920cf852851707e5a2bd11f50d4c)

## [6.2.1] - 2026-01-06

### Fixed

- Resolved the zizimor linting issue by persisting credentials in the actions/checkout step and updating the sh-checker action, without introducing breaking behavior or requiring migration steps. [3179385a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3179385a0dede46482b6320438e67cb952bcbd21)

## [6.2.0] - 2026-01-06

### Added

- Automated testing of the codebase is now enabled for maintainers through GitHub workflows triggered on push and pull requests. [abe0378d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/abe0378d55cde63f0250f8b99856745991b6c803)

## [6.1.1] - 2025-12-19

### Changed

- Upgraded several Kubernetes packages to their latest versions, including k8s.io/API v0.22.0 and others, which are expected to include bugfixes and new features without introducing any breaking changes or requiring migration efforts. [2eb806cf](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2eb806cf6c7e99cfb300a5801449f05111423be3)

## [6.1.0] - 2025-12-19

### Added

- Updated GitHub Actions to ignore commit hash updates for actions/attest-build-provenance and corrected the reviewdog/action-misspell commit hash in continuous integration workflows with no breaking behavior introduced. [d08c752d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d08c752d2217b116933008eb7e8069efd6476f23)

## [6.0.3] - 2025-12-19

### Changed

- Upgraded pre-commit configuration to require the latest version of the ai-prepare-commit-msg repository for proper commit message formatting. [de736cf5](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/de736cf5c2aea260ddf5b38801bc61d231f8bd60)

## [6.0.2] - 2025-12-19

### Changed

- Updated dependencies in several workflows to ensure compatibility and stability. [68b98274](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/68b982748be6282b55e13f42d203a4663e125cf6)

## [6.0.1] - 2025-12-19

### Changed

- Updated dependencies for GitHub Actions workflows by upgrading the version of actions/checkout from 6.0.0 to 6.0.1 in multiple workflows. [41ca18b4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/41ca18b410287b74d509020524e34ecdd8910c5d)

## [6.0.0] - 2025-11-29

### Removed

- Optimized build times by disabling pre-commit validation in the super-linter workflow without introducing any breaking behavior and maintaining existing validations for code quality. [eb595e97](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/eb595e97a221743ba3a81192149e42fa6dd45bc5)

## [5.7.1] - 2025-11-29

### Changed

- Enabled successful linter validation on Ubuntu platforms by installing OpenSSL libraries prior to running the linter in the CI workflow. [c4ba49c8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/c4ba49c8868583f346d7663e191387dda439a308)

## [5.7.0] - 2025-11-28

### Added

- Enabled the option to disable charset checking in .editorconfig-checker.json configurations for users who no longer rely on this feature to enforce character encoding standards without introducing any breaking behavior or requiring migration steps. [e0e658f1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e0e658f19bbab60a5e6d9cfab4682f79e05c9b20)

## [5.6.1] - 2025-11-28

### Changed

- Updated version dependencies in GitHub Actions workflows to 6.x for several packages and pre-commit configuration to include new repositories and hooks. [d97c90c4](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d97c90c4fc3d1de3a025adc9932203eddef9cdd5)

## [5.6.0] - 2025-11-27

### Added

- Enabled pre-commit and prepare-commit-msg hooks for the AI package to run checks during development automatically. [82c01dae](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/82c01dae6c3279ab45011594bb9aa1de830e86c3)

## [5.5.1] - 2025-10-26

### Changed

- Optimized linter analysis accuracy by switching to the Ministral-3B model from mistral-medium-2505 without introducing any breaking behavior and preserving an unchanged configuration schema. [8268c245](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8268c245cbe59384567baac60c65830a263c48d8)

## [5.5.0] - 2025-10-26

### Added

- Enabled Go linting validation to be skipped in GitHub Actions workflow by setting the `VALIDATE_GO` flag to false and updating the Makefile's lint target accordingly. [cc1e99cc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/cc1e99cc60590d846fe162935db35e76dca0c59f)

## [5.4.8] - 2025-10-26

### Changed

- Enabled automated analysis of linter failures through AI-powered diagnosis and proposed fixes in GitHub workflows. [be2469be](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/be2469be09612f45982376d93e47aac6de811ce8)

## [5.4.7] - 2025-10-23

### Changed

- Resolved the dependency version mismatch in the reviewdog/action-misspell action to ensure uninterrupted spell checking functionality. [6789b5ec](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6789b5ecbf6f0ddd26113a91d71ff96f889dfe5e)

## [5.4.6] - 2025-10-23

### Changed

- Updated GitHub workflows to use newer versions of actions including markdown-link-check at version 1.1.1 and misspell at version 1.27.0 with no breaking behavior introduced. [bccb011e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/bccb011e1e731341ff8b1808bc4e7da9de1d04a0)

## [5.4.5] - 2025-10-16

### Changed

- Upgraded super-linter version to 8.2.1 in GitHub workflows, providing users with the latest security fixes and improvements without introducing any breaking behavior or requiring migration steps. [6636e9d9](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6636e9d9a4235635814e4d95e82cd30a9ec7c681)

## [5.4.4] - 2025-10-16

### Changed

- Updated versions files to ensure users have access to the latest features and bugfixes from Docker login-action version 3.6.0 and super-linter version 8.2.0, with no migration steps required but potentially impacting builds that rely on these actions. [6589a248](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6589a2489048d73f300f511749194b364c517f73)

## [5.4.3] - 2025-09-25

### Changed

- Updated the Docker build environment to utilize the newer Git package version 2.49.1-r0, potentially necessitating users who have pinned dependencies to review and update their configurations accordingly. [d539a614](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d539a614051d24f106aee9336d8dd2dcf152908c)

## [5.4.2] - 2025-09-25

### Changed

- Optimized Docker CI job permissions to allow creating pull requests while maintaining security by removing the `contents: write` permission and adding the `pull-requests: write` permission. [4373a749](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4373a7491b0c683ab9b104167ef74f6f65b7c066)

## [5.4.1] - 2025-09-25

### Changed

- Upgraded Go setup and attest-build-provenance actions to versions 6.0.0 and 3.0.0 respectively, potentially requiring migration steps for developers building with these workflows. [08d737b2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/08d737b265f8be9fa98d6e19e811281e782726c7)

## [5.4.0] - 2025-09-25

### Added

- Enabled notifications for failed jobs via issue creation action, which requires write permissions to issues and runs on Ubuntu Linux. [7c9fc228](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7c9fc2283ce0ab1a2c00d3931e24f5f220278b17)

## [5.3.1] - 2025-09-25

### Changed

- Enabled GitHub users who manage workflows to persist credentials in workflow jobs by default, requiring adjustments for existing PATs and workflows with specific scopes. [08968a3a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/08968a3a13b13a7cda85e6f7fc43202f9785cf86)

## [5.3.0] - 2025-09-21

### Added

- Enabled developers to set up a development environment without manual configuration through the addition of devcontainer support. [783eee16](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/783eee16aa533d5efb93a5e5108051f6631c0f76)

## [5.2.0] - 2025-09-21

### Added

- Enabled pre-commit support for automatic execution of code quality checks before commits are made including hooks for checking trailing whitespace YAML syntax shell scripts Markdown formatting and YAML formatting with configuration specified in the `.pre-config.yaml` file. [d2d4df16](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d2d4df16d9eb3a6c6c79997dd285464b593ecf31)

## [5.1.6] - 2025-08-27

### Changed

- Resolved compatibility issues identified by Trivy through updated Go dependencies and versions. [99e90559](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/99e90559e12681fbb26b5a09308931908f4ab8a8)

## [5.1.5] - 2025-08-24

### Changed

- Updated GitHub workflows to leverage newer versions of actions and tools, including the checkout action, enabling compatibility with the latest GitHub features without introducing any breaking behavior. [01b3f61b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/01b3f61b6fb11b48bf1ed64872ca9ec0d4d6e9c6)

## [5.1.4] - 2025-08-24

### Changed

- The CI update script has been optimized to automatically install Go if it's missing and update GitHub Action commit hashes, ensuring consistent code formatting through the addition of a `make fmt` command on exit. [db5bff4d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/db5bff4d5bf5082d87c06a638e4ae50d1418b80f)

## [5.1.3] - 2025-08-18

### Changed

- Updated the CI job to correctly retrieve and utilize the current Go version from the official site, resolving an issue that could cause build failures due to outdated versions. [4617acff](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4617acff424704f404746fcfe5913dbe7d3419fc)

## [5.1.2] - 2025-08-18

### Changed

- Updated the editor configuration to use spaces for indentation in Bash files, changing the default behavior of shfmt and affecting the Makefile's fmt target and ci/update_versions.sh script. [7d533f3c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7d533f3c22e69a71d33e3d49d89ca0c83184a243)

## [5.1.1] - 2025-08-07

### Changed

- Upgraded several actions in GitHub workflows to their latest versions, including the actions/cache action from 4.2.3 to 4.2.4, docker/login-action from 3.4.0 to 3.5.0, and docker/metadata-action from 5.7.0 to 5.8.0, ensuring continued correct function of the workflows with the latest versions. [05ae6e94](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/05ae6e94ae3cbeeb3eb9a17dc5cfe8c1522b24ee)

## [5.1.0] - 2025-08-07

### Added

- Enabled users to specify a default branch for linter runs in GitHub workflows without modifying existing configurations. [6a520d1f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6a520d1f0cb34df425c4d0349ca6a39c4a5aac84)

## [5.0.4] - 2025-07-18

### Changed

- Updated dependencies in GitHub Actions workflows to super-linter 8.0.0 and pyspelling-any 1.0.5, requiring manual review of existing spell check results due to potential changes in detection rules or behavior. [5ee77a6d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5ee77a6d975ba7e66b436468975066e1add81c1e)

## [5.0.3] - 2025-07-18

### Changed

- Upgraded versions files to ensure compatibility with latest super-linter action configurations. [59536498](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/59536498a1c90b81bb7fbd8aad03ef17ab92280e)

## [5.0.2] - 2025-07-04

### Changed

- The pyspelling-any action has been optimized by upgrading its version from 1.0.4 to 1.0.5 without introducing any breaking behavior and preserving the existing functionality of the spell checking workflow. [2739d1a6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2739d1a6c31d6a4302c6b6cf864ecedcb53ce4d3)

## [5.0.1] - 2025-06-27

### Changed

- Enabled robust monitoring of Virtual Machine Instance status changes through improved watch resources functionality that now correctly handles various VMI phases without introducing breaking behavior or API/CLI contract changes and requires migration to the `github.com/matryer/resync` package for synchronization. [08ef6fa6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/08ef6fa60abab64a76c01860727e6f258479aa61)

## [5.0.0] - 2025-06-27

### Removed

- Eliminated reliance on obsolete current status variable, allowing users to directly access VM instance phase through virtClient without modification. [9e09a462](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/9e09a462c7ff0372478e0dacfffd6c6f2b101e74)

## [4.0.5] - 2025-06-27

### Changed

- Enabled developers to work with application context by introducing an AppContext struct that stores Virtual Machine Instance and Data Volume names. [82bcdde1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/82bcdde13f4bd89127d718a0f38ffc7730ec6e83)

## [4.0.4] - 2025-06-27

### Changed

- Reduced the default cleanup timeout to 5 minutes, which may require users to update their scripts or configurations to accommodate the new value and potentially break existing workflows relying on the longer timeout. [b55b780e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b55b780ec51322bb9e310d2f658a344fc775b30d)

## [4.0.3] - 2025-06-27

### Changed

- Updated build information is now available in the Docker image due to modifications that include changing the working directory and adding Git for version control without affecting the API or CLI contract. [2b296ea6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2b296ea68358509d52850d7dc433d71604f61e56)

## [4.0.2] - 2025-06-27

### Changed

- Updated the `kubevirt` client-go dependencies to new versions, including `k8s.io/kube-openapi v0.0.0-20250610211856-8b98d1ed966a`, `kubevirt.io/client-go v1.4.1`, and `kubevirt.io/containerized-data-importer-api v1.58.0`. [b5a272a8](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b5a272a8ffa2e864317fe2525b4f3e3003c8bb45)

## [4.0.1] - 2025-06-27

### Changed

- Corrected log messages to accurately reflect successful resource creation, improving debugging clarity for users who run the Virtual Machine runner without affecting API or CLI contracts. [c289626f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/c289626ff1b073248d5fc1a1c3499c433a4f4aa3)

## [4.0.0] - 2025-06-27

### Removed

- Eliminated periodic status updates for virtual machine instances, reducing logging output and requiring users to remove any remaining references to reporting functionality. [f9d61fdc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f9d61fdce8b904bafbfa50aea820ecfd19ea6adc)

## [3.4.4] - 2025-06-27

### Changed

- Enabled write access for creating pull requests and pushing code in the workflow's permissions section, allowing the technote-space/create-pr-action to perform these actions without issue. [7a068441](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/7a068441fd3ffaadd290d48b925d422506f3d561)

## [3.4.3] - 2025-06-26

### Changed

- Optimized cleanup operations by allowing users to set custom timeouts via environment variable KAR_CLEANUP_TIMEOUT and ensuring long-running operations can be terminated within a specified time frame preventing resource accumulation issues. [8dae66f1](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8dae66f17bd3c43e9caf343c61339a3b72abff44)

## [3.4.2] - 2025-06-26

### Changed

- Updated the linter configuration in `.golangci.yml` to require explicit rule settings and package-specific `gomoddirectives`, necessitating migration steps for customized configurations. [2383f2e7](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/2383f2e7701b0a8729377513652d1e3599d4f6d7)

## [3.4.1] - 2025-06-26

### Changed

- The cleanup process now continues even if it encounters non-existent runner instances and data volumes, allowing for smoother deletion operations without interruption. [1a8e24b6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1a8e24b6f9cb211e2bc60a152eac6dbbdfa4290d)

## [3.4.0] - 2025-06-26

### Added

- Enabled log messages at key points in the Virtual Machine runner workflow, providing visibility into successful resource creation, completion, and deletion without introducing any breaking behavior or migration requirements. [ecf9573c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ecf9573ce29fe666c9d9d18298fd473e93e1d178)

## [3.3.1] - 2025-06-24

### Changed

- The documentation for the `kubevirt-actions-runner` has been updated to better reflect its capabilities and usage. [73162ddb](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/73162ddb3679559f5318a20b9792d327529d57b7)

## [3.3.0] - 2025-06-24

### Added

- Enabled users to ignore missing schemas in Kubernetes linter configuration through an updated option in the GitHub Actions workflow file. [3f923521](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/3f9235215110b301105107945f9dacf3bb056f93)

## [3.2.1] - 2025-06-24

### Changed

- Updated the Docker image's base to Alpine 3.22, requiring users who have customized their images to rebuild them due to the change in the underlying OS version. [1c9bd97c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1c9bd97ca445825fe0780ff678ad2c3794682602)

## [3.2.0] - 2025-06-24

### Added

- Enabled relaxed security checks for container images by allowing users to opt out of the HEALTHCHECK checkov requirement during image build. [899b9719](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/899b9719c809eec22c79aa1dfb4bf1d8c22c1177)

## [3.1.1] - 2025-06-24

### Changed

- Upgraded GitHub Action versions to improve performance and security by leveraging the latest features in setup-go, cache, upload-artifact, docker/login-action, metadata-action, build-push-action, attest-build-provenance, markdown-link-check, super-linter, and misspell actions. [e5af489c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e5af489c754ce9150f9f143aa322ad9700a9ee85)

## [3.1.0] - 2025-04-09

### Added

- Enabled the project to bypass go linter checks by disabling it in the configuration, impacting users who relied on it for coding standards enforcement and requiring no migration steps due to unchanged API contracts. [8990975d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8990975d60f45959dd4fcfa10ae16362a586aa49)

## [3.0.4] - 2025-04-09

### Changed

- Updated the GitHub Actions workflow to utilize the super-linter/super-linter repository, ensuring uninterrupted linter feedback and no breaking behavior or API contract changes. [ac7cfeb6](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/ac7cfeb6501ede91e2621a6d6168a549936d5f69)

## [3.0.3] - 2025-02-14

### Changed

- Upgraded the Go version to 1.24.0, which affects various workflows and configurations including Docker images, go.mod files, and GitHub Actions setup scripts for build and update workflows with no breaking behavior or migration requirements. [1cb578cc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/1cb578cc1ea6026e6e9d3f9dea4926b6e9cbfafb)

## [3.0.2] - 2025-02-12

### Changed

- The linter configuration was optimized to address deprecated rules and unused settings, requiring users to review their project configurations for potential issues. [45950e4a](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/45950e4aff5278292376cd90142f1d87929b7a57)

## [3.0.1] - 2025-02-07

### Changed

- Runner methods now return explicit error types directly instead of using deferred functions and boolean flags to indicate failure. [d7477d74](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/d7477d740dfe95d36f97c4101667ed6602eed4e1)

## [3.0.0] - 2025-02-04

### Removed

- The data volume name parameter is now required for successful deletion of resources by the DeleteResources method. [cdb9f5b3](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/cdb9f5b3ed6d9be034c162fa5db7ec31707ef1dd)

## [2.4.1] - 2025-02-04

### Changed

- Enabled code coverage checks for Go changes via GitHub Actions with a threshold of 40%, requiring unit-test job dependencies to meet this new requirement. [b6d9ca9f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/b6d9ca9ffa129df18f92dbd08d648f200e587286)

## [2.4.0] - 2025-02-04

### Added

- Enabled secure authentication for GitHub API interactions by introducing a WORKFLOW_TOKEN secret token that replaces the need for personal access tokens with fine-grained permissions, ensuring sensitive credentials are never exposed in plain text and requiring no migration steps from users. [40abec1b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/40abec1b5cefadc27fdf97c8476aa9967b645a5c)

## [2.3.3] - 2025-02-03

### Changed

- Improved validation for the create resources use case by introducing unit tests that verify correct behavior when providing empty values for VM template, runner name, and JIT configuration. [e4c33b8d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/e4c33b8d4b06d8d2f33da2882d91840ba171b5bf)

## [2.3.2] - 2025-01-31

### Changed

- Updated the minimum supported Go version in workflows from 1.20.0 to 1.23 following an upgrade of the `actions/setup-go` action from 5.2.0 to 5.3.0. [54039b3e](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/54039b3e12294e55ea51036ba6f5ed67f25deab1)

## [2.3.1] - 2025-01-31

### Changed

- Enabled initial unit tests for the Kubevirt Actions Runner to cover test cases for the root command's initialization process without introducing any breaking behavior or migration requirements. [01742283](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/0174228332f9223279ae9954c9073102e1a88978)

## [2.3.0] - 2025-01-23

### Added

- Enabled logging of build information during startup, including commit hash, tree modification status, build date, and Go version. [5adf666b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/5adf666ba433c054e448e17270838cbf5a2d07dc)

## [2.2.2] - 2025-01-23

### Changed

- Improved handling of VMI failures now enables the runner to accurately report job results by properly indicating and responding to failed VMIs without introducing breaking behavior or modifying API or CLI contracts. [18208ce2](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/18208ce291ef079420a4f8bb534958841851a8ee)

## [2.2.1] - 2025-01-23

### Changed

- Optimized logging intervals to reduce log message frequency by outputting every 5 minutes instead of the default rate. [8f5d519b](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/8f5d519b5bba5b6018a25a6a909d15f22f5f3e57)

## [2.2.0] - 2025-01-23

### Added

- Enabled automatic updates to the project's container image by incorporating an external action that updates the Dockerfile as part of the PR process. [231e2a64](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/231e2a64dcdd209ebb1f9511435b58018ff4d11e)

## [2.1.2] - 2025-01-23

### Changed

- The super-linter image reference has been updated to ghcr.io from github.com requiring users to re-pull the image if they have previously pulled from the old repository. [edbf223c](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/edbf223c93a3e06b68032cba69f8dc8881239489)

## [2.1.1] - 2025-01-08

### Changed

- Modernized GH action versions to 4.2.2 and 1.26.1 for actions/checkout and reviewdog/action-misspell respectively, with setup-go updated to 5.2.0, ensuring correct dependency installation for Go version 1.20.0 and above. [484c6225](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/484c622500c0e48a9e113a9c3468b373536a850b)

## [2.1.0] - 2025-01-08

### Added

- The update version script now ensures the Go module is in a consistent state by running `go mod tidy`, which affects users who rely on automated updates and maintainers managing project dependencies without introducing any breaking behavior or API changes. [6e24be6f](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6e24be6f8ab5c5d9d4c40e75ee773de42fffbec2)

## [2.0.0] - 2025-01-08

### Removed

- Simplified the GoLangCI configuration format to remove unnecessary lines ensuring a more streamlined setup for users without introducing any breaking behavior or requiring migration. [4f0f484d](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/4f0f484d78b60f5ed26ca532a53736ce6002971c)

## [1.1.0] - 2025-01-08

### Added

- Automated Docker image creation and publication are now streamlined through the addition of GitHub Actions for building and pushing Docker images. [6ac394bc](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/6ac394bc31f56208e573db40c4fc7db4a904db5b)

## [1.0.2] - 2024-12-19

### Changed

- Enabled automatic deletion of associated DataVolumes by including an OwnerReference to the VirtualMachineInstance during DataVolume creation. [a3df2c68](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/a3df2c688cec490f8615ac5c882b25defe375ec4)

## [1.0.1] - 2024-12-19

### Changed

- The NewRootCommand function now requires users to pass a Runner instance instead of command options, affecting API usage and potentially necessitating migration steps for existing scripts. [350bb8ac](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/350bb8acd40d839229bf366e5cee03fd53e24b90)

## [1.0.0] - 2024-12-18

### Added

- Introduced support for spawning ephemeral virtual machines for jobs using KubeVirt in the Kubernetes-based runner image. [f571e6da](https://github.com/electrocucaracha/kubevirt-actions-runner/commit/f571e6da5f69cb100d5dc85ef41672814bb51279)
