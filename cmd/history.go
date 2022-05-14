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
var historyCmd = &cobra.Command{
	Use:   "history <server> <item>",
	Short: "Get historical market listings for the specified item",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		hq, _ := cmd.Flags().GetBool("hq")
		limit, _ := cmd.Flags().GetInt("limit")

		serverName := args[0]
		itemName := strings.Join(args[1:], " ")
		searchData := backend.FetchSearch(itemName)

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

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Quality", "Price", "Quantity", "Total", "Buyer", "Sold At"})
		table.SetFooter([]string{
			fmt.Sprintf("%s (%d)", resultData.Name, marketData.ItemID),
			fmt.Sprintf("Avg: %v", marketData.AveragePrice),
			" ",
			" ",
			" ",
			time.Unix(marketData.LastUploadTime/1000, 0).Format("2006-01-02 15:04:05"),
		})

		// Format and display the data
		for _, listing := range marketData.RecentHistory {
			quality := strconv.FormatBool(listing.Hq)

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
				listing.BuyerName,
				time.Unix(int64(listing.Timestamp), 0).Format("2006-01-02 15:04:05"),
			})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().Bool("hq", false, "Only fetch high quality listings")
	historyCmd.Flags().IntP("limit", "l", 5, "Limit the number of listings to show")
}
