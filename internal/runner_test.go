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

	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	k8sv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "kubevirt.io/api/core/v1"
	cdifake "kubevirt.io/client-go/containerizeddataimporter/fake"
	"kubevirt.io/client-go/kubecli"
	kubevirtfake "kubevirt.io/client-go/kubevirt/fake"
)

var _ = Describe("Runner", func() {
	var virtClient *kubecli.MockKubevirtClient
	var virtClientset *kubevirtfake.Clientset
	var karRunner runner.Runner

	BeforeEach(func() {
		cdiClientset := cdifake.NewSimpleClientset()
		virtClient = kubecli.NewMockKubevirtClient(gomock.NewController(GinkgoT()))

		virtClient.EXPECT().CdiClient().Return(cdiClientset).AnyTimes()

		karRunner = runner.NewRunner(k8sv1.NamespaceDefault, virtClient)
	})

	DescribeTable("create resources", func(shouldSucceed bool, vmTemplate, runnerName, jitConfig string) {
		vm := NewVirtualMachine(vmTemplate)
		virtClientset = kubevirtfake.NewSimpleClientset(vm)

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
			Expect(karRunner.GetVMIName()).Should(Equal(runnerName))
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
		Entry("when the valid information is provided", true, "vmTemplate", "runnerName", "jitConfig"),
		Entry("when empty vm template is provided", false, "", "runnerName", "jitConfig"),
		Entry("when empty runner name is provided", false, "vmTemplate", "", "jitConfig"),
		Entry("when empty jit config is provided", false, "vmTemplate", "runnerName", ""),
	)
})

func NewVirtualMachine(name string) *v1.VirtualMachine {
	return &v1.VirtualMachine{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: k8sv1.NamespaceDefault, ResourceVersion: "1", UID: "vm-uid"},
		Spec: v1.VirtualMachineSpec{
			Template: &v1.VirtualMachineInstanceTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
			},
		},
	}
}
