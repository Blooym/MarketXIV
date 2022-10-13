/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"os"

	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "marketxiv",
	Short: "MarketXIV is a command line tool for interacting with the FFXIV Online market via Universalis.",
	Long: `MarketXIV is a command line tool for interacting with the FFXIV Online market via Universalis. 
You can view the current market tax rates, or search for an item and view its listings.`,
	Version: backend.Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Set the default configuration file
	configDir, _ := os.UserConfigDir()
	viper.SetConfigName("config")
	viper.AddConfigPath(configDir + "/marketxiv")
	viper.SetConfigType("json")

	// Set the default values for the configuration
	viper.SetDefault("app.verbose", false)

	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// If the configuration file does not exist, create it here.
			os.MkdirAll(configDir+"/marketxiv", os.ModePerm)
			viper.SafeWriteConfig()
		} else {
			// If another error occured while trying to read the configuration file, exit.
			backend.Fatal(err)
		}
	}
}
