package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logoutCmd)
}

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of the node, lock the wallet, remove any stored password",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := blackboxClient.Logout(); err != nil {
			fatal(err)
		}

		fmt.Println("successfully logged out")
	},
}
