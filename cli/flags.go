package cli

import (
	"fmt"
)

var (
	Version   string = "Non-Direct Binary"
	BuildDate string = "Non-Direct Binary"
	BuildBy   string = "Non-Direct Binary"
)

func ShowHelpPage() {
	fmt.Println("MarketXIV Manual")
	fmt.Println("  ==================")
	fmt.Println("  About:")
	fmt.Println("    MarketXIV is a command line tool for interacting with the FFXIV Online market via Universalis.")
	fmt.Println("    You can view the current market tax rates, or search for an item and view its listings.")
	fmt.Println("  ==================")
	fmt.Println("  Usage:")
	fmt.Println("    marketxiv tax <server>, get the current market tax rates for the specified server.")
	fmt.Println("    marketxiv item <server> <item>, get the current market listings for the specified item.")
	fmt.Println("  ==================")
	fmt.Println("  Options:")
	fmt.Println("    --help, shows this help page.")
	fmt.Println("    --info, shows build info for marketxiv.")
}

func ShowQuickHelp() {
	fmt.Println("marketxiv: try 'marketxiv --help' for more information.")
}

func ShowInfo() {
	fmt.Println("MarketXIV Build Information:")
	fmt.Println("  Version: " + Version)
	fmt.Println("  Build Date: " + BuildDate)
	fmt.Println("  Build By: " + BuildBy)
}
