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
	Short: "Logout of the node and lock the wallet",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		defer handle(&err)
		check(blackboxClient.Logout())
		fmt.Println("successfully logged out")
	},
}
