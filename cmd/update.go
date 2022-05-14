/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update to the latest version of the app if available",
	Run: func(cmd *cobra.Command, args []string) {
		backend.Update(rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
