package defs

type ShopDef struct {
	ExclusionTags   ShopDef_sub1   `json:"ExclusionTags"`
	ID              string         `json:"ID"`
	Inventory       []interface{}  `json:"Inventory"`
	RequirementTags ShopDef_sub1   `json:"RequirementTags"`
	Specials        []ShopDef_sub2 `json:"Specials"`
}

type ShopDef_sub2 struct {
	Count            int64   `json:"Count"`
	DiscountModifier float64 `json:"DiscountModifier"`
	ID               string  `json:"ID"`
	Type             string  `json:"Type"`
	Weight           float64 `json:"Weight"`
}

type ShopDef_sub1 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}
