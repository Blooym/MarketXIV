/*
Copyright © 2022 BitsOfAByte

*/

package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
func sendGetRequest(url string, params ...interface{}) []byte {

	queryString := formatParams(params...)

	Debug(fmt.Sprintf("GET \"%s?%s\"", url, queryString))

	resp, err := http.Get(fmt.Sprintf("%s?%s", url, queryString))

	if err != nil {
		Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return bodyBytes
}

// Fetches the current tax rates for the specified server.
func FetchTaxRates(server string) structures.ApiTaxRegions {
	response := sendGetRequest("https://universalis.app/api/tax-rates", fmt.Sprintf("world=%s", server))
	response = []byte(strings.Replace(string(response), "Ul'dah", "Uldah", -1))
	var taxRateData structures.ApiTaxRegions
	json.Unmarshal(response, &taxRateData)
	return taxRateData
}

// Searches for the given item from the API.
func FetchSearch(query string, indexes ...string) structures.ApiSearchResult {
	response := sendGetRequest("https://xivapi.com/search", fmt.Sprintf("string=%s", query), "limit=1", fmt.Sprintf("indexes=%s", strings.Join(indexes, ",")))
	var itemData structures.ApiSearchResult
	json.Unmarshal(response, &itemData)
	return itemData
}

// Fetches information about the given item from the API.
func FetchItem(itemID int) structures.ApiItem {
	response := sendGetRequest(fmt.Sprintf("https://xivapi.com/item/%v", itemID))
	var itemData structures.ApiItem
	json.Unmarshal(response, &itemData)
	return itemData
}

// Fetches information about the given item from the market of the given server.
func FetchMarketItem(server string, itemID int, limit int, hq string) structures.ApiMarketItem {
	response := sendGetRequest(fmt.Sprintf("https://universalis.app/api/%s/%d", server, itemID), fmt.Sprintf("listings=%d&hq=%s", limit, hq))
	var marketItemData structures.ApiMarketItem
	json.Unmarshal(response, &marketItemData)
	return marketItemData
}
