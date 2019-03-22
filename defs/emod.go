package defs

type EngineDef struct {
	AllowedLocations    string            `json:"AllowedLocations"`
	BattleValue         int64             `json:"BattleValue"`
	BonusValueA         string            `json:"BonusValueA"`
	BonusValueB         string            `json:"BonusValueB"`
	ComponentSubType    string            `json:"ComponentSubType"`
	ComponentTags       EngineDef_sub1    `json:"ComponentTags"`
	ComponentType       string            `json:"ComponentType"`
	CriticalComponent   bool              `json:"CriticalComponent"`
	Custom              EngineDef_sub14   `json:"Custom"`
	Description         EngineDef_sub15   `json:"Description"`
	DisallowedLocations string            `json:"DisallowedLocations"`
	DissipationCapacity int64             `json:"DissipationCapacity"`
	InventorySize       int64             `json:"InventorySize"`
	PrefabIdentifier    string            `json:"PrefabIdentifier"`
	Tonnage             float64           `json:"Tonnage"`
	StatusEffects       []EngineDef_sub20 `json:"statusEffects"`
}

type EngineDef_sub14 struct {
	BonusDescriptions EngineDef_sub2   `json:"BonusDescriptions"`
	Category          []EngineDef_sub3 `json:"Category"`
	Cooling           EngineDef_sub4   `json:"Cooling"`
	EngineCore        EngineDef_sub5   `json:"EngineCore"`
	EngineHeatBlock   EngineDef_sub6   `json:"EngineHeatBlock"`
	Flags             EngineDef_sub7   `json:"Flags"`
	InventorySorter   EngineDef_sub8   `json:"InventorySorter"`
	Linked            EngineDef_sub10  `json:"Linked"`
	Weights           EngineDef_sub11  `json:"Weights"`
	WorkOrderCosts    EngineDef_sub13  `json:"WorkOrderCosts"`
}

type EngineDef_sub2 struct {
	Bonuses []string `json:"Bonuses"`
}

type EngineDef_sub12 struct {
	CBillCost string `json:"CBillCost"`
	TechCost  string `json:"TechCost"`
}

type EngineDef_sub3 struct {
	CategoryID string `json:"CategoryID"`
}

type EngineDef_sub9 struct {
	ComponentDefID string `json:"ComponentDefId"`
	Location       string `json:"Location"`
}

type EngineDef_sub15 struct {
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

type EngineDef_sub13 struct {
	Default EngineDef_sub12 `json:"Default"`
	Install EngineDef_sub12 `json:"Install"`
}

type EngineDef_sub20 struct {
	Description   EngineDef_sub16 `json:"Description"`
	DurationData  EngineDef_sub17 `json:"durationData"`
	EffectType    string          `json:"effectType"`
	Nature        string          `json:"nature"`
	StatisticData EngineDef_sub18 `json:"statisticData"`
	TargetingData EngineDef_sub19 `json:"targetingData"`
}

type EngineDef_sub16 struct {
	Details string `json:"Details"`
	Icon    string `json:"Icon"`
	ID      string `json:"Id"`
	Name    string `json:"Name"`
}

type EngineDef_sub17 struct {
	Duration   int64 `json:"duration"`
	StackLimit int64 `json:"stackLimit"`
}

type EngineDef_sub19 struct {
	EffectTargetType  string `json:"effectTargetType"`
	EffectTriggerType string `json:"effectTriggerType"`
}

type EngineDef_sub11 struct {
	EngineFactor  float64 `json:"EngineFactor"`
	ReservedSlots int64   `json:"ReservedSlots"`
}

type EngineDef_sub7 struct {
	Flags []string `json:"flags"`
}

type EngineDef_sub6 struct {
	HeatSinkCount int64 `json:"HeatSinkCount"`
}

type EngineDef_sub4 struct {
	HeatSinkDefID string `json:"HeatSinkDefId"`
}

type EngineDef_sub1 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}

type EngineDef_sub10 struct {
	Links []EngineDef_sub9 `json:"Links"`
}

type EngineDef_sub18 struct {
	ModType   string `json:"modType"`
	ModValue  string `json:"modValue"`
	Operation string `json:"operation"`
	StatName  string `json:"statName"`
}

type EngineDef_sub5 struct {
	Rating string `json:"Rating"`
}

type EngineDef_sub8 struct {
	SortKey string `json:"SortKey"`
}
