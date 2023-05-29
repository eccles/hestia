// Widgets create subcommand

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
)

//nolint:gochecknoglobals  // this is just the way cobra/viper works
var (
	widgetsFindCmd = &cobra.Command{
		Use:   "find",
		Short: "Widgets find",
		Long:  `Widgets find`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseWidgetsFindCmd()
			defer deferWidgetsFindCmd()

			rootLogger.Info("Find Widgets", "configFile", cfgFile)
			ctx := context.Background()
			resp, err := widgetsClient.FindByID(ctx, &widgetsAPI.FindRequest{})
			if err != nil {
				rootLogger.Info("Find Widget", "Error", err)

				return err
			}
			rootLogger.Info("Find Widget", "Response", resp)

			return nil
		},
	}
)

func parseWidgetsFindCmd() {
	parseWidgetsCmd()
}

func deferWidgetsFindCmd() {
	deferWidgetsCmd()
}

func widgetsFindNew() {
	widgetsCmd.AddCommand(widgetsFindCmd)
}
