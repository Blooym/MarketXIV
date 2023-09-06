package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Blooym/marketxiv/backend"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// itemCmd represents the item command
var listingsCmd = &cobra.Command{
	Use:   "listings <server> <item>",
	Short: "Get the current market listings for the specified item",
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
		table.SetHeader([]string{"Quality", "Price", "Quantity", "Total", "Retainer", "World"})

		// Format and display the data
		for _, listing := range marketData.Listings {
			world := listing.WorldName

			if world == "" {
				world = serverName
			}

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
				listing.RetainerName,
				world,
			})
		}

		fmt.Println("Showing listings for Item ID " + strconv.Itoa(itemID))
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(listingsCmd)
	listingsCmd.Flags().Bool("hq", false, "Only fetch high quality listings")
	listingsCmd.Flags().IntP("limit", "l", 5, "Limit the number of listings to show")
}
