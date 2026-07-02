/* jscpd:ignore-start */
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
/* jscpd:ignore-end */

package runner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	k8scorev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/types"
	k8swatch "k8s.io/apimachinery/pkg/watch"
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"
	"kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

const (
	tracerName                   = "kubevirt-actions-runner/runner"
	runnerInfoAnnotation  string = "electrocucaracha.kubevirt-actions-runner/runner-info"
	runnerInfoVolume      string = "runner-info"
	runnerInfoPath        string = "runner-info.json"
	watchReconnectBackoff        = time.Second
	watchChannelClosedMsg        = "watch channel closed unexpectedly"
)

var errWaitTimeout = errors.New("timeout while waiting for the virtual machine instance")

type Runner interface {
	CreateResources(ctx context.Context,
		vmTemplate string,
		vmTemplateNamespace string,
		runnerName string,
		jitConfig string,
	) error
	WaitForVirtualMachineInstance(ctx context.Context) error
	DeleteResources(ctx context.Context) error
}

type KubevirtRunner struct {
	virtClient  kubecli.KubevirtClient
	namespace   string
	waitTimeout time.Duration
}

var _ Runner = (*KubevirtRunner)(nil)

func NewRunner(namespace string, virtClient kubecli.KubevirtClient, waitTimeout time.Duration) *KubevirtRunner {
	return &KubevirtRunner{
		namespace:   namespace,
		virtClient:  virtClient,
		waitTimeout: waitTimeout,
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
	vmTemplate, vmTemplateNamespace, runnerName, jitConfig string,
) error {
	tracer := otel.Tracer(tracerName)

	if vmTemplateNamespace == "" {
		vmTemplateNamespace = k8scorev1.NamespaceDefault
	}

	ctx, span := tracer.Start(ctx, "CreateResources",
		trace.WithAttributes(
			attribute.String("vmTemplate", vmTemplate),
			attribute.String("vmTemplateNamespace", vmTemplateNamespace),
			attribute.String("runnerName", runnerName),
			attribute.String("namespace", rc.namespace),
		),
	)
	defer span.End()

	err := rc.validateResourceInputs(vmTemplate, runnerName, jitConfig, span)
	if err != nil {
		return err
	}

	virtualMachineInstance, dataVolume, err := rc.getResources(
		ctx,
		vmTemplate,
		vmTemplateNamespace,
		runnerName,
		jitConfig,
	)
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

	vmi, err := rc.createVMI(ctx, virtualMachineInstance, span, spanCreateVMI)
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
	tracer := otel.Tracer(tracerName)

	ctx, cancel := context.WithTimeout(ctx, rc.waitTimeout)
	defer cancel()

	ctx, span := tracer.Start(ctx, "WaitForVirtualMachineInstance")
	defer span.End()

	log := utils.GetLogger()
	vmiName := GetAppContext().GetVMIName()

	log.Printf("Watching %s Virtual Machine Instance\n", vmiName)
	span.SetAttributes(attribute.String("vmiName", vmiName))

	var (
		currentStatus v1.VirtualMachineInstancePhase
		readyReported bool
	)

	vmiInterface := rc.virtClient.VirtualMachineInstance(rc.namespace)

	for {
		done, resourceVersion, terminalErr := rc.refreshVMIStatus(
			ctx, span, vmiInterface, vmiName, &currentStatus, &readyReported)
		if done {
			return terminalErr
		}

		watch, watchErr := vmiInterface.Watch(ctx, watchOptions(vmiName, resourceVersion))
		if watchErr != nil {
			span.RecordError(watchErr)

			return fmt.Errorf("failed to watch the virtual machine instance: %w", watchErr)
		}

		done, watchResultErr := watchVMIEvents(ctx, span, watch, vmiName, &currentStatus, &readyReported)
		watch.Stop()

		if done {
			return watchResultErr
		}

		if watchResultErr == nil {
			log.Printf("Watch stream closed for %s Virtual Machine Instance; reconnecting\n", vmiName)
			span.AddEvent("watch_reconnect", trace.WithAttributes(attribute.String("reason", watchChannelClosedMsg)))

			timer := time.NewTimer(watchReconnectBackoff)
			select {
			case <-ctx.Done():
				timer.Stop()

				return errWaitTimeout
			case <-timer.C:
			}

			continue
		}

		return watchResultErr
	}
}

func watchOptions(vmiName, resourceVersion string) k8smetav1.ListOptions {
	return k8smetav1.ListOptions{
		FieldSelector:   fields.OneTermEqualSelector("metadata.name", vmiName).String(),
		ResourceVersion: resourceVersion,
	}
}

func watchVMIEvents(
	ctx context.Context,
	span trace.Span,
	watch k8swatch.Interface,
	vmiName string,
	currentStatus *v1.VirtualMachineInstancePhase,
	readyReported *bool,
) (bool, error) {
	for {
		select {
		case <-ctx.Done():
			return true, errWaitTimeout
		case event, watchOpen := <-watch.ResultChan():
			if !watchOpen {
				return false, nil
			}

			done, skip, err := handleWatchEvent(span, vmiName, event, currentStatus, readyReported)
			if skip {
				continue
			}

			if done {
				return true, err
			}
		}
	}
}

// handleWatchEvent processes one event from the VMI watch channel.
// It returns (done=true, _, err) when the watch loop should exit,
// (false, skip=true, nil) when the event should be ignored, or
// (false, false, nil) to continue watching.
func handleWatchEvent(
	span trace.Span,
	vmiName string,
	event k8swatch.Event,
	currentStatus *v1.VirtualMachineInstancePhase,
	readyReported *bool,
) (bool, bool, error) {
	vmi, isVMI := event.Object.(*v1.VirtualMachineInstance)
	if !isVMI || vmi.Name != vmiName {
		return false, true, nil
	}

	reportReadyMilestone(span, vmiName, vmi, readyReported)

	if vmi.Status.Phase == *currentStatus {
		return false, false, nil
	}

	done, err := handleVMIPhase(span, vmiName, vmi.Status.Phase)
	*currentStatus = vmi.Status.Phase

	return done, false, err
}

// handleVMIPhase processes a VMI phase transition. It returns (true, err) when a
// terminal state (Succeeded or Failed) is reached, or (false, nil) for non-terminal phases.
func handleVMIPhase(span trace.Span, vmiName string, phase v1.VirtualMachineInstancePhase) (bool, error) {
	log := utils.GetLogger()

	switch phase {
	case v1.Succeeded:
		log.Printf("%s has successfully completed\n", vmiName)
		span.SetAttributes(attribute.String("phase", "Succeeded"))

		return true, nil
	case v1.Failed:
		log.Printf("%s has failed\n", vmiName)
		span.SetAttributes(attribute.String("phase", "Failed"))

		return true, ErrRunnerFailed
	case v1.VmPhaseUnset, v1.Pending, v1.Scheduling, v1.Scheduled, v1.Running, v1.Unknown, v1.WaitingForSync:
		log.Printf("%s has transitioned to %s phase \n", vmiName, phase)
		span.AddEvent("phase_transition", trace.WithAttributes(
			attribute.String("phase", string(phase)),
		))

		return false, nil
	default:
		log.Printf("%s encountered an unrecognized phase: %s\n", vmiName, phase)
		span.AddEvent("phase_unhandled", trace.WithAttributes(
			attribute.String("phase", string(phase)),
		))

		return false, nil
	}
}

// isVMIReady reports whether the VMI has the Ready condition set to True.
func isVMIReady(vmi *v1.VirtualMachineInstance) bool {
	for _, cond := range vmi.Status.Conditions {
		if cond.Type == v1.VirtualMachineInstanceReady && cond.Status == k8scorev1.ConditionTrue {
			return true
		}
	}

	return false
}

// reportReadyMilestone logs and records a span attribute when a VMI first reaches Running+Ready.
// It is a no-op if readyReported is already true.
func reportReadyMilestone(span trace.Span, vmiName string, vmi *v1.VirtualMachineInstance, readyReported *bool) {
	if vmi.Status.Phase == v1.Running && isVMIReady(vmi) && !*readyReported {
		utils.GetLogger().Printf("%s is Running and Ready\n", vmiName)
		span.SetAttributes(attribute.String("phase", "Running+Ready"))

		*readyReported = true
	}
}

func (rc *KubevirtRunner) DeleteResources(ctx context.Context) error {
	tracer := otel.Tracer(tracerName)

	ctx, span := tracer.Start(ctx, "DeleteResources")
	defer span.End()

	if !HasAppContext() {
		return nil
	}

	log := utils.GetLogger()
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

func (rc *KubevirtRunner) refreshVMIStatus(
	ctx context.Context,
	span trace.Span,
	vmiInterface kubecli.VirtualMachineInstanceInterface,
	vmiName string,
	currentStatus *v1.VirtualMachineInstancePhase,
	readyReported *bool,
) (bool, string, error) {
	if ctx.Err() != nil {
		return true, "", errWaitTimeout
	}

	vmi, err := vmiInterface.Get(ctx, vmiName, k8smetav1.GetOptions{})
	if err != nil {
		if ctx.Err() != nil {
			return true, "", errWaitTimeout
		}

		span.RecordError(err)

		return true, "", fmt.Errorf("failed to get the virtual machine instance %q: %w", vmiName, err)
	}

	reportReadyMilestone(span, vmiName, vmi, readyReported)

	if vmi.Status.Phase == *currentStatus {
		return false, vmi.ResourceVersion, nil
	}

	done, err := handleVMIPhase(span, vmiName, vmi.Status.Phase)
	*currentStatus = vmi.Status.Phase

	return done, vmi.ResourceVersion, err
}

func (rc *KubevirtRunner) validateResourceInputs(vmTemplate, runnerName, jitConfig string, span trace.Span) error {
	if vmTemplate == "" {
		span.SetAttributes(attribute.String("error", "empty vm template"))

		return ErrEmptyVMTemplate
	}

	if runnerName == "" {
		span.SetAttributes(attribute.String("error", "empty runner name"))

		return ErrEmptyRunnerName
	}

	if jitConfig == "" {
		span.SetAttributes(attribute.String("error", "empty jit config"))

		return ErrEmptyJitConfig
	}

	return nil
}

func (rc *KubevirtRunner) createVMI(
	ctx context.Context,
	vmi *v1.VirtualMachineInstance,
	span, spanCreateVMI trace.Span,
) (*v1.VirtualMachineInstance, error) {
	log := utils.GetLogger()
	log.Printf("Creating %s Virtual Machine Instance\n", vmi.Name)

	createdVMI, err := rc.virtClient.VirtualMachineInstance(rc.namespace).Create(ctx,
		vmi, k8smetav1.CreateOptions{})
	if err != nil {
		spanCreateVMI.RecordError(err)
		span.RecordError(err)

		return nil, fmt.Errorf("failed to create runner instance: %w", err)
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
	log := utils.GetLogger()
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
			Controller: new(bool),
		},
	}

	_, err := rc.virtClient.CdiClient().CdiV1beta1().DataVolumes(
		rc.namespace).Create(ctx, dataVolume, k8smetav1.CreateOptions{})
	if err != nil {
		spanCreateDV.RecordError(err)
		span.RecordError(err)

		return fmt.Errorf("cannot create data volume: %w", err)
	}

	return nil
}

func (rc *KubevirtRunner) getResources(
	ctx context.Context,
	vmTemplate, vmTemplateNamespace, runnerName, jitConfig string,
) (
	*v1.VirtualMachineInstance, *v1beta1.DataVolume, error,
) {
	virtualMachine, err := rc.virtClient.VirtualMachine(vmTemplateNamespace).Get(
		ctx, vmTemplate, k8smetav1.GetOptions{})
	if err != nil {
		return nil, nil, fmt.Errorf(
			"failed to get KubeVirt virtual machine template %q in namespace %q: %w",
			vmTemplate,
			vmTemplateNamespace,
			err,
		)
	}

	virtualMachineInstance := v1.NewVMIReferenceFromNameWithNS(rc.namespace, runnerName)
	virtualMachineInstance.Spec = virtualMachine.Spec.Template.Spec

	if virtualMachineInstance.Annotations == nil {
		virtualMachineInstance.Annotations = make(map[string]string)
	}

	runnerInfo := map[string]string{
		"jitconfig": jitConfig,
	}

	out, err := json.Marshal(runnerInfo)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot marshal runner info annotation payload: %w", err)
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
