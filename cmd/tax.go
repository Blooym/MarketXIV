/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// taxCmd represents the tax command
var taxCmd = &cobra.Command{
	Use:        "tax",
	Short:      "Get the current market tax rates for the specified server",
	Args:       cobra.ExactArgs(1),
	SuggestFor: []string{"rates", "taxes"},
	Run: func(cmd *cobra.Command, args []string) {

		taxServer := args[0]
		taxData := backend.FetchTaxRates(taxServer)

		// If the starting cities have rates equal to zero then there is no tax data.
		if taxData.LimsaLominsa == 0 && taxData.UlDah == 0 && taxData.Gridania == 0 {
			fmt.Println("No tax rates found for", taxServer)
			return
		}

		// Prepare the table
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"City", "Tax Rate"})

		// Add the data to the table
		table.Append([]string{"Limsa Lominsa", strconv.Itoa(taxData.LimsaLominsa) + "%"})
		table.Append([]string{"Ul'dah", strconv.Itoa(taxData.UlDah) + "%"})
		table.Append([]string{"Gridania", strconv.Itoa(taxData.Gridania) + "%"})
		table.Append([]string{"Ishgard", strconv.Itoa(taxData.Ishgard) + "%"})
		table.Append([]string{"Kugane", strconv.Itoa(taxData.Kugane) + "%"})
		table.Append([]string{"The Crystarium", strconv.Itoa(taxData.Crystarium) + "%"})
		table.Append([]string{"Old Sharlayan", strconv.Itoa(taxData.OldSharlayan) + "%"})

		// Show the table
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(taxCmd)
}
