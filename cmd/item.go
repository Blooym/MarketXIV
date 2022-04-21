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

		hq, _ := cmd.Flags().GetBool("hq")
		limit, _ := cmd.Flags().GetInt("limit")

		serverName := args[0]
		itemName := strings.Join(args[1:], " ")
		itemData := backend.FetchItem(itemName)

		// Check to see if the item exists
		if len(itemData.Results) == 0 {
			fmt.Println("No results found for " + itemName)
			return
		}

		resultData := itemData.Results[0]
		marketData := backend.FetchMarketItem(serverName, resultData.ID, limit, strconv.FormatBool(hq))

		if len(marketData.Listings) == 0 {
			fmt.Println("No listings found for " + itemName)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Quality", "Price", "Quantity", "Total", "Retainer", "World"})
		table.SetFooter([]string{
			"Information",
			fmt.Sprintf("Item: %s", resultData.Name),
			fmt.Sprintf("ID: %d", marketData.ItemID),
			fmt.Sprintf("%v Shown", len(marketData.Listings)),
			fmt.Sprintf("Avg Price: %v", marketData.AveragePrice),
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
