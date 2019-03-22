package defs

type VehicleChassisDef struct {
	BattleValue               int64                    `json:"BattleValue"`
	Description               VehicleChassisDef_sub1   `json:"Description"`
	HardpointDataDefID        string                   `json:"HardpointDataDefID"`
	LOSSourcePositions        []VehicleChassisDef_sub2 `json:"LOSSourcePositions"`
	LOSTargetPositions        []VehicleChassisDef_sub2 `json:"LOSTargetPositions"`
	Locations                 []VehicleChassisDef_sub4 `json:"Locations"`
	MovementCapDefID          string                   `json:"MovementCapDefID"`
	PathingCapDefID           string                   `json:"PathingCapDefID"`
	PrefabBase                string                   `json:"PrefabBase"`
	PrefabIdentifier          string                   `json:"PrefabIdentifier"`
	Radius                    int64                    `json:"Radius"`
	SensorRangeMultiplier     float64                  `json:"SensorRangeMultiplier"`
	Signature                 float64                  `json:"Signature"`
	SpotterDistanceMultiplier float64                  `json:"SpotterDistanceMultiplier"`
	Tonnage                   int64                    `json:"Tonnage"`
	TopSpeed                  int64                    `json:"TopSpeed"`
	TurnRadius                int64                    `json:"TurnRadius"`
	VisibilityMultiplier      float64                  `json:"VisibilityMultiplier"`
	MovementType              string                   `json:"movementType"`
	WeightClass               string                   `json:"weightClass"`
}

type VehicleChassisDef_sub1 struct {
	Cost        int64  `json:"Cost"`
	Details     string `json:"Details"`
	Icon        string `json:"Icon"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Purchasable bool   `json:"Purchasable"`
	Rarity      int64  `json:"Rarity"`
}

type VehicleChassisDef_sub4 struct {
	Hardpoints        []VehicleChassisDef_sub3 `json:"Hardpoints"`
	InternalStructure int64                    `json:"InternalStructure"`
	InventorySlots    int64                    `json:"InventorySlots"`
	Location          string                   `json:"Location"`
	MaxArmor          int64                    `json:"MaxArmor"`
	Tonnage           int64                    `json:"Tonnage"`
}

type VehicleChassisDef_sub3 struct {
	Omni        bool   `json:"Omni"`
	WeaponMount string `json:"WeaponMount"`
}

type VehicleChassisDef_sub2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
