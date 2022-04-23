/*
Copyright © 2022 BitsOfAByte

*/

package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// itemCmd represents the item command
var itemCmd = &cobra.Command{
	Use:        "item <server> <item>",
	Deprecated: "Use 'listings' instead",
	Short:      "Get the current market listings for the specified item",
	Args:       cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		hq, _ := cmd.Flags().GetBool("hq")
		limit, _ := cmd.Flags().GetInt("limit")

		serverName := args[0]
		itemName := strings.Join(args[1:], " ")
		searchData := backend.FetchSearch(itemName, "item")

		// Check to see if the item exists
		if len(searchData.Results) == 0 {
			fmt.Println("No results found for " + itemName)
			return
		}

		resultData := searchData.Results[0]
		marketData := backend.FetchMarketItem(serverName, resultData.ID, limit, strconv.FormatBool(hq))

		if len(marketData.Listings) == 0 {
			fmt.Println("No listings found for " + itemName)
			return
		}

		itemData := backend.FetchItem(resultData.ID)

		if itemData.PriceMid <= marketData.Listings[0].PricePerUnit && itemData.PriceMid != 0 {
			fmt.Println("Note: This item may be cheaper to be at a vendor instead of the market.")
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Quality", "Price", "Quantity", "Total", "Retainer", "World"})
		table.SetFooter([]string{
			fmt.Sprintf("%s (%d)", resultData.Name, marketData.ItemID),
			fmt.Sprintf("Avg: %v", marketData.AveragePrice),
			" ",
			" ",
			" ",
			time.Unix(marketData.LastUploadTime/1000, 0).Format("2006-01-02 15:04:05"),
		})

		// Format and display the data
		for _, listing := range marketData.Listings {
			world := listing.WorldName
			quality := strconv.FormatBool(listing.Hq)

			if world == "" {
				world = serverName
			}

			if quality == "false" {
				quality = "Normal"
			} else {
				quality = "High"
			}

			table.Append([]string{
				quality,
				strconv.Itoa(listing.PricePerUnit),
				strconv.Itoa(listing.Quantity),
				strconv.Itoa(listing.Total),
				listing.RetainerName,
				world,
			})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(itemCmd)

	itemCmd.Flags().Bool("hq", false, "Only fetch high quality listings")
	itemCmd.Flags().IntP("limit", "l", 5, "Limit the number of listings to show")
}
