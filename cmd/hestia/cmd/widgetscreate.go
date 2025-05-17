// .*@mycompany\.com MY COMPANY 2025
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	widgetsapi "github.com/eccles/hestia/apis/widgets"
)

//nolint:gochecknoglobals  // this is just the way cobra/viper works
var (
	widgetsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Widgets create",
		Long:  `Widgets creation`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseWidgetsCreateCmd()
			defer deferWidgetsCreateCmd()

			rootLogger.Info("Create Widgets", "ConfigFile", cfgFile)
			ctx := context.Background()
			resp, err := widgetsClient.Create(ctx, &widgetsapi.CreateRequest{})
			if err != nil {
				rootLogger.Info("Create Widget", "Error", err)
				return err
			}
			rootLogger.Info("Create Widget", "Response", resp)

			return nil
		},
	}
)

func parseWidgetsCreateCmd() {
	parseWidgetsCmd()
}

func deferWidgetsCreateCmd() {
	deferWidgetsCmd()
}

func widgetsCreateNew() {
	widgetsCmd.AddCommand(widgetsCreateCmd)
}
