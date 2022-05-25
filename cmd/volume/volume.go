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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// volumeCmd represents the volume command

// var volumeCmd = NewVolumeCmd()

// func GetVolumeCommand() *cobra.Command {
// 	return volumeCmd
// }

var volumeCmd *cobra.Command

const VOLUMEIDKEY = "volume-id"

func NewVolumeCmd() *cobra.Command {

	volumeCmd = &cobra.Command{
		Use:   "volume",
		Short: "Access information about volumes",
		Long: `Access information about volumes.
		
		A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("volume called")
		// },
	}

	// volumeCmd.Flags().IntVar(&id, "id", -1, "The ComicVine ID of the volume")

	// register all the sub-commands here

	volumeCmd.AddCommand(NewVolumeGetCmd())
	volumeCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// volumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	volumeCmd.PersistentFlags().Int(VOLUMEIDKEY, 0, "The ComicVine ID of the volume")
	volumeCmd.MarkPersistentFlagRequired(VOLUMEIDKEY)
	viper.BindPFlag(VOLUMEIDKEY, volumeCmd.PersistentFlags().Lookup(VOLUMEIDKEY))

	return volumeCmd

}

func init() {

}
