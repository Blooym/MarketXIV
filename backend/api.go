package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/BitsOfAByte/marketxiv/structures"
)

// Formats a list of parameters into a query string.
func formatParams(params ...interface{}) string {
	var formattedParams string

	for _, param := range params {

		if strings.Contains(fmt.Sprintf("%v", param), " ") {
			param = strings.Replace(fmt.Sprintf("%v", param), " ", "%20", -1)
		}

		formattedParams += strings.ToLower(fmt.Sprintf("%v&", param))
	}
	return formattedParams
}

// GET data from the URL with params params, response is given as a Byte array.
func SendGetRequest(url string, params ...interface{}) []byte {

	queryString := formatParams(params...)

	verboseLog(fmt.Sprintf("GET %s?%s", url, queryString))

	resp, err := http.Get(fmt.Sprintf("%s?%s", url, queryString))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return bodyBytes
}

// Fetches the current tax rates for the specified server.
func FetchTaxRates(server string) structures.ApiTaxRegions {
	response := SendGetRequest("https://universalis.app/api/tax-rates", fmt.Sprintf("world=%s", server))
	response = []byte(strings.Replace(string(response), "Ul'dah", "Uldah", -1))
	var taxRateData structures.ApiTaxRegions
	json.Unmarshal(response, &taxRateData)
	return taxRateData
}

// Fetches information about the given item from the API.
func FetchItem(itemName string) structures.ApiItem {
	response := SendGetRequest("https://xivapi.com/search", fmt.Sprintf("string=%s", itemName), "limit=1")
	var itemData structures.ApiItem
	json.Unmarshal(response, &itemData)
	return itemData
}

// Fethces information about the given item from the market of the given server.
func FetchMarketItem(server string, itemID int, limit int) structures.ApiMarketItem {
	response := SendGetRequest(fmt.Sprintf("https://universalis.app/api/%s/%d", server, itemID), fmt.Sprintf("listings=%d", limit))
	var marketItemData structures.ApiMarketItem
	json.Unmarshal(response, &marketItemData)
	return marketItemData
}
