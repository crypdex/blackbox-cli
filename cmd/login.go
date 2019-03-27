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
	"io/ioutil"
	"os"

	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

type conf struct {
	Token string `yaml:"token"`
	Host  string `yaml:"host"`
	Chain string `yaml:"chain"`
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login and unlock a Blackbox",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		defer handle(&err)

		prompt := promptui.Prompt{
			Label: "password",
			Mask:  '*',
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("Password cannot be blank")
				}
				return nil
			},
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		save := cmd.Flag("save").Value.String() == "true"

		response, err := blackboxClient.Login(result, save)
		check(err)

		fmt.Println(response.JWT)
		path, err := homedir.Dir()
		check(err)

		c := new(conf)

		yamlPath := path + "/.crypdex/blackbox.yaml"
		if _, err := os.Stat(yamlPath); !os.IsNotExist(err) {
			yamlFile, err := ioutil.ReadFile(yamlPath)
			check(err)
			err = yaml.Unmarshal(yamlFile, c)
			check(err)
		}

		c.Token = response.JWT
		out, err := yaml.Marshal(c)
		check(err)

		err = ioutil.WriteFile(yamlPath, out, 0644)
		check(err)

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loginCmd.Flags().BoolP("save", "s", false, "[DEV ONLY] Save the password on the device.")
}
