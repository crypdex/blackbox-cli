// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/crypdex/blackbox-cli/blackbox"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Blackbox",
	Long:  `This command initializes a Blackbox including the wallet. You can force it to reinitialize with the '-f' flag.`,

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		defer handle(&err)

		var prompt promptui.Prompt

		prompt = promptui.Prompt{
			Label: "password",
			Mask:  '*',
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("Password cannot be blank")
				}
				return nil
			}}
		password, err := prompt.Run()
		check(err)

		prompt = promptui.Prompt{Label: "email (optional)"}
		email, err := prompt.Run()
		check(err)

		prompt = promptui.Prompt{Label: "mnemonic (optional)"}
		mnemonic, err := prompt.Run()
		check(err)

		response, err := blackboxClient.Init(blackbox.InitRequest{
			Email:    email,
			Password: password,
			Mnemonic: mnemonic,
			Force:    force,
		})

		check(err)

		fmt.Println("Write down this mnemonic and store it in a safe place:")
		fmt.Println(aurora.Green(response.Mnemonic))
	},
}

var force bool

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force re-initialization")

}
