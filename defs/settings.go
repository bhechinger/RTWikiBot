package defs

type Settings struct {
	AllowMixingHeatSinkTypes                     bool             `json:"AllowMixingHeatSinkTypes"`
	ArmorStructureRatioEnforcement               bool             `json:"ArmorStructureRatioEnforcement"`
	AutoFixArmActuatorCategorizer                interface{}      `json:"AutoFixArmActuatorCategorizer"`
	AutoFixArmActuatorSlotChange                 interface{}      `json:"AutoFixArmActuatorSlotChange"`
	AutoFixArmorCategorizer                      interface{}      `json:"AutoFixArmorCategorizer"`
	AutoFixChassisDefInitialToTotalTonnageFactor float64          `json:"AutoFixChassisDefInitialToTotalTonnageFactor"`
	AutoFixChassisDefInitialTonnage              bool             `json:"AutoFixChassisDefInitialTonnage"`
	AutoFixChassisDefMaxJumpjets                 bool             `json:"AutoFixChassisDefMaxJumpjets"`
	AutoFixChassisDefMaxJumpjetsCount            int64            `json:"AutoFixChassisDefMaxJumpjetsCount"`
	AutoFixChassisDefMaxJumpjetsRating           int64            `json:"AutoFixChassisDefMaxJumpjetsRating"`
	AutoFixChassisDefSkip                        []interface{}    `json:"AutoFixChassisDefSkip"`
	AutoFixChassisDefSlots                       interface{}      `json:"AutoFixChassisDefSlots"`
	AutoFixChassisDefSlotsChanges                interface{}      `json:"AutoFixChassisDefSlotsChanges"`
	AutoFixCockpitCategorizer                    interface{}      `json:"AutoFixCockpitCategorizer"`
	AutoFixCockpitSlotChange                     interface{}      `json:"AutoFixCockpitSlotChange"`
	AutoFixCockpitTonnageChange                  interface{}      `json:"AutoFixCockpitTonnageChange"`
	AutoFixGyroCategorizer                       interface{}      `json:"AutoFixGyroCategorizer"`
	AutoFixGyroSlotChange                        interface{}      `json:"AutoFixGyroSlotChange"`
	AutoFixLegUpgradesCategorizer                interface{}      `json:"AutoFixLegUpgradesCategorizer"`
	AutoFixLegUpgradesSlotChange                 interface{}      `json:"AutoFixLegUpgradesSlotChange"`
	AutoFixMechDefArmActuator                    bool             `json:"AutoFixMechDefArmActuator"`
	AutoFixMechDefArmorAdder                     interface{}      `json:"AutoFixMechDefArmorAdder"`
	AutoFixMechDefCockpitAdder                   interface{}      `json:"AutoFixMechDefCockpitAdder"`
	AutoFixMechDefCoolingDef                     string           `json:"AutoFixMechDefCoolingDef"`
	AutoFixMechDefEngine                         bool             `json:"AutoFixMechDefEngine"`
	AutoFixMechDefEngineTypeDef                  string           `json:"AutoFixMechDefEngineTypeDef"`
	AutoFixMechDefGyroAdder                      interface{}      `json:"AutoFixMechDefGyroAdder"`
	AutoFixMechDefHeatBlockDef                   string           `json:"AutoFixMechDefHeatBlockDef"`
	AutoFixMechDefSkip                           []interface{}    `json:"AutoFixMechDefSkip"`
	AutoFixMechDefStructureAdder                 interface{}      `json:"AutoFixMechDefStructureAdder"`
	AutoFixStructureCategorizer                  interface{}      `json:"AutoFixStructureCategorizer"`
	AutoFixUpgradeDefSkip                        []string         `json:"AutoFixUpgradeDefSkip"`
	AutoFixWeaponDefSlotsChanges                 interface{}      `json:"AutoFixWeaponDefSlotsChanges"`
	BonusDescriptions                            []Settings_sub1  `json:"BonusDescriptions"`
	BonusDescriptionsTooltipTitle                string           `json:"BonusDescriptionsTooltipTitle"`
	CBTMovement                                  Settings_sub2    `json:"CBTMovement"`
	Categories                                   []Settings_sub16 `json:"Categories"`
	DefaultEngineHeatSinkID                      string           `json:"DefaultEngineHeatSinkId"`
	DynamicSlotsValidateDropEnabled              bool             `json:"DynamicSlotsValidateDropEnabled"`
	EnforceRulesForAdditionalInternalHeatSinks   bool             `json:"EnforceRulesForAdditionalInternalHeatSinks"`
	EngineCriticalHitStates                      Settings_sub22   `json:"EngineCriticalHitStates"`
	EngineMissingFallbackHeatSinkCapacity        int64            `json:"EngineMissingFallbackHeatSinkCapacity"`
	FractionalAccountingPrecision                float64          `json:"FractionalAccountingPrecision"`
	HardpointFix                                 Settings_sub23   `json:"HardpointFix"`
	HeatDamageInjuryEnabled                      bool             `json:"HeatDamageInjuryEnabled"`
	LogLevels                                    []Settings_sub24 `json:"LogLevels"`
	MinimumHeatSinksOnMech                       int64            `json:"MinimumHeatSinksOnMech"`
	PerformanceEnableAvailabilityChecks          bool             `json:"PerformanceEnableAvailabilityChecks"`
	SaveMechDefOnMechLabConfirm                  bool             `json:"SaveMechDefOnMechLabConfirm"`
	ShutdownInjuryEnabled                        bool             `json:"ShutdownInjuryEnabled"`
	TagRestrictions                              []Settings_sub25 `json:"TagRestrictions"`
	TagRestrictionsUseDescriptionIds             bool             `json:"TagRestrictionsUseDescriptionIds"`
	TagRestrictionsValidateDropEnabled           bool             `json:"TagRestrictionsValidateDropEnabled"`
}

