package defs

type TurretDef struct {
	AssignedArmor            int64            `json:"AssignedArmor"`
	ChassisID                string           `json:"ChassisID"`
	CurrentArmor             int64            `json:"CurrentArmor"`
	CurrentInternalStructure int64            `json:"CurrentInternalStructure"`
	DamageLevel              string           `json:"DamageLevel"`
	Description              TurretDef_sub1   `json:"Description"`
	TurretTags               TurretDef_sub2   `json:"TurretTags"`
	Version                  int64            `json:"Version"`
	Inventory                []TurretDef_sub3 `json:"inventory"`
}

type TurretDef_sub3 struct {
	ComponentDefID   string `json:"ComponentDefID"`
	ComponentDefType string `json:"ComponentDefType"`
	DamageLevel      string `json:"DamageLevel"`
	HardpointSlot    int64  `json:"HardpointSlot"`
}

type TurretDef_sub1 struct {
	Cost        int64  `json:"Cost"`
	Details     string `json:"Details"`
	Icon        string `json:"Icon"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Purchasable bool   `json:"Purchasable"`
	Rarity      int64  `json:"Rarity"`
}

type TurretDef_sub2 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}
