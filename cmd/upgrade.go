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

	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the host's blackboxd",

	Run: func(cmd *cobra.Command, args []string) {
		log("info", "Checking for updates ...")
		out, err := blackboxClient.UpgradeStatus()
		if err != nil {
			fatal(err)
		}

		if !out.Upgradeable {
			log("info", "You are all up to date")
			log("info", out.Installed)
			return
		}

		log("warn", fmt.Sprintf("An update is available .. %s => %s", out.Installed, out.Candidate))
		log("warn", "Updating ...")

		_, err = blackboxClient.Upgrade()
		if err != nil {
			fatal(err)
		}

		log("info", "Success.")
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
