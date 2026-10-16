#!/bin/bash
# SPDX-license-identifier: Apache-2.0
##############################################################################
# Copyright (c) 2024
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Apache License, Version 2.0
# which accompanies this distribution, and is available at
# http://www.apache.org/licenses/LICENSE-2.0
##############################################################################

set -o errexit
set -o pipefail
if [[ ${DEBUG:-false} == "true" ]]; then
    set -o xtrace
fi

has_errors=false
failed_steps=()

format_changes() {
    local status=$?
    # Formatting is best-effort here so earlier successful updates are not discarded.
    if ! make fmt; then
        echo "WARNING: formatting failed; keeping generated changes in place" >&2
    fi
    exit "${status}"
}

run_best_effort() {
    local description=$1
    shift

    # Record non-critical failures so successful updates can still be committed.
    if ! "$@"; then
        echo "WARNING: ${description} failed; keeping generated changes in place" >&2
        has_errors=true
        failed_steps+=("${description}")
    fi
}

update_github_action_hashes() {
    uv run --no-project --with git+https://github.com/electrocucaracha/gh-updater.git gh-updater "$PWD"
}

latest_semver_tag() {
    local repository=$1

    git ls-remote --tags "https://github.com/${repository}" |
        awk '$2 !~ /\^\{\}$/ {
            tag = $2
            sub(/^refs\/tags\//, "", tag)
            if (tag ~ /^v[0-9]+(\.[0-9]+)*$/) {
                sortkey = tag
                sub(/^v/, "", sortkey)
                print sortkey "\t" tag
            }
        }' |
        sort -V |
        tail -1 |
        awk -F'\t' '{ print $2 }'
}

update_action_version() {
    local repository=$1
    local action_marker=$2
    local version file

    version=$(latest_semver_tag "$repository")

    if [[ -z $version ]]; then
        echo "WARNING: unable to resolve latest ${repository} version; skipping" >&2
        return 1
    fi

    while IFS= read -r -d '' file; do
        sed -i "s|^\([[:space:]]*version:[[:space:]]*\)v[0-9][^ ]*|\1${version}|" "$file"
    done < <(grep -ElRZ "$action_marker" .github/workflows/)
}

update_rtk_version() {
    local version file

    version=$(latest_semver_tag "rtk-ai/rtk")

    if [[ -z $version ]]; then
        echo "WARNING: unable to resolve latest rtk version; skipping" >&2
        return 1
    fi

    while IFS= read -r -d '' file; do
        sed -i "s|rtk-ai/rtk/refs/tags/v[0-9][^/]*|rtk-ai/rtk/refs/tags/${version}|g" "$file"
    done < <(grep -ElRZ "rtk-ai/rtk/refs/tags/" .github/workflows/)
}

update_dockerfile_base_image() {
    local go_docker_tag

    # The || true is intentional: a failed API call or missing jq leaves go_docker_tag empty
    # and the subsequent if-check skips the update without aborting the script.
    go_docker_tag=$(curl -sL "https://hub.docker.com/v2/repositories/library/golang/tags?page_size=100&name=${go_version}-alpine" |
        jq -r '[.results[].name | select(test("^[0-9]+\\.[0-9]+-alpine[0-9]+\\.[0-9]+$"))] | sort | last // empty' 2>/dev/null || true)
    if [[ -n $go_docker_tag ]]; then
        sed -i "s|^FROM golang:[^[:space:]]*[[:space:]]AS build|FROM golang:${go_docker_tag} AS build|" Dockerfile
    fi
}

update_agents() {
    local repo="https://github.com/electrocucaracha/gh-workflows/tree/main/.github/agents"
    local raw="https://raw.githubusercontent.com/electrocucaracha/gh-workflows/main/.github/agents"

    mkdir -p "$DEST"

    curl -fsSL "$repo" |
        grep -oE 'href="/electrocucaracha/gh-workflows/blob/main/\.github/agents/[^"]+"' |
        cut -d '"' -f 2 |
        sed 's|.*/||' |
        sort -u |
        while IFS= read -r file; do
            curl -fsSL "$raw/$file" -o ".github/agents/$file"
            echo "Updated: $file"
        done
}

