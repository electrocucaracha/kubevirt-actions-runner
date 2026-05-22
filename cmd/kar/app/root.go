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

package app

import (
	"context"
	"fmt"

	runner "github.com/electrocucaracha/kubevirt-actions-runner/internal"
	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
	"github.com/spf13/cobra"
)

func NewRootCommand(ctx context.Context, runner runner.Runner, opts Opts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kar",
		Short: "Tool that creates a GitHub Self-Host runner with Kubevirt Virtual Machine Instance",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return initializeConfig(cmd)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return run(ctx, runner, opts)
		},
	}

	installFlags(cmd.Flags(), &opts)

	return cmd
}

func run(ctx context.Context, runner runner.Runner, opts Opts) error {
	log := utils.GetLogger()

	err := runner.CreateResources(ctx, opts.VMTemplate, opts.RunnerName, opts.JitConfig)
	if err != nil {
		return fmt.Errorf("fail to create resources: %w", err)
	}

	log.Println("Virtual Machine runner resources created successfully")

	err = runner.WaitForVirtualMachineInstance(ctx)
	if err != nil {
		return fmt.Errorf("fail to wait for resources: %w", err)
	}

	log.Println("Virtual Machine runner completed successfully")

	err = runner.DeleteResources(ctx)
	if err != nil {
		return fmt.Errorf("fail to delete resources: %w", err)
	}

	log.Println("Virtual Machine runner deleted successfully")

	return nil
}
