// Widgets create subcommand

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
)

//nolint:gochecknoglobals  // this is just the way cobra/viper works
var (
	widgetsDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Widgets delete",
		Long:  `Widgets delete`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseWidgetsDeleteCmd()
			defer deferWidgetsDeleteCmd()

			rootLogger.Info().Msgf("Delete Widgets: %s", cfgFile)
			ctx := context.Background()
			resp, err := widgetsClient.Delete(ctx, &widgetsAPI.DeleteRequest{})
			if err != nil {
				rootLogger.Info().Msgf("Delete Widget Error: %v", err)

				return err
			}
			rootLogger.Info().Msgf("Delete Widget: %v", resp)

			return nil
		},
	}
)

func parseWidgetsDeleteCmd() {
	parseWidgetsCmd()
}

func deferWidgetsDeleteCmd() {
	deferWidgetsCmd()
}

func widgetsDeleteNew() {
	widgetsCmd.AddCommand(widgetsDeleteCmd)
}
