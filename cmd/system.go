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

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("system called")
	},
}

var systemStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of the system software",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := blackboxClient.SystemStatus()
		check(err)

		fmt.Println(response)
	},
}

var systemUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the system software",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := blackboxClient.SystemUpdate()
		check(err)

		fmt.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)
	systemCmd.AddCommand(systemStatusCmd)
	systemCmd.AddCommand(systemUpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// systemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// systemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
