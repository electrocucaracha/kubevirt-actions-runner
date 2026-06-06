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

had_errors=false

format_changes() {
    local status=$?
    if ! make fmt; then
        echo "WARNING: formatting failed; keeping generated changes in place" >&2
    fi
    exit "${status}"
}

run_best_effort() {
    local description=$1
    shift

    if ! "$@"; then
        echo "WARNING: ${description} failed; keeping generated changes in place" >&2
        had_errors=true
    fi
}

update_github_action_hashes() {
    local gh_actions action is_pinned pinned commit_hash

    gh_actions=$(grep -r "uses: [A-Za-z0-9_.-]*/[\_a-z\-]*@" .github/ | sed 's/@.*//' | awk -F ': ' '{ print $3 }' | sort -u)
    exceptions=('reviewdog/action-misspell' 'actions/attest-build-provenance' 'GrantBirki/git-diff-action' 'golangci/golangci-lint-action' 'actions/checkout')
    # Actions pinned to a specific version and excluded from auto-updates.
    # Remove an entry only once the underlying issue is confirmed resolved.
    # austenstone/copilot-cli: v3.0+ depends on actions/setup-copilot@v0 which does
    # not yet exist publicly; keep at v2.0 until that action is released.
    readonly pinned_actions=('austenstone/copilot-cli')
    for action in $gh_actions; do
        is_pinned=false
        for pinned in "${pinned_actions[@]}"; do
            if [[ $action == "$pinned" ]]; then
                is_pinned=true
                break
            fi
        done
        if [[ $is_pinned == "true" ]]; then
            echo "Skipping auto-update for pinned action: $action"
            continue
        fi
        if [[ ${exceptions[*]} =~ (^|[^[:alpha:]])$action([^[:alpha:]]|$) ]]; then
            commit_hash=$(git ls-remote "https://github.com/$action" | grep 'refs/tags/[v]\?[0-9][0-9\.]*\^{}$' | sed 's|refs/tags/[vV]\?[\.]\?||g; s|\^{}$||g' | sort -u -k2 -V | tail -1 | awk '{ printf "%s # %s\n",$1,$2 }')
        else
            commit_hash=$(git ls-remote "https://github.com/$action" | grep 'refs/tags/[v]\?[0-9][0-9\.]*$' | sed 's|refs/tags/[vV]\?[\.]\?||g' | sort -u -k2 -V | tail -1 | awk '{ printf "%s # %s\n",$1,$2 }')
        fi
        if [[ -z $commit_hash ]]; then
            echo "WARNING: unable to resolve a tag for $action; skipping update" >&2
            had_errors=true
            continue
        fi
        # shellcheck disable=SC2267
        grep -ElRZ "uses: $action@" .github/ | xargs -0 -l sed -i -e "s|uses: $action@.*|uses: $action@$commit_hash|g"
    done
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

trap format_changes EXIT

if ! command -v go >/dev/null; then
    curl -fsSL http://bit.ly/install_pkg | PKG=go-lang bash
    # shellcheck disable=SC1091
    source /etc/profile.d/path.sh
fi

full_go_version="$(curl -sL https://golang.org/VERSION?m=text | head -n1 | sed 's/^go//')"
go_version="$(printf '%s' "$full_go_version" | awk -F. '{print $1"."$2}')"
if [[ ! $go_version =~ ^[0-9]+\.[0-9]+$ ]]; then
    echo "ERROR: could not extract a valid Go major.minor version (got: '${go_version}')" >&2
    exit 1
fi
# Get direct modules from go.mod and upgrade them
mapfile -t direct_modules < <(
    go list -m -f '{{if not .Indirect}}{{.Path}}{{end}}' all |
        grep -v '^$' |
        grep -v "$(go list -m)"
)

if ((${#direct_modules[@]})); then
    run_best_effort "updating direct Go modules" go get -u "${direct_modules[@]}"
fi

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
    had_errors=true
fi
run_best_effort "updating GitHub Action commit hashes" update_github_action_hashes
run_best_effort "updating the Dockerfile base image" update_dockerfile_base_image

if [[ $had_errors == "true" ]]; then
    if git diff --quiet; then
        echo "ERROR: update steps failed and no file changes were produced" >&2
        exit 1
    fi
    echo "WARNING: some update steps failed, but generated changes were kept" >&2
fi
