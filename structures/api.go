// Structures used by the app to handle data
package structures

// Format the data returned from the tax API
type ApiTaxRegions struct {
	LimsaLominsa int `json:"Limsa Lominsa"`
	Gridania     int `json:"Gridania"`
	UlDah        int `json:"Uldah"`
	Ishgard      int `json:"Ishgard"`
	Kugane       int `json:"Kugane"`
	Crystarium   int `json:"Crystarium"`
	OldSharlayan int `json:"Old Sharlayan"`
}

// Stucture the information returned about an in-game item
type ApiItem struct {
	Pagination struct {
		Page           int         `json:"Page"`
		PageNext       int         `json:"PageNext"`
		PagePrev       interface{} `json:"PagePrev"`
		PageTotal      int         `json:"PageTotal"`
		Results        int         `json:"Results"`
		ResultsPerPage int         `json:"ResultsPerPage"`
		ResultsTotal   int         `json:"ResultsTotal"`
	} `json:"Pagination"`
	Results []struct {
		ID            int    `json:"ID"`
		Icon          string `json:"Icon"`
		Name          string `json:"Name"`
		URL           string `json:"Url"`
		URLType       string `json:"UrlType"`
		NAMING_FAILED string `json:"_"`
		Score         int    `json:"_Score"`
	} `json:"Results"`
	SpeedMs int `json:"SpeedMs"`
}

// Structure data returned about an item on the in-game market.
type ApiMarketItem struct {
	ItemID         int   `json:"itemID"`
	LastUploadTime int64 `json:"lastUploadTime"`
	Listings       []struct {
		LastReviewTime int           `json:"lastReviewTime"`
		PricePerUnit   int           `json:"pricePerUnit"`
		Quantity       int           `json:"quantity"`
		StainID        int           `json:"stainID"`
		WorldName      string        `json:"worldName"`
		WorldID        int           `json:"worldID"`
		CreatorName    string        `json:"creatorName"`
		CreatorID      string        `json:"creatorID"`
		Hq             bool          `json:"hq"`
		IsCrafted      bool          `json:"isCrafted"`
		ListingID      interface{}   `json:"listingID"`
		Materia        []interface{} `json:"materia"`
		OnMannequin    bool          `json:"onMannequin"`
		RetainerCity   int           `json:"retainerCity"`
		RetainerID     string        `json:"retainerID"`
		RetainerName   string        `json:"retainerName"`
		SellerID       string        `json:"sellerID"`
		Total          int           `json:"total"`
	} `json:"listings"`
	RecentHistory []struct {
		Hq           bool   `json:"hq"`
		PricePerUnit int    `json:"pricePerUnit"`
		Quantity     int    `json:"quantity"`
		Timestamp    int    `json:"timestamp"`
		WorldName    string `json:"worldName"`
		WorldID      int    `json:"worldID"`
		BuyerName    string `json:"buyerName"`
		Total        int    `json:"total"`
	} `json:"recentHistory"`
	DcName                string  `json:"dcName"`
	CurrentAveragePrice   float64 `json:"currentAveragePrice"`
	CurrentAveragePriceNQ float64 `json:"currentAveragePriceNQ"`
	CurrentAveragePriceHQ float64 `json:"currentAveragePriceHQ"`
	RegularSaleVelocity   float64 `json:"regularSaleVelocity"`
	NqSaleVelocity        float64 `json:"nqSaleVelocity"`
	HqSaleVelocity        float64 `json:"hqSaleVelocity"`
	AveragePrice          float64 `json:"averagePrice"`
	AveragePriceNQ        float64 `json:"averagePriceNQ"`
	AveragePriceHQ        float64 `json:"averagePriceHQ"`
	MinPrice              int     `json:"minPrice"`
	MinPriceNQ            int     `json:"minPriceNQ"`
	MinPriceHQ            int     `json:"minPriceHQ"`
	MaxPrice              int     `json:"maxPrice"`
	MaxPriceNQ            int     `json:"maxPriceNQ"`
	MaxPriceHQ            int     `json:"maxPriceHQ"`
	StackSizeHistogram    struct {
		Num1 int `json:"1"`
	} `json:"stackSizeHistogram"`
	StackSizeHistogramNQ struct {
		Num1 int `json:"1"`
	} `json:"stackSizeHistogramNQ"`
	StackSizeHistogramHQ struct {
		Num1 int `json:"1"`
	} `json:"stackSizeHistogramHQ"`
	WorldUploadTimes struct {
		Num33 int64 `json:"33"`
		Num36 int64 `json:"36"`
		Num42 int64 `json:"42"`
		Num56 int64 `json:"56"`
		Num66 int64 `json:"66"`
		Num67 int64 `json:"67"`
	} `json:"worldUploadTimes"`
}
