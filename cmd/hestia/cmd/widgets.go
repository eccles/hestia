// Widgets subcommand

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	widgetsapi "github.com/eccles/hestia/pkg/apis/widgets"
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
