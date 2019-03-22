package defs

type HardPointDataDef struct {
	HardpointData []HardPointDataDef_sub1 `json:"HardpointData"`
	ID            string                  `json:"ID"`
}

type HardPointDataDef_sub1 struct {
	Blanks         []interface{} `json:"blanks"`
	Location       string        `json:"location"`
	MountingPoints []interface{} `json:"mountingPoints"`
	Weapons        [][]string    `json:"weapons"`
}
