package defs

type LanceDef struct {
	Description LanceDef_sub1   `json:"Description"`
	Difficulty  int64           `json:"Difficulty"`
	LanceTags   LanceDef_sub2   `json:"LanceTags"`
	LanceUnits  []LanceDef_sub3 `json:"LanceUnits"`
}

type LanceDef_sub1 struct {
	Cost        int64  `json:"Cost"`
	Details     string `json:"Details"`
	Icon        string `json:"Icon"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Purchasable bool   `json:"Purchasable"`
	Rarity      int64  `json:"Rarity"`
}

type LanceDef_sub3 struct {
	ExcludedPilotTagSet LanceDef_sub2 `json:"excludedPilotTagSet"`
	ExcludedUnitTagSet  LanceDef_sub2 `json:"excludedUnitTagSet"`
	PilotID             string        `json:"pilotId"`
	PilotTagSet         LanceDef_sub2 `json:"pilotTagSet"`
	UnitID              string        `json:"unitId"`
	UnitTagSet          LanceDef_sub2 `json:"unitTagSet"`
	UnitType            string        `json:"unitType"`
}

type LanceDef_sub2 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}
