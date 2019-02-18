package defs

type Settings struct {
	AllowMixingHeatSinkTypes                     bool          `json:"AllowMixingHeatSinkTypes"`
	ArmorStructureRatioEnforcement               bool          `json:"ArmorStructureRatioEnforcement"`
	AutoFixArmActuatorCategorizer                interface{}   `json:"AutoFixArmActuatorCategorizer"`
	AutoFixArmActuatorSlotChange                 interface{}   `json:"AutoFixArmActuatorSlotChange"`
	AutoFixArmorCategorizer                      interface{}   `json:"AutoFixArmorCategorizer"`
	AutoFixChassisDefInitialToTotalTonnageFactor float64       `json:"AutoFixChassisDefInitialToTotalTonnageFactor"`
	AutoFixChassisDefInitialTonnage              bool          `json:"AutoFixChassisDefInitialTonnage"`
	AutoFixChassisDefMaxJumpjets                 bool          `json:"AutoFixChassisDefMaxJumpjets"`
	AutoFixChassisDefMaxJumpjetsCount            int64         `json:"AutoFixChassisDefMaxJumpjetsCount"`
	AutoFixChassisDefMaxJumpjetsRating           int64         `json:"AutoFixChassisDefMaxJumpjetsRating"`
	AutoFixChassisDefSkip                        []interface{} `json:"AutoFixChassisDefSkip"`
	AutoFixChassisDefSlots                       interface{}   `json:"AutoFixChassisDefSlots"`
	AutoFixChassisDefSlotsChanges                interface{}   `json:"AutoFixChassisDefSlotsChanges"`
	AutoFixCockpitCategorizer                    interface{}   `json:"AutoFixCockpitCategorizer"`
	AutoFixCockpitSlotChange                     interface{}   `json:"AutoFixCockpitSlotChange"`
	AutoFixCockpitTonnageChange                  interface{}   `json:"AutoFixCockpitTonnageChange"`
	AutoFixGyroCategorizer                       interface{}   `json:"AutoFixGyroCategorizer"`
	AutoFixGyroSlotChange                        interface{}   `json:"AutoFixGyroSlotChange"`
	AutoFixLegUpgradesCategorizer                interface{}   `json:"AutoFixLegUpgradesCategorizer"`
	AutoFixLegUpgradesSlotChange                 interface{}   `json:"AutoFixLegUpgradesSlotChange"`
	AutoFixMechDefArmActuator                    bool          `json:"AutoFixMechDefArmActuator"`
	AutoFixMechDefArmorAdder                     interface{}   `json:"AutoFixMechDefArmorAdder"`
	AutoFixMechDefCockpitAdder                   interface{}   `json:"AutoFixMechDefCockpitAdder"`
	AutoFixMechDefCoolingDef                     string        `json:"AutoFixMechDefCoolingDef"`
	AutoFixMechDefEngine                         bool          `json:"AutoFixMechDefEngine"`
	AutoFixMechDefEngineTypeDef                  string        `json:"AutoFixMechDefEngineTypeDef"`
	AutoFixMechDefGyroAdder                      interface{}   `json:"AutoFixMechDefGyroAdder"`
	AutoFixMechDefHeatBlockDef                   string        `json:"AutoFixMechDefHeatBlockDef"`
	AutoFixMechDefSkip                           []interface{} `json:"AutoFixMechDefSkip"`
	AutoFixMechDefStructureAdder                 interface{}   `json:"AutoFixMechDefStructureAdder"`
	AutoFixStructureCategorizer                  interface{}   `json:"AutoFixStructureCategorizer"`
	AutoFixUpgradeDefSkip                        []string      `json:"AutoFixUpgradeDefSkip"`
	AutoFixWeaponDefSlotsChanges                 interface{}   `json:"AutoFixWeaponDefSlotsChanges"`
	BonusDescriptions                            []struct {
		Bonus string `json:"Bonus"`
		Full  string `json:"Full"`
		Long  string `json:"Long"`
		Short string `json:"Short"`
	} `json:"BonusDescriptions"`
	BonusDescriptionsTooltipTitle string `json:"BonusDescriptionsTooltipTitle"`
	CBTMovement                   struct {
		JJRoundUp          bool  `json:"JJRoundUp"`
		TTSprintMultiplier int64 `json:"TTSprintMultiplier"`
		TTWalkMultiplier   int64 `json:"TTWalkMultiplier"`
		UseGameWalkValues  bool  `json:"UseGameWalkValues"`
	} `json:"CBTMovement"`
	Categories []struct {
		AddAlreadyEquiped         string `json:"AddAlreadyEquiped"`
		AddAlreadyEquipedLocation string `json:"AddAlreadyEquipedLocation"`
		AddMaximumLocationReached string `json:"AddMaximumLocationReached"`
		AddMaximumReached         string `json:"AddMaximumReached"`
		AddMixed                  string `json:"AddMixed"`
		AllowMixTags              bool   `json:"AllowMixTags"`
		AutoReplace               bool   `json:"AutoReplace"`
		DefaultCustoms            struct {
			Case struct {
				MaximumDamage int64 `json:"MaximumDamage"`
			} `json:"CASE"`
			CriticalHitStates struct {
				DeathMethod string `json:"DeathMethod"`
				HitEffects  []struct {
					State        int64 `json:"State"`
					StatusEffect struct {
						Description struct {
							Details string `json:"Details"`
							Icon    string `json:"Icon"`
							ID      string `json:"Id"`
							Name    string `json:"Name"`
						} `json:"Description"`
						DurationData  struct{} `json:"durationData"`
						EffectType    string   `json:"effectType"`
						Nature        string   `json:"nature"`
						StatisticData struct {
							ModType   string `json:"modType"`
							ModValue  string `json:"modValue"`
							Operation string `json:"operation"`
							StatName  string `json:"statName"`
						} `json:"statisticData"`
						TargetingData struct {
							EffectTargetType     string `json:"effectTargetType"`
							EffectTargetsCreator bool   `json:"effectTargetsCreator"`
							EffectTriggerType    string `json:"effectTriggerType"`
						} `json:"targetingData"`
					} `json:"StatusEffect"`
				} `json:"HitEffects"`
				MaxStates int64 `json:"MaxStates"`
			} `json:"CriticalHitStates"`
			Flags struct {
				Flags []string `json:"Flags"`
			} `json:"Flags"`
			InventorySorter struct {
				SortKey string `json:"SortKey"`
			} `json:"InventorySorter"`
			Replace struct {
				ComponentDefID string `json:"ComponentDefId"`
				Location       string `json:"Location"`
			} `json:"Replace"`
			Sorter struct {
				Order int64 `json:"Order"`
			} `json:"Sorter"`
		} `json:"DefaultCustoms"`
		DisplayName             string   `json:"DisplayName"`
		Forbidden               []string `json:"Forbidden"`
		MaxEquiped              int64    `json:"MaxEquiped"`
		MaxEquipedPerLocation   int64    `json:"MaxEquipedPerLocation"`
		MinEquiped              int64    `json:"MinEquiped"`
		Name                    string   `json:"Name"`
		ReplaceAnyLocation      bool     `json:"ReplaceAnyLocation"`
		Required                bool     `json:"Required"`
		Unique                  bool     `json:"Unique"`
		UniqueForLocation       bool     `json:"UniqueForLocation"`
		ValidateForbidden       string   `json:"ValidateForbidden"`
		ValidateMaximum         string   `json:"ValidateMaximum"`
		ValidateMaximumLocation string   `json:"ValidateMaximumLocation"`
		ValidateMinimum         string   `json:"ValidateMinimum"`
		ValidateMixed           string   `json:"ValidateMixed"`
		ValidateRequred         string   `json:"ValidateRequred"`
		ValidateUnique          string   `json:"ValidateUnique"`
		ValidateUniqueLocation  string   `json:"ValidateUniqueLocation"`
	} `json:"Categories"`
	DefaultEngineHeatSinkID                    string `json:"DefaultEngineHeatSinkId"`
	DynamicSlotsValidateDropEnabled            bool   `json:"DynamicSlotsValidateDropEnabled"`
	EnforceRulesForAdditionalInternalHeatSinks bool   `json:"EnforceRulesForAdditionalInternalHeatSinks"`
	EngineCriticalHitStates                    struct {
		DeathMethod string `json:"DeathMethod"`
		HitEffects  []struct {
			State        int64 `json:"State"`
			StatusEffect struct {
				Description struct {
					Details string `json:"Details"`
					Icon    string `json:"Icon"`
					ID      string `json:"Id"`
					Name    string `json:"Name"`
				} `json:"Description"`
				DurationData struct {
					ClearedWhenAttacked    bool  `json:"clearedWhenAttacked"`
					Duration               int64 `json:"duration"`
					StackLimit             int64 `json:"stackLimit"`
					TicksOnActivations     bool  `json:"ticksOnActivations"`
					TicksOnEndOfRound      bool  `json:"ticksOnEndOfRound"`
					TicksOnMovements       bool  `json:"ticksOnMovements"`
					UseActivationsOfTarget bool  `json:"useActivationsOfTarget"`
				} `json:"durationData"`
				EffectType    string `json:"effectType"`
				Nature        string `json:"nature"`
				StatisticData struct {
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
				TargetingData struct {
					EffectTargetType        string `json:"effectTargetType"`
					EffectTargetsCreator    bool   `json:"effectTargetsCreator"`
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
			} `json:"StatusEffect"`
		} `json:"HitEffects"`
		MaxStates int64 `json:"MaxStates"`
	} `json:"EngineCriticalHitStates"`
	EngineMissingFallbackHeatSinkCapacity int64   `json:"EngineMissingFallbackHeatSinkCapacity"`
	FractionalAccountingPrecision         float64 `json:"FractionalAccountingPrecision"`
	HardpointFix                          struct {
		AllowDefaultLoadoutWeapons             bool     `json:"AllowDefaultLoadoutWeapons"`
		AllowLRMInLargerSlotsForAll            bool     `json:"AllowLRMInLargerSlotsForAll"`
		AllowLRMInSmallerSlotsForAll           bool     `json:"AllowLRMInSmallerSlotsForAll"`
		AllowLRMInSmallerSlotsForMechs         []string `json:"AllowLRMInSmallerSlotsForMechs"`
		AutoFixChassisDefWeaponHardpointCounts bool     `json:"AutoFixChassisDefWeaponHardpointCounts"`
		EnforceHardpointLimits                 bool     `json:"EnforceHardpointLimits"`
	} `json:"HardpointFix"`
	HeatDamageInjuryEnabled bool `json:"HeatDamageInjuryEnabled"`
	LogLevels               []struct {
		Level string `json:"Level"`
		Name  string `json:"Name"`
	} `json:"LogLevels"`
	MinimumHeatSinksOnMech              int64 `json:"MinimumHeatSinksOnMech"`
	PerformanceEnableAvailabilityChecks bool  `json:"PerformanceEnableAvailabilityChecks"`
	SaveMechDefOnMechLabConfirm         bool  `json:"SaveMechDefOnMechLabConfirm"`
	ShutdownInjuryEnabled               bool  `json:"ShutdownInjuryEnabled"`
	TagRestrictions                     []struct {
		IncompatibleTags []string `json:"IncompatibleTags"`
		Tag              string   `json:"Tag"`
	} `json:"TagRestrictions"`
	TagRestrictionsUseDescriptionIds   bool `json:"TagRestrictionsUseDescriptionIds"`
	TagRestrictionsValidateDropEnabled bool `json:"TagRestrictionsValidateDropEnabled"`
}
