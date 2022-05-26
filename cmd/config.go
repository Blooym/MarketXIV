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

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "View the directory where the configuraiton file is stored",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.ConfigFileUsed())
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

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(verboseCmd)
	configCmd.AddCommand(dirCmd)
}
