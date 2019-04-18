// Copyrigh
package cmd

import (
	"fmt"

	"github.com/crypdex/blackbox-cli/blackbox"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Wallet related commands",
}

var banner = `
█▀▀▄ █░░ █▀▀█ █▀▀ █░█ █▀▀▄ █▀▀█ █░█ 
█▀▀▄ █░░ █▄▄█ █░░ █▀▄ █▀▀▄ █░░█ ▄▀▄ 
▀▀▀░ ▀▀▀ ▀░░▀ ▀▀▀ ▀░▀ ▀▀▀░ ▀▀▀▀ ▀░▀ 
`

var instructions = `
Okay, let's initialize the wallet.

First, choose a secure password and remember it.
If you forget it or lose it we cannot recover it.
`

var instructions2 = `
If you have a mnemonic you would like to use, enter it now. Otherwise one will be generated for you.
`

var instructions3 = `
Write down this mnemonic and store it in a safe place. With this phrase and your password, you can recreate your wallet if you lose this device.
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the wallet",
	Long:  `This command initializes a Blackbox including the wallet. You can force it to reinitialize with the '-f' flag.`,

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		defer handle(&err)

		var prompt promptui.Prompt

		fmt.Println(aurora.Black(banner))
		fmt.Println(aurora.Green(instructions))

		prompt = promptui.Prompt{
			Label: "Password",
			Mask:  '*',
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("Password cannot be blank")
				}
				return nil
			}}
		password, err := prompt.Run()
		check(err)

		// prompt = promptui.Prompt{Label: "email (optional)"}
		// email, err := prompt.Run()
		// check(err)
		fmt.Println(aurora.Green(instructions2))

		prompt = promptui.Prompt{Label: "mnemonic (optional)"}
		mnemonic, err := prompt.Run()
		check(err)

		response, err := blackboxClient.Init(blackbox.InitRequest{
			// Email:    email,
			Password: password,
			Mnemonic: mnemonic,
			Force:    force,
		})

		check(err)

		fmt.Println(aurora.Green(instructions3))
		fmt.Println(aurora.Cyan(response.Mnemonic))
	},
}

var force bool

func init() {
	rootCmd.AddCommand(walletCmd)
	walletCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force re-initialization")
}
