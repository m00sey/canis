/*
Copyright © 2020 Scoir, Inc <phil@scoir.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/scoir/canis/pkg/steward"
)

var cfgFile string

var config *steward.Config

var rootCmd = &cobra.Command{
	Use:   "steward",
	Short: "The canis steward orchestration service.",
	Long: `"The canis steward orchestration service.".

 Find more information at: https://canis.io/docs/reference/canis/overview`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.canis.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cfgFile = strings.Join([]string{home, ".canis"}, string(os.PathSeparator))
	}

	// If a config file is found, read it in.
	f, err := os.Open(cfgFile)
	if err != nil {
		fmt.Println("unable to read config:", cfgFile, err)
		os.Exit(1)
	}

	config = steward.NewConfig()
	err = yaml.NewDecoder(f).Decode(config)
	if err != nil {
		fmt.Println("failed to unmarshal config", err)
		os.Exit(1)
	}

}
