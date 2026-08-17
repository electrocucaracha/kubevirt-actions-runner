/* jscpd:ignore-start */
/*
Copyright © 2026

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/* jscpd:ignore-end */

// This file exercises unexported seams (e.g. marshalJSON) that cannot be
// reached from the external runner_test package, so it lives in package
// runner as a white-box test.
package runner

import (
	"context"
	"errors"
	"testing"

	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/mock/gomock"
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"
)

var (
	errSimulatedMarshalFailure        = errors.New("simulated marshal failure")
	errSimulatedResourceFailure       = errors.New("simulated resource creation failure")
	errSimulatedStdoutExporterFailure = errors.New("simulated stdout exporter creation failure")
)

// TestGetResourcesMarshalJSONError exercises the previously-uncovered error
// path in getResources where encoding the runner-info annotation payload
// fails. It swaps the marshalJSON seam to force the failure deterministically.
func TestGetResourcesMarshalJSONError(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	const (
		namespace  = "default"
		vmTemplate = "vm-template"
		runnerName = "runner-xyz123"
		jitConfig  = "jit-config"
	)

	virtClient := kubecli.NewMockKubevirtClient(mockCtrl)
	vmInterface := kubecli.NewMockVirtualMachineInterface(mockCtrl)

	virtClient.EXPECT().VirtualMachine(namespace).Return(vmInterface)
	vmInterface.EXPECT().Get(gomock.Any(), vmTemplate, gomock.Any()).Return(&v1.VirtualMachine{
		Spec: v1.VirtualMachineSpec{
			Template: &v1.VirtualMachineInstanceTemplateSpec{},
		},
	}, nil)

	runner := &KubevirtRunner{virtClient: virtClient, namespace: namespace}

	originalMarshal := marshalJSON

	defer func() { marshalJSON = originalMarshal }()

	marshalJSON = func(_ any) ([]byte, error) {
		return nil, errSimulatedMarshalFailure
	}

	vmi, dataVolume, err := runner.getResources(context.Background(), vmTemplate, namespace, runnerName, jitConfig)
	if err == nil {
		t.Fatal("expected an error when marshalling the runner info annotation payload fails")
	}

	if !errors.Is(err, errSimulatedMarshalFailure) {
		t.Fatalf("expected error to wrap %v, got %v", errSimulatedMarshalFailure, err)
	}

	if vmi != nil {
		t.Fatalf("expected nil VirtualMachineInstance, got %+v", vmi)
	}

	if dataVolume != nil {
		t.Fatalf("expected nil DataVolume, got %+v", dataVolume)
	}
}

// TestInitializeTelemetryResourceCreationError exercises the previously
// uncovered error path in InitializeTelemetry where building the OpenTelemetry
// resource fails. It swaps the newResource seam to force the failure
// deterministically.
func TestInitializeTelemetryResourceCreationError(t *testing.T) {
	t.Setenv("KAR_TELEMETRY_ENABLED", "true")
	t.Setenv("KAR_TELEMETRY_EXPORT_TYPE", "stdout")

	originalNewResource := newResource

	defer func() { newResource = originalNewResource }()

	newResource = func(_ context.Context, _ ...resource.Option) (*resource.Resource, error) {
		return nil, errSimulatedResourceFailure
	}

	shutdown, err := InitializeTelemetry(context.Background())
	if err == nil {
		t.Fatal("expected an error when resource creation fails")
	}

	if !errors.Is(err, errSimulatedResourceFailure) {
		t.Fatalf("expected error to wrap %v, got %v", errSimulatedResourceFailure, err)
	}

	if shutdown != nil {
		t.Fatal("expected nil shutdown function when resource creation fails")
	}
}

// TestCreateExporterStdoutCreationError exercises the previously uncovered
// error path in createExporter where building the stdout exporter fails. It
// swaps the newStdoutExporter seam to force the failure deterministically.
func TestCreateExporterStdoutCreationError(t *testing.T) {
	t.Parallel()

	originalNewStdoutExporter := newStdoutExporter

	defer func() { newStdoutExporter = originalNewStdoutExporter }()

	newStdoutExporter = func() (trace.SpanExporter, error) {
		return nil, errSimulatedStdoutExporterFailure
	}

	exporter, err := createExporter(context.Background(), "stdout")
	if err == nil {
		t.Fatal("expected an error when stdout exporter creation fails")
	}

	if !errors.Is(err, errSimulatedStdoutExporterFailure) {
		t.Fatalf("expected error to wrap %v, got %v", errSimulatedStdoutExporterFailure, err)
	}

	if exporter != nil {
		t.Fatal("expected nil exporter when stdout exporter creation fails")
	}
}
