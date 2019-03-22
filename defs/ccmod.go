package defs

type CCMod struct {
	Author        string        `json:"Author"`
	ConflictsWith []string      `json:"ConflictsWith"`
	Dll           string        `json:"DLL"`
	DLLEntryPoint string        `json:"DLLEntryPoint"`
	DependsOn     []interface{} `json:"DependsOn"`
	Enabled       bool          `json:"Enabled"`
	Name          string        `json:"Name"`
	Settings      CCMod_sub6    `json:"Settings"`
	Website       string        `json:"Website"`
}

type CCMod_sub2 struct {
	A int64 `json:"A"`
	B int64 `json:"B"`
	G int64 `json:"G"`
	R int64 `json:"R"`
}

type CCMod_sub3 struct {
	AddIfNotPresent bool   `json:"AddIfNotPresent"`
	AnyLocation     bool   `json:"AnyLocation"`
	CategoryID      string `json:"CategoryID"`
	DefID           string `json:"DefID"`
	Location        string `json:"Location"`
	Type            string `json:"Type"`
}

type CCMod_sub5 struct {
	AddIfNotPresent bool   `json:"AddIfNotPresent"`
	CategoryID      string `json:"CategoryID"`
	DefID           string `json:"DefID"`
	Location        string `json:"Location"`
	Tag             string `json:"Tag"`
	Type            string `json:"Type"`
}

type CCMod_sub1 struct {
	AddMaximumReached     string `json:"AddMaximumReached"`
	AutoReplace           string `json:"AutoReplace"`
	MaxEquiped            string `json:"MaxEquiped"`
	MaxEquipedPerLocation string `json:"MaxEquipedPerLocation"`
	Name                  string `json:"Name"`
	DisplayName           string `json:"displayName"`
}

type CCMod_sub6 struct {
	BaseRecoveryChance                         float64       `json:"BaseRecoveryChance"`
	Categories                                 []CCMod_sub1  `json:"Categories"`
	CenterTorsoDestroyedParts                  int64         `json:"CenterTorsoDestroyedParts"`
	ColorTags                                  []interface{} `json:"ColorTags"`
	DebugInfo                                  string        `json:"DebugInfo"`
	DefaultFlagBackgroundColor                 string        `json:"DefaultFlagBackgroundColor"`
	DefaultFlagOverlayCColor                   CCMod_sub2    `json:"DefaultFlagOverlayCColor"`
	Defaults                                   []CCMod_sub3  `json:"Defaults"`
	EjectRecoveryBonus                         float64       `json:"EjectRecoveryBonus"`
	FixDeletedComponents                       bool          `json:"FixDeletedComponents"`
	FixSaveGameMech                            bool          `json:"FixSaveGameMech"`
	HeadRecoveryPenaly                         int64         `json:"HeadRecoveryPenaly"`
	InvalidFlagBackgroundColor                 string        `json:"InvalidFlagBackgroundColor"`
	LimbRecoveryPenalty                        float64       `json:"LimbRecoveryPenalty"`
	LogLevel                                   string        `json:"LogLevel"`
	NoLootCTDestroyed                          bool          `json:"NoLootCTDestroyed"`
	OmniTechCostBySize                         bool          `json:"OmniTechCostBySize"`
	OmniTechFlag                               string        `json:"OmniTechFlag"`
	OmniTechInstallCost                        int64         `json:"OmniTechInstallCost"`
	OverrideMechPartCalculation                bool          `json:"OverrideMechPartCalculation"`
	OverrideRecoveryChance                     bool          `json:"OverrideRecoveryChance"`
	OverrideSalvageGeneration                  bool          `json:"OverrideSalvageGeneration"`
	PreinstalledOverlayCColor                  CCMod_sub2    `json:"PreinstalledOverlayCColor"`
	RunAutofixer                               bool          `json:"RunAutofixer"`
	SalvageArmWeight                           float64       `json:"SalvageArmWeight"`
	SalvageHeadWeight                          float64       `json:"SalvageHeadWeight"`
	SalvageLegWeight                           float64       `json:"SalvageLegWeight"`
	SalvageTorsoWeight                         int64         `json:"SalvageTorsoWeight"`
	SalvageUnrecoveredMech                     bool          `json:"SalvageUnrecoveredMech"`
	TagRestrictionDropValidateIncompatibleTags bool          `json:"TagRestrictionDropValidateIncompatibleTags"`
	TagRestrictionDropValidateRequiredTags     bool          `json:"TagRestrictionDropValidateRequiredTags"`
	TagRestrictions                            []CCMod_sub4  `json:"TagRestrictions"`
	TaggedDefaults                             []CCMod_sub5  `json:"TaggedDefaults"`
	TestEnableAllTags                          bool          `json:"TestEnableAllTags"`
	TorsoRecoveryPenalty                       float64       `json:"TorsoRecoveryPenalty"`
	UseDefaultFixedColor                       bool          `json:"UseDefaultFixedColor"`
}

type CCMod_sub4 struct {
	IncompatibleTags []string `json:"IncompatibleTags"`
	RequiredTags     []string `json:"RequiredTags"`
	Tag              string   `json:"Tag"`
}
