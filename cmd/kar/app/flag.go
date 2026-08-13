/* jscpd:ignore-start */
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
/* jscpd:ignore-end */

package app

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func installFlags(flags *pflag.FlagSet, cmdOptions *Opts) {
	flags.StringVarP(&cmdOptions.VMTemplate, "kubevirt-vm-template", "t", "vm-template",
		"The VirtualMachine resource to use as the template.")
	flags.StringVarP(&cmdOptions.VMTemplateNamespace, "kubevirt-vm-template-namespace", "n", "default",
		"The namespace where the VirtualMachine template resource exists.")
	flags.StringVarP(&cmdOptions.RunnerName, "runner-name", "r", "runner",
		"The name of the runner.")
	flags.StringVarP(&cmdOptions.JitConfig, "actions-runner-input-jitconfig", "c", "",
		"The opaque JIT runner config.")
}

func initializeConfig(cmd *cobra.Command) error {
	// Load environment variables into flags if not already set.
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Changed {
			return
		}

		// Convert flag name to env var: replace "-" with "_" and uppercase.
		envKey := strings.ToUpper(strings.ReplaceAll(flag.Name, "-", "_"))
		if val, ok := os.LookupEnv(envKey); ok {
			_ = cmd.Flags().Set(flag.Name, val)
		}
	})

	return nil
}
