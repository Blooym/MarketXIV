/*
Copyright © 2022 Blooym

MIT License, see the LICENSE file for more information.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Blooym/marketxiv/backend"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update to the latest version of the app if available",
	Run: func(cmd *cobra.Command, args []string) {
		forceFlag := cmd.Flags().Lookup("force").Value.String()
		if forceFlag == "true" {
			backend.Update(RootCmd.Version)
		} else {
			fmt.Println("WARNING! You should not use the app-update command unless you have a manual installation.")
			fmt.Println("If you are trying to update the app and have installed it with a package manager, use that instead.")
			fmt.Println("If you are ABSOLUTELY sure you want to update, use the --force flag.")
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)

	updateCmd.Flags().BoolP("force", "f", false, "Force updater to run")
}
