package defs

type AmmoDef struct {
	AllowedLocations string `json:"AllowedLocations"`
	AmmoID           string `json:"AmmoID"`
	BattleValue      int64  `json:"BattleValue"`
	BonusValueA      string `json:"BonusValueA"`
	BonusValueB      string `json:"BonusValueB"`
	Capacity         int64  `json:"Capacity"`
	ComponentSubType string `json:"ComponentSubType"`
	ComponentTags    struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"ComponentTags"`
	ComponentType     string `json:"ComponentType"`
	CriticalComponent bool   `json:"CriticalComponent"`
	Custom            struct {
		BonusDescriptions struct {
			Bonuses []string `json:"Bonuses"`
		} `json:"BonusDescriptions"`
		Category []struct {
			CategoryID string `json:"CategoryID"`
		} `json:"Category"`
		Flags struct {
			Flags []string `json:"flags"`
		} `json:"Flags"`
		InventorySorter struct {
			SortKey string `json:"SortKey"`
		} `json:"InventorySorter"`
		Lootable struct {
			ItemID string `json:"ItemID"`
		} `json:"Lootable"`
		Tag string `json:"Tag"`
	} `json:"Custom"`
	Description struct {
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
	} `json:"Description"`
	DisallowedLocations string        `json:"DisallowedLocations"`
	InventorySize       int64         `json:"InventorySize"`
	PrefabIdentifier    string        `json:"PrefabIdentifier"`
	Tonnage             float64       `json:"Tonnage"`
	StatusEffects       []interface{} `json:"statusEffects"`
}
