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
	"fmt"
	"strings"

	"github.com/electrocucaracha/kubevirt-actions-runner/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func installFlags(flags *pflag.FlagSet, cmdOptions *Opts) {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	flags.StringVarP(&cmdOptions.VMTemplate, "kubevirt-vm-template", "t", "vm-template",
		"The VirtualMachine resource to use as the template.")
	flags.StringVarP(&cmdOptions.RunnerName, "runner-name", "r", "runner",
		"The name of the runner.")
	flags.StringVarP(&cmdOptions.JitConfig, "actions-runner-input-jitconfig", "c", "",
		"The opaque JIT runner config.")
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, viperInstance *viper.Viper) {
	log := utils.GetLogger()

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		configName := flag.Name

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !flag.Changed && viperInstance.IsSet(configName) {
			val := viperInstance.Get(configName)

			err := cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", val))
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	})
}
