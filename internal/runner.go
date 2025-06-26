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
	"log"
	"time"

	"github.com/pkg/errors"
	k8scorev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	WaitForVirtualMachineInstance(ctx context.Context, vmi string) error
	DeleteResources(ctx context.Context, vmi string, dv string) error
	GetVMIName() string
	GetDataVolumeName() string
}

type KubevirtRunner struct {
	virtClient             kubecli.KubevirtClient
	namespace              string
	dataVolume             string
	virtualMachineInstance string
	currentStatus          v1.VirtualMachineInstancePhase
}

var _ Runner = (*KubevirtRunner)(nil)

func (rc *KubevirtRunner) GetVMIName() string {
	return rc.virtualMachineInstance
}

func (rc *KubevirtRunner) GetDataVolumeName() string {
	return rc.dataVolume
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

	jri := map[string]interface{}{
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
	if len(vmTemplate) == 0 {
		return ErrEmptyVMTemplate
	}

	if len(runnerName) == 0 {
		return ErrEmptyRunnerName
	}

	if len(jitConfig) == 0 {
		return ErrEmptyJitConfig
	}

	virtualMachineInstance, dataVolume, err := rc.getResources(ctx, vmTemplate, runnerName, jitConfig)
	if err != nil {
		return err
	}

	log.Printf("Creating %s Virtual Machine Instance\n", virtualMachineInstance.Name)

	vmi, err := rc.virtClient.VirtualMachineInstance(rc.namespace).Create(ctx,
		virtualMachineInstance, k8smetav1.CreateOptions{})
	if err != nil {
		return errors.Wrap(err, "fail to create runner instance")
	}

	rc.virtualMachineInstance = virtualMachineInstance.Name

	if dataVolume != nil {
		log.Printf("Creating %s Data Volume\n", dataVolume.Name)

		dataVolume.OwnerReferences = []k8smetav1.OwnerReference{
			{
				APIVersion: "kubevirt.io/v1",
				Kind:       "VirtualMachineInstance",
				Name:       vmi.Name,
				UID:        vmi.UID,
				Controller: ptr.To(false),
			},
		}

		if _, err := rc.virtClient.CdiClient().CdiV1beta1().DataVolumes(
			rc.namespace).Create(ctx, dataVolume, k8smetav1.CreateOptions{}); err != nil {
			return errors.Wrap(err, "cannot create data volume")
		}

		rc.dataVolume = dataVolume.Name
	}

	return nil
}

func (rc *KubevirtRunner) WaitForVirtualMachineInstance(ctx context.Context, virtualMachineInstance string) error {
	log.Printf("Watching %s Virtual Machine Instance\n", virtualMachineInstance)

	const reportingElapse = 5.0

	watch, err := rc.virtClient.VirtualMachineInstance(rc.namespace).Watch(ctx, k8smetav1.ListOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to watch the virtual machine instance")
	}
	defer watch.Stop()

	lastTimeChecked := time.Now()

	for event := range watch.ResultChan() {
		vmi, ok := event.Object.(*v1.VirtualMachineInstance)
		if ok && vmi.Name == rc.virtualMachineInstance {
			if vmi.Status.Phase != rc.currentStatus {
				rc.currentStatus = vmi.Status.Phase
				lastTimeChecked = time.Now()

				switch rc.currentStatus {
				case v1.Succeeded:
					log.Printf("%s has successfully completed\n", virtualMachineInstance)

					return nil
				case v1.Failed:
					log.Printf("%s has failed\n", virtualMachineInstance)

					return ErrRunnerFailed
				default:
					log.Printf("%s has transitioned to %s phase \n", virtualMachineInstance, rc.currentStatus)
				}
			} else if time.Since(lastTimeChecked).Minutes() > reportingElapse {
				log.Printf("%s is in %s phase \n", virtualMachineInstance, rc.currentStatus)

				lastTimeChecked = time.Now()
			}
		}
	}

	return nil
}

func (rc *KubevirtRunner) DeleteResources(ctx context.Context, virtualMachineInstance, dataVolume string) error {
	log.Printf("Cleaning %s Virtual Machine Instance resources\n",
		virtualMachineInstance)

	if err := rc.virtClient.VirtualMachineInstance(rc.namespace).Delete(
		ctx, virtualMachineInstance, k8smetav1.DeleteOptions{}); err != nil {
		if !k8serrors.IsNotFound(err) {
			log.Printf("fail to delete runner instance %s: %v", virtualMachineInstance, err)
		}
	}

	if len(dataVolume) > 0 {
		if err := rc.virtClient.CdiClient().CdiV1beta1().DataVolumes(rc.namespace).Delete(
			ctx, dataVolume, k8smetav1.DeleteOptions{}); err != nil {
			if !k8serrors.IsNotFound(err) {
				log.Printf("fail to delete runner data volume %s: %v", dataVolume, err)
			}
		}
	}

	return nil
}

func NewRunner(namespace string, virtClient kubecli.KubevirtClient) *KubevirtRunner {
	return &KubevirtRunner{
		namespace:  namespace,
		virtClient: virtClient,
	}
}