type Settings_sub16 struct {
	AddAlreadyEquiped         string         `json:"AddAlreadyEquiped"`
	AddAlreadyEquipedLocation string         `json:"AddAlreadyEquipedLocation"`
	AddMaximumLocationReached string         `json:"AddMaximumLocationReached"`
	AddMaximumReached         string         `json:"AddMaximumReached"`
	AddMixed                  string         `json:"AddMixed"`
	AllowMixTags              bool           `json:"AllowMixTags"`
	AutoReplace               bool           `json:"AutoReplace"`
	DefaultCustoms            Settings_sub15 `json:"DefaultCustoms"`
	DisplayName               string         `json:"DisplayName"`
	Forbidden                 []string       `json:"Forbidden"`
	MaxEquiped                int64          `json:"MaxEquiped"`
	MaxEquipedPerLocation     int64          `json:"MaxEquipedPerLocation"`
	MinEquiped                int64          `json:"MinEquiped"`
	Name                      string         `json:"Name"`
	ReplaceAnyLocation        bool           `json:"ReplaceAnyLocation"`
	Required                  bool           `json:"Required"`
	Unique                    bool           `json:"Unique"`
	UniqueForLocation         bool           `json:"UniqueForLocation"`
	ValidateForbidden         string         `json:"ValidateForbidden"`
	ValidateMaximum           string         `json:"ValidateMaximum"`
	ValidateMaximumLocation   string         `json:"ValidateMaximumLocation"`
	ValidateMinimum           string         `json:"ValidateMinimum"`
	ValidateMixed             string         `json:"ValidateMixed"`
	ValidateRequred           string         `json:"ValidateRequred"`
	ValidateUnique            string         `json:"ValidateUnique"`
	ValidateUniqueLocation    string         `json:"ValidateUniqueLocation"`
}

