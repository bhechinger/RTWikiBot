package main

type StarSystemDef struct {
	BlackMarketShopItems interface{} `json:"BlackMarketShopItems"`
	ContractEmployers    []string    `json:"ContractEmployers"`
	ContractTargets      []string    `json:"ContractTargets"`
	CoreSystemID         string      `json:"CoreSystemID"`
	DefaultDifficulty    int64       `json:"DefaultDifficulty"`
	Depletable           bool        `json:"Depletable"`
	Description          struct {
		Details string `json:"Details"`
		Icon    string `json:"Icon"`
		ID      string `json:"Id"`
		Name    string `json:"Name"`
	} `json:"Description"`
	DifficultyList   []interface{} `json:"DifficultyList"`
	DifficultyModes  []interface{} `json:"DifficultyModes"`
	FactionShopItems []string      `json:"FactionShopItems"`
	FactionShopOwner string        `json:"FactionShopOwner"`
	FuelingStation   bool          `json:"FuelingStation"`
	JumpDistance     int64         `json:"JumpDistance"`
	MapExcludedTags  struct {
		Items            []interface{} `json:"items"`
		TagSetSourceFile string        `json:"tagSetSourceFile"`
	} `json:"MapExcludedTags"`
	MapRequiredTags struct {
		Items            []interface{} `json:"items"`
		TagSetSourceFile string        `json:"tagSetSourceFile"`
	} `json:"MapRequiredTags"`
	MaxContractOverride int64  `json:"MaxContractOverride"`
	Owner               string `json:"Owner"`
	Position            struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z int64   `json:"z"`
	} `json:"Position"`
	RoninHiringChance   int64       `json:"RoninHiringChance"`
	ShopMaxInventory    int64       `json:"ShopMaxInventory"`
	ShopMaxSpecials     int64       `json:"ShopMaxSpecials"`
	ShopRefreshRate     int64       `json:"ShopRefreshRate"`
	StarPosition        interface{} `json:"StarPosition"`
	StarType            string      `json:"StarType"`
	StartingSystemModes []string    `json:"StartingSystemModes"`
	SupportedBiomes     []string    `json:"SupportedBiomes"`
	SystemInfluence     []struct {
		Faction   string `json:"Faction"`
		Influence int64  `json:"Influence"`
	} `json:"SystemInfluence"`
	SystemShopItems []string `json:"SystemShopItems"`
	Tags            struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"Tags"`
	TravelRequirements []struct {
		ExclusionTags struct {
			Items            []interface{} `json:"items"`
			TagSetSourceFile string        `json:"tagSetSourceFile"`
		} `json:"ExclusionTags"`
		RequirementComparisons []interface{} `json:"RequirementComparisons"`
		RequirementTags        struct {
			Items            []string `json:"items"`
			TagSetSourceFile string   `json:"tagSetSourceFile"`
		} `json:"RequirementTags"`
		Scope string `json:"Scope"`
	} `json:"TravelRequirements"`
	UseMaxContractOverride     bool `json:"UseMaxContractOverride"`
	UseSystemRoninHiringChance bool `json:"UseSystemRoninHiringChance"`
}
