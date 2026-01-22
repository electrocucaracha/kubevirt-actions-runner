/*
Copyright © 2024

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

package runner

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	k8scorev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"
	"kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

const (
	runnerInfoAnnotation string = "electrocucaracha.kubevirt-actions-runner/runner-info"
	runnerInfoVolume     string = "runner-info"
	runnerInfoPath       string = "runner-info.json"
)

type Runner interface {
	CreateResources(ctx context.Context, vmTemplate string, runnerName string, jitConfig string) error
	WaitForVirtualMachineInstance(ctx context.Context) error
	DeleteResources(ctx context.Context) error
}

type KubevirtRunner struct {
	virtClient kubecli.KubevirtClient
	namespace  string
}

var _ Runner = (*KubevirtRunner)(nil)

func NewRunner(namespace string, virtClient kubecli.KubevirtClient) *KubevirtRunner {
	return &KubevirtRunner{
		namespace:  namespace,
		virtClient: virtClient,
	}
}

func generateRunnerInfoVolume() v1.Volume {
	return v1.Volume{
		Name: runnerInfoVolume,
		VolumeSource: v1.VolumeSource{
			DownwardAPI: &v1.DownwardAPIVolumeSource{
				Fields: []k8scorev1.DownwardAPIVolumeFile{
					{
						Path: runnerInfoPath,
						FieldRef: &k8scorev1.ObjectFieldSelector{
							FieldPath: fmt.Sprintf("metadata.annotations['%s']", runnerInfoAnnotation),
						},
					},
				},
			},
		},
	}
}

func (rc *KubevirtRunner) CreateResources(ctx context.Context,
	vmTemplate, runnerName, jitConfig string,
) error {
	tracer := otel.Tracer("kubevirt-actions-runner/runner")

	ctx, span := tracer.Start(ctx, "CreateResources",
		trace.WithAttributes(
			attribute.String("vmTemplate", vmTemplate),
			attribute.String("runnerName", runnerName),
			attribute.String("namespace", rc.namespace),
		),
	)
	defer span.End()

	err := rc.validateResourceInputs(vmTemplate, runnerName, jitConfig, span)
	if err != nil {
		return err
	}

	virtualMachineInstance, dataVolume, err := rc.getResources(ctx, vmTemplate, runnerName, jitConfig)
	if err != nil {
		span.RecordError(err)

		return err
	}

	_, spanCreateVMI := tracer.Start(ctx, "CreateVMI",
		trace.WithAttributes(
			attribute.String("vmiName", virtualMachineInstance.Name),
		),
	)
	defer spanCreateVMI.End()

	vmi, err := rc.createVMI(ctx, tracer, virtualMachineInstance, span, spanCreateVMI)
	if err != nil {
		return err
	}

	dataVolumeName := ""
	if dataVolume != nil {
		dataVolumeName = dataVolume.Name

		err := rc.createDataVolume(ctx, tracer, dataVolume, vmi.Name, vmi.UID, span)
		if err != nil {
			return err
		}
	}

	NewAppContext(virtualMachineInstance.Name, dataVolumeName)

	return nil
}

func (rc *KubevirtRunner) WaitForVirtualMachineInstance(ctx context.Context) error {
	tracer := otel.Tracer("kubevirt-actions-runner/runner")

	ctx, span := tracer.Start(ctx, "WaitForVirtualMachineInstance")
	defer span.End()

	log := GetLogger()
	appCtx := GetAppContext()

	log.Printf("Watching %s Virtual Machine Instance\n", appCtx.GetVMIName())
	span.SetAttributes(attribute.String("vmiName", appCtx.GetVMIName()))

	watch, err := rc.virtClient.VirtualMachineInstance(rc.namespace).Watch(ctx, k8smetav1.ListOptions{})
	if err != nil {
		span.RecordError(err)

		return errors.Wrap(err, "failed to watch the virtual machine instance")
	}

	defer watch.Stop()

	var currentStatus v1.VirtualMachineInstancePhase

	for event := range watch.ResultChan() {
		vmi, ok := event.Object.(*v1.VirtualMachineInstance)
		if ok && vmi.Name == appCtx.GetVMIName() {
			if vmi.Status.Phase != currentStatus {
				switch vmi.Status.Phase {
				case v1.Succeeded:
					log.Printf("%s has successfully completed\n", appCtx.GetVMIName())
					span.SetAttributes(attribute.String("phase", "Succeeded"))

					return nil
				case v1.Failed:
					log.Printf("%s has failed\n", appCtx.GetVMIName())
					span.SetAttributes(attribute.String("phase", "Failed"))

					return ErrRunnerFailed
				case v1.VmPhaseUnset, v1.Pending, v1.Scheduling, v1.Scheduled, v1.Running, v1.Unknown, v1.WaitingForSync:
					log.Printf("%s has transitioned to %s phase \n", appCtx.GetVMIName(), vmi.Status.Phase)
					span.AddEvent("phase_transition", trace.WithAttributes(
						attribute.String("phase", string(vmi.Status.Phase)),
					))

					currentStatus = vmi.Status.Phase
				}
			}
		}
	}

	return nil
}

func (rc *KubevirtRunner) DeleteResources(ctx context.Context) error {
	tracer := otel.Tracer("kubevirt-actions-runner/runner")

	ctx, span := tracer.Start(ctx, "DeleteResources")
	defer span.End()

	log := GetLogger()
	appCtx := GetAppContext()

	log.Printf("Cleaning %s Virtual Machine Instance resources\n",
		appCtx.GetVMIName())
	span.SetAttributes(attribute.String("vmiName", appCtx.GetVMIName()))

	err := rc.virtClient.VirtualMachineInstance(rc.namespace).Delete(
		ctx, appCtx.GetVMIName(), k8smetav1.DeleteOptions{})
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			log.Printf("fail to delete runner instance %s: %v", appCtx.GetVMIName(), err)
			span.RecordError(err)
		}
	}

	if len(appCtx.GetDataVolumeName()) > 0 {
		_, spanDeleteDV := tracer.Start(ctx, "DeleteDataVolume",
			trace.WithAttributes(
				attribute.String("dataVolumeName", appCtx.GetDataVolumeName()),
			),
		)

		err := rc.virtClient.CdiClient().CdiV1beta1().DataVolumes(rc.namespace).Delete(
			ctx, appCtx.GetDataVolumeName(), k8smetav1.DeleteOptions{})
		if err != nil {
			if !k8serrors.IsNotFound(err) {
				log.Printf("fail to delete runner data volume %s: %v", appCtx.GetDataVolumeName(), err)
				spanDeleteDV.RecordError(err)
			}
		}

		spanDeleteDV.End()
	}

	return nil
}

func (rc *KubevirtRunner) validateResourceInputs(vmTemplate, runnerName, jitConfig string, span trace.Span) error {
	if len(vmTemplate) == 0 {
		span.SetAttributes(attribute.String("error", "empty vm template"))

		return ErrEmptyVMTemplate
	}

	if len(runnerName) == 0 {
		span.SetAttributes(attribute.String("error", "empty runner name"))

		return ErrEmptyRunnerName
	}

	if len(jitConfig) == 0 {
		span.SetAttributes(attribute.String("error", "empty jit config"))

		return ErrEmptyJitConfig
	}

	return nil
}

func (rc *KubevirtRunner) createVMI(
	ctx context.Context,
	_ trace.Tracer,
	vmi *v1.VirtualMachineInstance,
	span, spanCreateVMI trace.Span,
) (*v1.VirtualMachineInstance, error) {
	log := GetLogger()
	log.Printf("Creating %s Virtual Machine Instance\n", vmi.Name)

	createdVMI, err := rc.virtClient.VirtualMachineInstance(rc.namespace).Create(ctx,
		vmi, k8smetav1.CreateOptions{})
	if err != nil {
		spanCreateVMI.RecordError(err)
		span.RecordError(err)

		return nil, errors.Wrap(err, "fail to create runner instance")
	}

	return createdVMI, nil
}

func (rc *KubevirtRunner) createDataVolume(
	ctx context.Context,
	tracer trace.Tracer,
	dataVolume *v1beta1.DataVolume,
	vmiName string,
	vmiUID types.UID,
	span trace.Span,
) error {
	if dataVolume == nil {
		return nil
	}

	log := GetLogger()
	log.Printf("Creating %s Data Volume\n", dataVolume.Name)

	_, spanCreateDV := tracer.Start(ctx, "CreateDataVolume",
		trace.WithAttributes(
			attribute.String("dataVolumeName", dataVolume.Name),
		),
	)
	defer spanCreateDV.End()

	dataVolume.OwnerReferences = []k8smetav1.OwnerReference{
		{
			APIVersion: "kubevirt.io/v1",
			Kind:       "VirtualMachineInstance",
			Name:       vmiName,
			UID:        vmiUID,
			Controller: ptr.To(false),
		},
	}

	_, err := rc.virtClient.CdiClient().CdiV1beta1().DataVolumes(
		rc.namespace).Create(ctx, dataVolume, k8smetav1.CreateOptions{})
	if err != nil {
		spanCreateDV.RecordError(err)
		span.RecordError(err)

		return errors.Wrap(err, "cannot create data volume")
	}

	return nil
}

func (rc *KubevirtRunner) getResources(ctx context.Context, vmTemplate, runnerName, jitConfig string) (
	*v1.VirtualMachineInstance, *v1beta1.DataVolume, error,
) {
	virtualMachine, err := rc.virtClient.VirtualMachine(rc.namespace).Get(
		ctx, vmTemplate, k8smetav1.GetOptions{})
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot obtain KubeVirt vm list")
	}

	virtualMachineInstance := v1.NewVMIReferenceFromNameWithNS(rc.namespace, runnerName)
	virtualMachineInstance.Spec = virtualMachine.Spec.Template.Spec

	if virtualMachineInstance.Annotations == nil {
		virtualMachineInstance.Annotations = make(map[string]string)
	}

	jri := map[string]any{
		"jitconfig": jitConfig,
	}

	out, err := json.Marshal(jri)
	if err != nil {
		return nil, nil, errors.Wrap(err, "cannot marshal jitConfig")
	}

	virtualMachineInstance.Annotations[runnerInfoAnnotation] = string(out)

	var dataVolume *v1beta1.DataVolume

	for _, dvt := range virtualMachine.Spec.DataVolumeTemplates {
		for _, volume := range virtualMachineInstance.Spec.Volumes {
			if volume.DataVolume != nil && volume.DataVolume.Name == dvt.Name {
				dataVolume = &v1beta1.DataVolume{
					ObjectMeta: k8smetav1.ObjectMeta{
						Name: fmt.Sprintf("%s-%s", dvt.Name, runnerName),
					},
					Spec: dvt.Spec,
				}

				volume.DataVolume.Name = dataVolume.Name

				break
			}
		}
	}

	virtualMachineInstance.Spec.Volumes = append(virtualMachineInstance.Spec.Volumes, generateRunnerInfoVolume())

	return virtualMachineInstance, dataVolume, nil
}