type Settings_sub18 struct {
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

type Settings_sub23 struct {
	AllowDefaultLoadoutWeapons             bool     `json:"AllowDefaultLoadoutWeapons"`
	AllowLRMInLargerSlotsForAll            bool     `json:"AllowLRMInLargerSlotsForAll"`
	AllowLRMInSmallerSlotsForAll           bool     `json:"AllowLRMInSmallerSlotsForAll"`
	AllowLRMInSmallerSlotsForMechs         []string `json:"AllowLRMInSmallerSlotsForMechs"`
	AutoFixChassisDefWeaponHardpointCounts bool     `json:"AutoFixChassisDefWeaponHardpointCounts"`
	EnforceHardpointLimits                 bool     `json:"EnforceHardpointLimits"`
}

type Settings_sub1 struct {
	Bonus string `json:"Bonus"`
	Full  string `json:"Full"`
	Long  string `json:"Long"`
	Short string `json:"Short"`
}

type Settings_sub15 struct {
	Case              Settings_sub3  `json:"CASE"`
	CriticalHitStates Settings_sub10 `json:"CriticalHitStates"`
	Flags             Settings_sub11 `json:"Flags"`
	InventorySorter   Settings_sub12 `json:"InventorySorter"`
	Replace           Settings_sub13 `json:"Replace"`
	Sorter            Settings_sub14 `json:"Sorter"`
}

type Settings_sub17 struct {
	ClearedWhenAttacked    bool  `json:"clearedWhenAttacked"`
	Duration               int64 `json:"duration"`
	StackLimit             int64 `json:"stackLimit"`
	TicksOnActivations     bool  `json:"ticksOnActivations"`
	TicksOnEndOfRound      bool  `json:"ticksOnEndOfRound"`
	TicksOnMovements       bool  `json:"ticksOnMovements"`
	UseActivationsOfTarget bool  `json:"useActivationsOfTarget"`
}

type Settings_sub13 struct {
	ComponentDefID string `json:"ComponentDefId"`
	Location       string `json:"Location"`
}

type Settings_sub22 struct {
	DeathMethod string           `json:"DeathMethod"`
	HitEffects  []Settings_sub21 `json:"HitEffects"`
	MaxStates   int64            `json:"MaxStates"`
}

type Settings_sub10 struct {
	DeathMethod string          `json:"DeathMethod"`
	HitEffects  []Settings_sub9 `json:"HitEffects"`
	MaxStates   int64           `json:"MaxStates"`
}

type Settings_sub20 struct {
	Description   Settings_sub4  `json:"Description"`
	DurationData  Settings_sub17 `json:"durationData"`
	EffectType    string         `json:"effectType"`
	Nature        string         `json:"nature"`
	StatisticData Settings_sub18 `json:"statisticData"`
	TargetingData Settings_sub19 `json:"targetingData"`
}

type Settings_sub8 struct {
	Description   Settings_sub4 `json:"Description"`
	DurationData  Settings_sub5 `json:"durationData"`
	EffectType    string        `json:"effectType"`
	Nature        string        `json:"nature"`
	StatisticData Settings_sub6 `json:"statisticData"`
	TargetingData Settings_sub7 `json:"targetingData"`
}

type Settings_sub4 struct {
	Details string `json:"Details"`
	Icon    string `json:"Icon"`
	ID      string `json:"Id"`
	Name    string `json:"Name"`
}

type Settings_sub19 struct {
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
}

type Settings_sub7 struct {
	EffectTargetType     string `json:"effectTargetType"`
	EffectTargetsCreator bool   `json:"effectTargetsCreator"`
	EffectTriggerType    string `json:"effectTriggerType"`
}

type Settings_sub11 struct {
	Flags []string `json:"Flags"`
}

type Settings_sub25 struct {
	IncompatibleTags []string `json:"IncompatibleTags"`
	Tag              string   `json:"Tag"`
}

type Settings_sub2 struct {
	JJRoundUp          bool  `json:"JJRoundUp"`
	TTSprintMultiplier int64 `json:"TTSprintMultiplier"`
	TTWalkMultiplier   int64 `json:"TTWalkMultiplier"`
	UseGameWalkValues  bool  `json:"UseGameWalkValues"`
}

type Settings_sub24 struct {
	Level string `json:"Level"`
	Name  string `json:"Name"`
}

type Settings_sub3 struct {
	MaximumDamage int64 `json:"MaximumDamage"`
}

type Settings_sub6 struct {
	ModType   string `json:"modType"`
	ModValue  string `json:"modValue"`
	Operation string `json:"operation"`
	StatName  string `json:"statName"`
}

type Settings_sub14 struct {
	Order int64 `json:"Order"`
}

type Settings_sub12 struct {
	SortKey string `json:"SortKey"`
}

type Settings_sub21 struct {
	State        int64          `json:"State"`
	StatusEffect Settings_sub20 `json:"StatusEffect"`
}

type Settings_sub9 struct {
	State        int64         `json:"State"`
	StatusEffect Settings_sub8 `json:"StatusEffect"`
}

type Settings_sub5 struct{}
