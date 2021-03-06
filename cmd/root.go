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
	"net"
	"os"
	"strings"

	"github.com/goware/urlx"
	. "github.com/logrusorgru/aurora"

	"github.com/crypdex/blackbox-cli/blackbox"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var chain string
var host string
var blackboxClient *blackbox.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "blackbox-cli",
	Short: "A command line utility for interfacing with devices running the BlackboxOS",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var debug bool

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initClient)
	cobra.OnInitialize(initChain)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blackbox.yaml)")

	rootCmd.PersistentFlags().StringVarP(&host, "address", "a", "", "blackbox node address")

	rootCmd.PersistentFlags().StringVarP(&chain, "chain", "c", "", "selected chain")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".blackbox-cli" (without extension).
		viper.AddConfigPath(home + "/.crypdex")
		viper.SetConfigName("blackbox")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}

func initChain() {
	if chain == "" {
		chain = viper.GetString("chain")
	}
}

func initClient() {
	var err error
	// Has the host been set with the global flag?
	// Is the host saved in a config?
	if host == "" {
		host = viper.GetString("host")
	}

	if host == "" {
		log("debug", "No host given, defaulting to crypdex.local")
		host = "crypdex.local"
	}

	u, err := urlx.Parse(host)
	host, port, _ := urlx.SplitHostPort(u)
	addrs, err := net.LookupHost(host)
	if err != nil {
		fatal(err)
	}
	for _, addr := range addrs {
		if IsIPv4(addr) {
			host = u.Scheme + "://" + addr + ":" + port
		}
	}

	log("debug", fmt.Sprintf("Using host at %s", host))

	blackboxClient, err = blackbox.NewClient(host, viper.GetString("token"), debug)
	if err != nil {
		fatal(err)
	}
}

func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

func log(level string, msg string) {

	if level == "debug" && debug {
		fmt.Println("[debug]", msg)
	} else if level == "info" {
		fmt.Println(Green(msg))
	} else if level == "warn" {
		fmt.Println(Cyan(msg))
	} else {
		fmt.Println(msg)
	}

}
