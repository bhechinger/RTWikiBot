package main

type HeatSinkDef struct {
	AllowedLocations string `json:"AllowedLocations"`
	BattleValue      int64  `json:"BattleValue"`
	BonusValueA      string `json:"BonusValueA"`
	BonusValueB      string `json:"BonusValueB"`
	ComponentSubType string `json:"ComponentSubType"`
	ComponentTags    struct {
		Items            []string `json:"items"`
		TagSetSourceFile string   `json:"tagSetSourceFile"`
	} `json:"ComponentTags"`
	ComponentType     string `json:"ComponentType"`
	CriticalComponent bool   `json:"CriticalComponent"`
	Custom            struct {
		BonusDescriptions struct {
			Bonuses []string `json:"Bonuses"`
		} `json:"BonusDescriptions"`
		Category []struct {
			CategoryID string `json:"CategoryID"`
		} `json:"Category"`
		ComponentExplosion struct {
			ExplosionDamage int64 `json:"ExplosionDamage"`
			HeatDamage      int64 `json:"HeatDamage"`
			StabilityDamage int64 `json:"StabilityDamage"`
		} `json:"ComponentExplosion"`
		InventorySorter struct {
			SortKey string `json:"SortKey"`
		} `json:"InventorySorter"`
	} `json:"Custom"`
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
	DisallowedLocations string  `json:"DisallowedLocations"`
	DissipationCapacity int64   `json:"DissipationCapacity"`
	InventorySize       int64   `json:"InventorySize"`
	PrefabIdentifier    string  `json:"PrefabIdentifier"`
	Tonnage             float64 `json:"Tonnage"`
	StatusEffects       []struct {
		Description struct {
			Details string `json:"Details"`
			Icon    string `json:"Icon"`
			ID      string `json:"Id"`
			Name    string `json:"Name"`
		} `json:"Description"`
		ActorBurningData interface{} `json:"actorBurningData"`
		DurationData     struct {
			ClearedWhenAttacked    bool  `json:"clearedWhenAttacked"`
			Duration               int64 `json:"duration"`
			StackLimit             int64 `json:"stackLimit"`
			TicksOnActivations     bool  `json:"ticksOnActivations"`
			TicksOnEndOfRound      bool  `json:"ticksOnEndOfRound"`
			TicksOnMovements       bool  `json:"ticksOnMovements"`
			UseActivationsOfTarget bool  `json:"useActivationsOfTarget"`
		} `json:"durationData"`
		EffectType                 string      `json:"effectType"`
		FloatieData                interface{} `json:"floatieData"`
		InstantModData             interface{} `json:"instantModData"`
		Nature                     string      `json:"nature"`
		PoorlyMaintainedEffectData interface{} `json:"poorlyMaintainedEffectData"`
		StatisticData              struct {
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
		} `json:"statisticData"`
		TagData       interface{} `json:"tagData"`
		TargetingData struct {
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
		} `json:"targetingData"`
		VfxData interface{} `json:"vfxData"`
	} `json:"statusEffects"`
}
