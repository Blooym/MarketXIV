/*
Copyright © 2022 Blooym

MIT License, see the LICENSE file for more information.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and edit the configuration for MarketXIV",
}

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "View the directory where the configuration file is stored",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.ConfigFileUsed())
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "View the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		settingsMap := viper.AllSettings()
		for _, value := range settingsMap {
			for key2, value2 := range value.(map[string]interface{}) {
				fmt.Println(key2, "=", value2)
			}
		}
	},
}

var verboseCmd = &cobra.Command{
	Use:       "verbose <boolean>",
	Short:     "Enable or disable verbose output",
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"true", "false"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "true":
			viper.Set("app.verbose", true)
			err := viper.WriteConfig()

			if err != nil {
				fmt.Println("Error writing config file: ", err)
			}

			fmt.Println("Verbose output enabled")
		case "false":
			viper.Set("app.verbose", false)
			err := viper.WriteConfig()

			if err != nil {
				fmt.Println("Error writing config file: ", err)
			}

			fmt.Println("Verbose output disabled")
		}
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.AddCommand(verboseCmd)
	configCmd.AddCommand(dirCmd)
	configCmd.AddCommand(showCmd)
}
