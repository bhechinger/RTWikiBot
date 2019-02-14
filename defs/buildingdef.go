package defs

type BuildingDef struct {
	Description struct {
		Cost        int64  `json:"Cost"`
		Details     string `json:"Details"`
		Icon        string `json:"Icon"`
		ID          string `json:"Id"`
		Name        string `json:"Name"`
		Purchasable bool   `json:"Purchasable"`
		Rarity      int64  `json:"Rarity"`
	} `json:"Description"`
	Destructible    bool  `json:"Destructible"`
	Opacity         int64 `json:"Opacity"`
	StructurePoints int64 `json:"StructurePoints"`
}
