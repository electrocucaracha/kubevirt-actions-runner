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
#set -o nounset
if [[ ${DEBUG:-false} == "true" ]]; then
    set -o xtrace
fi

# shellcheck source=scripts/_common.sh
source _common.sh
# shellcheck source=./scripts/_utils.sh
source _utils.sh

# NOTE: this env var is used by kind tool
export KIND_CLUSTER_NAME=k8s

# Install KubeVirt (https://kubevirt.io/quickstart_kind/)

function _create_cluster {
    if ! sudo "$(command -v kind)" get clusters | grep -e "$KIND_CLUSTER_NAME"; then
        sudo -E kind create cluster
        mkdir -p "$HOME/.kube"
        sudo chown -R "$USER": "$HOME/.kube"
        sudo -E kind get kubeconfig | tee "$HOME/.kube/config"
    fi
}

function _deploy_kubevirt {
    attempt_counter=0
    max_attempts=10
    VERSION=$(curl -s https://storage.googleapis.com/kubevirt-prow/release/kubevirt/kubevirt/stable.txt)

    kubectl create -f "https://github.com/kubevirt/kubevirt/releases/download/${VERSION}/kubevirt-operator.yaml"
    kubectl create -f "https://github.com/kubevirt/kubevirt/releases/download/${VERSION}/kubevirt-cr.yaml"

    info "Wait for Kubevirt resources to be ready"
    kubectl rollout status deployment/virt-operator -n kubevirt --timeout=5m
    until kubectl logs -n kubevirt -l kubevirt.io=virt-operator | grep "All KubeVirt components ready"; do
        if [ ${attempt_counter} -eq ${max_attempts} ]; then
            error "Max attempts reached"
        fi
        attempt_counter=$((attempt_counter + 1))
        sleep $((attempt_counter * 15))
    done
}

function main {
    _create_cluster
    _deploy_kubevirt
    kubectl apply -f test-data/vm.yaml
}

if [[ ${__name__:-"__main__"} == "__main__" ]]; then
    main
fi
