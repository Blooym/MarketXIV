package cli

import (
	"fmt"
	"strings"

	"github.com/BitsOfAByte/marketxiv/backend"
)

// Information about an item that is marketable.
type MarketItem struct {
	ItemName string
	Server   string
}

// Get information about the item from the market.
func (i MarketItem) Info() {
	itemID := backend.FetchItem(i.ItemName).Results[0].ID
	data := backend.FetchMarketItem(i.Server, itemID, 10)

	if len(data.Listings) == 0 {
		fmt.Println(fmt.Sprintf("marketxiv: No listings found for %s on %s.", i.ItemName, i.Server))
		return
	}

	// Print information about item
	fmt.Println(fmt.Sprintf("Showing listings for %s (%d) on %s", strings.Title(strings.ToLower(i.ItemName)), data.ItemID, strings.Title(strings.ToLower(i.Server))))
	fmt.Println(fmt.Sprintf("  %d listings found", len(data.Listings)))
	fmt.Println(fmt.Sprintf("    Lowest: %d Gil | Highest: %d Gil | Average: %e Gil", data.MinPrice, data.MaxPrice, data.AveragePrice))
	for _, listing := range data.Listings {
		var quality string
		if listing.Hq {
			quality = "High"
		} else {
			quality = "Regular"
		}
		if listing.WorldName == "" {
			fmt.Println(fmt.Sprintf("      (%v)  Quantity: %d | Total: %d Gil | Price-Per-Unit: %d Gil | Retainer: %s", quality, listing.Quantity, listing.Total, listing.PricePerUnit, listing.RetainerName))
		} else {
			fmt.Println(fmt.Sprintf("      (%v)  Quantity: %d | Total: %d Gil | Price-Per-Unit: %d Gil | Retainer: %s (%s)", quality, listing.Quantity, listing.Total, listing.PricePerUnit, listing.RetainerName, listing.WorldName))
		}
	}
}
