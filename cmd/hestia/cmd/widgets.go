// Widgets subcommand

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
)

const (
	widgetsAddressLabel = "widgetsaddress"
)

//nolint:gochecknoglobals  // this is just the way cobra/viper works
var (
	widgetsAddress string
	widgetsClient  widgetsAPI.WidgetsClient
	widgetsCmd     = &cobra.Command{
		Use:   "widgets",
		Short: "Widgets client",
		Long:  `Widgets communication`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			parseWidgetsCmd()
			defer deferWidgetsCmd()

			rootLogger.Info().Msgf("Widgets: %s", cfgFile)

			return nil
		},
	}
)

func parseWidgetsCmd() {
	parseRootCmd()
	conn, err := grpc.Dial(
		widgetsAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		rootLogger.Panic().Msgf("Error creating Widgets connection: %v", err)
	}

	widgetsClient = widgetsAPI.NewWidgetsClient(conn)
}

func deferWidgetsCmd() {
	rootLogger.Close()
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