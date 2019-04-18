package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
}

// configCmd represents the config command
var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show the current saved config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config get called")
		fmt.Println(viper.AllSettings())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	// configCmd.AddCommand(configGetCmd)
}
