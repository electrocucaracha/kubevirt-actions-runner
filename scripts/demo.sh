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

function exit_trap {
    printf "CPU usage: "
    grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$4+$5)} END {print usage " %"}'
    printf "Memory free(Kb):"
    awk -v low="$(grep low /proc/zoneinfo | awk '{k+=$2}END{print k}')" '{a[$1]=$2}  END{ print a["MemFree:"]+a["Active(file):"]+a["Inactive(file):"]+a["SReclaimable:"]-(12*low);}' /proc/meminfo
    echo "Environment variables:"
    env | grep "KRD"
    if command -v kubectl; then
        echo "Kubernetes Events:"
        kubectl get events -A --sort-by=".metadata.managedFields[0].time"
        echo "Kubernetes Resources:"
        kubectl get all -A -o wide
        echo "Kubernetes Pods:"
        kubectl describe pods
        echo "Kubernetes Nodes:"
        kubectl describe nodes
    fi
}

trap exit_trap ERR

info "Running a alpine demo instance"
# shellcheck disable=SC1091
[ -f /etc/profile.d/path.sh ] && . /etc/profile.d/path.sh
timeout 5m go run "$(git rev-parse --show-toplevel)/cmd/kar/main.go" -c test-data/runner-info.json -t testvm -r test
info "Demo completed"
