/*
Copyright © 2022 BitsOfAByte

MIT License, see the LICENSE file for more information.
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
var historyCmd = &cobra.Command{
	Use:   "history <server> <item>",
	Short: "Get historical market listings for the specified item",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		hq, _ := cmd.Flags().GetBool("hq")
		limit, _ := cmd.Flags().GetInt("limit")

		if limit > 100 {
			fmt.Println("Limit must be less than 100")
			os.Exit(1)
		}

		serverName := args[0]
		itemName := strings.Join(args[1:], " ")

		var itemID int

		if _, err := strconv.Atoi(itemName); err != nil {

			searchData := backend.FetchSearch(itemName, "item")

			// Check to see if the item exists
			if len(searchData.Results) == 0 {
				fmt.Println("No results found for " + itemName)
				return
			}

			resultData := searchData.Results[0]
			itemID = resultData.ID
		} else {
			itemID, _ = strconv.Atoi(itemName)
		}

		marketData := backend.FetchMarketItem(serverName, itemID, limit, strconv.FormatBool(hq))

		if len(marketData.Listings) == 0 {
			fmt.Println("No listings found for " + itemName)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Quality", "Price", "Quantity", "Total", "Buyer", "Sold At"})

		// Format and display the data
		for _, listing := range marketData.RecentHistory {
			quality := "Normal"
			switch listing.Hq {
			case true:
				quality = "HQ"
			case false:
				quality = "Normal"
			}

			table.Append([]string{
				quality,
				strconv.Itoa(listing.PricePerUnit),
				strconv.Itoa(listing.Quantity),
				strconv.Itoa(listing.Total),
				listing.BuyerName,
				time.Unix(int64(listing.Timestamp), 0).Format("02-01-06 15:04"),
			})
		}

		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(historyCmd)
	historyCmd.Flags().Bool("hq", false, "Only fetch high quality listings")
	historyCmd.Flags().IntP("limit", "l", 5, "Limit the number of listings to show")
}
