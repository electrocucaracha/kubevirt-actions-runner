/* jscpd:ignore-start */
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
/* jscpd:ignore-end */

package runner_test

import (
	"context"
	"errors"
	"time"

	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	k8sv1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime" //nolint:depguard // required by fake reactor signature
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	k8stesting "k8s.io/client-go/testing" //nolint:depguard // required by fake reactor signature
	v1 "kubevirt.io/api/core/v1"
	cdifake "kubevirt.io/client-go/containerizeddataimporter/fake"
	"kubevirt.io/client-go/kubecli"
	kubevirtfake "kubevirt.io/client-go/kubevirt/fake"
	"kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

var errSimulatedDataVolumeCreateFailure = errors.New("simulated data volume create failure")

var _ = Describe("Runner", func() {
	var virtClient *kubecli.MockKubevirtClient

	var virtClientset *kubevirtfake.Clientset

	var karRunner runner.Runner

	var mockCtrl *gomock.Controller

	const (
		defaultWaitTimeout = 5 * time.Minute
		consistencyTimeout = 100 * time.Millisecond
		eventuallyTimeout  = time.Second
		vmTemplate         = "vm-template"
		vmInstance         = "runner-xyz123"
		dataVolume         = "dv-xyz123"
		kubevirtGroup      = "kubevirt.io"
		vmiResource        = "virtualmachineinstances"
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		virtClient = kubecli.NewMockKubevirtClient(mockCtrl)
		virtClientset = kubevirtfake.NewSimpleClientset(NewVirtualMachineInstance(vmInstance), NewVirtualMachine(vmTemplate))
		cdiClientset := cdifake.NewSimpleClientset(NewDataVolume(dataVolume))

		virtClient.EXPECT().CdiClient().Return(cdiClientset).AnyTimes()

		karRunner = runner.NewRunner(k8sv1.NamespaceDefault, virtClient, defaultWaitTimeout)
	})

	AfterEach(func() {
		mockCtrl.Finish()
		runner.CancelAppContext()
	})

	startVMIWatcherWithGet := func(
		karRunner runner.Runner,
		getVMI func() (*v1.VirtualMachineInstance, error),
		watchers ...*watch.FakeWatcher,
	) chan error {
		vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(mockCtrl)
		vmiInterface.EXPECT().Get(gomock.Any(), vmInstance, gomock.Any()).DoAndReturn(
			func(_ context.Context, _ string, _ metav1.GetOptions) (*v1.VirtualMachineInstance, error) {
				return getVMI()
			}).AnyTimes()

		for _, watcher := range watchers {
			vmiInterface.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(watcher, nil).Times(1)
		}

		virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface).AnyTimes()
		runner.NewAppContext(vmInstance, "")

		errChan := make(chan error, 1)

		go func() {
			errChan <- karRunner.WaitForVirtualMachineInstance(context.TODO())

			close(errChan)
		}()

		return errChan
	}

	startVMIWatcher := func(karRunner runner.Runner) (*watch.FakeWatcher, chan error) {
		fakeWatcher := watch.NewFake()
		errChan := startVMIWatcherWithGet(karRunner, func() (*v1.VirtualMachineInstance, error) {
			return NewVirtualMachineInstance(vmInstance), nil
		}, fakeWatcher)

		return fakeWatcher, errChan
	}

	startReconnectVMIWatcher := func(karRunner runner.Runner) (*watch.FakeWatcher, *watch.FakeWatcher, chan error) {
		firstWatcher := watch.NewFake()
		secondWatcher := watch.NewFakeWithChanSize(1, false)
		errChan := startVMIWatcherWithGet(karRunner, func() (*v1.VirtualMachineInstance, error) {
			return NewVirtualMachineInstance(vmInstance), nil
		}, firstWatcher, secondWatcher)

		return firstWatcher, secondWatcher, errChan
	}

	waitForWatchCompletion := func(
		errChan chan error,
		timeout time.Duration,
		readyVMI *v1.VirtualMachineInstance,
		emitEvent func(*v1.VirtualMachineInstance),
	) {
		Consistently(errChan, consistencyTimeout).ShouldNot(Receive())

		readyVMI.Status.Phase = v1.Succeeded
		emitEvent(readyVMI)

		Eventually(errChan, timeout).Should(Receive(BeNil()))
	}

	expectDefaultNamespaceClients := func(vmiInterface kubecli.VirtualMachineInstanceInterface) {
		virtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
			virtClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault),
		)
		virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface)
	}

	DescribeTable("create resources", func(shouldSucceed bool, vmTemplate, runnerName, jitConfig string) {
		if shouldSucceed {
			expectDefaultNamespaceClients(virtClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault))
		}

		err := karRunner.CreateResources(context.TODO(), vmTemplate, k8sv1.NamespaceDefault, runnerName, jitConfig)

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

	It("delete resources does nothing when AppContext is not initialized", func() {
		// Ensure AppContext is not initialized (AfterEach calls CancelAppContext,
		// but be explicit here for clarity).
		runner.CancelAppContext()

		err := karRunner.DeleteResources(context.TODO())

		Expect(err).NotTo(HaveOccurred())
	})

	DescribeTable("watch resources", func(shouldSucceed bool, lastPhase v1.VirtualMachineInstancePhase) {
		const timeout = eventuallyTimeout

		fakeWatcher, errChan := startVMIWatcher(karRunner)

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

	It("logs Running+Ready as a milestone and succeeds when VMI reaches Succeeded", func() {
		const timeout = eventuallyTimeout

		fakeWatcher, errChan := startVMIWatcher(karRunner)

		vmi := NewVirtualMachineInstance(vmInstance)
		for _, phase := range []v1.VirtualMachineInstancePhase{v1.Pending, v1.Scheduling, v1.Scheduled} {
			vmi.Status.Phase = phase
			fakeWatcher.Add(vmi)
		}

		vmi.Status.Phase = v1.Running
		fakeWatcher.Add(vmi)

		readyVMI := NewVirtualMachineInstanceReady(vmInstance)
		fakeWatcher.Modify(readyVMI)

		// Running+Ready is only a milestone; the watcher must continue until Succeeded.
		waitForWatchCompletion(errChan, timeout, readyVMI, func(vmi *v1.VirtualMachineInstance) {
			fakeWatcher.Modify(vmi)
		})
	})

	It("times out when no terminal VMI phase is observed within the wait timeout", func() {
		const waitTimeout = 100 * time.Millisecond

		const timeout = eventuallyTimeout

		shortTimeoutRunner := runner.NewRunner(k8sv1.NamespaceDefault, virtClient, waitTimeout)
		fakeWatcher, errChan := startVMIWatcher(shortTimeoutRunner)

		vmi := NewVirtualMachineInstance(vmInstance)
		vmi.Status.Phase = v1.Running
		fakeWatcher.Add(vmi)

		Eventually(errChan, timeout).Should(Receive(MatchError("timeout while waiting for the virtual machine instance")))
	})

	It("re-establishes the VMI watch when the watch stream closes", func() {
		const timeout = 3 * time.Second

		firstWatcher, secondWatcher, errChan := startReconnectVMIWatcher(karRunner)

		vmi := NewVirtualMachineInstance(vmInstance)
		vmi.Status.Phase = v1.Running
		firstWatcher.Add(vmi)
		firstWatcher.Stop()

		Consistently(errChan, consistencyTimeout).ShouldNot(Receive())

		vmi.Status.Phase = v1.Succeeded
		secondWatcher.Modify(vmi)

		Eventually(errChan, timeout).Should(Receive(BeNil()))
	})

	It("re-establishes the VMI watch after the Running+Ready milestone", func() {
		const timeout = 3 * time.Second

		firstWatcher, secondWatcher, errChan := startReconnectVMIWatcher(karRunner)

		readyVMI := NewVirtualMachineInstanceReady(vmInstance)
		firstWatcher.Modify(readyVMI)
		firstWatcher.Stop()

		Consistently(errChan, consistencyTimeout).ShouldNot(Receive())

		readyVMI.Status.Phase = v1.Succeeded
		secondWatcher.Modify(readyVMI)

		Eventually(errChan, timeout).Should(Receive(BeNil()))
	})

	It("uses the wait timeout as the upper bound for closed watch streams", func() {
		const waitTimeout = 10 * time.Millisecond

		const timeout = eventuallyTimeout

		shortTimeoutRunner := runner.NewRunner(k8sv1.NamespaceDefault, virtClient, waitTimeout)
		vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(mockCtrl)
		vmiInterface.EXPECT().Get(gomock.Any(), vmInstance, gomock.Any()).DoAndReturn(
			func(_ context.Context, _ string, _ metav1.GetOptions) (*v1.VirtualMachineInstance, error) {
				vmi := NewVirtualMachineInstance(vmInstance)
				vmi.Status.Phase = v1.Running

				return vmi, nil
			}).AnyTimes()
		vmiInterface.EXPECT().Watch(gomock.Any(), gomock.Any()).DoAndReturn(
			func(_ context.Context, _ metav1.ListOptions) (watch.Interface, error) {
				fakeWatcher := watch.NewFake()
				fakeWatcher.Stop()

				return fakeWatcher, nil
			}).AnyTimes()

		virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface).AnyTimes()
		runner.NewAppContext(vmInstance, "")

		errChan := make(chan error, 1)
		go func() {
			errChan <- shortTimeoutRunner.WaitForVirtualMachineInstance(context.TODO())

			close(errChan)
		}()

		Eventually(errChan, timeout).Should(Receive(MatchError("timeout while waiting for the virtual machine instance")))
	})

	It("exits immediately when the context is already cancelled on entry", func() {
		vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(mockCtrl)
		virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface)
		runner.NewAppContext(vmInstance, "")

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := karRunner.WaitForVirtualMachineInstance(ctx)

		Expect(err).To(MatchError("timeout while waiting for the virtual machine instance"))
	})

	It("reports Running+Ready milestone when initial VMI Get returns a ready VMI", func() {
		const timeout = eventuallyTimeout

		fakeWatcher := watch.NewFake()
		readyVMI := NewVirtualMachineInstanceReady(vmInstance)

		errChan := startVMIWatcherWithGet(karRunner, func() (*v1.VirtualMachineInstance, error) {
			return readyVMI, nil
		}, fakeWatcher)

		waitForWatchCompletion(errChan, timeout, readyVMI, func(vmi *v1.VirtualMachineInstance) {
			fakeWatcher.Modify(vmi)
		})
	})

	It("ignores non-VMI events in the watch stream", func() {
		const timeout = eventuallyTimeout

		fakeWatcher, errChan := startVMIWatcher(karRunner)

		// A Pod event should be skipped; the watcher must still process the VMI.
		pod := &k8sv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "some-pod", Namespace: k8sv1.NamespaceDefault}}
		fakeWatcher.Add(pod)

		vmi := NewVirtualMachineInstance(vmInstance)
		vmi.Status.Phase = v1.Succeeded
		fakeWatcher.Add(vmi)

		Eventually(errChan, timeout).Should(Receive(BeNil()))
	})

	It("ignores watch events from VMIs with a different name", func() {
		const timeout = eventuallyTimeout

		fakeWatcher, errChan := startVMIWatcher(karRunner)

		// An event for a different VMI must not affect the watch outcome.
		otherVMI := NewVirtualMachineInstance("other-vmi-name")
		otherVMI.Status.Phase = v1.Failed
		fakeWatcher.Add(otherVMI)

		vmi := NewVirtualMachineInstance(vmInstance)
		vmi.Status.Phase = v1.Succeeded
		fakeWatcher.Add(vmi)

		Eventually(errChan, timeout).Should(Receive(BeNil()))
	})

	It("handles unrecognized VMI phases as a no-op and continues watching", func() {
		const timeout = eventuallyTimeout

		fakeWatcher, errChan := startVMIWatcher(karRunner)

		vmi := NewVirtualMachineInstance(vmInstance)
		vmi.Status.Phase = v1.VirtualMachineInstancePhase("UnrecognizedPhase")
		fakeWatcher.Add(vmi)

		vmi.Status.Phase = v1.Succeeded
		fakeWatcher.Add(vmi)

		Eventually(errChan, timeout).Should(Receive(BeNil()))
	})

	It("returns an error when VMI creation fails", func() {
		mockVMIInterface := kubecli.NewMockVirtualMachineInstanceInterface(mockCtrl)
		mockVMIInterface.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(
			nil, k8serrors.NewServiceUnavailable("simulated create failure"))

		expectDefaultNamespaceClients(mockVMIInterface)

		err := karRunner.CreateResources(context.TODO(), vmTemplate, k8sv1.NamespaceDefault, "runner-new", "jitConfig")

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(ContainSubstring("failed to create runner instance")))
	})

	It("defaults the vm template namespace when it is empty", func() {
		expectDefaultNamespaceClients(virtClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault))

		err := karRunner.CreateResources(context.TODO(), vmTemplate, "", "runner-default-ns", "jitConfig")

		Expect(err).NotTo(HaveOccurred())

		appCtx := runner.GetAppContext()
		Expect(appCtx.GetVMIName()).Should(Equal("runner-default-ns"))
	})

	It("succeeds when the VMI already exists", func() {
		mockVMIInterface := kubecli.NewMockVirtualMachineInstanceInterface(mockCtrl)
		mockVMIInterface.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(
			nil, k8serrors.NewAlreadyExists(
				schema.GroupResource{Group: kubevirtGroup, Resource: vmiResource}, "runner-existing"))

		expectDefaultNamespaceClients(mockVMIInterface)

		err := karRunner.CreateResources(context.TODO(), vmTemplate, k8sv1.NamespaceDefault, "runner-existing", "jitConfig")

		Expect(err).NotTo(HaveOccurred())

		appCtx := runner.GetAppContext()
		Expect(appCtx.GetVMIName()).Should(Equal("runner-existing"))
	})

	It("returns an error when the data volume creation fails", func() {
		const dvTemplateName = "boot-disk"

		const runnerWithDV = "runner-with-dv-failure"

		dvVM := NewVirtualMachineWithDataVolume(vmTemplate, dvTemplateName)
		dvClientset := kubevirtfake.NewSimpleClientset(dvVM)
		failingCdiClientset := cdifake.NewSimpleClientset()
		failingCdiClientset.PrependReactor("create", "datavolumes", func(_ k8stesting.Action) (bool, runtime.Object, error) {
			return true, nil, errSimulatedDataVolumeCreateFailure
		})

		failingVirtClient := kubecli.NewMockKubevirtClient(mockCtrl)
		failingVirtClient.EXPECT().CdiClient().Return(failingCdiClientset).AnyTimes()
		failingVirtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
			dvClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault))
		failingVirtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(
			dvClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault))

		failingRunner := runner.NewRunner(k8sv1.NamespaceDefault, failingVirtClient, defaultWaitTimeout)

		err := failingRunner.CreateResources(context.TODO(), vmTemplate, k8sv1.NamespaceDefault, runnerWithDV, "jitConfig")

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(ContainSubstring("cannot create data volume")))
	})

	It("creates resources that include a data volume template", func() {
		const dvTemplateName = "boot-disk"

		const runnerWithDV = "runner-with-dv"

		dvVM := NewVirtualMachineWithDataVolume(vmTemplate, dvTemplateName)
		dvClientset := kubevirtfake.NewSimpleClientset(dvVM)

		virtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
			dvClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault))
		virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(
			virtClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault))

		err := karRunner.CreateResources(context.TODO(), vmTemplate, k8sv1.NamespaceDefault, runnerWithDV, "jitConfig")

		Expect(err).NotTo(HaveOccurred())

		appCtx := runner.GetAppContext()
		Expect(appCtx.GetDataVolumeName()).To(ContainSubstring(dvTemplateName))
	})

	It("logs but does not return an error when VMI delete fails with a non-NotFound error", func() {
		forbiddenErr := k8serrors.NewForbidden(
			schema.GroupResource{Group: kubevirtGroup, Resource: vmiResource},
			vmInstance, nil)

		mockVMIInterface := kubecli.NewMockVirtualMachineInstanceInterface(mockCtrl)
		mockVMIInterface.EXPECT().Delete(gomock.Any(), vmInstance, gomock.Any()).Return(forbiddenErr)
		virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(mockVMIInterface)
		runner.NewAppContext(vmInstance, "")

		err := karRunner.DeleteResources(context.TODO())

		Expect(err).NotTo(HaveOccurred())
	})

	It("fails when the VMI disappears before a watch can be re-established", func() {
		firstWatcher := watch.NewFake()
		getCalls := 0
		errChan := startVMIWatcherWithGet(karRunner, func() (*v1.VirtualMachineInstance, error) {
			getCalls++
			if getCalls > 1 {
				err := k8serrors.NewNotFound(
					schema.GroupResource{Group: "kubevirt.io", Resource: "virtualmachineinstances"}, vmInstance)

				return nil, err
			}

			return NewVirtualMachineInstance(vmInstance), nil
		}, firstWatcher)

		firstWatcher.Stop()

		Eventually(errChan, 3*time.Second).Should(
			Receive(MatchError(ContainSubstring("failed to get the virtual machine instance"))))
	})
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

func NewVirtualMachineInstanceReady(name string) *v1.VirtualMachineInstance {
	return &v1.VirtualMachineInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: k8sv1.NamespaceDefault,
		},
		Status: v1.VirtualMachineInstanceStatus{
			Phase: v1.Running,
			Conditions: []v1.VirtualMachineInstanceCondition{
				{
					Type:   v1.VirtualMachineInstanceReady,
					Status: k8sv1.ConditionTrue,
				},
			},
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

func NewVirtualMachineWithDataVolume(name, dvName string) *v1.VirtualMachine {
	return &v1.VirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: k8sv1.NamespaceDefault,
		},
		Spec: v1.VirtualMachineSpec{
			DataVolumeTemplates: []v1.DataVolumeTemplateSpec{
				{
					ObjectMeta: metav1.ObjectMeta{Name: dvName},
				},
			},
			Template: &v1.VirtualMachineInstanceTemplateSpec{
				Spec: v1.VirtualMachineInstanceSpec{
					Volumes: []v1.Volume{
						{
							Name: "disk0",
							VolumeSource: v1.VolumeSource{
								DataVolume: &v1.DataVolumeSource{
									Name: dvName,
								},
							},
						},
					},
				},
			},
		},
	}
}
