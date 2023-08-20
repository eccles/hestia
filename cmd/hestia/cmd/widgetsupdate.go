// Widgets create subcommand

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	widgetsapi "github.com/eccles/hestia/pkg/apis/widgets"
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

			rootLogger.Info("Update Widgets", "ConfigFile", cfgFile)
			ctx := context.Background()
			resp, err := widgetsClient.Update(ctx, &widgetsapi.UpdateRequest{})
			if err != nil {
				rootLogger.Info("Update Widget", "Error", err)

				return err
			}
			rootLogger.Info("Update Widget", "Response", resp)

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
