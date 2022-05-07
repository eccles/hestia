// Entrypoint for hestia command (no subcommands)

package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/eccles/hestia/pkg/logger"
)

const (
	cmd           = "hestia"
	logLevelLabel = "loglevel"
)

//nolint:gochecknoglobals  // this is just the way cobra/viper works
var (
	cfgFile    string
	logLevel   string
	rootLogger logger.Logger
	rootCmd    = &cobra.Command{
		Use:   cmd,
		Short: "Manages hestia system",
		Long:  `Hestia manages the hestia subsystem.`,
	}
)

func parseRootCmd() {
	logLevel = viper.GetString(logLevelLabel)
	rootLogger = logger.Logger{
		Console:     true,
		Level:       logLevel,
		ServiceName: cmd,
	}
	_ = rootLogger.Open()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootNew()
	widgetsNew()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootNew() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(
		&logLevel,
		logLevelLabel,
		"l",
		"INFO",
		"logging level (INFO or DEBUG)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&cfgFile,
		"config",
		"f",
		"",
		"config file (default is $PWD/.hestia.yaml)",
	)

	_ = viper.BindPFlag(logLevelLabel, rootCmd.PersistentFlags().Lookup(logLevelLabel))

	rand.Seed(time.Now().UnixNano())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}// else {
	//	viper.SetConfigName("." + rootCmd.Use)
	//	viper.AddConfigPath("/etc/" + rootCmd.Use)
	//	viper.AddConfigPath("/$HOME/" + rootCmd.Use)
	//	viper.AddConfigPath(".")
	//}

	viper.SetEnvPrefix(rootCmd.Use)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	if err := viper.WriteConfigAs("." + rootCmd.Use + ".yaml"); err != nil {
		fmt.Println("Cannot write config file:", err)
	}
}
