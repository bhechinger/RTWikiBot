package defs

type Weapon struct {
	AOECapable                 bool          `json:"AOECapable"`
	AccuracyModifier           int64         `json:"AccuracyModifier"`
	AllowedLocations           string        `json:"AllowedLocations"`
	AmmoCategory               string        `json:"AmmoCategory"`
	AttackRecoil               int64         `json:"AttackRecoil"`
	BattleValue                int64         `json:"BattleValue"`
	BonusValueA                string        `json:"BonusValueA"`
	BonusValueB                string        `json:"BonusValueB"`
	Category                   string        `json:"Category"`
	ComponentSubType           string        `json:"ComponentSubType"`
	ComponentTags              Weapon_sub1   `json:"ComponentTags"`
	ComponentType              string        `json:"ComponentType"`
	CriticalChanceMultiplier   float64       `json:"CriticalChanceMultiplier"`
	CriticalComponent          bool          `json:"CriticalComponent"`
	Custom                     Weapon_sub4   `json:"Custom"`
	Damage                     int64         `json:"Damage"`
	DamageVariance             int64         `json:"DamageVariance"`
	Description                Weapon_sub5   `json:"Description"`
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

type Weapon_sub4 struct {
	BonusDescriptions Weapon_sub2 `json:"BonusDescriptions"`
	InventorySorter   Weapon_sub3 `json:"InventorySorter"`
}

type Weapon_sub2 struct {
	Bonuses []string `json:"Bonuses"`
}

type Weapon_sub5 struct {
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
}

type Weapon_sub1 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}

type Weapon_sub3 struct {
	SortKey string `json:"SortKey"`
}
