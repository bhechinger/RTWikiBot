package defs

type AmmoDef struct {
	AllowedLocations    string        `json:"AllowedLocations"`
	AmmoID              string        `json:"AmmoID"`
	BattleValue         int64         `json:"BattleValue"`
	BonusValueA         string        `json:"BonusValueA"`
	BonusValueB         string        `json:"BonusValueB"`
	Capacity            int64         `json:"Capacity"`
	ComponentSubType    string        `json:"ComponentSubType"`
	ComponentTags       AmmoDef_sub1  `json:"ComponentTags"`
	ComponentType       string        `json:"ComponentType"`
	CriticalComponent   bool          `json:"CriticalComponent"`
	Custom              AmmoDef_sub7  `json:"Custom"`
	Description         AmmoDef_sub8  `json:"Description"`
	DisallowedLocations string        `json:"DisallowedLocations"`
	InventorySize       int64         `json:"InventorySize"`
	PrefabIdentifier    string        `json:"PrefabIdentifier"`
	Tonnage             float64       `json:"Tonnage"`
	StatusEffects       []interface{} `json:"statusEffects"`
}

type AmmoDef_sub7 struct {
	BonusDescriptions AmmoDef_sub2   `json:"BonusDescriptions"`
	Category          []AmmoDef_sub3 `json:"Category"`
	Flags             AmmoDef_sub4   `json:"Flags"`
	InventorySorter   AmmoDef_sub5   `json:"InventorySorter"`
	Lootable          AmmoDef_sub6   `json:"Lootable"`
	Tag               string         `json:"Tag"`
}

type AmmoDef_sub2 struct {
	Bonuses []string `json:"Bonuses"`
}

type AmmoDef_sub3 struct {
	CategoryID string `json:"CategoryID"`
}

type AmmoDef_sub8 struct {
	Cost         int64  `json:"Cost"`
	Details      string `json:"Details"`
	Icon         string `json:"Icon"`
	ID           string `json:"Id"`
	Manufacturer string `json:"Manufacturer"`
	Model        string `json:"Model"`
	Name         string `json:"Name"`
	Purchasable  bool   `json:"Purchasable"`
	Rarity       int64  `json:"Rarity"`
	UIName       string `json:"UIName"`
}

type AmmoDef_sub4 struct {
	Flags []string `json:"flags"`
}

type AmmoDef_sub6 struct {
	ItemID string `json:"ItemID"`
}

type AmmoDef_sub1 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}

type AmmoDef_sub5 struct {
	SortKey string `json:"SortKey"`
}
