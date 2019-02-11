package main

type HardPointDataDef struct {
	HardpointData []struct {
		Blanks         []interface{} `json:"blanks"`
		Location       string        `json:"location"`
		MountingPoints []interface{} `json:"mountingPoints"`
		Weapons        [][]string    `json:"weapons"`
	} `json:"HardpointData"`
	ID string `json:"ID"`
}
