// Widgets create subcommand

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	widgetsapi "github.com/eccles/hestia/pkg/apis/widgets"
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
