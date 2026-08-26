/* jscpd:ignore-start */
/*
Copyright © 2025

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
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
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

const (
	defaultWaitTimeout = 5 * time.Minute
	consistencyTimeout = 100 * time.Millisecond
	eventuallyTimeout  = time.Second
	vmTemplateName     = "vm-template"
	vmInstanceName     = "runner-xyz123"
	dataVolumeName     = "dv-xyz123"
	kubevirtGroupName  = "kubevirt.io"
	vmiResourceName    = "virtualmachineinstances"
)

var (
	errSimulatedDataVolumeCreateFailure = errors.New("simulated data volume create failure")
	errSimulatedWatchFailure            = errors.New("simulated watch failure")
	errSimulatedTransientGetFailure     = errors.New("simulated transient get failure")
)

// panicReporter turns gomock expectation violations into panics, which godog
// recovers and reports as failed steps.
type panicReporter struct{}

func (panicReporter) Errorf(format string, args ...interface{}) { panic(fmt.Sprintf(format, args...)) }
func (panicReporter) Fatalf(format string, args ...interface{}) { panic(fmt.Sprintf(format, args...)) }

// runnerCtx holds the shared state for a single scenario.
type runnerCtx struct {
	mockCtrl      *gomock.Controller
	virtClient    *kubecli.MockKubevirtClient
	virtClientset *kubevirtfake.Clientset
	karRunner     runner.Runner

	firstWatcher  *watch.FakeWatcher
	secondWatcher *watch.FakeWatcher
	errChan       chan error
	vmi           *v1.VirtualMachineInstance

	lastRunnerName string
	result         error
}

func (rc *runnerCtx) reset(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
	rc.mockCtrl = gomock.NewController(panicReporter{})
	rc.virtClient = kubecli.NewMockKubevirtClient(rc.mockCtrl)
	rc.virtClientset = kubevirtfake.NewSimpleClientset(NewVirtualMachineInstance(vmInstanceName), NewVirtualMachine(vmTemplateName))
	cdiClientset := cdifake.NewSimpleClientset(NewDataVolume(dataVolumeName))

	rc.virtClient.EXPECT().CdiClient().Return(cdiClientset).AnyTimes()

	rc.karRunner = runner.NewRunner(k8sv1.NamespaceDefault, rc.virtClient, defaultWaitTimeout)

	rc.firstWatcher = nil
	rc.secondWatcher = nil
	rc.errChan = nil
	rc.vmi = nil
	rc.lastRunnerName = ""
	rc.result = nil

	return ctx, nil
}

func (rc *runnerCtx) after(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
	rc.mockCtrl.Finish()
	runner.CancelAppContext()

	return ctx, err
}

func (rc *runnerCtx) expectVirtualMachineAndInstance() {
	rc.virtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
		rc.virtClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault),
	)
	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(
		rc.virtClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault),
	)
}

func (rc *runnerCtx) expectVirtualMachineWithVMIInterface(vmiInterface kubecli.VirtualMachineInstanceInterface) {
	rc.virtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
		rc.virtClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault),
	)
	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface)
}

func (rc *runnerCtx) startVMIWatcherWithGet(
	getVMI func() (*v1.VirtualMachineInstance, error),
	watchers ...*watch.FakeWatcher,
) chan error {
	vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	vmiInterface.EXPECT().Get(gomock.Any(), vmInstanceName, gomock.Any()).DoAndReturn(
		func(_ context.Context, _ string, _ metav1.GetOptions) (*v1.VirtualMachineInstance, error) {
			return getVMI()
		}).AnyTimes()

	for _, watcher := range watchers {
		vmiInterface.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(watcher, nil).Times(1)
	}

	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface).AnyTimes()
	runner.NewAppContext(vmInstanceName, "")

	errChan := make(chan error, 1)
	karRunner := rc.karRunner

	go func() {
		errChan <- karRunner.WaitForVirtualMachineInstance(context.Background())

		close(errChan)
	}()

	return errChan
}

func (rc *runnerCtx) startVMIWatcher() (*watch.FakeWatcher, chan error) {
	fakeWatcher := watch.NewFake()
	errChan := rc.startVMIWatcherWithGet(func() (*v1.VirtualMachineInstance, error) {
		return NewVirtualMachineInstance(vmInstanceName), nil
	}, fakeWatcher)

	return fakeWatcher, errChan
}

func (rc *runnerCtx) startReconnectVMIWatcher() (*watch.FakeWatcher, *watch.FakeWatcher, chan error) {
	firstWatcher := watch.NewFake()
	secondWatcher := watch.NewFakeWithChanSize(1, false)
	errChan := rc.startVMIWatcherWithGet(func() (*v1.VirtualMachineInstance, error) {
		return NewVirtualMachineInstance(vmInstanceName), nil
	}, firstWatcher, secondWatcher)

	return firstWatcher, secondWatcher, errChan
}

func waitNoReceive(errChan chan error, d time.Duration) error {
	select {
	case err, ok := <-errChan:
		return fmt.Errorf("expected no result yet, got err=%v ok=%v", err, ok)
	case <-time.After(d):
		return nil
	}
}

func (rc *runnerCtx) receiveResult(timeout time.Duration) error {
	select {
	case err, ok := <-rc.errChan:
		if !ok {
			return errors.New("errChan closed unexpectedly without a value")
		}

		rc.result = err

		return nil
	case <-time.After(timeout):
		return fmt.Errorf("timed out after %s waiting for a watch result", timeout)
	}
}

// Background.
func (rc *runnerCtx) freshRunnerWithDefaultTimeout() error {
	return nil
}

// Create resources.
func (rc *runnerCtx) createResourcesWith(vmTemplate, runnerName, jitConfig string) error {
	if vmTemplate != "" && runnerName != "" && jitConfig != "" {
		rc.expectVirtualMachineAndInstance()
	}

	rc.lastRunnerName = runnerName
	rc.result = rc.karRunner.CreateResources(context.Background(), vmTemplate, k8sv1.NamespaceDefault, runnerName, jitConfig)

	return nil
}

func (rc *runnerCtx) createResultShouldBe(outcome string) error {
	if outcome == "success" {
		if rc.result != nil {
			return fmt.Errorf("expected success, got error: %w", rc.result)
		}

		if rc.lastRunnerName != "" {
			appCtx := runner.GetAppContext()
			if appCtx.GetVMIName() != rc.lastRunnerName {
				return fmt.Errorf("expected VMI name %q, got %q", rc.lastRunnerName, appCtx.GetVMIName())
			}
		}

		return nil
	}

	if rc.result == nil || rc.result.Error() != outcome {
		return fmt.Errorf("expected error %q, got %v", outcome, rc.result)
	}

	return nil
}

func (rc *runnerCtx) createWithNonexistentTemplate() error {
	rc.virtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
		rc.virtClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault))

	rc.result = rc.karRunner.CreateResources(
		context.Background(), "nonexistent-template", k8sv1.NamespaceDefault, "runnerName", "jitConfig")

	return nil
}

func (rc *runnerCtx) createButVMICreationFails() error {
	mockVMIInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	mockVMIInterface.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		nil, k8serrors.NewServiceUnavailable("simulated create failure"))

	rc.expectVirtualMachineWithVMIInterface(mockVMIInterface)

	rc.result = rc.karRunner.CreateResources(context.Background(), vmTemplateName, k8sv1.NamespaceDefault, "runner-new", "jitConfig")

	return nil
}

func (rc *runnerCtx) createWithEmptyNamespace() error {
	rc.expectVirtualMachineAndInstance()

	rc.lastRunnerName = "runner-default-ns"
	rc.result = rc.karRunner.CreateResources(context.Background(), vmTemplateName, "", "runner-default-ns", "jitConfig")

	return nil
}

func (rc *runnerCtx) createButVMIAlreadyExists() error {
	mockVMIInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	mockVMIInterface.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		nil, k8serrors.NewAlreadyExists(
			schema.GroupResource{Group: kubevirtGroupName, Resource: vmiResourceName}, "runner-existing"))

	rc.expectVirtualMachineWithVMIInterface(mockVMIInterface)

	rc.lastRunnerName = "runner-existing"
	rc.result = rc.karRunner.CreateResources(context.Background(), vmTemplateName, k8sv1.NamespaceDefault, "runner-existing", "jitConfig")

	return nil
}

func (rc *runnerCtx) createWithDataVolumeCreationFailure() error {
	const dvTemplateName = "boot-disk"

	const runnerWithDV = "runner-with-dv-failure"

	dvVM := NewVirtualMachineWithDataVolume(vmTemplateName, dvTemplateName)
	dvClientset := kubevirtfake.NewSimpleClientset(dvVM)
	failingCdiClientset := cdifake.NewSimpleClientset()
	failingCdiClientset.PrependReactor("create", "datavolumes", func(_ k8stesting.Action) (bool, runtime.Object, error) {
		return true, nil, errSimulatedDataVolumeCreateFailure
	})

	failingVirtClient := kubecli.NewMockKubevirtClient(rc.mockCtrl)
	failingVirtClient.EXPECT().CdiClient().Return(failingCdiClientset).AnyTimes()
	failingVirtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
		dvClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault))
	failingVirtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(
		dvClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault))

	failingRunner := runner.NewRunner(k8sv1.NamespaceDefault, failingVirtClient, defaultWaitTimeout)

	rc.result = failingRunner.CreateResources(context.Background(), vmTemplateName, k8sv1.NamespaceDefault, runnerWithDV, "jitConfig")

	return nil
}

func (rc *runnerCtx) createWithDataVolumeSuccess() error {
	const dvTemplateName = "boot-disk"

	const runnerWithDV = "runner-with-dv"

	dvVM := NewVirtualMachineWithDataVolume(vmTemplateName, dvTemplateName)
	dvClientset := kubevirtfake.NewSimpleClientset(dvVM)

	rc.virtClient.EXPECT().VirtualMachine(k8sv1.NamespaceDefault).Return(
		dvClientset.KubevirtV1().VirtualMachines(k8sv1.NamespaceDefault))
	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(
		rc.virtClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault))

	rc.result = rc.karRunner.CreateResources(context.Background(), vmTemplateName, k8sv1.NamespaceDefault, runnerWithDV, "jitConfig")

	return nil
}

func (rc *runnerCtx) appContextDataVolumeShouldContain(sub string) error {
	appCtx := runner.GetAppContext()
	if !strings.Contains(appCtx.GetDataVolumeName(), sub) {
		return fmt.Errorf("expected data volume name to contain %q, got %q", sub, appCtx.GetDataVolumeName())
	}

	return nil
}

// Delete resources.
func (rc *runnerCtx) appContextInitialized(vmi, dataVolume string) error {
	runner.NewAppContext(vmi, dataVolume)

	return nil
}

func (rc *runnerCtx) appContextInitializedNoDV(vmi string) error {
	return rc.appContextInitialized(vmi, "")
}

func (rc *runnerCtx) appContextNotInitialized() error {
	runner.CancelAppContext()

	return nil
}

func (rc *runnerCtx) resourcesAreDeleted() error {
	if runner.HasAppContext() {
		rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(
			rc.virtClientset.KubevirtV1().VirtualMachineInstances(k8sv1.NamespaceDefault),
		)
	}

	rc.result = rc.karRunner.DeleteResources(context.Background())

	return nil
}

func (rc *runnerCtx) deleteButVMIDeleteForbidden() error {
	forbiddenErr := k8serrors.NewForbidden(
		schema.GroupResource{Group: kubevirtGroupName, Resource: vmiResourceName}, vmInstanceName, nil)

	mockVMIInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	mockVMIInterface.EXPECT().Delete(gomock.Any(), vmInstanceName, gomock.Any()).Return(forbiddenErr)
	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(mockVMIInterface)

	rc.result = rc.karRunner.DeleteResources(context.Background())

	return nil
}

func (rc *runnerCtx) deleteButDataVolumeDeleteForbidden() error {
	forbiddenErr := k8serrors.NewForbidden(
		schema.GroupResource{Group: "cdi.kubevirt.io", Resource: "datavolumes"}, dataVolumeName, nil)

	failingCdiClientset := cdifake.NewSimpleClientset(NewDataVolume(dataVolumeName))
	failingCdiClientset.PrependReactor("delete", "datavolumes", func(_ k8stesting.Action) (bool, runtime.Object, error) {
		return true, nil, forbiddenErr
	})

	mockVMIInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	mockVMIInterface.EXPECT().Delete(gomock.Any(), vmInstanceName, gomock.Any()).Return(nil)

	failingVirtClient := kubecli.NewMockKubevirtClient(rc.mockCtrl)
	failingVirtClient.EXPECT().CdiClient().Return(failingCdiClientset).AnyTimes()
	failingVirtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(mockVMIInterface)

	failingRunner := runner.NewRunner(k8sv1.NamespaceDefault, failingVirtClient, defaultWaitTimeout)

	rc.result = failingRunner.DeleteResources(context.Background())

	return nil
}

// Generic assertions.
func (rc *runnerCtx) operationShouldSucceed() error {
	if rc.result != nil {
		return fmt.Errorf("expected success, got error: %w", rc.result)
	}

	return nil
}

func (rc *runnerCtx) operationShouldFailContaining(sub string) error {
	if rc.result == nil {
		return fmt.Errorf("expected error containing %q, got nil", sub)
	}

	if !strings.Contains(rc.result.Error(), sub) {
		return fmt.Errorf("expected error containing %q, got %v", sub, rc.result)
	}

	return nil
}

// Watching resources.
func (rc *runnerCtx) givenWatchedVMI() error {
	fakeWatcher, errChan := rc.startVMIWatcher()
	rc.firstWatcher = fakeWatcher
	rc.errChan = errChan

	return nil
}

func (rc *runnerCtx) vmiReachesPhase(lastPhase string) error {
	phase := v1.VirtualMachineInstancePhase(lastPhase)
	vmi := NewVirtualMachineInstance(vmInstanceName)

	for _, p := range []v1.VirtualMachineInstancePhase{v1.Pending, v1.Scheduling, v1.Scheduled, v1.Running, phase} {
		vmi.Status.Phase = p
		rc.firstWatcher.Add(vmi)
		time.Sleep(10 * time.Millisecond)
	}

	return rc.receiveResult(eventuallyTimeout)
}

func (rc *runnerCtx) watchResultShouldBe(outcome string) error {
	if outcome == "success" {
		if rc.result != nil {
			return fmt.Errorf("expected nil error, got %v", rc.result)
		}

		return nil
	}

	if rc.result == nil || rc.result.Error() != runner.ErrRunnerFailed.Error() {
		return fmt.Errorf("expected runner failed error, got %v", rc.result)
	}

	return nil
}

func (rc *runnerCtx) givenWatchedVMIBecomesRunningAndReady() error {
	fakeWatcher, errChan := rc.startVMIWatcher()
	rc.firstWatcher = fakeWatcher
	rc.errChan = errChan

	vmi := NewVirtualMachineInstance(vmInstanceName)
	for _, p := range []v1.VirtualMachineInstancePhase{v1.Pending, v1.Scheduling, v1.Scheduled} {
		vmi.Status.Phase = p
		fakeWatcher.Add(vmi)
	}

	vmi.Status.Phase = v1.Running
	fakeWatcher.Add(vmi)

	rc.vmi = NewVirtualMachineInstanceReady(vmInstanceName)
	fakeWatcher.Modify(rc.vmi)

	return nil
}

func (rc *runnerCtx) givenInitialGetReportsReady() error {
	fakeWatcher := watch.NewFake()
	readyVMI := NewVirtualMachineInstanceReady(vmInstanceName)

	errChan := rc.startVMIWatcherWithGet(func() (*v1.VirtualMachineInstance, error) {
		return readyVMI, nil
	}, fakeWatcher)

	rc.firstWatcher = fakeWatcher
	rc.errChan = errChan
	rc.vmi = readyVMI

	return nil
}

// vmiBecomesSucceededMilestone verifies that Running+Ready is only a
// milestone before Succeeded, which is what actually ends the watch.
func (rc *runnerCtx) vmiBecomesSucceededMilestone() error {
	if err := waitNoReceive(rc.errChan, consistencyTimeout); err != nil {
		return err
	}

	rc.vmi.Status.Phase = v1.Succeeded
	rc.firstWatcher.Modify(rc.vmi)

	return rc.receiveResult(eventuallyTimeout)
}

func (rc *runnerCtx) vmiTransitionsToSucceeded() error {
	vmi := NewVirtualMachineInstance(vmInstanceName)
	vmi.Status.Phase = v1.Succeeded
	rc.firstWatcher.Add(vmi)

	return rc.receiveResult(eventuallyTimeout)
}

func (rc *runnerCtx) emitUnrelatedPodEvent() error {
	pod := &k8sv1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "some-pod", Namespace: k8sv1.NamespaceDefault}}
	rc.firstWatcher.Add(pod)

	return nil
}

func (rc *runnerCtx) emitDifferentNamedFailedEvent() error {
	otherVMI := NewVirtualMachineInstance("other-vmi-name")
	otherVMI.Status.Phase = v1.Failed
	rc.firstWatcher.Add(otherVMI)

	return nil
}

func (rc *runnerCtx) emitUnrecognizedPhase() error {
	vmi := NewVirtualMachineInstance(vmInstanceName)
	vmi.Status.Phase = v1.VirtualMachineInstancePhase("UnrecognizedPhase")
	rc.firstWatcher.Add(vmi)

	return nil
}

func (rc *runnerCtx) givenShortWaitTimeoutRunner(kind string) error {
	waitTimeout := 100 * time.Millisecond
	if kind == "very short" {
		waitTimeout = 10 * time.Millisecond
	}

	rc.karRunner = runner.NewRunner(k8sv1.NamespaceDefault, rc.virtClient, waitTimeout)

	return nil
}

func (rc *runnerCtx) givenWatchedVMIStaysRunning() error {
	fakeWatcher, errChan := rc.startVMIWatcher()
	rc.firstWatcher = fakeWatcher
	rc.errChan = errChan

	vmi := NewVirtualMachineInstance(vmInstanceName)
	vmi.Status.Phase = v1.Running
	fakeWatcher.Add(vmi)

	return nil
}

func (rc *runnerCtx) whenWaitTimeoutElapses() error {
	return rc.receiveResult(eventuallyTimeout)
}

func (rc *runnerCtx) watchShouldTimeOut() error {
	if rc.result == nil {
		return errors.New("expected a timeout error, got nil")
	}

	if rc.result.Error() != "timeout while waiting for the virtual machine instance" {
		return fmt.Errorf("expected timeout error, got %v", rc.result)
	}

	return nil
}

func (rc *runnerCtx) givenAlwaysClosedWatchForRunningVMI() error {
	vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	vmiInterface.EXPECT().Get(gomock.Any(), vmInstanceName, gomock.Any()).DoAndReturn(
		func(_ context.Context, _ string, _ metav1.GetOptions) (*v1.VirtualMachineInstance, error) {
			vmi := NewVirtualMachineInstance(vmInstanceName)
			vmi.Status.Phase = v1.Running

			return vmi, nil
		}).AnyTimes()
	vmiInterface.EXPECT().Watch(gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, _ metav1.ListOptions) (watch.Interface, error) {
			fakeWatcher := watch.NewFake()
			fakeWatcher.Stop()

			return fakeWatcher, nil
		}).AnyTimes()

	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface).AnyTimes()
	runner.NewAppContext(vmInstanceName, "")

	errChan := make(chan error, 1)
	karRunner := rc.karRunner

	go func() {
		errChan <- karRunner.WaitForVirtualMachineInstance(context.Background())

		close(errChan)
	}()
	rc.errChan = errChan

	return nil
}

func (rc *runnerCtx) waitInvokedWithCancelledContext() error {
	vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	rc.result = rc.karRunner.WaitForVirtualMachineInstance(ctx)

	return nil
}

func (rc *runnerCtx) waitButWatchFails() error {
	vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	vmiInterface.EXPECT().Get(gomock.Any(), vmInstanceName, gomock.Any()).Return(
		NewVirtualMachineInstance(vmInstanceName), nil).AnyTimes()
	vmiInterface.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(nil, errSimulatedWatchFailure)

	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface).AnyTimes()

	rc.result = rc.karRunner.WaitForVirtualMachineInstance(context.Background())

	return nil
}

func (rc *runnerCtx) waitContextCancelledDuringFailingGet() error {
	ctx, cancel := context.WithCancel(context.Background())

	vmiInterface := kubecli.NewMockVirtualMachineInstanceInterface(rc.mockCtrl)
	vmiInterface.EXPECT().Get(gomock.Any(), vmInstanceName, gomock.Any()).DoAndReturn(
		func(_ context.Context, _ string, _ metav1.GetOptions) (*v1.VirtualMachineInstance, error) {
			// Simulate the context being cancelled concurrently with the API
			// call failing, exercising the race between ctx.Done() and the
			// Get error path.
			cancel()

			return nil, errSimulatedTransientGetFailure
		}).AnyTimes()

	rc.virtClient.EXPECT().VirtualMachineInstance(k8sv1.NamespaceDefault).Return(vmiInterface).AnyTimes()

	rc.result = rc.karRunner.WaitForVirtualMachineInstance(ctx)

	return nil
}

func (rc *runnerCtx) givenReconnectWatch() error {
	first, second, errChan := rc.startReconnectVMIWatcher()
	rc.firstWatcher = first
	rc.secondWatcher = second
	rc.errChan = errChan

	return nil
}

func (rc *runnerCtx) firstStreamClosesAfterRunning() error {
	vmi := NewVirtualMachineInstance(vmInstanceName)
	vmi.Status.Phase = v1.Running
	rc.firstWatcher.Add(vmi)
	rc.firstWatcher.Stop()

	return waitNoReceive(rc.errChan, consistencyTimeout)
}

func (rc *runnerCtx) runningReadyThenFirstStreamCloses() error {
	readyVMI := NewVirtualMachineInstanceReady(vmInstanceName)
	rc.firstWatcher.Modify(readyVMI)
	rc.firstWatcher.Stop()
	rc.vmi = readyVMI

	return waitNoReceive(rc.errChan, consistencyTimeout)
}

func (rc *runnerCtx) secondStreamVMIBecomesSucceeded() error {
	vmi := rc.vmi
	if vmi == nil {
		vmi = NewVirtualMachineInstance(vmInstanceName)
	}

	vmi.Status.Phase = v1.Succeeded
	rc.secondWatcher.Modify(vmi)

	return rc.receiveResult(3 * time.Second)
}

func (rc *runnerCtx) vmiDisappearsBeforeReestablish() error {
	firstWatcher := watch.NewFake()
	getCalls := 0
	errChan := rc.startVMIWatcherWithGet(func() (*v1.VirtualMachineInstance, error) {
		getCalls++
		if getCalls > 1 {
			return nil, k8serrors.NewNotFound(
				schema.GroupResource{Group: "kubevirt.io", Resource: "virtualmachineinstances"}, vmInstanceName)
		}

		return NewVirtualMachineInstance(vmInstanceName), nil
	}, firstWatcher)

	firstWatcher.Stop()
	rc.errChan = errChan

	return rc.receiveResult(3 * time.Second)
}

//nolint:funlen // Registration wires every step of the runner feature file.
func InitializeScenario(sc *godog.ScenarioContext) {
	rc := &runnerCtx{}

	sc.Before(rc.reset)
	sc.After(rc.after)

	sc.Step(`^a fresh KubeVirt runner with the default wait timeout$`, rc.freshRunnerWithDefaultTimeout)

	sc.Step(`^resources are created with vm template "([^"]*)", runner name "([^"]*)" and jit config "([^"]*)"$`,
		rc.createResourcesWith)
	sc.Step(`^the create result should be "([^"]*)"$`, rc.createResultShouldBe)
	sc.Step(`^resources are created referencing a nonexistent vm template$`, rc.createWithNonexistentTemplate)
	sc.Step(`^resources are created but the VMI creation call fails$`, rc.createButVMICreationFails)
	sc.Step(`^resources are created with an empty vm template namespace$`, rc.createWithEmptyNamespace)
	sc.Step(`^resources are created but the VMI already exists$`, rc.createButVMIAlreadyExists)
	sc.Step(`^resources with a data volume template are created but the data volume creation fails$`,
		rc.createWithDataVolumeCreationFailure)
	sc.Step(`^resources with a data volume template are created successfully$`, rc.createWithDataVolumeSuccess)
	sc.Step(`^the created app context data volume name should contain "([^"]*)"$`, rc.appContextDataVolumeShouldContain)

	sc.Step(`^the app context is initialized with vmi "([^"]*)" and data volume "([^"]*)"$`, rc.appContextInitialized)
	sc.Step(`^the app context is initialized with vmi "([^"]*)" and no data volume$`, rc.appContextInitializedNoDV)
	sc.Step(`^the app context is not initialized$`, rc.appContextNotInitialized)
	sc.Step(`^resources are deleted$`, rc.resourcesAreDeleted)
	sc.Step(`^resources are deleted but the VMI delete call fails with a forbidden error$`, rc.deleteButVMIDeleteForbidden)
	sc.Step(`^resources are deleted but the data volume delete call fails with a forbidden error$`,
		rc.deleteButDataVolumeDeleteForbidden)

	sc.Step(`^the operation should succeed$`, rc.operationShouldSucceed)
	sc.Step(`^the operation should fail with an error containing "([^"]*)"$`, rc.operationShouldFailContaining)

	sc.Step(`^a watched vmi$`, rc.givenWatchedVMI)
	sc.Step(`^the vmi reaches phase "([^"]*)"$`, rc.vmiReachesPhase)
	sc.Step(`^the watch result should be "([^"]*)"$`, rc.watchResultShouldBe)
	sc.Step(`^an unrelated pod event is emitted$`, rc.emitUnrelatedPodEvent)
	sc.Step(`^a failed event for a different vmi name is emitted$`, rc.emitDifferentNamedFailedEvent)
	sc.Step(`^the vmi reports an unrecognized phase$`, rc.emitUnrecognizedPhase)
	sc.Step(`^the vmi transitions to Succeeded$`, rc.vmiTransitionsToSucceeded)

	sc.Step(`^a watched vmi that becomes Running and Ready$`, rc.givenWatchedVMIBecomesRunningAndReady)
	sc.Step(`^a watched vmi whose initial Get already reports Running and Ready$`, rc.givenInitialGetReportsReady)
	sc.Step(`^the vmi becomes Succeeded$`, rc.vmiBecomesSucceededMilestone)

	sc.Step(`^a runner with a (short|very short) wait timeout$`, rc.givenShortWaitTimeoutRunner)
	sc.Step(`^a watched vmi that stays Running$`, rc.givenWatchedVMIStaysRunning)
	sc.Step(`^the wait timeout elapses$`, rc.whenWaitTimeoutElapses)
	sc.Step(`^the watch should time out$`, rc.watchShouldTimeOut)
	sc.Step(`^a watch that always returns an already closed stream for a running vmi$`, rc.givenAlwaysClosedWatchForRunningVMI)
	sc.Step(`^the wait is invoked with an already cancelled context$`, rc.waitInvokedWithCancelledContext)
	sc.Step(`^the wait is invoked but the Watch call fails$`, rc.waitButWatchFails)
	sc.Step(`^the wait context is cancelled during a failing Get call$`, rc.waitContextCancelledDuringFailingGet)

	sc.Step(`^a watch that will be re-established after the first stream closes$`, rc.givenReconnectWatch)
	sc.Step(`^the first watch stream closes after the vmi becomes Running$`, rc.firstStreamClosesAfterRunning)
	sc.Step(`^the vmi becomes Running and Ready then the first watch stream closes$`, rc.runningReadyThenFirstStreamCloses)
	sc.Step(`^the vmi becomes Succeeded on the second watch stream$`, rc.secondStreamVMIBecomesSucceeded)
	sc.Step(`^the vmi becomes Running then the first watch stream closes and the vmi is no longer found$`,
		rc.vmiDisappearsBeforeReestablish)
}

func TestFeatures(t *testing.T) {
	t.Parallel()

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

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
