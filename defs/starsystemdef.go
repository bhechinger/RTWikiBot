package defs

type StarSystemDef struct {
	BlackMarketShopItems       interface{}          `json:"BlackMarketShopItems"`
	ContractEmployers          []string             `json:"ContractEmployers"`
	ContractTargets            []string             `json:"ContractTargets"`
	CoreSystemID               string               `json:"CoreSystemID"`
	DefaultDifficulty          int64                `json:"DefaultDifficulty"`
	Depletable                 bool                 `json:"Depletable"`
	Description                StarSystemDef_sub1   `json:"Description"`
	DifficultyList             []interface{}        `json:"DifficultyList"`
	DifficultyModes            []interface{}        `json:"DifficultyModes"`
	FactionShopItems           []string             `json:"FactionShopItems"`
	FactionShopOwner           string               `json:"FactionShopOwner"`
	FuelingStation             bool                 `json:"FuelingStation"`
	JumpDistance               int64                `json:"JumpDistance"`
	MapExcludedTags            StarSystemDef_sub2   `json:"MapExcludedTags"`
	MapRequiredTags            StarSystemDef_sub2   `json:"MapRequiredTags"`
	MaxContractOverride        int64                `json:"MaxContractOverride"`
	Owner                      string               `json:"Owner"`
	Position                   StarSystemDef_sub3   `json:"Position"`
	RoninHiringChance          int64                `json:"RoninHiringChance"`
	ShopMaxInventory           int64                `json:"ShopMaxInventory"`
	ShopMaxSpecials            int64                `json:"ShopMaxSpecials"`
	ShopRefreshRate            int64                `json:"ShopRefreshRate"`
	StarPosition               interface{}          `json:"StarPosition"`
	StarType                   string               `json:"StarType"`
	StartingSystemModes        []string             `json:"StartingSystemModes"`
	SupportedBiomes            []string             `json:"SupportedBiomes"`
	SystemInfluence            []StarSystemDef_sub4 `json:"SystemInfluence"`
	SystemShopItems            []string             `json:"SystemShopItems"`
	Tags                       StarSystemDef_sub5   `json:"Tags"`
	TravelRequirements         []StarSystemDef_sub6 `json:"TravelRequirements"`
	UseMaxContractOverride     bool                 `json:"UseMaxContractOverride"`
	UseSystemRoninHiringChance bool                 `json:"UseSystemRoninHiringChance"`
}

type StarSystemDef_sub1 struct {
	Details string `json:"Details"`
	Icon    string `json:"Icon"`
	ID      string `json:"Id"`
	Name    string `json:"Name"`
}

type StarSystemDef_sub6 struct {
	ExclusionTags          StarSystemDef_sub2 `json:"ExclusionTags"`
	RequirementComparisons []interface{}      `json:"RequirementComparisons"`
	RequirementTags        StarSystemDef_sub5 `json:"RequirementTags"`
	Scope                  string             `json:"Scope"`
}

type StarSystemDef_sub4 struct {
	Faction   string `json:"Faction"`
	Influence int64  `json:"Influence"`
}

type StarSystemDef_sub2 struct {
	Items            []interface{} `json:"items"`
	TagSetSourceFile string        `json:"tagSetSourceFile"`
}

type StarSystemDef_sub5 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}

type StarSystemDef_sub3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
