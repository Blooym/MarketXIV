package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BitsOfAByte/marketxiv/backend"
	"github.com/BitsOfAByte/marketxiv/cli"
)

func main() {

	if len(os.Args) < 2 {
		cli.ShowQuickHelp()
		return
	}

	switch os.Args[1] {
	case "--help":
		cli.ShowHelpPage()
	case "--info":
		cli.ShowInfo()
	case "--update":
		/* If the version is N/A, it likely means the binary was installed with `go install`
		In this case, the best bet is to just update the binary anyway to avoid future issues */
		if cli.Version == "N/A" {
			backend.Update("v0.0.0")
		} else {
			backend.Update(cli.Version)
		}
	case "item":
		ItemLookup()
	case "tax":
		MarketTax()
	default:
		cli.ShowQuickHelp()
	}
}

func ItemLookup() {
	if len(os.Args) < 4 {
		fmt.Println("Specify a server and item name. Use --help for more information.")
		return
	}

	item := cli.MarketItem{ItemName: strings.Join(os.Args[3:], " "), Server: os.Args[2]}
	item.ShowListings(10)
}

func MarketTax() {
	if len(os.Args) < 3 {
		fmt.Printf("Specify a server. Use --help for more information.")
		return
	}

	market := cli.Market{Server: os.Args[2]}
	market.TaxRates()
}
