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

package app_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/electrocucaracha/kubevirt-actions-runner/cmd/kar/app"
	"github.com/spf13/cobra"
)

var errExpectedFailure = errors.New("failure")

var (
	errRootCommandAssertion     = errors.New("root command assertion failed")
	errCreateResourcesNotCalled = errors.New("CreateResources was not called")
	errWaitForVMINotCalled      = errors.New("WaitForVirtualMachineInstance was not called")
	errDeleteResourcesNotCalled = errors.New("DeleteResources was not called")
	errExpectedCommandFailure   = errors.New("expected command failure")
)

// mock implements the runner.Runner interface for exercising the root command.
type mock struct {
	createErr    error
	deleteErr    error
	waitErr      error
	createCalled bool
	waitCalled   bool
	deleteCalled bool
	vmTemplate   string
	vmTemplateNS string
	runnerName   string
	jitConfig    string
}

func (m *mock) CreateResources(
	_ context.Context,
	vmTemplate,
	vmTemplateNamespace,
	runnerName,
	jitConfig string,
) error {
	m.vmTemplate = vmTemplate
	m.vmTemplateNS = vmTemplateNamespace
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

type rootCmdCtx struct {
	runner *mock
	cmd    *cobra.Command
	result error
}

func (rc *rootCmdCtx) reset(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
	rc.runner = new(mock)

	var opts app.Opts

	rc.cmd = app.NewRootCommand(context.TODO(), rc.runner, opts)
	rc.result = nil

	return ctx, nil
}

func (rc *rootCmdCtx) aMockRunner() error {
	return nil
}

//nolint:cyclop,funlen // Mirrors the assertions of the original DescribeTable entries.
func (rc *rootCmdCtx) executeRootCommand(flag, value, failure string) error {
	var args []string
	if flag != "" {
		args = []string{flag, value}
	}

	rc.cmd.SetArgs(args)

	switch failure {
	case "create":
		rc.runner.createErr = errExpectedFailure
	case "delete":
		rc.runner.deleteErr = errExpectedFailure
	case "wait":
		rc.runner.waitErr = errExpectedFailure
	}

	rc.result = rc.cmd.Execute()

	switch flag {
	case "-c":
		if rc.runner.jitConfig != value {
			return fmt.Errorf("%w: JIT config mismatch: want %q, got %q", errRootCommandAssertion, value, rc.runner.jitConfig)
		}
	case "-r":
		if rc.runner.runnerName != value {
			return fmt.Errorf("%w: runner name mismatch: want %q, got %q", errRootCommandAssertion, value, rc.runner.runnerName)
		}
	case "-t":
		if rc.runner.vmTemplate != value {
			return fmt.Errorf("%w: VM template mismatch: want %q, got %q", errRootCommandAssertion, value, rc.runner.vmTemplate)
		}
	case "-n":
		if rc.runner.vmTemplateNS != value {
			return fmt.Errorf(
				"%w: VM template namespace mismatch: want %q, got %q",
				errRootCommandAssertion,
				value,
				rc.runner.vmTemplateNS,
			)
		}
	}

	if flag != "-n" && rc.runner.vmTemplateNS != "default" {
		return fmt.Errorf(
			"%w: expected default VM template namespace, got %q",
			errRootCommandAssertion,
			rc.runner.vmTemplateNS,
		)
	}

	if !rc.runner.createCalled {
		return errCreateResourcesNotCalled
	}

	if failure == "create" {
		return nil
	}

	if !rc.runner.waitCalled {
		return errWaitForVMINotCalled
	}

	if failure == "wait" {
		return nil
	}

	if !rc.runner.deleteCalled {
		return errDeleteResourcesNotCalled
	}

	return nil
}

func (rc *rootCmdCtx) commandExecutionShould(outcome string) error {
	switch outcome {
	case "succeed":
		if rc.result != nil {
			return fmt.Errorf("expected command to succeed, but it failed: %w", rc.result)
		}
	case "fail":
		if rc.result == nil {
			return fmt.Errorf("%w, but it succeeded", errExpectedCommandFailure)
		}
	}

	return nil
}

func InitializeScenario(scenarioCtx *godog.ScenarioContext) {
	rootCtx := new(rootCmdCtx)

	scenarioCtx.Before(rootCtx.reset)

	scenarioCtx.Step(`^a mock runner$`, rootCtx.aMockRunner)
	scenarioCtx.Step(`^the root command is executed with flag "([^"]*)" value "([^"]*)" and induced failure "([^"]*)"$`,
		rootCtx.executeRootCommand)
	scenarioCtx.Step(`^the command execution should "([^"]*)"$`, rootCtx.commandExecutionShould)
}

func TestFeatures(t *testing.T) {
	t.Parallel()

	options := new(godog.Options)
	options.Format = "pretty"
	options.Paths = []string{"features"}
	options.TestingT = t

	suite := new(godog.TestSuite)
	suite.ScenarioInitializer = InitializeScenario
	suite.Options = options

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
