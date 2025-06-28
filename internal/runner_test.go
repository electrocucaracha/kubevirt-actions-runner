/*
Copyright © 2023

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

package runner_test

import (
	"context"
	"time"

	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	k8sv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	v1 "kubevirt.io/api/core/v1"
	cdifake "kubevirt.io/client-go/containerizeddataimporter/fake"
	"kubevirt.io/client-go/kubecli"
	kubevirtfake "kubevirt.io/client-go/kubevirt/fake"
	"kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

var _ = Describe("Runner", func() {
	var virtClient *kubecli.MockKubevirtClient
	var virtClientset *kubevirtfake.Clientset
	var karRunner runner.Runner
	var mockCtrl *gomock.Controller

	const (
		vmTemplate = "vm-template"
		vmInstance = "runner-xyz123"
		dataVolume = "dv-xyz123"
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		virtClient = kubecli.NewMockKubevirtClient(mockCtrl)
		virtClientset = kubevirtfake.NewSimpleClientset(NewVirtualMachineInstance(vmInstance), NewVirtualMachine(vmTemplate))
		cdiClientset := cdifake.NewSimpleClientset(NewDataVolume(dataVolume))

		virtClient.EXPECT().CdiClient().Return(cdiClientset).AnyTimes()

		karRunner = runner.NewRunner(k8sv1.NamespaceDefault, virtClient)
	})

	AfterEach(func() {
		mockCtrl.Finish()
		runner.CancelAppContext()
	})

	DescribeTable("create resources", func(shouldSucceed bool, vmTemplate, runnerName, jitConfig string) {
		if shouldSucceed {
			virtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
				virtClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault),
			)
			virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(
				virtClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault),
			)
		}

		err := karRunner.CreateResources(context.TODO(), vmTemplate, runnerName, jitConfig)

		if shouldSucceed {
			Expect(err).NotTo(HaveOccurred())
			appCtx := runner.GetAppContext()
			Expect(appCtx.GetVMIName()).Should(Equal(runnerName))
		} else {
			Expect(err).To(HaveOccurred())
			if len(vmTemplate) == 0 {
				Expect(err).Should(Equal(runner.ErrEmptyVMTemplate))
			}
			if len(runnerName) == 0 {
				Expect(err).Should(Equal(runner.ErrEmptyRunnerName))
			}
			if len(jitConfig) == 0 {
				Expect(err).Should(Equal(runner.ErrEmptyJitConfig))
			}
		}
	},
		Entry("when the valid information is provided", true, vmTemplate, "runnerName", "jitConfig"),
		Entry("when empty vm template is provided", false, "", "runnerName", "jitConfig"),
		Entry("when empty runner name is provided", false, vmTemplate, "", "jitConfig"),
		Entry("when empty jit config is provided", false, vmTemplate, "runnerName", ""),
	)

	DescribeTable("delete resources", func(vmInstance, dataVolume string) {
		virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(
			virtClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault),
		)
		runner.NewAppContext(vmInstance, dataVolume)

		err := karRunner.DeleteResources(context.TODO())

		Expect(err).NotTo(HaveOccurred())
	},
		Entry("when the runner has a data volume", vmInstance, dataVolume),
		Entry("when the runner doesn't have data volumes", vmInstance, ""),
		Entry("when the runner doesn't exist", "runner-abc098", ""),
		Entry("when the data volume doesn't exist", vmInstance, "dv-abc098"),
	)

	DescribeTable("watch resources", func(shouldSucceed bool, lastPhase v1.VirtualMachineInstancePhase) {
		const timeout = 1 * time.Second
		fakeWatcher := watch.NewFake()
		vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(mockCtrl)
		vmiInterface.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(fakeWatcher, nil).MinTimes(1)
		virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface).AnyTimes()
		runner.NewAppContext(vmInstance, "")

		errChan := make(chan error, 1)
		go func() {
			errChan <- karRunner.WaitForVirtualMachineInstance(context.TODO())
			close(errChan)
		}()

		phases := [5]v1.VirtualMachineInstancePhase{v1.Pending, v1.Scheduling, v1.Scheduled, v1.Running, lastPhase}
		vmi := NewVirtualMachineInstance(vmInstance)
		for _, phase := range phases {
			vmi.Status.Phase = phase
			fakeWatcher.Add(vmi)
			time.Sleep(10 * time.Millisecond)
		}

		if shouldSucceed {
			Eventually(errChan, timeout).Should(Receive(BeNil()))
		} else {
			Eventually(errChan, timeout).Should(Receive(Equal(runner.ErrRunnerFailed)))
		}
	},
		Entry("when the runner completes successfully", true, v1.Succeeded),
		Entry("when the runner completes unsuccessfully", false, v1.Failed),
	)
})

func NewVirtualMachine(name string) *v1.VirtualMachine {
	return &v1.VirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: k8sv1.NamespaceDefault,
		},
		Spec: v1.VirtualMachineSpec{
			Template: &v1.VirtualMachineInstanceTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
			},
		},
	}
}

func NewVirtualMachineInstance(name string) *v1.VirtualMachineInstance {
	return &v1.VirtualMachineInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: k8sv1.NamespaceDefault,
		},
	}
}

func NewDataVolume(name string) *v1beta1.DataVolume {
	return &v1beta1.DataVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: k8sv1.NamespaceDefault,
		},
	}
}
