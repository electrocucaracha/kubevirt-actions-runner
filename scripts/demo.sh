#!/bin/bash
# SPDX-license-identifier: Apache-2.0
##############################################################################
# Copyright (c) 2025
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Apache License, Version 2.0
# which accompanies this distribution, and is available at
# http://www.apache.org/licenses/LICENSE-2.0
##############################################################################

set -o pipefail
set -o errexit
set -o nounset

# shellcheck source=scripts/_common.sh
source _common.sh
# shellcheck source=./scripts/_utils.sh
source _utils.sh

export KAR_TELEMETRY_ENABLED=true
export KAR_TELEMETRY_EXPORT_TYPE=stdout

readonly VM_TEMPLATE_NAMESPACE="${KAR_VM_TEMPLATE_NAMESPACE:-default}"
readonly RUNNER_NAMESPACE="${KAR_RUNNER_NAMESPACE:-kar-runner-demo}"

function ensure_distinct_namespaces {
    if [[ ${VM_TEMPLATE_NAMESPACE} == "${RUNNER_NAMESPACE}" ]]; then
        error "KAR_VM_TEMPLATE_NAMESPACE (${VM_TEMPLATE_NAMESPACE}) and KAR_RUNNER_NAMESPACE (${RUNNER_NAMESPACE}) must be different"
    fi
}

function ensure_namespace_exists {
    local namespace="$1"

    if [[ ${namespace} == "default" ]]; then
        return
    fi

    kubectl create namespace "${namespace}" --dry-run=client -o yaml | kubectl apply -f -
}

function prepare_demo_namespaces {
    ensure_namespace_exists "${VM_TEMPLATE_NAMESPACE}"
    ensure_namespace_exists "${RUNNER_NAMESPACE}"

    # The VM template can be created in a dedicated namespace and retrieved via -n.
    kubectl apply -f test-data/vm.yaml -n "${VM_TEMPLATE_NAMESPACE}"
}

trap get_status ERR

info "Running a alpine demo instance"
ensure_distinct_namespaces
prepare_demo_namespaces
# shellcheck disable=SC1091
[ -f /etc/profile.d/path.sh ] && . /etc/profile.d/path.sh
kar_dir="$(mktemp -d)"
kar_bin="$kar_dir/kar"
trap 'rm -rf "$kar_dir"' EXIT
go build -o "$kar_bin" "$(git rev-parse --show-toplevel)/cmd/kar"
# Run the runner in a namespace different from the VM template namespace.
kar_kubeconfig="${kar_dir}/config"
cp "${KUBECONFIG:-$HOME/.kube/config}" "${kar_kubeconfig}"
KUBECONFIG="${kar_kubeconfig}" kubectl config set-context --current --namespace="${RUNNER_NAMESPACE}" >/dev/null

KUBECONFIG="${kar_kubeconfig}" timeout 5m "$kar_bin" \
    -c test-data/runner-info.json \
    -t testvm \
    -n "${VM_TEMPLATE_NAMESPACE}" \
    -r test
info "Demo completed"
