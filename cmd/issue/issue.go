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

var issueCmd *cobra.Command

const ISSUEIDKEY = "issue-id"
const ISSUENUMBERKEY = "issue"
const VOLUMEIDKEY = "volume-id"

func NewIssueCmd() *cobra.Command {

	issueCmd = &cobra.Command{
		Use:   "issue",
		Short: "Access information about issues",
		Long: `Access information about issues.
		
		A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("issue called")
		// },
	}

	// issueCmd.Flags().IntVar(&id, "id", -1, "The ComicVine ID of the issue")

	// register all the sub-commands here

	issueCmd.AddCommand(NewIssueGetCmd())
	// issueCmd.AddCommand(NewissueSearchCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	issueCmd.PersistentFlags().Int(ISSUEIDKEY, 0, "The ComicVine ID of the issue")
	viper.BindPFlag(ISSUEIDKEY, issueCmd.PersistentFlags().Lookup(ISSUEIDKEY))

	issueCmd.PersistentFlags().Int(VOLUMEIDKEY, 0, "The ComicVine ID of the volume")
	issueCmd.MarkPersistentFlagRequired(VOLUMEIDKEY)
	viper.BindPFlag(VOLUMEIDKEY, issueCmd.PersistentFlags().Lookup(VOLUMEIDKEY))
	return issueCmd

}

func init() {

}
