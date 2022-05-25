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
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jaevans/comic-finder/pkg/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// getCmd represents the get command
var getCmd *cobra.Command

// func NewVolumeGetCmd(c *types.ComicVineClient) *cobra.Command {
func NewIssueGetCmd() *cobra.Command {
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get information on a single issue",
		Long:  `Retrieve information about a single volume, within the provided volume ID`,
		Run:   Run,
	}
	return getCmd
}

func printIssue(issue *types.Issue) {
	switch viper.GetString("format") {
	case "text":
		fmt.Printf("comicvine:%d - %s Issue #%d [Publisher: %s]\n", issue.Id, issue.Volume.Name, issue.IssueNumber, issue.Volume.Publisher.Name)
	case "yaml":
		encoder := yaml.NewEncoder(os.Stdout)
		err := encoder.Encode(issue)
		if err != nil {
			log.Fatal(fmt.Errorf("error encoding output to yaml: %w", err))
			// fmt.E(os.Stderr, "Error encoding output to stdout: %w", err)
		}
	case "json", "prettyjson":
		var data []byte
		var err error
		if viper.GetString("format") == "prettyjson" {
			data, err = json.MarshalIndent(issue, "", "  ")
		} else {
			data, err = json.Marshal(issue)
		}
		if err != nil {
			log.Fatal(fmt.Errorf("error encoding output to json: %w", err))
		}
		fmt.Println(string(data))
	}

}

func Run(command *cobra.Command, args []string) {
	client := types.NewComicVineClient(viper.GetString("api-key"))
	issueId := viper.GetInt("issue-id")
	if issueId != 0 {
		issue, err := client.GetIssueById(issueId, types.GetOptions{})
		if err != nil {
			log.Fatal("error retrieving issue by id:", err)
		}
		printIssue(issue)
		return
	}

	issueNumber := viper.GetInt("issue-number")
	volumeId := viper.GetInt("volume-id")
	if issueNumber != 0 {
		issues, err := client.GetIssuesByVolume()
	}

	// if err != nil {
	// 	panic(err)
	// }

}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
