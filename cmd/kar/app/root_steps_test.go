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
	rc.runner = &mock{}
	rc.cmd = app.NewRootCommand(context.TODO(), rc.runner, app.Opts{})
	rc.result = nil

	return ctx, nil
}

func (rc *rootCmdCtx) aMockRunner() error {
	return nil
}

//nolint:cyclop // Mirrors the assertions of the original DescribeTable entries.
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
			return fmt.Errorf("JIT config mismatch: want %q, got %q", value, rc.runner.jitConfig)
		}
	case "-r":
		if rc.runner.runnerName != value {
			return fmt.Errorf("runner name mismatch: want %q, got %q", value, rc.runner.runnerName)
		}
	case "-t":
		if rc.runner.vmTemplate != value {
			return fmt.Errorf("VM template mismatch: want %q, got %q", value, rc.runner.vmTemplate)
		}
	case "-n":
		if rc.runner.vmTemplateNS != value {
			return fmt.Errorf("VM template namespace mismatch: want %q, got %q", value, rc.runner.vmTemplateNS)
		}
	}

	if flag != "-n" && rc.runner.vmTemplateNS != "default" {
		return fmt.Errorf("expected default VM template namespace, got %q", rc.runner.vmTemplateNS)
	}

	if !rc.runner.createCalled {
		return errors.New("CreateResources was not called")
	}

	if failure == "create" {
		return nil
	}

	if !rc.runner.waitCalled {
		return errors.New("WaitForVirtualMachineInstance was not called")
	}

	if failure == "wait" {
		return nil
	}

	if !rc.runner.deleteCalled {
		return errors.New("DeleteResources was not called")
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
			return errors.New("expected command to fail, but it succeeded")
		}
	}

	return nil
}

func InitializeScenario(sc *godog.ScenarioContext) {
	rc := &rootCmdCtx{}

	sc.Before(rc.reset)

	sc.Step(`^a mock runner$`, rc.aMockRunner)
	sc.Step(`^the root command is executed with flag "([^"]*)" value "([^"]*)" and induced failure "([^"]*)"$`,
		rc.executeRootCommand)
	sc.Step(`^the command execution should "([^"]*)"$`, rc.commandExecutionShould)
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
