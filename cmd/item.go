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
	Use:   "item <server> <item>",
	Short: "Get the current market listings for the specified item",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		server := args[0]
		itemName := strings.Join(args[1:], " ")

		itemApiResponse := backend.FetchItem(itemName)

		if len(itemApiResponse.Results) == 0 {
			fmt.Println("No listings found for", itemName)
			return
		}

		itemData := itemApiResponse.Results[0]
		marketItemData := backend.FetchMarketItem(server, itemData.ID, 10)

		if len(marketItemData.Listings) == 0 {
			fmt.Println(itemName, "is not a marketable item.")
			return
		}

		// Prepare a table to display the data
		listingData := make([][]string, len(marketItemData.Listings))
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Quality", "Price", "Quantity", "Total", "Retainer", "World"})

		// Loop through the listings and add them their data to the table
		for _, listing := range marketItemData.Listings {

			// Handle the quality boolean and convert it to a string.
			var quality string
			switch listing.Hq {
			case true:
				quality = "HQ"
			default:
				quality = "NQ"
			}

			// When there is no world name, assume the world is only the server entered.
			if listing.WorldName == "" {
				listing.WorldName = strings.ToLower(server)
			}

			// Add the listing to the table
			listingData = append(listingData, []string{
				quality,
				strconv.Itoa(listing.PricePerUnit),
				strconv.Itoa(listing.Quantity),
				strconv.Itoa(listing.Total),
				listing.RetainerName,
				listing.WorldName,
			})
		}

		// Show the table
		table.AppendBulk(listingData)
		fmt.Println("Showing listings for", itemName, "| Uploaded:", time.Unix(marketItemData.LastUploadTime/1000, 0).Format(time.RFC1123))
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(itemCmd)
}
