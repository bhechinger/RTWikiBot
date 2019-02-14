package defs

type HeraldryDef struct {
	Description struct {
		Cost         int64  `json:"Cost"`
		Details      string `json:"Details"`
		Icon         string `json:"Icon"`
		ID           string `json:"Id"`
		Manufacturer string `json:"Manufacturer"`
		Model        string `json:"Model"`
		Name         string `json:"Name"`
		Purchasable  bool   `json:"Purchasable"`
		Rarity       int64  `json:"Rarity"`
		UIName       string `json:"UIName"`
	} `json:"Description"`
	PrimaryMechColorID   string `json:"primaryMechColorID"`
	SecondaryMechColorID string `json:"secondaryMechColorID"`
	TertiaryMechColorID  string `json:"tertiaryMechColorID"`
	TextureLogoID        string `json:"textureLogoID"`
}
