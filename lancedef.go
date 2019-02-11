package main

type LanceDef struct {
	Description struct {
		Cost        int64  `json:"Cost"`
		Details     string `json:"Details"`
		Icon        string `json:"Icon"`
		ID          string `json:"Id"`
		Name        string `json:"Name"`
		Purchasable bool   `json:"Purchasable"`
		Rarity      int64  `json:"Rarity"`
	} `json:"Description"`
	Difficulty int64 `json:"Difficulty"`
	LanceTags  struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"LanceTags"`
	LanceUnits []struct {
		ExcludedPilotTagSet struct {
			Items            []string `json:"items"`
			TagSetSourceFile string   `json:"tagSetSourceFile"`
		} `json:"excludedPilotTagSet"`
		ExcludedUnitTagSet struct {
			Items            []string `json:"items"`
			TagSetSourceFile string   `json:"tagSetSourceFile"`
		} `json:"excludedUnitTagSet"`
		PilotID     string `json:"pilotId"`
		PilotTagSet struct {
			Items            []string `json:"items"`
			TagSetSourceFile string   `json:"tagSetSourceFile"`
		} `json:"pilotTagSet"`
		UnitID     string `json:"unitId"`
		UnitTagSet struct {
			Items            []string `json:"items"`
			TagSetSourceFile string   `json:"tagSetSourceFile"`
		} `json:"unitTagSet"`
		UnitType string `json:"unitType"`
	} `json:"LanceUnits"`
}
