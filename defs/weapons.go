package defs

type Weapon struct {
	AOECapable       bool   `json:"AOECapable"`
	AccuracyModifier int64  `json:"AccuracyModifier"`
	AllowedLocations string `json:"AllowedLocations"`
	AmmoCategory     string `json:"AmmoCategory"`
	AttackRecoil     int64  `json:"AttackRecoil"`
	BattleValue      int64  `json:"BattleValue"`
	BonusValueA      string `json:"BonusValueA"`
	BonusValueB      string `json:"BonusValueB"`
	Category         string `json:"Category"`
	ComponentSubType string `json:"ComponentSubType"`
	ComponentTags    struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"ComponentTags"`
	ComponentType            string  `json:"ComponentType"`
	CriticalChanceMultiplier float64 `json:"CriticalChanceMultiplier"`
	CriticalComponent        bool    `json:"CriticalComponent"`
	Custom                   struct {
		BonusDescriptions struct {
			Bonuses []string `json:"Bonuses"`
		} `json:"BonusDescriptions"`
		InventorySorter struct {
			SortKey string `json:"SortKey"`
		} `json:"InventorySorter"`
	} `json:"Custom"`
	Damage         int64 `json:"Damage"`
	DamageVariance int64 `json:"DamageVariance"`
	Description    struct {
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
	DisallowedLocations        string        `json:"DisallowedLocations"`
	EvasiveDamageMultiplier    float64       `json:"EvasiveDamageMultiplier"`
	EvasivePipsIgnored         int64         `json:"EvasivePipsIgnored"`
	HeatDamage                 int64         `json:"HeatDamage"`
	HeatGenerated              int64         `json:"HeatGenerated"`
	IndirectFireCapable        bool          `json:"IndirectFireCapable"`
	Instability                float64       `json:"Instability"`
	InventorySize              int64         `json:"InventorySize"`
	MaxRange                   int64         `json:"MaxRange"`
	MinRange                   int64         `json:"MinRange"`
	OverheatedDamageMultiplier float64       `json:"OverheatedDamageMultiplier"`
	PrefabIdentifier           string        `json:"PrefabIdentifier"`
	ProjectilesPerShot         int64         `json:"ProjectilesPerShot"`
	RangeSplit                 []int64       `json:"RangeSplit"`
	RefireModifier             int64         `json:"RefireModifier"`
	ShotsWhenFired             int64         `json:"ShotsWhenFired"`
	StartingAmmoCapacity       int64         `json:"StartingAmmoCapacity"`
	Tonnage                    float64       `json:"Tonnage"`
	Type                       string        `json:"Type"`
	WeaponEffectID             string        `json:"WeaponEffectID"`
	WeaponSubType              string        `json:"WeaponSubType"`
	StatusEffects              []interface{} `json:"statusEffects"`
}
