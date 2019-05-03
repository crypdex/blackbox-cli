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

var banner = `
$$\       $$\                     $$\       $$\                           
$$ |      $$ |                    $$ |      $$ |                          
$$$$$$$\  $$ | $$$$$$\   $$$$$$$\ $$ |  $$\ $$$$$$$\   $$$$$$\  $$\   $$\ 
$$  __$$\ $$ | \____$$\ $$  _____|$$ | $$  |$$  __$$\ $$  __$$\ \$$\ $$  |
$$ |  $$ |$$ | $$$$$$$ |$$ /      $$$$$$  / $$ |  $$ |$$ /  $$ | \$$$$  / 
$$ |  $$ |$$ |$$  __$$ |$$ |      $$  _$$<  $$ |  $$ |$$ |  $$ | $$  $$<  
$$$$$$$  |$$ |\$$$$$$$ |\$$$$$$$\ $$ | \$$\ $$$$$$$  |\$$$$$$  |$$  /\$$\ 
\_______/ \__| \_______| \_______|\__|  \__|\_______/  \______/ \__/  \__|
`

var instructions = `
Okay, let's initialize the wallet.

A mnemonic phrase will be generated for you.
Please enter a password to secure your wallet.
`

var instructions2 = `
If you have a mnemonic you would like to use, enter it now. 
Otherwise one will be generated for you.
`

var instructions3 = `
IMPORTANT: Write down this mnemonic and store it in a safe place.
With this phrase and your password, you can recreate your wallet if you lose this device.
`

var force bool
var initPassword bool

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force re-initialization")
	initCmd.Flags().BoolVarP(&initPassword, "password", "p", false, "Allows for initialization with a mnemonic password")
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the wallet",
	Long:  `This command initializes a App including the wallet. You can force it to reinitialize with the '-f' flag.`,

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		defer handle(&err)

		var password string
		var passwordConfirm string
		var mnemonicPassword string

		var prompt promptui.Prompt

		fmt.Println(aurora.Black(banner))
		log("info", instructions)
		if force {
			log("warn", "Since you are forcing initialization,\nthis password must match your existing password.")
		}

		prompt = promptui.Prompt{
			Label: "Password ",
			Mask:  '*',
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("Password cannot be blank")
				}
				return nil
			}}
		password, err = prompt.Run()
		check(err)

		prompt = promptui.Prompt{
			Label: "Re-enter Password ",
			Mask:  '*',
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("Password cannot be blank")
				}
				return nil
			}}
		passwordConfirm, err = prompt.Run()
		check(err)

		// Verify that the password and the confirm match please
		if password != passwordConfirm {
			fatal(errors.New("the passwords you gave do not match"))
		}

		if initPassword {
			prompt = promptui.Prompt{
				Label: "Mnemonic Password (optional) ",
				Mask:  '*',
			}
			mnemonicPassword, err = prompt.Run()
			check(err)
		}

		response, err := blackboxClient.Init(blackbox.InitRequest{
			// Email:    email,
			Password:         password,
			MnemonicPassword: mnemonicPassword,
			Force:            force,
		})

		if err != nil && err.Error() == "already initialized" {
			log("warn", "This wallet is already initialized.\nYou may re-rerun this command with the -f flag to force it.")
			return
		} else {
			check(err)
		}

		fmt.Println(aurora.Green(instructions3))
		fmt.Println(aurora.Cyan(response.Mnemonic))
		fmt.Println()
	},
}
