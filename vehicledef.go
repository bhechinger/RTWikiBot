package main

type VehicleDef struct {
	ChassisID   string `json:"ChassisID"`
	Description struct {
		Cost        int64  `json:"Cost"`
		Details     string `json:"Details"`
		Icon        string `json:"Icon"`
		ID          string `json:"Id"`
		Name        string `json:"Name"`
		Purchasable bool   `json:"Purchasable"`
		Rarity      int64  `json:"Rarity"`
	} `json:"Description"`
	Locations []struct {
		AssignedArmor            int64  `json:"AssignedArmor"`
		CurrentArmor             int64  `json:"CurrentArmor"`
		CurrentInternalStructure int64  `json:"CurrentInternalStructure"`
		DamageLevel              string `json:"DamageLevel"`
		Location                 string `json:"Location"`
	} `json:"Locations"`
	VehicleTags struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"VehicleTags"`
	Version   int64 `json:"Version"`
	Inventory []struct {
		ComponentDefID   string `json:"ComponentDefID"`
		ComponentDefType string `json:"ComponentDefType"`
		DamageLevel      string `json:"DamageLevel"`
		HardpointSlot    int64  `json:"HardpointSlot"`
		MountedLocation  string `json:"MountedLocation"`
	} `json:"inventory"`
}
