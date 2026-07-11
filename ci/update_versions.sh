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
    local gh_actions action is_pinned pinned commit_hash file

    gh_actions=$(grep -rhoE 'uses: [^@]+@' .github |
        sed -E 's/uses: ([^@]+)@/\1/' |
        sort -u)

    readonly exceptions=(
        'reviewdog/action-misspell'
        'actions/attest-build-provenance'
        'GrantBirki/git-diff-action'
        'golangci/golangci-lint-action'
        'actions/checkout'
        'actions/upload-artifact'
        'tcort/github-action-markdown-link-check'
    )

    readonly pinned_actions=()

    for action in $gh_actions; do
        is_pinned=false
        for pinned in "${pinned_actions[@]}"; do
            if [[ $action == "$pinned" ]]; then
                is_pinned=true
                break
            fi
        done

        if [[ $is_pinned == true ]]; then
            echo "Skipping auto-update for pinned action: $action"
            continue
        fi

        is_exception=false
        for ex in "${exceptions[@]}"; do
            if [[ $action == "$ex" ]]; then
                is_exception=true
                break
            fi
        done

        if [[ $is_exception == true ]]; then
            continue
        fi

        commit_hash=$(
            git ls-remote --tags "https://github.com/$action" |
                awk '
            {
                sha=$1
                ref=$2

                if (ref ~ /\^\{\}$/) {
                    tag=ref
                    sub(/\^\{\}$/, "", tag)
                    commits[tag]=sha
                } else {
                    tags[ref]=sha
                }
            }
            END {
                for (ref in tags) {
                    sha = (ref in commits ? commits[ref] : tags[ref])

                    tag = ref
                    sub(/^refs\/tags\//, "", tag)

                    # semver only
                    if (tag ~ /^v?[0-9]+(\.[0-9]+)*$/) {
                        sortkey = tag
                        sub(/^v/, "", sortkey)
                        print sortkey "\t" sha "\t" tag
                    }
                }
            }' |
                sort -V |
                tail -1 |
                awk -F'\t' '{ printf "%s # %s\n", $2, $3 }'
        )

        if [[ -z $commit_hash ]]; then
            echo "WARNING: unable to resolve tag for $action; skipping" >&2
            continue
        fi

        while IFS= read -r -d '' file; do
            sed -i -e "s|uses: $action@.*|uses: $action@$commit_hash|g" "$file"
        done < <(grep -ElRZ "uses: $action@" .github/)
    done
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

update_golangci_lint_version() {
    local version file

    version=$(latest_semver_tag "golangci/golangci-lint")

    if [[ -z $version ]]; then
        echo "WARNING: unable to resolve latest golangci-lint version; skipping" >&2
        return 1
    fi

    while IFS= read -r -d '' file; do
        sed -i "s|^\([[:space:]]*version:[[:space:]]*\)v[0-9][^ ]*|\1${version}|" "$file"
    done < <(grep -ElRZ "golangci/golangci-lint-action" .github/workflows/)
}

update_gremlins_version() {
    local version file

    version=$(latest_semver_tag "go-gremlins/gremlins")

    if [[ -z $version ]]; then
        echo "WARNING: unable to resolve latest gremlins version; skipping" >&2
        return 1
    fi

    while IFS= read -r -d '' file; do
        sed -i "s|^\([[:space:]]*version:[[:space:]]*\)v[0-9][^ ]*|\1${version}|" "$file"
    done < <(grep -ElRZ "go-gremlins/gremlins-action" .github/workflows/)
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
    has_errors=true
    failed_steps+=("updating pre-commit hooks")
fi
run_best_effort "updating GitHub Action commit hashes" update_github_action_hashes
run_best_effort "updating the Dockerfile base image" update_dockerfile_base_image
run_best_effort "updating golangci-lint version" update_golangci_lint_version
run_best_effort "updating gremlins version" update_gremlins_version
run_best_effort "updating rtk version" update_rtk_version

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
