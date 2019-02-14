package defs

type ShopDef struct {
	ExclusionTags struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"ExclusionTags"`
	ID              string        `json:"ID"`
	Inventory       []interface{} `json:"Inventory"`
	RequirementTags struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"RequirementTags"`
	Specials []struct {
		Count            int64   `json:"Count"`
		DiscountModifier float64 `json:"DiscountModifier"`
		ID               string  `json:"ID"`
		Type             string  `json:"Type"`
		Weight           float64 `json:"Weight"`
	} `json:"Specials"`
}
