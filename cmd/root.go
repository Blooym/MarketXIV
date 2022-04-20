/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Version string = "v0.0.0"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "marketxiv",
	Short: "MarketXIV is a command line tool for interacting with the FFXIV Online market via Universalis.",
	Long: `MarketXIV is a command line tool for interacting with the FFXIV Online market via Universalis. 
You can view the current market tax rates, or search for an item and view its listings.`,
	Version: Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
