package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Token string `yaml:"token"`
	Host  string `yaml:"host"`
	Chain string `yaml:"chain"`
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login and unlock a Blackbox",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		defer handle(&err)

		prompt := promptui.Prompt{
			Label: "password",
			Mask:  '*',
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("Password cannot be blank")
				}
				return nil
			},
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		save := cmd.Flag("save").Value.String() == "true"

		response, err := blackboxClient.Login(result, save)
		check(err)

		// fmt.Println(response.JWT)
		path, err := homedir.Dir()
		check(err)

		// Does the CLI directory exist?
		configpath := filepath.Join(path, ".crypdex")
		if _, err := os.Stat(configpath); os.IsNotExist(err) {
			fmt.Println(".crypdex does not exist. Creating ...")

			os.MkdirAll(configpath, os.ModePerm)
		}

		// Does a stored config file exist?
		configfile := filepath.Join(configpath, "blackbox.yml")
		if _, err := os.Stat(configfile); os.IsNotExist(err) {
			fmt.Println(".crypdex/blackbox.yml does not exist. Creating ...")

			emptyFile, err := os.Create(configfile)
			check(err)

			check(emptyFile.Close())
		}

		c := new(conf)

		// if _, err := os.Stat(yamlPath); !os.IsNotExist(err) {
		yamlFile, err := ioutil.ReadFile(configfile)
		check(err)
		err = yaml.Unmarshal(yamlFile, c)
		check(err)

		c.Token = response.JWT
		out, err := yaml.Marshal(c)
		check(err)

		err = ioutil.WriteFile(configfile, out, 0644)
		check(err)

		fmt.Println(aurora.Green("Success."))
		fmt.Println("A token has been saved to ~/.crypdex/blackbox.yml")
		fmt.Println("You may now use privileged commands.")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	loginCmd.Flags().BoolP("save", "s", false, "[DEV ONLY] Save the password on the device.")
}
