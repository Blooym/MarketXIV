/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View and edit the configuration for MarketXIV",
}

// analyticsCmd represents the analytics command
var analyticsCmd = &cobra.Command{
	Use:       "analytics <boolean>",
	Short:     "Enable or disable analytics collection",
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"true", "false"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "true":
			viper.Set("analytics.enabled", true)
			err := viper.WriteConfig()

			if err != nil {
				fmt.Println("Error writing config file: ", err)
			}

			fmt.Println("Analytics collection enabled")
		case "false":
			viper.Set("analytics.enabled", false)
			err := viper.WriteConfig()

			if err != nil {
				fmt.Println("Error writing config file: ", err)
			}

			fmt.Println("Analytics collection disabled")
		}
	},
}

// verboseCmd represents the verbose command
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

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "View your unique analytics ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("analytics.uuid"))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(analyticsCmd)
	configCmd.AddCommand(verboseCmd)
	configCmd.AddCommand(uuidCmd)
}
