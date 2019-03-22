package defs

type MechDef struct {
	ChassisID           string         `json:"ChassisID"`
	Description         MechDef_sub1   `json:"Description"`
	HeraldryID          interface{}    `json:"HeraldryID"`
	Locations           []MechDef_sub2 `json:"Locations"`
	MechTags            MechDef_sub3   `json:"MechTags"`
	Version             int64          `json:"Version"`
	Inventory           []MechDef_sub4 `json:"inventory"`
	SimGameMechPartCost float64        `json:"simGameMechPartCost"`
}

type MechDef_sub2 struct {
	AssignedArmor            int64  `json:"AssignedArmor"`
	AssignedRearArmor        int64  `json:"AssignedRearArmor"`
	CurrentArmor             int64  `json:"CurrentArmor"`
	CurrentInternalStructure int64  `json:"CurrentInternalStructure"`
	CurrentRearArmor         int64  `json:"CurrentRearArmor"`
	DamageLevel              string `json:"DamageLevel"`
	Location                 string `json:"Location"`
}

type MechDef_sub4 struct {
	ComponentDefID   string      `json:"ComponentDefID"`
	ComponentDefType string      `json:"ComponentDefType"`
	DamageLevel      string      `json:"DamageLevel"`
	GUID             interface{} `json:"GUID"`
	HardpointSlot    int64       `json:"HardpointSlot"`
	MountedLocation  string      `json:"MountedLocation"`
	SimGameUID       string      `json:"SimGameUID"`
	HasPrefabName    bool        `json:"hasPrefabName"`
	PrefabName       string      `json:"prefabName"`
}

type MechDef_sub1 struct {
	Cost         int64       `json:"Cost"`
	Details      string      `json:"Details"`
	Icon         string      `json:"Icon"`
	ID           string      `json:"Id"`
	Manufacturer interface{} `json:"Manufacturer"`
	Model        interface{} `json:"Model"`
	Name         string      `json:"Name"`
	Purchasable  bool        `json:"Purchasable"`
	Rarity       int64       `json:"Rarity"`
	UIName       string      `json:"UIName"`
}

type MechDef_sub3 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}
