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
type ApiSearchResult struct {
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

type ApiItem struct {
	AdditionalData struct {
		Color    int    `json:"Color"`
		ID       int    `json:"ID"`
		Name     string `json:"Name"`
		NameDe   string `json:"Name_de"`
		NameEn   string `json:"Name_en"`
		NameFr   string `json:"Name_fr"`
		NameJa   string `json:"Name_ja"`
		Shade    int    `json:"Shade"`
		SubOrder int    `json:"SubOrder"`
	} `json:"AdditionalData"`
	AdditionalDataTarget      string      `json:"AdditionalDataTarget"`
	AdditionalDataTargetID    int         `json:"AdditionalDataTargetID"`
	Adjective                 int         `json:"Adjective"`
	AetherialReduce           int         `json:"AetherialReduce"`
	AlwaysCollectable         int         `json:"AlwaysCollectable"`
	Article                   int         `json:"Article"`
	BaseParam0                interface{} `json:"BaseParam0"`
	BaseParam0Target          string      `json:"BaseParam0Target"`
	BaseParam0TargetID        int         `json:"BaseParam0TargetID"`
	BaseParam1                interface{} `json:"BaseParam1"`
	BaseParam1Target          string      `json:"BaseParam1Target"`
	BaseParam1TargetID        int         `json:"BaseParam1TargetID"`
	BaseParam2                interface{} `json:"BaseParam2"`
	BaseParam2Target          string      `json:"BaseParam2Target"`
	BaseParam2TargetID        int         `json:"BaseParam2TargetID"`
	BaseParam3                interface{} `json:"BaseParam3"`
	BaseParam3Target          string      `json:"BaseParam3Target"`
	BaseParam3TargetID        int         `json:"BaseParam3TargetID"`
	BaseParam4                interface{} `json:"BaseParam4"`
	BaseParam4Target          string      `json:"BaseParam4Target"`
	BaseParam4TargetID        int         `json:"BaseParam4TargetID"`
	BaseParam5                interface{} `json:"BaseParam5"`
	BaseParam5Target          string      `json:"BaseParam5Target"`
	BaseParam5TargetID        int         `json:"BaseParam5TargetID"`
	BaseParamModifier         int         `json:"BaseParamModifier"`
	BaseParamSpecial0         interface{} `json:"BaseParamSpecial0"`
	BaseParamSpecial0Target   string      `json:"BaseParamSpecial0Target"`
	BaseParamSpecial0TargetID int         `json:"BaseParamSpecial0TargetID"`
	BaseParamSpecial1         interface{} `json:"BaseParamSpecial1"`
	BaseParamSpecial1Target   string      `json:"BaseParamSpecial1Target"`
	BaseParamSpecial1TargetID int         `json:"BaseParamSpecial1TargetID"`
	BaseParamSpecial2         interface{} `json:"BaseParamSpecial2"`
	BaseParamSpecial2Target   string      `json:"BaseParamSpecial2Target"`
	BaseParamSpecial2TargetID int         `json:"BaseParamSpecial2TargetID"`
	BaseParamSpecial3         interface{} `json:"BaseParamSpecial3"`
	BaseParamSpecial3Target   string      `json:"BaseParamSpecial3Target"`
	BaseParamSpecial3TargetID int         `json:"BaseParamSpecial3TargetID"`
	BaseParamSpecial4         interface{} `json:"BaseParamSpecial4"`
	BaseParamSpecial4Target   string      `json:"BaseParamSpecial4Target"`
	BaseParamSpecial4TargetID int         `json:"BaseParamSpecial4TargetID"`
	BaseParamSpecial5         interface{} `json:"BaseParamSpecial5"`
	BaseParamSpecial5Target   string      `json:"BaseParamSpecial5Target"`
	BaseParamSpecial5TargetID int         `json:"BaseParamSpecial5TargetID"`
	BaseParamValue0           int         `json:"BaseParamValue0"`
	BaseParamValue1           int         `json:"BaseParamValue1"`
	BaseParamValue2           int         `json:"BaseParamValue2"`
	BaseParamValue3           int         `json:"BaseParamValue3"`
	BaseParamValue4           int         `json:"BaseParamValue4"`
	BaseParamValue5           int         `json:"BaseParamValue5"`
	BaseParamValueSpecial0    int         `json:"BaseParamValueSpecial0"`
	BaseParamValueSpecial1    int         `json:"BaseParamValueSpecial1"`
	BaseParamValueSpecial2    int         `json:"BaseParamValueSpecial2"`
	BaseParamValueSpecial3    int         `json:"BaseParamValueSpecial3"`
	BaseParamValueSpecial4    int         `json:"BaseParamValueSpecial4"`
	BaseParamValueSpecial5    int         `json:"BaseParamValueSpecial5"`
	Block                     int         `json:"Block"`
	BlockRate                 int         `json:"BlockRate"`
	CanBeHq                   int         `json:"CanBeHq"`
	CastTimeS                 int         `json:"CastTimeS"`
	ClassJobCategory          interface{} `json:"ClassJobCategory"`
	ClassJobCategoryTarget    string      `json:"ClassJobCategoryTarget"`
	ClassJobCategoryTargetID  int         `json:"ClassJobCategoryTargetID"`
	ClassJobRepair            interface{} `json:"ClassJobRepair"`
	ClassJobRepairTarget      string      `json:"ClassJobRepairTarget"`
	ClassJobRepairTargetID    int         `json:"ClassJobRepairTargetID"`
	ClassJobUse               interface{} `json:"ClassJobUse"`
	ClassJobUseTarget         string      `json:"ClassJobUseTarget"`
	ClassJobUseTargetID       int         `json:"ClassJobUseTargetID"`
	CooldownS                 int         `json:"CooldownS"`
	DamageMag                 int         `json:"DamageMag"`
	DamagePhys                int         `json:"DamagePhys"`
	DefenseMag                int         `json:"DefenseMag"`
	DefensePhys               int         `json:"DefensePhys"`
	DelayMs                   int         `json:"DelayMs"`
	Description               string      `json:"Description"`
	DescriptionJSON           []string    `json:"DescriptionJSON"`
	DescriptionJSONCn         interface{} `json:"DescriptionJSON_cn"`
	DescriptionJSONDe         []string    `json:"DescriptionJSON_de"`
	DescriptionJSONEn         []string    `json:"DescriptionJSON_en"`
	DescriptionJSONFr         []string    `json:"DescriptionJSON_fr"`
	DescriptionJSONJa         []string    `json:"DescriptionJSON_ja"`
	DescriptionJSONKr         interface{} `json:"DescriptionJSON_kr"`
	DescriptionCn             interface{} `json:"Description_cn"`
	DescriptionDe             string      `json:"Description_de"`
	DescriptionEn             string      `json:"Description_en"`
	DescriptionFr             string      `json:"Description_fr"`
	DescriptionJa             string      `json:"Description_ja"`
	DescriptionKr             interface{} `json:"Description_kr"`
	Desynth                   int         `json:"Desynth"`
	EquipRestriction          int         `json:"EquipRestriction"`
	EquipSlotCategory         interface{} `json:"EquipSlotCategory"`
	EquipSlotCategoryTarget   string      `json:"EquipSlotCategoryTarget"`
	EquipSlotCategoryTargetID int         `json:"EquipSlotCategoryTargetID"`
	FilterGroup               int         `json:"FilterGroup"`
	GameContentLinks          struct {
		GilShopItem struct {
			Item []string `json:"Item"`
		} `json:"GilShopItem"`
		Quest struct {
			ItemReward0         []int `json:"ItemReward0"`
			ItemReward00        []int `json:"ItemReward00"`
			ItemReward10        []int `json:"ItemReward10"`
			ItemReward11        []int `json:"ItemReward11"`
			OptionalItemReward1 []int `json:"OptionalItemReward1"`
		} `json:"Quest"`
		Recipe struct {
			ItemResult []int `json:"ItemResult"`
		} `json:"Recipe"`
		StainTransient struct {
			Item1 []int `json:"Item1"`
		} `json:"StainTransient"`
	} `json:"GameContentLinks"`
	GamePatch struct {
		Banner      string `json:"Banner"`
		ExName      string `json:"ExName"`
		ExVersion   int    `json:"ExVersion"`
		ID          int    `json:"ID"`
		Name        string `json:"Name"`
		NameCn      string `json:"Name_cn"`
		NameDe      string `json:"Name_de"`
		NameEn      string `json:"Name_en"`
		NameFr      string `json:"Name_fr"`
		NameJa      string `json:"Name_ja"`
		NameKr      string `json:"Name_kr"`
		ReleaseDate int    `json:"ReleaseDate"`
		Version     string `json:"Version"`
	} `json:"GamePatch"`
	GrandCompany               interface{} `json:"GrandCompany"`
	GrandCompanyTarget         string      `json:"GrandCompanyTarget"`
	GrandCompanyTargetID       int         `json:"GrandCompanyTargetID"`
	ID                         int         `json:"ID"`
	Icon                       string      `json:"Icon"`
	IconHD                     string      `json:"IconHD"`
	IconID                     int         `json:"IconID"`
	IsAdvancedMeldingPermitted int         `json:"IsAdvancedMeldingPermitted"`
	IsCollectable              int         `json:"IsCollectable"`
	IsCrestWorthy              int         `json:"IsCrestWorthy"`
	IsDyeable                  int         `json:"IsDyeable"`
	IsGlamourous               int         `json:"IsGlamourous"`
	IsIndisposable             int         `json:"IsIndisposable"`
	IsPvP                      int         `json:"IsPvP"`
	IsUnique                   int         `json:"IsUnique"`
	IsUntradable               int         `json:"IsUntradable"`
	ItemAction                 interface{} `json:"ItemAction"`
	ItemActionTarget           string      `json:"ItemActionTarget"`
	ItemActionTargetID         int         `json:"ItemActionTargetID"`
	ItemGlamour                interface{} `json:"ItemGlamour"`
	ItemGlamourTarget          string      `json:"ItemGlamourTarget"`
	ItemGlamourTargetID        int         `json:"ItemGlamourTargetID"`
	ItemKind                   struct {
		ID     int    `json:"ID"`
		Name   string `json:"Name"`
		NameCn string `json:"Name_cn"`
		NameDe string `json:"Name_de"`
		NameEn string `json:"Name_en"`
		NameFr string `json:"Name_fr"`
		NameJa string `json:"Name_ja"`
		NameKr string `json:"Name_kr"`
	} `json:"ItemKind"`
	ItemRepair         interface{} `json:"ItemRepair"`
	ItemRepairTarget   string      `json:"ItemRepairTarget"`
	ItemRepairTargetID int         `json:"ItemRepairTargetID"`
	ItemSearchCategory struct {
		Category         int         `json:"Category"`
		ClassJob         interface{} `json:"ClassJob"`
		ClassJobTarget   string      `json:"ClassJobTarget"`
		ClassJobTargetID int         `json:"ClassJobTargetID"`
		ID               int         `json:"ID"`
		Icon             string      `json:"Icon"`
		IconHD           string      `json:"IconHD"`
		IconID           int         `json:"IconID"`
		Name             string      `json:"Name"`
		NameDe           string      `json:"Name_de"`
		NameEn           string      `json:"Name_en"`
		NameFr           string      `json:"Name_fr"`
		NameJa           string      `json:"Name_ja"`
		Order            int         `json:"Order"`
	} `json:"ItemSearchCategory"`
	ItemSearchCategoryTarget   string      `json:"ItemSearchCategoryTarget"`
	ItemSearchCategoryTargetID int         `json:"ItemSearchCategoryTargetID"`
	ItemSeries                 interface{} `json:"ItemSeries"`
	ItemSeriesTarget           string      `json:"ItemSeriesTarget"`
	ItemSeriesTargetID         int         `json:"ItemSeriesTargetID"`
	ItemSortCategory           struct {
		ID    int `json:"ID"`
		Param int `json:"Param"`
	} `json:"ItemSortCategory"`
	ItemSortCategoryTarget   string      `json:"ItemSortCategoryTarget"`
	ItemSortCategoryTargetID int         `json:"ItemSortCategoryTargetID"`
	ItemSpecialBonus         interface{} `json:"ItemSpecialBonus"`
	ItemSpecialBonusParam    int         `json:"ItemSpecialBonusParam"`
	ItemSpecialBonusTarget   string      `json:"ItemSpecialBonusTarget"`
	ItemSpecialBonusTargetID int         `json:"ItemSpecialBonusTargetID"`
	ItemUICategory           struct {
		ID         int    `json:"ID"`
		Icon       string `json:"Icon"`
		IconHD     string `json:"IconHD"`
		IconID     int    `json:"IconID"`
		Name       string `json:"Name"`
		NameDe     string `json:"Name_de"`
		NameEn     string `json:"Name_en"`
		NameFr     string `json:"Name_fr"`
		NameJa     string `json:"Name_ja"`
		OrderMajor int    `json:"OrderMajor"`
		OrderMinor int    `json:"OrderMinor"`
	} `json:"ItemUICategory"`
	ItemUICategoryTarget   string      `json:"ItemUICategoryTarget"`
	ItemUICategoryTargetID int         `json:"ItemUICategoryTargetID"`
	LevelEquip             int         `json:"LevelEquip"`
	LevelItem              int         `json:"LevelItem"`
	Lot                    int         `json:"Lot"`
	Materia                interface{} `json:"Materia"`
	MateriaSlotCount       int         `json:"MateriaSlotCount"`
	MaterializeType        int         `json:"MaterializeType"`
	ModelMain              string      `json:"ModelMain"`
	ModelSub               string      `json:"ModelSub"`
	Name                   string      `json:"Name"`
	NameDe                 string      `json:"Name_de"`
	NameEn                 string      `json:"Name_en"`
	NameFr                 string      `json:"Name_fr"`
	NameJa                 string      `json:"Name_ja"`
	Patch                  int         `json:"Patch"`
	Plural                 string      `json:"Plural"`
	PluralDe               string      `json:"Plural_de"`
	PluralEn               string      `json:"Plural_en"`
	PluralFr               string      `json:"Plural_fr"`
	PluralJa               string      `json:"Plural_ja"`
	PossessivePronoun      int         `json:"PossessivePronoun"`
	PriceLow               int         `json:"PriceLow"`
	PriceMid               int         `json:"PriceMid"`
	Pronoun                int         `json:"Pronoun"`
	Rarity                 int         `json:"Rarity"`
	Recipes                []struct {
		ClassJobID int `json:"ClassJobID"`
		ID         int `json:"ID"`
		Level      int `json:"Level"`
	} `json:"Recipes"`
	Singular        string `json:"Singular"`
	SingularDe      string `json:"Singular_de"`
	SingularEn      string `json:"Singular_en"`
	SingularFr      string `json:"Singular_fr"`
	SingularJa      string `json:"Singular_ja"`
	StackSize       int    `json:"StackSize"`
	StartsWithVowel int    `json:"StartsWithVowel"`
	SubStatCategory int    `json:"SubStatCategory"`
	URL             string `json:"Url"`
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
