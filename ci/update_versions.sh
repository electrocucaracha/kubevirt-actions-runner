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

trap "make fmt" EXIT

if ! command -v go >/dev/null; then
    curl -fsSL http://bit.ly/install_pkg | PKG=go-lang bash
    # shellcheck disable=SC1091
    source /etc/profile.d/path.sh
fi

go_version="$(curl -sL https://golang.org/VERSION?m=text | sed -n 's/go//;s/\..$//;1p')"
go mod tidy -go="$go_version"
sed -i "s/go-version: .*/go-version: \"^$go_version\"/g" .github/workflows/update.yml

if ! command -v uvx >/dev/null; then
    curl -LsSf https://astral.sh/uv/install.sh | sh
fi
uvx pre-commit autoupdate

# Update GitHub Action commit hashes
gh_actions=$(grep -r "uses: [A-Za-z0-9_.-]*/[\_a-z\-]*@" .github/ | sed 's/@.*//' | awk -F ': ' '{ print $3 }' | sort -u)
exceptions=('reviewdog/action-misspell')
for action in $gh_actions; do
    if [[ ${exceptions[*]} =~ (^|[^[:alpha:]])$action([^[:alpha:]]|$) ]]; then
        commit_hash=$(git ls-remote "https://github.com/$action" | grep 'refs/tags/[v]\?[0-9][0-9\.]*\^{}$' | sed 's|refs/tags/[vV]\?[\.]\?||g; s|\^{}$||g' | sort -u -k2 -V | tail -1 | awk '{ printf "%s # %s\n",$1,$2 }')
    else
        commit_hash=$(git ls-remote "https://github.com/$action" | grep 'refs/tags/[v]\?[0-9][0-9\.]*$' | sed 's|refs/tags/[vV]\?[\.]\?||g' | sort -u -k2 -V | tail -1 | awk '{ printf "%s # %s\n",$1,$2 }')
    fi
    # shellcheck disable=SC2267
    grep -ElRZ "uses: $action@" .github/ | xargs -0 -l sed -i -e "s|uses: $action@.*|uses: $action@$commit_hash|g"
done
