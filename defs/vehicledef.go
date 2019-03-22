package defs

type VehicleDef struct {
	ChassisID   string            `json:"ChassisID"`
	Description VehicleDef_sub1   `json:"Description"`
	Locations   []VehicleDef_sub2 `json:"Locations"`
	VehicleTags VehicleDef_sub3   `json:"VehicleTags"`
	Version     int64             `json:"Version"`
	Inventory   []VehicleDef_sub4 `json:"inventory"`
}

type VehicleDef_sub2 struct {
	AssignedArmor            int64  `json:"AssignedArmor"`
	CurrentArmor             int64  `json:"CurrentArmor"`
	CurrentInternalStructure int64  `json:"CurrentInternalStructure"`
	DamageLevel              string `json:"DamageLevel"`
	Location                 string `json:"Location"`
}

type VehicleDef_sub4 struct {
	ComponentDefID   string `json:"ComponentDefID"`
	ComponentDefType string `json:"ComponentDefType"`
	DamageLevel      string `json:"DamageLevel"`
	HardpointSlot    int64  `json:"HardpointSlot"`
	MountedLocation  string `json:"MountedLocation"`
}

type VehicleDef_sub1 struct {
	Cost        int64  `json:"Cost"`
	Details     string `json:"Details"`
	Icon        string `json:"Icon"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Purchasable bool   `json:"Purchasable"`
	Rarity      int64  `json:"Rarity"`
}

type VehicleDef_sub3 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}
