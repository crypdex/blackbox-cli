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
	"strconv"

	"github.com/logrusorgru/aurora"

	"github.com/olekukonko/tablewriter"

	"github.com/crypdex/blackbox-cli/blackbox"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Manage addresses",
	Long:  ``,
}

var addressListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List known addresses",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateChain()
	},
	Run: func(cmd *cobra.Command, args []string) {
		addresses, err := blackboxClient.AddressList(chain)
		if err != nil {
			fatal(err)
		}

		if len(addresses) == 0 {
			log("info", fmt.Sprintf("There are no %s addresses yet", chain))
			return
		}

		log("info", fmt.Sprintf("Found %d %s addresses", len(addresses), chain))

		table := tablewriter.NewWriter(os.Stdout)
		// table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
		table.SetCenterSeparator("|")
		table.SetHeader([]string{
			"public key", "available", "pending", "locked",
		})
		for _, address := range addresses {
			table.AppendBulk([][]string{
				{
					aurora.Cyan(address.PublicKey).String(),
					aurora.Green(address.Balance.Available).String(),
					address.Balance.Pending.String(),
					aurora.Red(address.Balance.Locked).String(),
				},
			})
		}

		table.Render()
	},
}

var addressCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new address",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateChain()
	},

	Run: func(cmd *cobra.Command, args []string) {
		request := new(blackbox.CreateAddressRequest)

		rescan := cmd.Flag("rescan").Value.String() == "true"
		request.Rescan = rescan

		watchonly := cmd.Flag("watchonly").Value.String() == "true"
		request.Watchonly = watchonly

		response, err := blackboxClient.AddressCreate(chain, *request)
		if err != nil {
			fatal(err)
		}

		log("info", "Success. Your new address is:")
		log("warn", response.PublicKey)
	},
}

var addressRecreateCmd = &cobra.Command{
	Use:   "recreate",
	Short: "Re-create addresses",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateChain()
	},
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		count, err := strconv.Atoi(args[0])
		if err != nil {
			fatal(err)
		}

		response, err := blackboxClient.AddressRecreate(chain, count)
		if err != nil {
			fatal(err)
		}

		log("info", fmt.Sprintf("Successfully recreated %d addresses", count))
		log("info", "The RPC is effectively locked while this processes.")
		log("info", response["message"])
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)
	addressCmd.AddCommand(addressListCmd)
	addressCmd.AddCommand(addressCreateCmd)
	addressCmd.AddCommand(addressRecreateCmd)

	addressCreateCmd.Flags().BoolP("rescan", "r", false, "Rescan the blockchain for transactions")
	addressCreateCmd.Flags().BoolP("watchonly", "w", false, "Create as watchonly address. This is not recommended.")
}

func validateChain() {
	if chain == "" {
		fatal(errors.New("chain is required: use -c flag or add to ~/.blackbox.yaml"))
	}
}
