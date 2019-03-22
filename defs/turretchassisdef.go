package defs

type TurretChassisDef struct {
	BattleValue               int64                   `json:"BattleValue"`
	Description               TurretChassisDef_sub1   `json:"Description"`
	FiringArcDegrees          int64                   `json:"FiringArcDegrees"`
	HardpointDataDefID        string                  `json:"HardpointDataDefID"`
	Hardpoints                []TurretChassisDef_sub2 `json:"Hardpoints"`
	InventorySlots            int64                   `json:"InventorySlots"`
	LOSSourcePositions        []TurretChassisDef_sub3 `json:"LOSSourcePositions"`
	LOSTargetPositions        []TurretChassisDef_sub3 `json:"LOSTargetPositions"`
	MaxArmor                  int64                   `json:"MaxArmor"`
	MaxInternalStructure      int64                   `json:"MaxInternalStructure"`
	MovementCapDefID          string                  `json:"MovementCapDefID"`
	PathingCapDefID           string                  `json:"PathingCapDefID"`
	PrefabBase                string                  `json:"PrefabBase"`
	PrefabIdentifier          string                  `json:"PrefabIdentifier"`
	Radius                    int64                   `json:"Radius"`
	SensorRangeMultiplier     float64                 `json:"SensorRangeMultiplier"`
	Signature                 float64                 `json:"Signature"`
	SpotterDistanceMultiplier float64                 `json:"SpotterDistanceMultiplier"`
	Tonnage                   int64                   `json:"Tonnage"`
	VisibilityMultiplier      float64                 `json:"VisibilityMultiplier"`
	WeightClass               string                  `json:"weightClass"`
}

type TurretChassisDef_sub1 struct {
	Cost        int64  `json:"Cost"`
	Details     string `json:"Details"`
	Icon        string `json:"Icon"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Purchasable bool   `json:"Purchasable"`
	Rarity      int64  `json:"Rarity"`
}

type TurretChassisDef_sub2 struct {
	Omni        bool   `json:"Omni"`
	WeaponMount string `json:"WeaponMount"`
}

type TurretChassisDef_sub3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
