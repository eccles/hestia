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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	widgetsapi "github.com/eccles/hestia/apis/widgets"
)

const (
	widgetsAddressLabel = "widgetsaddress"
)

//nolint:gochecknoglobals  // this is just the way cobra/viper works
var (
	widgetsAddress string
	widgetsClient  widgetsapi.WidgetsClient
	widgetsCmd     = &cobra.Command{
		Use:   "widgets",
		Short: "Widgets client",
		Long:  `Widgets communication`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseWidgetsCmd()
			defer deferWidgetsCmd()

			rootLogger.Info("Widgets", "configFile", cfgFile)

			return nil
		},
	}
)

func parseWidgetsCmd() {
	parseRootCmd()

	conn, err := grpc.NewClient(
		widgetsAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		rootLogger.Info("creating Widgets connection", "err", err)
	}

	widgetsClient = widgetsapi.NewWidgetsClient(conn)
}

func deferWidgetsCmd() {
	widgetsClient = nil
}

func widgetsNew() {
	widgetsCreateNew()
	widgetsFindNew()
	widgetsListNew()
	widgetsUpdateNew()
	widgetsDeleteNew()
	widgetsCmd.PersistentFlags().StringVarP(
		&widgetsAddress,
		widgetsAddressLabel,
		"",
		"",
		"Address of widgets service",
	)

	_ = viper.BindPFlag(widgetsAddressLabel, widgetsCmd.PersistentFlags().Lookup(widgetsAddressLabel))

	rootCmd.AddCommand(widgetsCmd)
}
