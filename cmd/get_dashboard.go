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
	"github.com/spf13/cobra"
	"os"
)

// getDashboardCmd represents the get dashboard command
var getDashboardCmd = &cobra.Command{
	Use:   "dashboard [id]",
	Short: "Get a dashboard by ID",
	Long: `Get a dashboard by ID`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// get dashboard for specific IDs passed
		resp, r, err := client.DashboardsApi.GetDashboard(ddCtx(), args[0]).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboard``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}

		switch output {
		case "terraform":
			// Display terraform resources for dashboard
		default: // json
			dashboardJson, _ := resp.MarshalJSON()
			fmt.Println(prettyJson(dashboardJson))
		}
	},
}

func init() {
	getCmd.AddCommand(getDashboardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dashboardsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dashboardsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
