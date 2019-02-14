package defs

type EngineDef struct {
	AllowedLocations string `json:"AllowedLocations"`
	BattleValue      int64  `json:"BattleValue"`
	BonusValueA      string `json:"BonusValueA"`
	BonusValueB      string `json:"BonusValueB"`
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
		Cooling struct {
			HeatSinkDefID string `json:"HeatSinkDefId"`
		} `json:"Cooling"`
		EngineCore struct {
			Rating string `json:"Rating"`
		} `json:"EngineCore"`
		EngineHeatBlock struct {
			HeatSinkCount int64 `json:"HeatSinkCount"`
		} `json:"EngineHeatBlock"`
		Flags struct {
			Flags []string `json:"flags"`
		} `json:"Flags"`
		InventorySorter struct {
			SortKey string `json:"SortKey"`
		} `json:"InventorySorter"`
		Linked struct {
			Links []struct {
				ComponentDefID string `json:"ComponentDefId"`
				Location       string `json:"Location"`
			} `json:"Links"`
		} `json:"Linked"`
		Weights struct {
			EngineFactor  float64 `json:"EngineFactor"`
			ReservedSlots int64   `json:"ReservedSlots"`
		} `json:"Weights"`
		WorkOrderCosts struct {
			Default struct {
				CBillCost string `json:"CBillCost"`
				TechCost  string `json:"TechCost"`
			} `json:"Default"`
			Install struct {
				CBillCost string `json:"CBillCost"`
				TechCost  string `json:"TechCost"`
			} `json:"Install"`
		} `json:"WorkOrderCosts"`
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
	DisallowedLocations string  `json:"DisallowedLocations"`
	DissipationCapacity int64   `json:"DissipationCapacity"`
	InventorySize       int64   `json:"InventorySize"`
	PrefabIdentifier    string  `json:"PrefabIdentifier"`
	Tonnage             float64 `json:"Tonnage"`
	StatusEffects       []struct {
		Description struct {
			Details string `json:"Details"`
			Icon    string `json:"Icon"`
			ID      string `json:"Id"`
			Name    string `json:"Name"`
		} `json:"Description"`
		DurationData struct {
			Duration   int64 `json:"duration"`
			StackLimit int64 `json:"stackLimit"`
		} `json:"durationData"`
		EffectType    string `json:"effectType"`
		Nature        string `json:"nature"`
		StatisticData struct {
			ModType   string `json:"modType"`
			ModValue  string `json:"modValue"`
			Operation string `json:"operation"`
			StatName  string `json:"statName"`
		} `json:"statisticData"`
		TargetingData struct {
			EffectTargetType  string `json:"effectTargetType"`
			EffectTriggerType string `json:"effectTriggerType"`
		} `json:"targetingData"`
	} `json:"statusEffects"`
}
