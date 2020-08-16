/*
Copyright © 2020 Sarkis Varozian <svarozian@gmail.com>

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

	"github.com/spf13/cobra"
)

// listDashboardsCmd represents the dashboards command
var listDashboardsCmd = &cobra.Command{
	Use:   "dashboards",
	Short: "List all dashboards",
	Long: `List all dashboards`,
	Run: func(cmd *cobra.Command, args []string) {
		// get dashboards passed with no IDs so list all dashboards
		resp, r, err := client.DashboardsApi.ListDashboards(ddCtx()).Execute()
		if err != nil {
			// TODO: make this better
			fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ListDashboards``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		dashboardJson, _ := resp.MarshalJSON()
		fmt.Println(prettyJson(dashboardJson))
	},
}

func init() {
	listCmd.AddCommand(listDashboardsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dashboardsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dashboardsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
