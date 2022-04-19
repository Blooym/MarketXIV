package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/olekukonko/tablewriter"
)

type Market struct {
	Server string
}

// Outputs the current market tax rates.
func (m Market) TaxRates() {

	// Get the tax rates
	taxData := backend.FetchTaxRates(m.Server)
	taxServer := strings.ToTitle(strings.ToLower(m.Server))

	// If the starting cities have rates equal to zero then there is no tax data.
	if taxData.LimsaLominsa == 0 && taxData.UlDah == 0 && taxData.Gridania == 0 {
		fmt.Printf("marketxiv: No tax rates found for %s.", taxServer)
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
}
