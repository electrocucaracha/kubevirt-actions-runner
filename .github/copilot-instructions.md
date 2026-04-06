---
description: "Custom instructions for GitHub Copilot to ensure code quality and documentation standards in this repository."
applyTo: "**"
---

# GitHub Copilot Custom Instructions

This repository uses GitHub Copilot with the following custom instructions to ensure high-quality contributions and maintain consistency across the codebase.

## General Guidelines

1. **Testing and Validation:**
   - Every change must pass the following commands before being considered complete:
     - `make test`: Run the test suite to ensure all tests pass.
     - `make lint`: Check for code style and linting issues.
     - `make fmt`: Format the code to adhere to the repository's style guidelines.

1. **Documentation Standards:**
   - All documentation updates must follow the [Diátaxis Framework](https://diataxis.fr/), ensuring clarity, accuracy, and user-centricity.
   - Use [Google's Documentation Best Practices](https://developers.google.com/style/) for tone, terminology, and structure.
   - Write documentation with **semantic line breaks** to improve readability and version control. Refer to [Semantic Line Breaks](https://sembr.org/) for guidance.

## Workflow for Contributions

1. **Code Changes:**
   - Ensure all code changes are clean, maintainable, and follow the repository's coding standards.
   - Avoid introducing unnecessary complexity or breaking existing functionality.

1. **Testing:**
   - Run `make test` to verify that all tests pass.
   - If new functionality is added, include corresponding tests to ensure adequate coverage.

1. **Linting and Formatting:**
   - Run `make lint` to check for linting issues.
   - Run `make fmt` to format the code according to the repository's style.

1. **Documentation Updates:**
   - Identify the type of documentation required (Tutorial, How-to Guide, Reference, or Explanation) based on the [Diátaxis Framework](https://diataxis.fr/).
   - Write clear, concise, and user-focused documentation.
   - Use semantic line breaks to improve readability and version control.
   - Follow [Google's Style Guide](https://developers.google.com/style/) for tone and structure.

## Review Process

1. **Code Review:**
   - All pull requests will be reviewed for adherence to the repository's coding standards and guidelines.
   - Reviewers will ensure that `make test`, `make lint`, and `make fmt` have been run successfully.

1. **Documentation Review:**
   - Documentation will be reviewed for clarity, accuracy, and adherence to the Diátaxis Framework and Google's Style Guide.
   - Semantic line breaks will be checked for consistency.

## References

- [Diátaxis Framework](https://diataxis.fr/)
- [Google Documentation Style Guide](https://developers.google.com/style/)
- [Semantic Line Breaks](https://sembr.org/)
