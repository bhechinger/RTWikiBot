package defs

type ChassisDef struct {
	BattleValue int64 `json:"BattleValue"`
	ChassisTags struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"ChassisTags"`
	Custom struct {
		ArmActuatorSupport struct {
			LeftLimit  string `json:"LeftLimit"`
			RightLimit string `json:"RightLimit"`
		} `json:"ArmActuatorSupport"`
	} `json:"Custom"`
	DFADamage        int64 `json:"DFADamage"`
	DFAInstability   int64 `json:"DFAInstability"`
	DFASelfDamage    int64 `json:"DFASelfDamage"`
	DFAToHitModifier int64 `json:"DFAToHitModifier"`
	Description      struct {
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
	FixedEquipment []struct {
		ComponentDefID   string      `json:"ComponentDefID"`
		ComponentDefType string      `json:"ComponentDefType"`
		DamageLevel      string      `json:"DamageLevel"`
		GUID             interface{} `json:"GUID"`
		HardpointSlot    int64       `json:"HardpointSlot"`
		IsFixed          bool        `json:"IsFixed"`
		MountedLocation  string      `json:"MountedLocation"`
		SimGameUID       string      `json:"SimGameUID"`
		HasPrefabName    bool        `json:"hasPrefabName"`
		PrefabName       string      `json:"prefabName"`
	} `json:"FixedEquipment"`
	HardpointDataDefID string  `json:"HardpointDataDefID"`
	Heatsinks          int64   `json:"Heatsinks"`
	InitialTonnage     float64 `json:"InitialTonnage"`
	LOSSourcePositions []struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"LOSSourcePositions"`
	LOSTargetPositions []struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"LOSTargetPositions"`
	Locations []struct {
		Hardpoints []struct {
			Omni        bool   `json:"Omni"`
			WeaponMount string `json:"WeaponMount"`
		} `json:"Hardpoints"`
		InternalStructure int64  `json:"InternalStructure"`
		InventorySlots    int64  `json:"InventorySlots"`
		Location          string `json:"Location"`
		MaxArmor          int64  `json:"MaxArmor"`
		MaxRearArmor      int64  `json:"MaxRearArmor"`
		Tonnage           int64  `json:"Tonnage"`
	} `json:"Locations"`
	MaxJumpjets               int64   `json:"MaxJumpjets"`
	MeleeDamage               int64   `json:"MeleeDamage"`
	MeleeInstability          int64   `json:"MeleeInstability"`
	MeleeToHitModifier        int64   `json:"MeleeToHitModifier"`
	MovementCapDefID          string  `json:"MovementCapDefID"`
	PathingCapDefID           string  `json:"PathingCapDefID"`
	PrefabBase                string  `json:"PrefabBase"`
	PrefabIdentifier          string  `json:"PrefabIdentifier"`
	PunchesWithLeftArm        bool    `json:"PunchesWithLeftArm"`
	Radius                    int64   `json:"Radius"`
	SensorRangeMultiplier     float64 `json:"SensorRangeMultiplier"`
	Signature                 float64 `json:"Signature"`
	SpotterDistanceMultiplier float64 `json:"SpotterDistanceMultiplier"`
	Stability                 int64   `json:"Stability"`
	StabilityDefenses         []int64 `json:"StabilityDefenses"`
	StockRole                 string  `json:"StockRole"`
	Tonnage                   int64   `json:"Tonnage"`
	TopSpeed                  int64   `json:"TopSpeed"`
	TurnRadius                int64   `json:"TurnRadius"`
	VariantName               string  `json:"VariantName"`
	VisibilityMultiplier      float64 `json:"VisibilityMultiplier"`
	YangsThoughts             string  `json:"YangsThoughts"`
	WeightClass               string  `json:"weightClass"`
}
