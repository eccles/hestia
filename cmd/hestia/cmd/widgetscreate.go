// Widgets create subcommand

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
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

			rootLogger.Info().Msgf("Create Widgets: %s", cfgFile)
			ctx := context.Background()
			resp, err := widgetsClient.Create(ctx, &widgetsAPI.CreateRequest{})
			if err != nil {
				rootLogger.Info().Msgf("Create Widget Error: %v", err)

				return err
			}
			rootLogger.Info().Msgf("Create Widget: %v", resp)

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
