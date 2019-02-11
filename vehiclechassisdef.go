package main

type VehicleChassisDef struct {
	BattleValue int64 `json:"BattleValue"`
	Description struct {
		Cost        int64  `json:"Cost"`
		Details     string `json:"Details"`
		Icon        string `json:"Icon"`
		ID          string `json:"Id"`
		Name        string `json:"Name"`
		Purchasable bool   `json:"Purchasable"`
		Rarity      int64  `json:"Rarity"`
	} `json:"Description"`
	HardpointDataDefID string `json:"HardpointDataDefID"`
	LOSSourcePositions []struct {
		X int64 `json:"x"`
		Y int64 `json:"y"`
		Z int64 `json:"z"`
	} `json:"LOSSourcePositions"`
	LOSTargetPositions []struct {
		X int64 `json:"x"`
		Y int64 `json:"y"`
		Z int64 `json:"z"`
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
		Tonnage           int64  `json:"Tonnage"`
	} `json:"Locations"`
	MovementCapDefID          string `json:"MovementCapDefID"`
	PathingCapDefID           string `json:"PathingCapDefID"`
	PrefabBase                string `json:"PrefabBase"`
	PrefabIdentifier          string `json:"PrefabIdentifier"`
	Radius                    int64  `json:"Radius"`
	SensorRangeMultiplier     int64  `json:"SensorRangeMultiplier"`
	Signature                 int64  `json:"Signature"`
	SpotterDistanceMultiplier int64  `json:"SpotterDistanceMultiplier"`
	Tonnage                   int64  `json:"Tonnage"`
	TopSpeed                  int64  `json:"TopSpeed"`
	TurnRadius                int64  `json:"TurnRadius"`
	VisibilityMultiplier      int64  `json:"VisibilityMultiplier"`
	MovementType              string `json:"movementType"`
	WeightClass               string `json:"weightClass"`
}
