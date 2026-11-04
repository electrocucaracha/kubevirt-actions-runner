#!/bin/bash
# SPDX-license-identifier: Apache-2.0
##############################################################################
# Copyright (c) 2026
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Apache License, Version 2.0
# which accompanies this distribution, and is available at
# http://www.apache.org/licenses/LICENSE-2.0
##############################################################################

set -o pipefail
set -o errexit
set -o nounset

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly script_dir
repo_root="$(cd "${script_dir}/.." && pwd)"
readonly repo_root
readonly test_namespace="${KAR_CLEANUP_TEST_NAMESPACE:-}"
readonly template_name="${KAR_CLEANUP_TEST_TEMPLATE:-}"
template_namespace="${KAR_CLEANUP_TEST_TEMPLATE_NAMESPACE:-${test_namespace}}"
readonly template_namespace
readonly wait_timeout="${KAR_CLEANUP_TEST_WAIT_TIMEOUT:-2s}"
readonly cleanup_timeout="${KAR_CLEANUP_TEST_TIMEOUT:-30s}"

if [[ -z ${test_namespace} || -z ${template_name} ]]; then
    printf '%s\n' \
        'Set KAR_CLEANUP_TEST_NAMESPACE and KAR_CLEANUP_TEST_TEMPLATE to a' \
        'disposable namespace and an unschedulable VM template.' >&2
    exit 2
fi

command -v kubectl >/dev/null
command -v go >/dev/null

temp_dir="$(mktemp -d)"
trap 'rm -rf "${temp_dir}"' EXIT

readonly kubeconfig="${temp_dir}/kubeconfig"
readonly kar_binary="${temp_dir}/kar"
readonly kar_log="${temp_dir}/kar.log"
runner_name="cleanup-test-$(date +%s)-$$"
readonly runner_name

kubectl config view --raw --flatten >"${kubeconfig}"
chmod 600 "${kubeconfig}"
export KUBECONFIG="${kubeconfig}"
kubectl config set-context --current --namespace="${test_namespace}" >/dev/null

kubectl get namespace "${test_namespace}" >/dev/null
kubectl -n "${template_namespace}" get vm "${template_name}" >/dev/null

printf 'Running cleanup timeout scenario for runner %s\n' "${runner_name}"
go -C "${repo_root}" build -o "${kar_binary}" ./cmd/kar

if KAR_TELEMETRY_ENABLED=false \
    KAR_WAIT_TIMEOUT="${wait_timeout}" \
    KAR_CLEANUP_TIMEOUT="${cleanup_timeout}" \
    "${kar_binary}" \
    --kubevirt-vm-template "${template_name}" \
    --kubevirt-vm-template-namespace "${template_namespace}" \
    --runner-name "${runner_name}" \
    --actions-runner-input-jitconfig e30= >"${kar_log}" 2>&1; then
    cat "${kar_log}"
    printf 'Expected kar to fail after the VMI wait timeout.\n' >&2
    exit 1
else
    kar_exit_code=$?
fi

if [[ ${kar_exit_code} -ne 1 ]] || ! grep -Fq 'failed to wait for resources' "${kar_log}"; then
    cat "${kar_log}"
    printf '%s (exit code %s).\n' \
        'kar did not fail because of the expected VMI wait timeout' \
        "${kar_exit_code}" >&2
    exit 1
fi

if ! vmi_resource="$(kubectl -n "${test_namespace}" get vmi "${runner_name}" \
    -o name --ignore-not-found)"; then
    printf 'Could not verify VMI cleanup for %s.\n' \
        "${runner_name}" >&2
    exit 1
fi

if [[ -z ${vmi_resource} ]]; then
    printf 'PASS: VMI %s was deleted before verification.\n' "${runner_name}"
    exit 0
fi

deletion_timestamp="$(kubectl -n "${test_namespace}" get vmi "${runner_name}" \
    -o 'jsonpath={.metadata.deletionTimestamp}')"
if [[ -z ${deletion_timestamp} ]]; then
    printf 'VMI %s still exists without a deletion timestamp.\n' \
        "${runner_name}" >&2
    printf 'Inspect it with: kubectl -n %s get vmi %s -o yaml\n' \
        "${test_namespace}" "${runner_name}" >&2
    exit 1
fi

printf 'PASS: Kubernetes accepted deletion of VMI %s at %s.\n' \
    "${runner_name}" "${deletion_timestamp}"
