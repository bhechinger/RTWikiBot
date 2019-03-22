package defs

type Quirk struct {
	AbsoluteModifier    int64         `json:"AbsoluteModifier"`
	AllowedLocations    string        `json:"AllowedLocations"`
	BattleValue         int64         `json:"BattleValue"`
	BonusValueA         string        `json:"BonusValueA"`
	BonusValueB         string        `json:"BonusValueB"`
	ComponentSubType    string        `json:"ComponentSubType"`
	ComponentTags       Quirk_sub1    `json:"ComponentTags"`
	ComponentType       string        `json:"ComponentType"`
	CriticalComponent   bool          `json:"CriticalComponent"`
	Custom              Quirk_sub6    `json:"Custom"`
	Description         Quirk_sub7    `json:"Description"`
	DisallowedLocations string        `json:"DisallowedLocations"`
	InventorySize       int64         `json:"InventorySize"`
	PrefabIdentifier    string        `json:"PrefabIdentifier"`
	RelativeModifier    int64         `json:"RelativeModifier"`
	StatName            interface{}   `json:"StatName"`
	Tonnage             float64       `json:"Tonnage"`
	StatusEffects       []Quirk_sub12 `json:"statusEffects"`
}

type Quirk_sub10 struct {
	AdditionalRules                string `json:"additionalRules"`
	AppliesEachTick                bool   `json:"appliesEachTick"`
	EffectsPersistAfterDestruction bool   `json:"effectsPersistAfterDestruction"`
	ModType                        string `json:"modType"`
	ModValue                       string `json:"modValue"`
	Operation                      string `json:"operation"`
	StatName                       string `json:"statName"`
	TargetAmmoCategory             string `json:"targetAmmoCategory"`
	TargetCollection               string `json:"targetCollection"`
	TargetWeaponCategory           string `json:"targetWeaponCategory"`
	TargetWeaponSubType            string `json:"targetWeaponSubType"`
	TargetWeaponType               string `json:"targetWeaponType"`
}

type Quirk_sub6 struct {
	BonusDescriptions Quirk_sub2   `json:"BonusDescriptions"`
	Category          []Quirk_sub3 `json:"Category"`
	Flags             Quirk_sub4   `json:"Flags"`
	InventorySorter   Quirk_sub5   `json:"InventorySorter"`
}

type Quirk_sub2 struct {
	Bonuses []string `json:"Bonuses"`
}

type Quirk_sub3 struct {
	CategoryID string `json:"CategoryID"`
}

type Quirk_sub9 struct {
	ClearedWhenAttacked    bool  `json:"clearedWhenAttacked"`
	Duration               int64 `json:"duration"`
	StackLimit             int64 `json:"stackLimit"`
	TicksOnActivations     bool  `json:"ticksOnActivations"`
	TicksOnEndOfRound      bool  `json:"ticksOnEndOfRound"`
	TicksOnMovements       bool  `json:"ticksOnMovements"`
	UseActivationsOfTarget bool  `json:"useActivationsOfTarget"`
}

type Quirk_sub7 struct {
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

type Quirk_sub12 struct {
	Description                Quirk_sub8  `json:"Description"`
	ActorBurningData           interface{} `json:"actorBurningData"`
	DurationData               Quirk_sub9  `json:"durationData"`
	EffectType                 string      `json:"effectType"`
	FloatieData                interface{} `json:"floatieData"`
	InstantModData             interface{} `json:"instantModData"`
	Nature                     string      `json:"nature"`
	PoorlyMaintainedEffectData interface{} `json:"poorlyMaintainedEffectData"`
	StatisticData              Quirk_sub10 `json:"statisticData"`
	TagData                    interface{} `json:"tagData"`
	TargetingData              Quirk_sub11 `json:"targetingData"`
	VfxData                    interface{} `json:"vfxData"`
}

type Quirk_sub8 struct {
	Details string `json:"Details"`
	Icon    string `json:"Icon"`
	ID      string `json:"Id"`
	Name    string `json:"Name"`
}

type Quirk_sub11 struct {
	EffectTargetType        string `json:"effectTargetType"`
	EffectTriggerType       string `json:"effectTriggerType"`
	ExtendDurationOnTrigger int64  `json:"extendDurationOnTrigger"`
	ForcePathRebuild        bool   `json:"forcePathRebuild"`
	ForceVisRebuild         bool   `json:"forceVisRebuild"`
	Range                   int64  `json:"range"`
	ShowInStatusPanel       bool   `json:"showInStatusPanel"`
	ShowInTargetPreview     bool   `json:"showInTargetPreview"`
	SpecialRules            string `json:"specialRules"`
	TriggerLimit            int64  `json:"triggerLimit"`
}

type Quirk_sub4 struct {
	Flags []string `json:"flags"`
}

type Quirk_sub1 struct {
	Items            []string `json:"items"`
	TagSetSourceFile string   `json:"tagSetSourceFile"`
}

type Quirk_sub5 struct {
	SortKey string `json:"SortKey"`
}
