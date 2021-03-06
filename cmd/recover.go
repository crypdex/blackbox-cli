package cmd

import (
	"errors"

	"github.com/crypdex/blackbox-cli/blackbox"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(recoverCmd)
}

var recoverInstructions1 = `
To recover your wallet, enter your mnemonic phrase and password.
Be aware that the mnemonic is not currenly checked for validity.
`

// initCmd represents the init command
var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover a wallet",
	Long:  `This command recovers a App wallet.`,

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		defer handle(&err)

		var prompt promptui.Prompt

		log("info", "CAREFUL: This is recovery mode.\nRestoring a wallet will wipe the current wallet if it exists.")
		log("info", recoverInstructions1)

		prompt = promptui.Prompt{
			Label: "Mnemonic ",
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("mnemonic cannot be blank")
				}
				return nil
			}}
		mnemonic, err := prompt.Run()
		check(err)

		prompt = promptui.Prompt{
			Label: "Mnemonic Password (optional) ",
			Mask:  '*',
		}
		mnemonicPassword, err := prompt.Run()
		check(err)

		prompt = promptui.Prompt{
			Label: "Current Password ",
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("password cannot be blank")
				}
				return nil
			}}
		password, err := prompt.Run()
		check(err)

		_, err = blackboxClient.Init(blackbox.InitRequest{
			// Email:    email,
			Password:         password,
			MnemonicPassword: mnemonicPassword,
			Mnemonic:         mnemonic,
			Force:            true,
		})

		check(err)

		log("info", "Success.")
	},
}