# Requires Go and jq. Resolve the KubeVirt release before updating other modules.
update_go_modules() {
    local metadata version gomod dependencies replacements modules module replacement
    local targets=()

    metadata=$(go mod download -json kubevirt.io/client-go@latest) || return 1
    version=$(jq -er '.Version | select(startswith("v"))' <<<"$metadata") || return 1
    gomod=$(jq -er '.GoMod | select(length > 0)' <<<"$metadata") || return 1
    metadata=$(go mod edit -json "$gomod") || return 1

    # Published modules, unlike the monorepo, contain usable client-go requirements.
    # Keep KubeVirt and its Kubernetes dependencies together; -u would override them.
    dependencies=$(jq -er '
        [.Require[]? | select(.Path | startswith("k8s.io/") or startswith("kubevirt.io/"))]
        | if (map(.Path) | index("k8s.io/api")) != null
            and (map(.Path) | index("k8s.io/apimachinery")) != null
            and (map(.Path) | index("k8s.io/client-go")) != null
        then .[]
            | select(.Path != "kubevirt.io/api" and .Path != "kubevirt.io/client-go")
            | "\(.Path)@\(.Version)"
        else error("KubeVirt client module is missing Kubernetes requirements")
        end
    ' <<<"$metadata") || return 1
    replacements=$(jq -er '
        [.Replace[]? | select(.Old.Path | startswith("k8s.io/"))]
        | map(if .New.Version == null or .New.Version == ""
            then error("KubeVirt Kubernetes replacement has no module version")
            else "\(.Old.Path)=\(.New.Path)@\(.New.Version)" end)
        | join("\n")
    ' <<<"$metadata") || return 1
    modules=$(go list -m -f '{{if and (not .Main) (not .Indirect)}}{{.Path}}{{end}}' all) || return 1

    targets=("kubevirt.io/api@${version}" "kubevirt.io/client-go@${version}")
    while IFS= read -r module; do
        [[ -z $module ]] || targets+=("$module")
    done <<<"$dependencies"
    while IFS= read -r module; do
        case "$module" in
        "" | k8s.io/* | kubevirt.io/*) continue ;;
        esac
        targets+=("${module}@latest")
    done <<<"$modules"
    while IFS= read -r replacement; do
        [[ -z $replacement ]] || go mod edit "-replace=${replacement}" || return 1
    done <<<"$replacements"

    go get "${targets[@]}"
}

# Allow the module update to be tested without running unrelated maintenance.
if [[ ${BASH_SOURCE[0]} != "$0" ]]; then
    return 0
fi

trap format_changes EXIT

if ! command -v go >/dev/null; then
    curl -fsSL http://bit.ly/install_pkg | PKG=go-lang bash
    # shellcheck disable=SC1091 # file is generated at install time by install_pkg; not present for static analysis.
    source /etc/profile.d/path.sh
fi

full_go_version="$(curl -sL https://golang.org/VERSION?m=text | head -n1 | sed 's/^go//')"
go_version="$(printf '%s' "$full_go_version" | awk -F. '{print $1"."$2}')"
if [[ ! $go_version =~ ^[0-9]+\.[0-9]+$ ]]; then
    echo "ERROR: could not extract a valid Go major.minor version (got: '${go_version}')" >&2
    exit 1
fi
run_best_effort "updating Go modules with KubeVirt constraints" update_go_modules

echo "==> Tidying modules"
run_best_effort "tidying Go modules" go mod tidy -go="${go_version}"

# Exclude update.yml so its go-version stays "stable" (always installs the latest Go toolchain)
find .github/workflows -type f \( -name '*.yml' -o -name '*.yaml' \) ! -name 'update.yml' \
    -exec grep -l 'go-version:' {} + \
    -exec env go_version="${go_version}" bash -s {} + <<'EOF'
    for file; do
        sed -i \
            "s|^\([[:space:]]*go-version:[[:space:]]*\).*|\
\1\"^${go_version}\"|" \
            "${file}"
    done
EOF

if ! command -v uvx >/dev/null; then
    run_best_effort "installing uv" bash -c 'curl -LsSf https://astral.sh/uv/install.sh | sh'
fi
if command -v uvx >/dev/null; then
    run_best_effort "updating pre-commit hooks" uvx pre-commit autoupdate
else
    echo "WARNING: uvx is unavailable; skipping pre-commit updates" >&2
    has_errors=true
    failed_steps+=("updating pre-commit hooks")
fi
run_best_effort "updating GitHub Action commit hashes" update_github_action_hashes
run_best_effort "updating the Dockerfile base image" update_dockerfile_base_image
run_best_effort "updating golangci-lint version" update_action_version "golangci/golangci-lint" "golangci/golangci-lint-action"
run_best_effort "updating gremlins version" update_action_version "go-gremlins/gremlins" "go-gremlins/gremlins-action"
run_best_effort "updating rtk version" update_rtk_version
run_best_effort "updating agents definition" update_agents

if [[ $has_errors == "true" ]]; then
    failed_summary=$(
        IFS=', '
        echo "${failed_steps[*]}"
    )
    if git diff --quiet; then
        echo "ERROR: update steps failed (${failed_summary}) and no file changes were produced" >&2
        exit 1
    fi
    echo "WARNING: some update steps failed (${failed_summary}), but generated changes were kept" >&2
fi
