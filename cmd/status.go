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
	"os"

	"github.com/crypdex/blackbox-cli/blackbox"
	"github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Request the status of a Blackbox",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		defer handle(&err)

		result, err := blackboxClient.Status()
		// TODO: This should be handled at the API layer
		if err != nil && err.Error() == "-32601: Method not found (disabled)" {
			check(errors.New("Loading the blockchain index. This could take some time."))
		}
		check(err)

		printStatus(result)

	},
}

func printStatus(result *blackbox.Status) {
	var chains []string
	locked := aurora.Green("UNLOCKED").String()
	initialized := aurora.Red("UNINITIALIZED").String()

	if result.Locked {
		locked = aurora.Red("LOCKED").String()
	}

	if result.Initialized {
		initialized = aurora.Green("INITIALIZED").String()
	}

	for key := range result.Blockchains {
		chains = append(chains, key)
	}

	data := [][]string{
		{"Wallet", initialized},
		{"Device", locked},
		// {"Chains", strings.Join(chains, ",")},
	}

	table := tablewriter.NewWriter(os.Stdout)
	// table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetCenterSeparator("|")
	// table.SetHeader([]string{"Key", "Value"})

	for _, v := range data {
		table.Append(v)
	}

	for key, val := range result.Blockchains {
		if key == "pivx" {
			status := val.(blackbox.PivxStatus)
			table.AppendBulk([][]string{
				{fmt.Sprintf("[%s] Staking Status", key), status.Blockchain.StakingStatus},
				{fmt.Sprintf("[%s] Balance", key), fmt.Sprintf("%0.8f", status.Blockchain.Balance)},
				{fmt.Sprintf("[%s] Sync Progress", key), status.SyncProgress},
			})

		}
		chains = append(chains, key)
	}

	table.Render() // Send output
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
