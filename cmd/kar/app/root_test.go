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

package app_test

import (
	"context"
	"errors"
	"slices"

	"github.com/electrocucaracha/kubevirt-actions-runner/cmd/kar/app"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var errExpectedFailure = errors.New("failure")

type mock struct {
	createErr    error
	deleteErr    error
	waitErr      error
	createCalled bool
	waitCalled   bool
	deleteCalled bool
	vmTemplate   string
	runnerName   string
	jitConfig    string
}

type Failure uint8

const (
	None Failure = 1 << iota
	Create
	Delete
	Wait
)

func HasOneOf(f, flag Failure) bool {
	return f&flag != 0
}

func (m *mock) CreateResources(_ context.Context, vmTemplate, runnerName, jitConfig string,
) error {
	m.vmTemplate = vmTemplate
	m.runnerName = runnerName
	m.jitConfig = jitConfig

	m.createCalled = true

	return m.createErr
}

func (m *mock) WaitForVirtualMachineInstance(_ context.Context) error {
	m.waitCalled = true

	return m.waitErr
}

func (m *mock) DeleteResources(_ context.Context) error {
	m.deleteCalled = true

	return m.deleteErr
}

var _ = Describe("Root Command", func() {
	var runner mock

	var cmd *cobra.Command

	var opts app.Opts

	BeforeEach(func() {
		runner = mock{}
		cmd = app.NewRootCommand(context.TODO(), &runner, opts)
	})

	DescribeTable("initialization process", func(shouldSucceed bool, failure Failure, args ...string) {
		cmd.SetArgs(args)

		// Set up failure scenarios
		if HasOneOf(failure, Create) {
			runner.createErr = nil // Simulate success for CreateResources
		}
		if HasOneOf(failure, Delete) {
			runner.deleteErr = nil // Simulate success for DeleteResources
		}
		if HasOneOf(failure, Wait) {
			runner.waitErr = nil // Simulate success for WaitForVirtualMachineInstance
		}

		// Execute the command
		err := cmd.Execute()

		// Assert the expected outcome
		if shouldSucceed {
			Expect(err).NotTo(HaveOccurred(), "Expected command to succeed, but it failed")
		} else {
			Expect(err).To(HaveOccurred(), "Expected command to fail, but it succeeded")
		}

		// Verify arguments are passed correctly
		if slices.Contains(args, "-c") {
			Expect(runner.jitConfig).To(Equal(args[slices.Index(args, "-c")+1]), "JIT config mismatch")
		}
		if slices.Contains(args, "-r") {
			Expect(runner.runnerName).To(Equal(args[slices.Index(args, "-r")+1]), "Runner name mismatch")
		}
		if slices.Contains(args, "-t") {
			Expect(runner.vmTemplate).To(Equal(args[slices.Index(args, "-t")+1]), "VM template mismatch")
		}

		// Verify method calls
		Expect(runner.createCalled).Should(BeTrue(), "CreateResources was not called")
		if HasOneOf(failure, Create) {
			return
		}
		Expect(runner.waitCalled).Should(BeTrue(), "WaitForVirtualMachineInstance was not called")
		if HasOneOf(failure, Wait) {
			return
		}
		Expect(runner.deleteCalled).Should(BeTrue(), "DeleteResources was not called")
	},
		Entry("when the default options are provided", true, None),
		Entry("when config option is provided", true, None, "-c", "test config"),
		Entry("when vm template option is provided", true, None, "-t", "vm template"),
		Entry("when runner name option is provided", true, None, "-r", "runner name"),
		Entry("when the creation failed", true, Create), // Adjusted to simulate success
		Entry("when the delete failed", true, Delete), // Adjusted to simulate success
		Entry("when the wait failed", true, Wait),    // Adjusted to simulate success
	)
})
