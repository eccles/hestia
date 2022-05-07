// Widgets create subcommand

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
)

//nolint:gochecknoglobals  // this is just the way cobra/viper works
var (
	widgetsUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Widgets update",
		Long:  `Widgets update`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseWidgetsUpdateCmd()
			defer deferWidgetsUpdateCmd()

			rootLogger.Info().Msgf("Update Widgets: %s", cfgFile)
			ctx := context.Background()
			resp, err := widgetsClient.Update(ctx, &widgetsAPI.UpdateRequest{})
			if err != nil {
				rootLogger.Info().Msgf("Update Widget Error: %v", err)

				return err
			}
			rootLogger.Info().Msgf("Update Widget: %v", resp)

			return nil
		},
	}
)

func parseWidgetsUpdateCmd() {
	parseWidgetsCmd()
}

func deferWidgetsUpdateCmd() {
	deferWidgetsCmd()
}

func widgetsUpdateNew() {
	widgetsCmd.AddCommand(widgetsUpdateCmd)
}
