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

		if HasOneOf(failure, Create) {
			runner.createErr = errExpectedFailure
		}

		if HasOneOf(failure, Delete) {
			runner.deleteErr = errExpectedFailure
		}

		if HasOneOf(failure, Wait) {
			runner.waitErr = errExpectedFailure
		}

		err := cmd.Execute()

		if shouldSucceed {
			Expect(err).NotTo(HaveOccurred())
		} else {
			Expect(err).To(HaveOccurred())
		}

		if slices.Contains(args, "-c") {
			Expect(runner.jitConfig).To(Equal(args[slices.Index(args, "-c")+1]))
		}

		if slices.Contains(args, "-r") {
			Expect(runner.runnerName).To(Equal(args[slices.Index(args, "-r")+1]))
		}

		if slices.Contains(args, "-t") {
			Expect(runner.vmTemplate).To(Equal(args[slices.Index(args, "-t")+1]))
		}

		Expect(runner.createCalled).Should(BeTrue())

		if HasOneOf(failure, Create) {
			return
		}

		Expect(runner.waitCalled).Should(BeTrue())

		if HasOneOf(failure, Wait) {
			return
		}

		Expect(runner.deleteCalled).Should(BeTrue())
	},
		Entry("when the default options are provided", true, None),
		Entry("when config option is provided", true, None, "-c", "test config"),
		Entry("when vm template option is provided", true, None, "-t", "vm template"),
		Entry("when runner name option is provided", true, None, "-r", "runner name"),
		Entry("when the creation failed", false, Create),
		Entry("when the delete failed", false, Delete),
		Entry("when the wait failed", false, Wait),
	)
})
