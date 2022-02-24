// Formatting and Handling of the raw market data from the API.

package cli

import (
	"fmt"
	"strings"

	"github.com/BitsOfAByte/marketxiv/backend"
)

type Market struct {
	Server string
}

// Outputs the current market tax rates.
func (m Market) TaxRates() {
	// Get the tax rates
	data := backend.FetchTaxRates(m.Server)
	server := strings.Title(strings.ToLower(m.Server))

	// If the starting cities have rates equal to zero then there is no tax data.
	// This is a workaround for checking if the server is valid as starting cities
	// Do not typically have tax rates of zero.
	if data.LimsaLominsa == 0 && data.UlDah == 0 && data.Gridania == 0 {
		fmt.Println(fmt.Sprintf("marketxiv: No tax rates found for %s.", server))
		return
	}

	//  Print information about item
	fmt.Println(fmt.Sprintf("Tax Rates for %s", server))
	fmt.Println(fmt.Sprintf("  Limsa Lominsa: %d", data.LimsaLominsa) + "%")
	fmt.Println(fmt.Sprintf("  Gridania: %d", data.Gridania) + "%")
	fmt.Println(fmt.Sprintf("  Uldah: %d", data.UlDah) + "%")
	fmt.Println(fmt.Sprintf("  Ishgard: %d", data.Ishgard) + "%")
	fmt.Println(fmt.Sprintf("  Kugane: %d", data.Kugane) + "%")
	fmt.Println(fmt.Sprintf("  Crystarium: %d", data.Crystarium) + "%")
	fmt.Println(fmt.Sprintf("  Old Sharlayan: %d", data.OldSharlayan) + "%")
}
