/*
Copyright © 2022 James Evans

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

	issue "github.com/jaevans/comic-finder/cmd/issue"
	volume "github.com/jaevans/comic-finder/cmd/volume"

	"github.com/jaevans/comic-finder/pkg/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const APIOPTIONKEY = "api-key"

var cfgFile string
var apiKey string
var outputFormat string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "comic-finder",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var client *types.ComicVineClient

func GetAPIClient() *types.ComicVineClient {
	return client
}
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.comic-finder.yaml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().StringVar(&apiKey, APIOPTIONKEY, "", "API key for accessing ComicVine database. Can also be specified in the configuration file")
	viper.BindPFlag(APIOPTIONKEY, rootCmd.PersistentFlags().Lookup(APIOPTIONKEY))

	viper.SetDefault("format", "text")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "o", viper.GetString("format"), "The output format. One of text, json, or yaml")
	viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(volume.NewVolumeCmd())
	rootCmd.AddCommand(issue.NewIssueCmd())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".comic-finder" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".comic-finder")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if viper.GetString(APIOPTIONKEY) == "" {
		// log.Fatal("Required field api-key not found.")
		fmt.Fprintln(os.Stderr, "Required field api-key not found.")
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}
}
