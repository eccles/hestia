// Widgets create subcomman/List/List/g

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
)

//nolint:gochecknoglobals  // this is just the way cobra/viper works
var (
	widgetsListCmd = &cobra.Command{
		Use:   "list",
		Short: "Widgets list",
		Long:  `Widgets list`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseWidgetsListCmd()
			defer deferWidgetsListCmd()

			rootLogger.Info().Msgf("List Widgets: %s", cfgFile)
			ctx := context.Background()
			resp, err := widgetsClient.List(ctx, &widgetsAPI.ListRequest{})
			if err != nil {
				rootLogger.Info().Msgf("List Widget Error: %v", err)

				return err
			}
			rootLogger.Info().Msgf("List Widget: %v", resp)

			return nil
		},
	}
)

func parseWidgetsListCmd() {
	parseWidgetsCmd()
}

func deferWidgetsListCmd() {
	deferWidgetsCmd()
}

func widgetsListNew() {
	widgetsCmd.AddCommand(widgetsListCmd)
}
