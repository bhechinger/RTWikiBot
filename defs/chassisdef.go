package defs

type ChassisDef struct {
	BattleValue               int64             `json:"BattleValue"`
	ChassisTags               ChassisDef_sub1   `json:"ChassisTags"`
	Custom                    ChassisDef_sub4   `json:"Custom"`
	DFADamage                 int64             `json:"DFADamage"`
	DFAInstability            int64             `json:"DFAInstability"`
	DFASelfDamage             int64             `json:"DFASelfDamage"`
	DFAToHitModifier          int64             `json:"DFAToHitModifier"`
	Description               ChassisDef_sub5   `json:"Description"`
	FixedEquipment            []ChassisDef_sub6 `json:"FixedEquipment"`
	HardpointDataDefID        string            `json:"HardpointDataDefID"`
	Heatsinks                 int64             `json:"Heatsinks"`
	InitialTonnage            float64           `json:"InitialTonnage"`
	LOSSourcePositions        []ChassisDef_sub7 `json:"LOSSourcePositions"`
	LOSTargetPositions        []ChassisDef_sub7 `json:"LOSTargetPositions"`
	Locations                 []ChassisDef_sub9 `json:"Locations"`
	MaxJumpjets               int64             `json:"MaxJumpjets"`
	MeleeDamage               int64             `json:"MeleeDamage"`
	MeleeInstability          int64             `json:"MeleeInstability"`
	MeleeToHitModifier        int64             `json:"MeleeToHitModifier"`
	MovementCapDefID          string            `json:"MovementCapDefID"`
	PathingCapDefID           string            `json:"PathingCapDefID"`
	PrefabBase                string            `json:"PrefabBase"`
	PrefabIdentifier          string            `json:"PrefabIdentifier"`
	PunchesWithLeftArm        bool              `json:"PunchesWithLeftArm"`
	Radius                    int64             `json:"Radius"`
	SensorRangeMultiplier     float64           `json:"SensorRangeMultiplier"`
	Signature                 float64           `json:"Signature"`
	SpotterDistanceMultiplier float64           `json:"SpotterDistanceMultiplier"`
	Stability                 int64             `json:"Stability"`
	StabilityDefenses         []int64           `json:"StabilityDefenses"`
	StockRole                 string            `json:"StockRole"`
	Tonnage                   int64             `json:"Tonnage"`
	TopSpeed                  int64             `json:"TopSpeed"`
	TurnRadius                int64             `json:"TurnRadius"`
	VariantName               string            `json:"VariantName"`
	VisibilityMultiplier      float64           `json:"VisibilityMultiplier"`
	YangsThoughts             string            `json:"YangsThoughts"`
	WeightClass               string            `json:"weightClass"`
}

type ChassisDef_sub4 struct {
	ArmActuatorSupport ChassisDef_sub2   `json:"ArmActuatorSupport"`
	ChassisDefaults    []ChassisDef_sub3 `json:"ChassisDefaults"`
}

type ChassisDef_sub3 struct {
	CategoryID string `json:"CategoryID"`
	DefID      string `json:"DefID"`
	Location   string `json:"Location"`
	Type       string `json:"Type"`
}

type ChassisDef_sub6 struct {
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
}

type ChassisDef_sub5 struct {
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

type ChassisDef_sub9 struct {
	Hardpoints        []ChassisDef_sub8 `json:"Hardpoints"`
	InternalStructure int64             `json:"InternalStructure"`
	InventorySlots    int64             `json:"InventorySlots"`
	Location          string            `json:"Location"`
	MaxArmor          int64             `json:"MaxArmor"`
	MaxRearArmor      int64             `json:"MaxRearArmor"`
	Tonnage           int64             `json:"Tonnage"`
}

type ChassisDef_sub1 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}

type ChassisDef_sub2 struct {
	LeftLimit  string `json:"LeftLimit"`
	RightLimit string `json:"RightLimit"`
}

type ChassisDef_sub8 struct {
	Omni        bool   `json:"Omni"`
	WeaponMount string `json:"WeaponMount"`
}

type ChassisDef_sub7 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
