// .*@mycompany\.com MY COMPANY 2025
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/eccles/hestia/logger"
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
	logger.New(logLevel)
	rootLogger = logger.RootLogger
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
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} // else {
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
