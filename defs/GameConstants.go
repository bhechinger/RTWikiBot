package defs

type GameConstants struct {
	CameraConstants        GameConstants_sub1  `json:"CameraConstants"`
	CombatUIConstants      GameConstants_sub5  `json:"CombatUIConstants"`
	CombatValueMultipliers GameConstants_sub6  `json:"CombatValueMultipliers"`
	DynamicAIRoleConstants GameConstants_sub7  `json:"DynamicAIRoleConstants"`
	FuryConstants          GameConstants_sub12 `json:"FuryConstants"`
	Heat                   GameConstants_sub13 `json:"Heat"`
	HitTables              GameConstants_sub15 `json:"HitTables"`
	MoraleConstants        GameConstants_sub19 `json:"MoraleConstants"`
	MoveConstants          GameConstants_sub20 `json:"MoveConstants"`
	Phase                  GameConstants_sub21 `json:"Phase"`
	PilotingConstants      GameConstants_sub22 `json:"PilotingConstants"`
	ResolutionConstants    GameConstants_sub26 `json:"ResolutionConstants"`
	Skills                 GameConstants_sub27 `json:"Skills"`
	ToHit                  GameConstants_sub28 `json:"ToHit"`
	VFXNames               GameConstants_sub30 `json:"VFXNames"`
	Visibility             GameConstants_sub38 `json:"Visibility"`
}

type GameConstants_sub26 struct {
	AICritChanceBaseMultiplier          float64             `json:"AICritChanceBaseMultiplier"`
	AllowTotalNegativeModifier          bool                `json:"AllowTotalNegativeModifier"`
	BuildingDestructionForceMultiplier  float64             `json:"BuildingDestructionForceMultiplier"`
	ChargeDistanceThreshold             float64             `json:"ChargeDistanceThreshold"`
	ComponentDestructionForceMultiplier float64             `json:"ComponentDestructionForceMultiplier"`
	CoolantVentCooldownEffect           GameConstants_sub25 `json:"CoolantVentCooldownEffect"`
	CoolantVentEffect                   GameConstants_sub25 `json:"CoolantVentEffect"`
	DFAMissedPilotingModifier           float64             `json:"DFAMissedPilotingModifier"`
	DFASelfDamageMultiplier             float64             `json:"DFASelfDamageMultiplier"`
	DFASelfDamagePilotingDivisor        float64             `json:"DFASelfDamagePilotingDivisor"`
	DefaultUnsteadyThreshold            float64             `json:"DefaultUnsteadyThreshold"`
	EntrenchedMultiplier                float64             `json:"EntrenchedMultiplier"`
	EvasiveDodgeChance                  float64             `json:"EvasiveDodgeChance"`
	FlimsyDestructionForceMultiplier    float64             `json:"FlimsyDestructionForceMultiplier"`
	GlancingBlowDamageMultiplier        float64             `json:"GlancingBlowDamageMultiplier"`
	HeavyDamageThreshold                int64               `json:"HeavyDamageThreshold"`
	HeavyHitReactDamageThreshold        int64               `json:"HeavyHitReactDamageThreshold"`
	IneffectiveBlowDamageMultiplier     float64             `json:"IneffectiveBlowDamageMultiplier"`
	LowArmorThreshold                   float64             `json:"LowArmorThreshold"`
	MeleeCountersEvasive                bool                `json:"MeleeCountersEvasive"`
	MeleeCountersGuarded                bool                `json:"MeleeCountersGuarded"`
	MeleeDamageDFAMultiplier            float64             `json:"MeleeDamageDFAMultiplier"`
	MeleeDamageMultiplierTurret         float64             `json:"MeleeDamageMultiplierTurret"`
	MeleeDamageMultiplierVehicle        float64             `json:"MeleeDamageMultiplierVehicle"`
	MeleeDamageTonnageMultiplier        float64             `json:"MeleeDamageTonnageMultiplier"`
	MeleeInstabilityChargeMultiplier    float64             `json:"MeleeInstabilityChargeMultiplier"`
	MeleeInstabilityDFAMultiplier       float64             `json:"MeleeInstabilityDFAMultiplier"`
	MeleeInstabilityNormalMultiplier    float64             `json:"MeleeInstabilityNormalMultiplier"`
	MeleeKickWeighting                  float64             `json:"MeleeKickWeighting"`
	MeleePunchWeighting                 float64             `json:"MeleePunchWeighting"`
	MeleeTackleWeighting                float64             `json:"MeleeTackleWeighting"`
	MinCritChance                       float64             `json:"MinCritChance"`
	MissOffsetHorizontalMax             float64             `json:"MissOffsetHorizontalMax"`
	MissOffsetHorizontalMin             float64             `json:"MissOffsetHorizontalMin"`
	MissOffsetVerticalMax               float64             `json:"MissOffsetVerticalMax"`
	MissOffsetVerticalMin               float64             `json:"MissOffsetVerticalMin"`
	NormalBlowDamageMultiplier          float64             `json:"NormalBlowDamageMultiplier"`
	SearchForValidCritSlot              bool                `json:"SearchForValidCritSlot"`
	SensorLockCountersEvasive           bool                `json:"SensorLockCountersEvasive"`
	SensorLockCountersGuarded           bool                `json:"SensorLockCountersGuarded"`
	SolidBlowDamageMultiplier           float64             `json:"SolidBlowDamageMultiplier"`
	SolidDamageCountersEvasive          bool                `json:"SolidDamageCountersEvasive"`
	SolidDamageCountersGuarded          bool                `json:"SolidDamageCountersGuarded"`
	StabilityDumpAbilityDefer           int64               `json:"StabilityDumpAbilityDefer"`
	StabilityDumpBracing                int64               `json:"StabilityDumpBracing"`
	StabilityDumpDFA                    float64             `json:"StabilityDumpDFA"`
	StabilityDumpJumping                float64             `json:"StabilityDumpJumping"`
	StabilityDumpMoving                 int64               `json:"StabilityDumpMoving"`
	StabilityDumpStationary             int64               `json:"StabilityDumpStationary"`
	StabilityDumpStoodUp                int64               `json:"StabilityDumpStoodUp"`
	StabilityLevels                     int64               `json:"StabilityLevels"`
	UnsteadyCountersEvasive             bool                `json:"UnsteadyCountersEvasive"`
	UnsteadyCountersGuarded             bool                `json:"UnsteadyCountersGuarded"`
	VehicleEvasiveResultMultiplier      int64               `json:"VehicleEvasiveResultMultiplier"`
	VehiclesGetEvasive                  bool                `json:"VehiclesGetEvasive"`
	VehiclesGetGuarded                  bool                `json:"VehiclesGetGuarded"`
}

type GameConstants_sub5 struct {
	AlliedTurnText                     string               `json:"AlliedTurnText"`
	AttackSpeedUpFactor                float64              `json:"AttackSpeedUpFactor"`
	AutoSelectDuringInterleaved        bool                 `json:"AutoSelectDuringInterleaved"`
	AutoSelectDuringNonInterleaved     bool                 `json:"AutoSelectDuringNonInterleaved"`
	BackwardDescription                GameConstants_sub2   `json:"BackwardDescription"`
	BadFaithRetreatDesc                GameConstants_sub2   `json:"BadFaithRetreatDesc"`
	CameraStageOneTutorialDesc         GameConstants_sub2   `json:"CameraStageOneTutorialDesc"`
	CameraStageTwoTutorialDesc         GameConstants_sub2   `json:"CameraStageTwoTutorialDesc"`
	DangerousLocationDesc              GameConstants_sub2   `json:"DangerousLocationDesc"`
	DoneDescription                    GameConstants_sub2   `json:"DoneDescription"`
	DoneDescriptionBulwark             GameConstants_sub2   `json:"DoneDescriptionBulwark"`
	DoneDescriptionNoBrace             GameConstants_sub2   `json:"DoneDescriptionNoBrace"`
	DrophipLocationDesc                GameConstants_sub2   `json:"DrophipLocationDesc"`
	EjectDescription                   GameConstants_sub2   `json:"EjectDescription"`
	EnemyTurnText                      string               `json:"EnemyTurnText"`
	EvasiveTutorialDesc                GameConstants_sub2   `json:"EvasiveTutorialDesc"`
	FireDescription                    GameConstants_sub2   `json:"FireDescription"`
	FireMultiDescription               GameConstants_sub2   `json:"FireMultiDescription"`
	FiringTutorialDesc                 GameConstants_sub2   `json:"FiringTutorialDesc"`
	FuryBarDescription                 GameConstants_sub2   `json:"FuryBarDescription"`
	GoodFaithRetreatDesc               GameConstants_sub2   `json:"GoodFaithRetreatDesc"`
	InjuryStringAmmoExplosion          string               `json:"InjuryStringAmmoExplosion"`
	InjuryStringArtillery              string               `json:"InjuryStringArtillery"`
	InjuryStringDFAEnemy               string               `json:"InjuryStringDFAEnemy"`
	InjuryStringDFAEnemyGeneric        string               `json:"InjuryStringDFAEnemyGeneric"`
	InjuryStringDFASelf                string               `json:"InjuryStringDFASelf"`
	InjuryStringDFASelfGeneric         string               `json:"InjuryStringDFASelfGeneric"`
	InjuryStringDropShip               string               `json:"InjuryStringDropShip"`
	InjuryStringHeadHit                string               `json:"InjuryStringHeadHit"`
	InjuryStringHeadHitGeneric         string               `json:"InjuryStringHeadHitGeneric"`
	InjuryStringHeadHitMelee           string               `json:"InjuryStringHeadHitMelee"`
	InjuryStringKilledByEnemy          string               `json:"InjuryStringKilledByEnemy"`
	InjuryStringKilledByEnemyGeneric   string               `json:"InjuryStringKilledByEnemyGeneric"`
	InjuryStringKnockdownGeneric       string               `json:"InjuryStringKnockdownGeneric"`
	InjuryStringKnockedDownBy          string               `json:"InjuryStringKnockedDownBy"`
	InjuryStringMelee                  string               `json:"InjuryStringMelee"`
	InjuryStringMeleeGeneric           string               `json:"InjuryStringMeleeGeneric"`
	InjuryStringOverheatedEnemy        string               `json:"InjuryStringOverheatedEnemy"`
	InjuryStringOverheatedEnemyUnknown string               `json:"InjuryStringOverheatedEnemyUnknown"`
	InjuryStringOverheatedSelf         string               `json:"InjuryStringOverheatedSelf"`
	InjuryStringOverheatedSelfGeneric  string               `json:"InjuryStringOverheatedSelfGeneric"`
	InjuryStringSideTorso              string               `json:"InjuryStringSideTorso"`
	InjuryStringSideTorsoGeneric       string               `json:"InjuryStringSideTorsoGeneric"`
	InjuryStringUnknown                string               `json:"InjuryStringUnknown"`
	InjuryStringWeapon                 string               `json:"InjuryStringWeapon"`
	InterleaveIntroTutorialDesc        GameConstants_sub2   `json:"InterleaveIntroTutorialDesc"`
	JumpDescription                    GameConstants_sub2   `json:"JumpDescription"`
	MeleeTutorialDesc                  GameConstants_sub2   `json:"MeleeTutorialDesc"`
	MoraleAttackDescription            GameConstants_sub2   `json:"MoraleAttackDescription"`
	MoraleAttackTutorialDesc           GameConstants_sub2   `json:"MoraleAttackTutorialDesc"`
	MoraleBarDescription               GameConstants_sub2   `json:"MoraleBarDescription"`
	MoraleCostAttackDescription        GameConstants_sub2   `json:"MoraleCostAttackDescription"`
	MoraleCostAttackDescriptionFury    GameConstants_sub2   `json:"MoraleCostAttackDescriptionFury"`
	MoraleCostAttackDescriptionHigh    GameConstants_sub2   `json:"MoraleCostAttackDescriptionHigh"`
	MoraleCostAttackDescriptionLow     GameConstants_sub2   `json:"MoraleCostAttackDescriptionLow"`
	MoraleCostDefendDescription        GameConstants_sub2   `json:"MoraleCostDefendDescription"`
	MoraleCostDefendDescriptionHigh    GameConstants_sub2   `json:"MoraleCostDefendDescriptionHigh"`
	MoraleCostDefendDescriptionLow     GameConstants_sub2   `json:"MoraleCostDefendDescriptionLow"`
	MoraleDefendDescription            GameConstants_sub2   `json:"MoraleDefendDescription"`
	MoraleDefendTutorialDesc           GameConstants_sub2   `json:"MoraleDefendTutorialDesc"`
	MoraleGeneralStageTwoTutorialDesc  GameConstants_sub2   `json:"MoraleGeneralStageTwoTutorialDesc"`
	MoraleGeneralTutorialDesc          GameConstants_sub2   `json:"MoraleGeneralTutorialDesc"`
	MoveDescription                    GameConstants_sub2   `json:"MoveDescription"`
	MoveSpeedUpFactor                  float64              `json:"MoveSpeedUpFactor"`
	MovementStageTwoTutorialDesc       GameConstants_sub2   `json:"MovementStageTwoTutorialDesc"`
	MovementTutorialDesc               GameConstants_sub2   `json:"MovementTutorialDesc"`
	NeutralTurnText                    string               `json:"NeutralTurnText"`
	NonInterleaveIntroTutorialDesc     GameConstants_sub2   `json:"NonInterleaveIntroTutorialDesc"`
	PathingDeadZoneMaxDist             float64              `json:"PathingDeadZoneMaxDist"`
	PathingDeadZoneMinDist             float64              `json:"PathingDeadZoneMinDist"`
	PoorlyMaintained25Desc             GameConstants_sub3   `json:"PoorlyMaintained25Desc"`
	PoorlyMaintained50Desc             GameConstants_sub3   `json:"PoorlyMaintained50Desc"`
	PoorlyMaintained75Desc             GameConstants_sub3   `json:"PoorlyMaintained75Desc"`
	RangeStrings                       []GameConstants_sub4 `json:"RangeStrings"`
	ScreenShakeDFAAbsoluteMod          float64              `json:"ScreenShakeDFAAbsoluteMod"`
	ScreenShakeDFARelativeMod          float64              `json:"ScreenShakeDFARelativeMod"`
	ScreenShakeFootfallAbsoluteMod     float64              `json:"ScreenShakeFootfallAbsoluteMod"`
	ScreenShakeFootfallRelativeMod     float64              `json:"ScreenShakeFootfallRelativeMod"`
	ScreenShakeGroundImpactAbsoluteMod float64              `json:"ScreenShakeGroundImpactAbsoluteMod"`
	ScreenShakeGroundImpactRelativeMod float64              `json:"ScreenShakeGroundImpactRelativeMod"`
	ScreenShakeMax                     float64              `json:"ScreenShakeMax"`
	ScreenShakeMeleeDamageAbsoluteMod  float64              `json:"ScreenShakeMeleeDamageAbsoluteMod"`
	ScreenShakeMeleeDamageRelativeMod  float64              `json:"ScreenShakeMeleeDamageRelativeMod"`
	ScreenShakeRangedDamageAbsoluteMod float64              `json:"ScreenShakeRangedDamageAbsoluteMod"`
	ScreenShakeRangedDamageRelativeMod float64              `json:"ScreenShakeRangedDamageRelativeMod"`
	ShowDoneWithAllButton              bool                 `json:"ShowDoneWithAllButton"`
	SprintDescription                  GameConstants_sub2   `json:"SprintDescription"`
	StabilityDangerRatio               float64              `json:"StabilityDangerRatio"`
	StabilityWarningRatio              float64              `json:"StabilityWarningRatio"`
	YourTurnText                       string               `json:"YourTurnText"`
	CollapseWeaponTypesInTickMarks     bool                 `json:"collapseWeaponTypesInTickMarks"`
	FloatieLifetime                    float64              `json:"floatieLifetime"`
	FloatieSizeLarge                   int64                `json:"floatieSizeLarge"`
	FloatieSizeMedium                  int64                `json:"floatieSizeMedium"`
	FloatieSizeSmall                   int64                `json:"floatieSizeSmall"`
	ShowFLoatiesForMisses              bool                 `json:"showFLoatiesForMisses"`
	ShowFloatiesForDodges              bool                 `json:"showFloatiesForDodges"`
	ShowStateRemovals                  bool                 `json:"showStateRemovals"`
	SortWeaponsByAttackOrder           bool                 `json:"sortWeaponsByAttackOrder"`
}

type GameConstants_sub38 struct {
	AllowRearArcSpotting               bool                `json:"AllowRearArcSpotting"`
	BaseSensorDistance                 float64             `json:"BaseSensorDistance"`
	BaseSpotterDistance                float64             `json:"BaseSpotterDistance"`
	ExtendedSensorDistanceMultiplier   float64             `json:"ExtendedSensorDistanceMultiplier"`
	FiredWeaponsEffect                 GameConstants_sub31 `json:"FiredWeaponsEffect"`
	MaxShadowingDistance               float64             `json:"MaxShadowingDistance"`
	MinRatioFromActors                 float64             `json:"MinRatioFromActors"`
	NumSensorLockImpairedEffects       int64               `json:"NumSensorLockImpairedEffects"`
	NumSensorLockSteps                 int64               `json:"NumSensorLockSteps"`
	NumShadowingSteps                  int64               `json:"NumShadowingSteps"`
	ProneSpottingDistanceMultiplier    float64             `json:"ProneSpottingDistanceMultiplier"`
	RatioFullVis                       float64             `json:"RatioFullVis"`
	RatioObstructedVis                 float64             `json:"RatioObstructedVis"`
	SensorHysteresisAdditive           float64             `json:"SensorHysteresisAdditive"`
	SensorLockSingleStepEffect         GameConstants_sub34 `json:"SensorLockSingleStepEffect"`
	SensorsImpairedEffect              GameConstants_sub37 `json:"SensorsImpairedEffect"`
	ShowExtendedBlipsBeforeContact     bool                `json:"ShowExtendedBlipsBeforeContact"`
	ShutDownSignatureModifier          float64             `json:"ShutDownSignatureModifier"`
	ShutDownVisibilityModifier         float64             `json:"ShutDownVisibilityModifier"`
	ShutdownSpottingDistanceMultiplier float64             `json:"ShutdownSpottingDistanceMultiplier"`
	SpotterTacticsMultiplier           float64             `json:"SpotterTacticsMultiplier"`
	UseAsymmetricalSensors             bool                `json:"UseAsymmetricalSensors"`
}

type GameConstants_sub22 struct {
	ArmBlackRelativeInstability             float64  `json:"ArmBlackRelativeInstability"`
	CenterTorsoYellowRelativeInstability    float64  `json:"CenterTorsoYellowRelativeInstability"`
	DefaultMaxInjuries                      int64    `json:"DefaultMaxInjuries"`
	FixedInjuryPenalty                      int64    `json:"FixedInjuryPenalty"`
	GutsBaseFloor                           float64  `json:"GutsBaseFloor"`
	GutsDivisor                             int64    `json:"GutsDivisor"`
	GutsModifierDivisor                     int64    `json:"GutsModifierDivisor"`
	GutsPenaltyFromHeadDamageMajor          int64    `json:"GutsPenaltyFromHeadDamageMajor"`
	GutsPenaltyFromHeadDamageMinor          int64    `json:"GutsPenaltyFromHeadDamageMinor"`
	InjuryFromAmmoExplosion                 bool     `json:"InjuryFromAmmoExplosion"`
	InjuryFromSideTorsoDestruction          bool     `json:"InjuryFromSideTorsoDestruction"`
	InjuryPenaltyOnlyPastGuts               bool     `json:"InjuryPenaltyOnlyPastGuts"`
	InstabilityModifierDivisor              int64    `json:"InstabilityModifierDivisor"`
	InstabilityModifierFloor                float64  `json:"InstabilityModifierFloor"`
	InstabilityReductionAssault             float64  `json:"InstabilityReductionAssault"`
	InstabilityReductionHeavy               float64  `json:"InstabilityReductionHeavy"`
	InstabilityReductionLight               float64  `json:"InstabilityReductionLight"`
	InstabilityReductionMedium              float64  `json:"InstabilityReductionMedium"`
	LegDamageRelativeInstability            float64  `json:"LegDamageRelativeInstability"`
	LocationDestroyedPermanentStabilityLoss float64  `json:"LocationDestroyedPermanentStabilityLoss"`
	OnlyPermanentLossFromLegs               bool     `json:"OnlyPermanentLossFromLegs"`
	PilotingBaseFloor                       float64  `json:"PilotingBaseFloor"`
	PilotingDFAModifier                     int64    `json:"PilotingDFAModifier"`
	PilotingDamageCTPenalizedModifier       float64  `json:"PilotingDamageCTPenalizedModifier"`
	PilotingDamageLegNonfunctionalModifier  float64  `json:"PilotingDamageLegNonfunctionalModifier"`
	PilotingDamageLegPenalizedModifier      float64  `json:"PilotingDamageLegPenalizedModifier"`
	PilotingDivisor                         int64    `json:"PilotingDivisor"`
	PilotingModifierDivisor                 int64    `json:"PilotingModifierDivisor"`
	SideTorsoBlackRelativeInstability       float64  `json:"SideTorsoBlackRelativeInstability"`
	TurretAbilities                         []string `json:"TurretAbilities"`
	UseFixedInjuryPenalty                   bool     `json:"UseFixedInjuryPenalty"`
}

type GameConstants_sub6 struct {
	ArmorMultiplierAssault     int64   `json:"ArmorMultiplierAssault"`
	ArmorMultiplierHeavy       int64   `json:"ArmorMultiplierHeavy"`
	ArmorMultiplierLight       int64   `json:"ArmorMultiplierLight"`
	ArmorMultiplierMedium      int64   `json:"ArmorMultiplierMedium"`
	ArmorMultiplierVehicle     float64 `json:"ArmorMultiplierVehicle"`
	GlobalDamageMultiplier     float64 `json:"GlobalDamageMultiplier"`
	StructureMultiplierAssault int64   `json:"StructureMultiplierAssault"`
	StructureMultiplierHeavy   int64   `json:"StructureMultiplierHeavy"`
	StructureMultiplierLight   int64   `json:"StructureMultiplierLight"`
	StructureMultiplierMedium  int64   `json:"StructureMultiplierMedium"`
	StructureMultiplierVehicle float64 `json:"StructureMultiplierVehicle"`
	TESTMechScaleMultiplier    float64 `json:"TEST_MechScaleMultiplier"`
}

type GameConstants_sub30 struct {
	ArtilleryExplosion                GameConstants_sub29   `json:"artillery_explosion"`
	ArtilleryImpactBarrage            GameConstants_sub29   `json:"artillery_impact_barrage"`
	ArtilleryImpactSingle             GameConstants_sub29   `json:"artillery_impact_single"`
	ArtilleryIncoming                 GameConstants_sub29   `json:"artillery_incoming"`
	ArtilleryOrbitalPpc               GameConstants_sub29   `json:"artillery_orbital_ppc"`
	ArtilleryProjectileBarrage        GameConstants_sub29   `json:"artillery_projectile_barrage"`
	ArtilleryProjectileSingle         GameConstants_sub29   `json:"artillery_projectile_single"`
	ComponentDestructionA             GameConstants_sub29   `json:"componentDestruction_A"`
	ComponentDestructionAmmoExplosion GameConstants_sub29   `json:"componentDestruction_AmmoExplosion"`
	ComponentDestructionB             GameConstants_sub29   `json:"componentDestruction_B"`
	ComponentDestructionC             GameConstants_sub29   `json:"componentDestruction_C"`
	ComponentDestructionD             GameConstants_sub29   `json:"componentDestruction_D"`
	DeadMechLoop                      GameConstants_sub29   `json:"deadMechLoop"`
	DeadMechLoopA                     GameConstants_sub29   `json:"deadMechLoop_A"`
	DeadMechLoopB                     GameConstants_sub29   `json:"deadMechLoop_B"`
	DeadMechLoopC                     GameConstants_sub29   `json:"deadMechLoop_C"`
	DeadMechLoopD                     GameConstants_sub29   `json:"deadMechLoop_D"`
	DirectionalExplosionElectrical    GameConstants_sub29   `json:"directionalExplosion_electrical"`
	DirectionalExplosionEnergy        GameConstants_sub29   `json:"directionalExplosion_energy"`
	DirectionalExplosionLarge         GameConstants_sub29   `json:"directionalExplosion_large"`
	DirectionalExplosionMedium        GameConstants_sub29   `json:"directionalExplosion_medium"`
	DirectionalExplosionSmall         GameConstants_sub29   `json:"directionalExplosion_small"`
	FlamerImpact                      GameConstants_sub29   `json:"flamer_impact"`
	FlamerPersistent                  GameConstants_sub29   `json:"flamer_persistent"`
	FlamesLoopLarge                   GameConstants_sub29   `json:"flames_loop_large"`
	FlamesLoopMedium                  GameConstants_sub29   `json:"flames_loop_medium"`
	FlamesLoopSmall                   GameConstants_sub29   `json:"flames_loop_small"`
	FlamesSingleLarge                 GameConstants_sub29   `json:"flames_single_large"`
	FlamesSingleSmall                 GameConstants_sub29   `json:"flames_single_small"`
	FlamesTreesLarge                  GameConstants_sub29   `json:"flames_trees_large"`
	FlamesTreesSmall                  GameConstants_sub29   `json:"flames_trees_small"`
	FootfallBase                      GameConstants_sub29   `json:"footfallBase"`
	GroundImpactBase                  GameConstants_sub29   `json:"groundImpactBase"`
	HeatCoolantVent                   GameConstants_sub29   `json:"heat_coolantVent"`
	HeatHeatShutdown                  GameConstants_sub29   `json:"heat_heatShutdown"`
	HeatHeatVentLarge                 GameConstants_sub29   `json:"heat_heatVent_large"`
	HeatHeatVentShutdown              GameConstants_sub29   `json:"heat_heatVent_shutdown"`
	HeatHeatVentSmall                 GameConstants_sub29   `json:"heat_heatVent_small"`
	HeatHighHeatPersistent            GameConstants_sub29   `json:"heat_highHeat_persistent"`
	HeatMidHeatPersistent             GameConstants_sub29   `json:"heat_midHeat_persistent"`
	JumpjetLandBase                   GameConstants_sub29   `json:"jumpjetLandBase"`
	JumpjetLaunch                     GameConstants_sub29   `json:"jumpjet_launch"`
	MechDeathAmmoExplosionA           GameConstants_sub29   `json:"mechDeath_ammoExplosion_A"`
	MechDeathAmmoExplosionB           GameConstants_sub29   `json:"mechDeath_ammoExplosion_B"`
	MechDeathAmmoExplosionC           GameConstants_sub29   `json:"mechDeath_ammoExplosion_C"`
	MechDeathCenterTorsoA             GameConstants_sub29   `json:"mechDeath_centerTorso_A"`
	MechDeathCenterTorsoB             GameConstants_sub29   `json:"mechDeath_centerTorso_B"`
	MechDeathCenterTorsoC             GameConstants_sub29   `json:"mechDeath_centerTorso_C"`
	MechDeathCockpit                  GameConstants_sub29   `json:"mechDeath_cockpit"`
	MechDeathEngine                   GameConstants_sub29   `json:"mechDeath_engine"`
	MechDeathGyro                     GameConstants_sub29   `json:"mechDeath_gyro"`
	MechDeathLegs                     GameConstants_sub29   `json:"mechDeath_legs"`
	MechDeathPilotKilled              GameConstants_sub29   `json:"mechDeath_pilotKilled"`
	MechDeathVitalComponent           GameConstants_sub29   `json:"mechDeath_vitalComponent"`
	PersistentCrit                    []GameConstants_sub29 `json:"persistentCrit"`
	PersistentDamage                  []GameConstants_sub29 `json:"persistentDamage"`
	PilotEject                        GameConstants_sub29   `json:"pilotEject"`
	VehicleDeathA                     GameConstants_sub29   `json:"vehicleDeath_A"`
	VehicleDeathB                     GameConstants_sub29   `json:"vehicleDeath_B"`
}

type GameConstants_sub19 struct {
	AutoInspire                         bool                `json:"AutoInspire"`
	BaselineAddFromSimGame              bool                `json:"BaselineAddFromSimGame"`
	BaselineAddFromSimGameThresholds    []int64             `json:"BaselineAddFromSimGameThresholds"`
	BaselineAddFromSimGameValues        []int64             `json:"BaselineAddFromSimGameValues"`
	BaselineAddMoraleMods               bool                `json:"BaselineAddMoraleMods"`
	BaselineModsOnlyBiggest             bool                `json:"BaselineModsOnlyBiggest"`
	BaselineMultiplyByActiveElites      bool                `json:"BaselineMultiplyByActiveElites"`
	CameraHighlight                     bool                `json:"CameraHighlight"`
	CanAIBeInspired                     bool                `json:"CanAIBeInspired"`
	CanUseInspireLevel                  int64               `json:"CanUseInspireLevel"`
	ChangeAllyAmmoExplodes              int64               `json:"ChangeAllyAmmoExplodes"`
	ChangeAllyCrit                      int64               `json:"ChangeAllyCrit"`
	ChangeAllyDestroyedAssault          float64             `json:"ChangeAllyDestroyedAssault"`
	ChangeAllyDestroyedHeavy            int64               `json:"ChangeAllyDestroyedHeavy"`
	ChangeAllyDestroyedLight            int64               `json:"ChangeAllyDestroyedLight"`
	ChangeAllyDestroyedMedium           float64             `json:"ChangeAllyDestroyedMedium"`
	ChangeAllyDestroyedMeleeAdditional  int64               `json:"ChangeAllyDestroyedMeleeAdditional"`
	ChangeAllyKnockedDown               int64               `json:"ChangeAllyKnockedDown"`
	ChangeAllyLocationDestroyed         int64               `json:"ChangeAllyLocationDestroyed"`
	ChangeAllyMajorArmorDamage          int64               `json:"ChangeAllyMajorArmorDamage"`
	ChangeAllyMinorArmorDamage          int64               `json:"ChangeAllyMinorArmorDamage"`
	ChangeAllyWeaponDestroyed           int64               `json:"ChangeAllyWeaponDestroyed"`
	ChangeDFADealt                      int64               `json:"ChangeDFADealt"`
	ChangeDFAReceived                   int64               `json:"ChangeDFAReceived"`
	ChangeEnemyAmmoExplodes             int64               `json:"ChangeEnemyAmmoExplodes"`
	ChangeEnemyCrit                     int64               `json:"ChangeEnemyCrit"`
	ChangeEnemyDestroyedAssault         int64               `json:"ChangeEnemyDestroyedAssault"`
	ChangeEnemyDestroyedHeavy           int64               `json:"ChangeEnemyDestroyedHeavy"`
	ChangeEnemyDestroyedLight           int64               `json:"ChangeEnemyDestroyedLight"`
	ChangeEnemyDestroyedMedium          int64               `json:"ChangeEnemyDestroyedMedium"`
	ChangeEnemyDestroyedMeleeAdditional int64               `json:"ChangeEnemyDestroyedMeleeAdditional"`
	ChangeEnemyKnockedDown              int64               `json:"ChangeEnemyKnockedDown"`
	ChangeEnemyLocationDestroyed        int64               `json:"ChangeEnemyLocationDestroyed"`
	ChangeEnemyMajorArmorDamage         int64               `json:"ChangeEnemyMajorArmorDamage"`
	ChangeEnemyMinorArmorDamage         float64             `json:"ChangeEnemyMinorArmorDamage"`
	ChangeEnemyWeaponDestroyed          int64               `json:"ChangeEnemyWeaponDestroyed"`
	ChangeMajorityAttackingShotsHit     int64               `json:"ChangeMajorityAttackingShotsHit"`
	ChangeObjectiveCompleted            int64               `json:"ChangeObjectiveCompleted"`
	ChangeObjectiveFailed               int64               `json:"ChangeObjectiveFailed"`
	DefensivePushCost                   int64               `json:"DefensivePushCost"`
	DefensivePushHighMoraleCost         int64               `json:"DefensivePushHighMoraleCost"`
	DefensivePushLowMoraleCost          int64               `json:"DefensivePushLowMoraleCost"`
	FreeStandUpAction                   bool                `json:"FreeStandUpAction"`
	InspirationAccelerationMultiplier   float64             `json:"InspirationAccelerationMultiplier"`
	InspireCost                         int64               `json:"InspireCost"`
	InspiredEffect                      GameConstants_sub18 `json:"InspiredEffect"`
	MoraleBaselineGainPerRound          int64               `json:"MoraleBaselineGainPerRound"`
	MoraleBaselineMultiplyByActive      bool                `json:"MoraleBaselineMultiplyByActive"`
	MoraleInitialDefault                int64               `json:"MoraleInitialDefault"`
	MoraleMax                           int64               `json:"MoraleMax"`
	MoraleMin                           int64               `json:"MoraleMin"`
	OffensivePushCost                   int64               `json:"OffensivePushCost"`
	OffensivePushHighMoraleCost         int64               `json:"OffensivePushHighMoraleCost"`
	OffensivePushLowMoraleCost          int64               `json:"OffensivePushLowMoraleCost"`
	ThresholdMajorArmor                 float64             `json:"ThresholdMajorArmor"`
	ThresholdMajorityHit                float64             `json:"ThresholdMajorityHit"`
	ThresholdMajorityMiss               float64             `json:"ThresholdMajorityMiss"`
	ThresholdMinorArmor                 float64             `json:"ThresholdMinorArmor"`
	UseBaselineMoraleGainEachRound      bool                `json:"UseBaselineMoraleGainEachRound"`
	UseDefensivePush                    bool                `json:"UseDefensivePush"`
	UseOffensivePush                    bool                `json:"UseOffensivePush"`
}

type GameConstants_sub12 struct {
	AutoInspire                         bool                `json:"AutoInspire"`
	CameraHighlight                     bool                `json:"CameraHighlight"`
	CanAIBeInspired                     bool                `json:"CanAIBeInspired"`
	CanUseInspireLevel                  int64               `json:"CanUseInspireLevel"`
	ChangeAllyAmmoExplodes              int64               `json:"ChangeAllyAmmoExplodes"`
	ChangeAllyCrit                      int64               `json:"ChangeAllyCrit"`
	ChangeAllyDestroyedAssault          int64               `json:"ChangeAllyDestroyedAssault"`
	ChangeAllyDestroyedHeavy            int64               `json:"ChangeAllyDestroyedHeavy"`
	ChangeAllyDestroyedLight            int64               `json:"ChangeAllyDestroyedLight"`
	ChangeAllyDestroyedMedium           int64               `json:"ChangeAllyDestroyedMedium"`
	ChangeAllyDestroyedMeleeAdditional  int64               `json:"ChangeAllyDestroyedMeleeAdditional"`
	ChangeAllyKnockedDown               int64               `json:"ChangeAllyKnockedDown"`
	ChangeAllyLocationDestroyed         int64               `json:"ChangeAllyLocationDestroyed"`
	ChangeAllyMajorArmorDamage          int64               `json:"ChangeAllyMajorArmorDamage"`
	ChangeAllyMinorArmorDamage          int64               `json:"ChangeAllyMinorArmorDamage"`
	ChangeAllyWeaponDestroyed           int64               `json:"ChangeAllyWeaponDestroyed"`
	ChangeDFADealt                      int64               `json:"ChangeDFADealt"`
	ChangeDFAReceived                   int64               `json:"ChangeDFAReceived"`
	ChangeEnemyAmmoExplodes             int64               `json:"ChangeEnemyAmmoExplodes"`
	ChangeEnemyCrit                     int64               `json:"ChangeEnemyCrit"`
	ChangeEnemyDestroyedAssault         int64               `json:"ChangeEnemyDestroyedAssault"`
	ChangeEnemyDestroyedHeavy           int64               `json:"ChangeEnemyDestroyedHeavy"`
	ChangeEnemyDestroyedLight           int64               `json:"ChangeEnemyDestroyedLight"`
	ChangeEnemyDestroyedMedium          int64               `json:"ChangeEnemyDestroyedMedium"`
	ChangeEnemyDestroyedMeleeAdditional int64               `json:"ChangeEnemyDestroyedMeleeAdditional"`
	ChangeEnemyKnockedDown              int64               `json:"ChangeEnemyKnockedDown"`
	ChangeEnemyLocationDestroyed        int64               `json:"ChangeEnemyLocationDestroyed"`
	ChangeEnemyMajorArmorDamage         int64               `json:"ChangeEnemyMajorArmorDamage"`
	ChangeEnemyMinorArmorDamage         int64               `json:"ChangeEnemyMinorArmorDamage"`
	ChangeEnemyWeaponDestroyed          int64               `json:"ChangeEnemyWeaponDestroyed"`
	ChangeMajorityAttackingShotsHit     int64               `json:"ChangeMajorityAttackingShotsHit"`
	ChangeObjectiveCompleted            int64               `json:"ChangeObjectiveCompleted"`
	ChangeObjectiveFailed               int64               `json:"ChangeObjectiveFailed"`
	FreeStandUpAction                   bool                `json:"FreeStandUpAction"`
	InspirationAccelerationMultiplier   float64             `json:"InspirationAccelerationMultiplier"`
	InspireCost                         int64               `json:"InspireCost"`
	InspiredEffect                      GameConstants_sub11 `json:"InspiredEffect"`
	MoraleBaselineGainPerRound          int64               `json:"MoraleBaselineGainPerRound"`
	MoraleBaselineMultiplyByActive      bool                `json:"MoraleBaselineMultiplyByActive"`
	MoraleInitialDefault                int64               `json:"MoraleInitialDefault"`
	MoraleMax                           int64               `json:"MoraleMax"`
	MoraleMin                           int64               `json:"MoraleMin"`
	OffensivePushCost                   int64               `json:"OffensivePushCost"`
	OffensivePushHighMoraleCost         int64               `json:"OffensivePushHighMoraleCost"`
	OffensivePushLowMoraleCost          int64               `json:"OffensivePushLowMoraleCost"`
	ThresholdMajorArmor                 float64             `json:"ThresholdMajorArmor"`
	ThresholdMajorityHit                float64             `json:"ThresholdMajorityHit"`
	ThresholdMajorityMiss               float64             `json:"ThresholdMajorityMiss"`
	ThresholdMinorArmor                 float64             `json:"ThresholdMinorArmor"`
	UseBaselineMoraleGainEachRound      bool                `json:"UseBaselineMoraleGainEachRound"`
	UseDefensivePush                    bool                `json:"UseDefensivePush"`
	UseOffensivePush                    bool                `json:"UseOffensivePush"`
}

type GameConstants_sub7 struct {
	BrawlerMaxFrac       float64 `json:"brawlerMaxFrac"`
	BrawlerMinAbs        int64   `json:"brawlerMinAbs"`
	BrawlerTagMultiplier float64 `json:"brawlerTagMultiplier"`
	Hysteresis           float64 `json:"hysteresis"`
	SniperMaxFrac        float64 `json:"sniperMaxFrac"`
	SniperMinAbs         int64   `json:"sniperMinAbs"`
	SniperTagMultiplier  float64 `json:"sniperTagMultiplier"`
}

type GameConstants_sub29 struct {
	CacheCount int64  `json:"cacheCount"`
	Name       string `json:"name"`
}

type GameConstants_sub15 struct {
	CalledShotBonusDegradeLerpFactor float64               `json:"CalledShotBonusDegradeLerpFactor"`
	CalledShotBonusMultiplier        float64               `json:"CalledShotBonusMultiplier"`
	HitMechLocationFromArtillery     []GameConstants_sub14 `json:"HitMechLocationFromArtillery"`
	HitMechLocationFromBack          []GameConstants_sub14 `json:"HitMechLocationFromBack"`
	HitMechLocationFromFront         []GameConstants_sub14 `json:"HitMechLocationFromFront"`
	HitMechLocationFromLeft          []GameConstants_sub14 `json:"HitMechLocationFromLeft"`
	HitMechLocationFromRight         []GameConstants_sub14 `json:"HitMechLocationFromRight"`
	HitMechLocationFromTop           []GameConstants_sub14 `json:"HitMechLocationFromTop"`
	HitMechLocationProne             []GameConstants_sub14 `json:"HitMechLocationProne"`
	HitVehicleLocationFromArtillery  []GameConstants_sub14 `json:"HitVehicleLocationFromArtillery"`
	HitVehicleLocationFromBack       []GameConstants_sub14 `json:"HitVehicleLocationFromBack"`
	HitVehicleLocationFromFront      []GameConstants_sub14 `json:"HitVehicleLocationFromFront"`
	HitVehicleLocationFromLeft       []GameConstants_sub14 `json:"HitVehicleLocationFromLeft"`
	HitVehicleLocationFromRight      []GameConstants_sub14 `json:"HitVehicleLocationFromRight"`
	HitVehicleLocationFromTop        []GameConstants_sub14 `json:"HitVehicleLocationFromTop"`
}

type GameConstants_sub1 struct {
	CamDelayBuildingDeath        float64 `json:"CamDelayBuildingDeath"`
	CamDelayComponentDestroyed   float64 `json:"CamDelayComponentDestroyed"`
	CamDelayCriticalHit          float64 `json:"CamDelayCriticalHit"`
	CamDelayDeath                float64 `json:"CamDelayDeath"`
	CamDelayFall                 float64 `json:"CamDelayFall"`
	CamDelayLocationDestroyed    float64 `json:"CamDelayLocationDestroyed"`
	CamDelayMiscInfo             float64 `json:"CamDelayMiscInfo"`
	CamDelayPilotInjured         float64 `json:"CamDelayPilotInjured"`
	CamDelayPilotInspired        float64 `json:"CamDelayPilotInspired"`
	CamMultiplierSlowMove        float64 `json:"CamMultiplierSlowMove"`
	CutAngleTolerance            float64 `json:"CutAngleTolerance"`
	DefaultAngle                 float64 `json:"DefaultAngle"`
	DefaultFOV                   float64 `json:"DefaultFOV"`
	EdgeAutoZoomSize             float64 `json:"EdgeAutoZoomSize"`
	EdgePushSize                 float64 `json:"EdgePushSize"`
	FollowCamDelayTime           float64 `json:"FollowCamDelayTime"`
	FollowCamGroundMin           float64 `json:"FollowCamGroundMin"`
	FollowCamIdealHeight         float64 `json:"FollowCamIdealHeight"`
	FollowCamMaxDistMult         float64 `json:"FollowCamMaxDistMult"`
	FollowCamMinDistMult         float64 `json:"FollowCamMinDistMult"`
	FollowCamTime                float64 `json:"FollowCamTime"`
	FrameRightMultiplier         float64 `json:"FrameRightMultiplier"`
	FrameUpMultiplier            float64 `json:"FrameUpMultiplier"`
	FramingSafeMultiplier        float64 `json:"FramingSafeMultiplier"`
	FramingZoomMultiplier        float64 `json:"FramingZoomMultiplier"`
	MaxAngle                     float64 `json:"MaxAngle"`
	MaxHeightAboveTerrain        float64 `json:"MaxHeightAboveTerrain"`
	MaxSpeed                     float64 `json:"MaxSpeed"`
	MeleeAttackCamMinHeight      float64 `json:"MeleeAttackCamMinHeight"`
	MinAngle                     float64 `json:"MinAngle"`
	MinDistToFocalPoint          float64 `json:"MinDistToFocalPoint"`
	MinHeightAboveTerrain        float64 `json:"MinHeightAboveTerrain"`
	MoveFalloff                  float64 `json:"MoveFalloff"`
	MoveSpeed                    float64 `json:"MoveSpeed"`
	MovementCamTime              float64 `json:"MovementCamTime"`
	RestorePlayerCamTime         float64 `json:"RestorePlayerCamTime"`
	RotationSpeed                float64 `json:"RotationSpeed"`
	RotationSpeedKeyboard        float64 `json:"RotationSpeedKeyboard"`
	ShakyCamOffset               float64 `json:"ShakyCamOffset"`
	ShakyCamTime                 float64 `json:"ShakyCamTime"`
	ShakyMoveCamHeightMultiplier float64 `json:"ShakyMoveCamHeightMultiplier"`
	ShowAttackCamTime            float64 `json:"ShowAttackCamTime"`
	TargetCamPosMultiplierBack   float64 `json:"TargetCamPosMultiplierBack"`
	TargetCamPosMultiplierHeight float64 `json:"TargetCamPosMultiplierHeight"`
	TargetCamPosMultiplierSide   float64 `json:"TargetCamPosMultiplierSide"`
	TargetCamTime                float64 `json:"TargetCamTime"`
	TargetRotLerpFactor          float64 `json:"TargetRotLerpFactor"`
	TimeBetweenTutorials         float64 `json:"TimeBetweenTutorials"`
	TimeToShowNewBlips           float64 `json:"TimeToShowNewBlips"`
	TutorialMoveTime             float64 `json:"TutorialMoveTime"`
	TutorialRotateTime           float64 `json:"TutorialRotateTime"`
	ZoomSpeed                    float64 `json:"ZoomSpeed"`
}

type GameConstants_sub28 struct {
	ClusterChanceAdjacentMultiplier         int64         `json:"ClusterChanceAdjacentMultiplier"`
	ClusterChanceNeverClusterHead           bool          `json:"ClusterChanceNeverClusterHead"`
	ClusterChanceNeverMultiplyHead          bool          `json:"ClusterChanceNeverMultiplyHead"`
	ClusterChanceNonadjacentMultiplier      int64         `json:"ClusterChanceNonadjacentMultiplier"`
	ClusterChanceOriginalLocationMultiplier int64         `json:"ClusterChanceOriginalLocationMultiplier"`
	DamageResistanceIndirectFire            int64         `json:"DamageResistanceIndirectFire"`
	DamageResistanceObstructed              float64       `json:"DamageResistanceObstructed"`
	EvasivePipsMovingTarget                 []int64       `json:"EvasivePipsMovingTarget"`
	FiringArcDegrees                        float64       `json:"FiringArcDegrees"`
	SensorLockPipsStripped                  int64         `json:"SensorLockPipsStripped"`
	SensorLockStripsEvasivePips             bool          `json:"SensorLockStripsEvasivePips"`
	TacticsBaseFloor                        float64       `json:"TacticsBaseFloor"`
	TacticsDivisor                          int64         `json:"TacticsDivisor"`
	TacticsModifierDivisor                  int64         `json:"TacticsModifierDivisor"`
	ToHitAssault                            int64         `json:"ToHitAssault"`
	ToHitBaseFloor                          float64       `json:"ToHitBaseFloor"`
	ToHitCoverObstructed                    int64         `json:"ToHitCoverObstructed"`
	ToHitElevationApplyPenalties            bool          `json:"ToHitElevationApplyPenalties"`
	ToHitElevationLevelHeight               int64         `json:"ToHitElevationLevelHeight"`
	ToHitElevationModifierPerLevel          int64         `json:"ToHitElevationModifierPerLevel"`
	ToHitElevationUseMultipleLevels         bool          `json:"ToHitElevationUseMultipleLevels"`
	ToHitGunneryDivisor                     int64         `json:"ToHitGunneryDivisor"`
	ToHitHeavy                              int64         `json:"ToHitHeavy"`
	ToHitIndirect                           int64         `json:"ToHitIndirect"`
	ToHitLight                              int64         `json:"ToHitLight"`
	ToHitLongRange                          int64         `json:"ToHitLongRange"`
	ToHitMaximumRange                       int64         `json:"ToHitMaximumRange"`
	ToHitMedium                             int64         `json:"ToHitMedium"`
	ToHitMediumRange                        int64         `json:"ToHitMediumRange"`
	ToHitMinimumRange                       int64         `json:"ToHitMinimumRange"`
	ToHitModifierDivisor                    int64         `json:"ToHitModifierDivisor"`
	ToHitMovingPipUMs                       []int64       `json:"ToHitMovingPipUMs"`
	ToHitMovingTargetDistances              []int64       `json:"ToHitMovingTargetDistances"`
	ToHitMovingTargetStrings                []string      `json:"ToHitMovingTargetStrings"`
	ToHitOffensivePush                      int64         `json:"ToHitOffensivePush"`
	ToHitProneTargetBonus                   int64         `json:"ToHitProneTargetBonus"`
	ToHitProneTargetBonusMelee              int64         `json:"ToHitProneTargetBonusMelee"`
	ToHitSelfArmMountedWeapon               int64         `json:"ToHitSelfArmMountedWeapon"`
	ToHitSelfOverheated                     int64         `json:"ToHitSelfOverheated"`
	ToHitSelfSprinted                       int64         `json:"ToHitSelfSprinted"`
	ToHitSelfStoodUp                        int64         `json:"ToHitSelfStoodUp"`
	ToHitSelfWalkAssault                    int64         `json:"ToHitSelfWalkAssault"`
	ToHitSelfWalkHeavy                      int64         `json:"ToHitSelfWalkHeavy"`
	ToHitSelfWalkLight                      int64         `json:"ToHitSelfWalkLight"`
	ToHitSelfWalkMedium                     int64         `json:"ToHitSelfWalkMedium"`
	ToHitSelfWalkVehicle                    int64         `json:"ToHitSelfWalkVehicle"`
	ToHitSelfWeaponDamaged                  int64         `json:"ToHitSelfWeaponDamaged"`
	ToHitSelfWeaponLocationDamagedMajor     int64         `json:"ToHitSelfWeaponLocationDamagedMajor"`
	ToHitSelfWeaponLocationDamagedMinor     int64         `json:"ToHitSelfWeaponLocationDamagedMinor"`
	ToHitShortRange                         int64         `json:"ToHitShortRange"`
	ToHitShutdownTargetBonus                int64         `json:"ToHitShutdownTargetBonus"`
	ToHitShutdownTargetBonusMelee           int64         `json:"ToHitShutdownTargetBonusMelee"`
	ToHitStepThresholds                     []int64       `json:"ToHitStepThresholds"`
	ToHitStepValues                         []float64     `json:"ToHitStepValues"`
	ToHitTurretAssault                      int64         `json:"ToHitTurretAssault"`
	ToHitTurretHeavy                        int64         `json:"ToHitTurretHeavy"`
	ToHitTurretLight                        int64         `json:"ToHitTurretLight"`
	ToHitTurretMedium                       int64         `json:"ToHitTurretMedium"`
	ToHitUseSteppedAlgorithm                bool          `json:"ToHitUseSteppedAlgorithm"`
	ToHitVehicleAssault                     int64         `json:"ToHitVehicleAssault"`
	ToHitVehicleHeavy                       int64         `json:"ToHitVehicleHeavy"`
	ToHitVehicleLight                       int64         `json:"ToHitVehicleLight"`
	ToHitVehicleMedium                      int64         `json:"ToHitVehicleMedium"`
	UseMultipleGuardLevels                  bool          `json:"UseMultipleGuardLevels"`
	WeaponCategoriesAffectedByEvasive       []int64       `json:"WeaponCategoriesAffectedByEvasive"`
	WeaponTypesGuaranteedEvasiveStrip       []interface{} `json:"WeaponTypesGuaranteedEvasiveStrip"`
	WeaponsAffectedByEvasive                []int64       `json:"WeaponsAffectedByEvasive"`
}

type GameConstants_sub13 struct {
	CriticalHeatPerLocationAssault     int64   `json:"CriticalHeatPerLocationAssault"`
	CriticalHeatPerLocationHeavy       int64   `json:"CriticalHeatPerLocationHeavy"`
	CriticalHeatPerLocationLight       int64   `json:"CriticalHeatPerLocationLight"`
	CriticalHeatPerLocationMedium      int64   `json:"CriticalHeatPerLocationMedium"`
	DefaultHeatSinkDissipationCapacity int64   `json:"DefaultHeatSinkDissipationCapacity"`
	EngineDamageHeat                   int64   `json:"EngineDamageHeat"`
	GlobalHeatIncreaseMultiplier       float64 `json:"GlobalHeatIncreaseMultiplier"`
	GlobalHeatSinkMultiplier           float64 `json:"GlobalHeatSinkMultiplier"`
	InternalHeatSinkCount              int64   `json:"InternalHeatSinkCount"`
	JumpHeatMin                        int64   `json:"JumpHeatMin"`
	JumpHeatPerUnit                    float64 `json:"JumpHeatPerUnit"`
	JumpHeatUnitSize                   float64 `json:"JumpHeatUnitSize"`
	MaxHeat                            int64   `json:"MaxHeat"`
	OverheatLevel                      float64 `json:"OverheatLevel"`
	ShowPeakHeat                       bool    `json:"ShowPeakHeat"`
	ShutdownCausesInjury               bool    `json:"ShutdownCausesInjury"`
	SprintHeat                         int64   `json:"SprintHeat"`
	StartupSinkRatio                   float64 `json:"StartupSinkRatio"`
	WalkHeat                           int64   `json:"WalkHeat"`
}

type GameConstants_sub20 struct {
	DEBUGJumpjetDuration         float64 `json:"DEBUG_JumpjetDuration"`
	DEBUGUseJumpjetDuration      bool    `json:"DEBUG_UseJumpjetDuration"`
	DFAExecutionHeight           float64 `json:"DFAExecutionHeight"`
	DFAMinHeight                 float64 `json:"DFAMinHeight"`
	DefaultMeleeEngageRange      float64 `json:"DefaultMeleeEngageRange"`
	ExperimentalDotDistantColorB float64 `json:"ExperimentalDotDistantColorB"`
	ExperimentalDotDistantColorG float64 `json:"ExperimentalDotDistantColorG"`
	ExperimentalDotDistantColorR float64 `json:"ExperimentalDotDistantColorR"`
	ExperimentalDotForestColorB  float64 `json:"ExperimentalDotForestColorB"`
	ExperimentalDotForestColorG  float64 `json:"ExperimentalDotForestColorG"`
	ExperimentalDotForestColorR  float64 `json:"ExperimentalDotForestColorR"`
	ExperimentalDotWaterColorB   float64 `json:"ExperimentalDotWaterColorB"`
	ExperimentalDotWaterColorG   float64 `json:"ExperimentalDotWaterColorG"`
	ExperimentalDotWaterColorR   float64 `json:"ExperimentalDotWaterColorR"`
	ExperimentalFlag             bool    `json:"ExperimentalFlag"`
	ExperimentalGridDistance     float64 `json:"ExperimentalGridDistance"`
	ExperimentalHexRadius        int64   `json:"ExperimentalHexRadius"`
	JumpjetMinHeight             float64 `json:"JumpjetMinHeight"`
	JumpjetVelocityMultiplier    float64 `json:"JumpjetVelocityMultiplier"`
	LegDamageRedPenalty          float64 `json:"LegDamageRedPenalty"`
	LegDamageYellowPenalty       float64 `json:"LegDamageYellowPenalty"`
	LegDestroyedPenalty          float64 `json:"LegDestroyedPenalty"`
	MaxMeleeVerticalOffset       float64 `json:"MaxMeleeVerticalOffset"`
	MeleeDistance                float64 `json:"MeleeDistance"`
	MinMoveSpeed                 float64 `json:"MinMoveSpeed"`
	MoveTable                    []int64 `json:"MoveTable"`
	NumDFADestinationChoices     int64   `json:"NumDFADestinationChoices"`
	NumMeleeDestinationChoices   int64   `json:"NumMeleeDestinationChoices"`
	OverheatedMovePenalty        float64 `json:"OverheatedMovePenalty"`
	PathBlockerGradeMultiplier   float64 `json:"PathBlockerGradeMultiplier"`
	RotateInPlaceThreshold       float64 `json:"RotateInPlaceThreshold"`
	SortMeleeHexesByPathingCost  bool    `json:"SortMeleeHexesByPathingCost"`
}

type GameConstants_sub18 struct {
	Description   GameConstants_sub8  `json:"Description"`
	DurationData  GameConstants_sub16 `json:"durationData"`
	EffectType    string              `json:"effectType"`
	Nature        string              `json:"nature"`
	StatisticData GameConstants_sub17 `json:"statisticData"`
	TargetingData GameConstants_sub10 `json:"targetingData"`
}

type GameConstants_sub25 struct {
	Description   GameConstants_sub8  `json:"Description"`
	DurationData  GameConstants_sub23 `json:"durationData"`
	EffectType    string              `json:"effectType"`
	Nature        string              `json:"nature"`
	StatisticData GameConstants_sub17 `json:"statisticData"`
	TargetingData GameConstants_sub24 `json:"targetingData"`
}

type GameConstants_sub34 struct {
	Description   GameConstants_sub8  `json:"Description"`
	DurationData  GameConstants_sub32 `json:"durationData"`
	EffectType    string              `json:"effectType"`
	Nature        string              `json:"nature"`
	StatisticData GameConstants_sub17 `json:"statisticData"`
	TargetingData GameConstants_sub33 `json:"targetingData"`
}

type GameConstants_sub37 struct {
	Description   GameConstants_sub8  `json:"Description"`
	DurationData  GameConstants_sub35 `json:"durationData"`
	EffectType    string              `json:"effectType"`
	Nature        string              `json:"nature"`
	StatisticData GameConstants_sub17 `json:"statisticData"`
	TargetingData GameConstants_sub36 `json:"targetingData"`
}

type GameConstants_sub31 struct {
	Description   GameConstants_sub8  `json:"Description"`
	DurationData  GameConstants_sub9  `json:"durationData"`
	EffectType    string              `json:"effectType"`
	Nature        string              `json:"nature"`
	StatisticData GameConstants_sub17 `json:"statisticData"`
	TargetingData GameConstants_sub10 `json:"targetingData"`
}

type GameConstants_sub11 struct {
	Description   GameConstants_sub8  `json:"Description"`
	DurationData  GameConstants_sub9  `json:"durationData"`
	EffectType    string              `json:"effectType"`
	Nature        string              `json:"nature"`
	TargetingData GameConstants_sub10 `json:"targetingData"`
}

type GameConstants_sub8 struct {
	Details string `json:"Details"`
	Icon    string `json:"Icon"`
	ID      string `json:"Id"`
	Name    string `json:"Name"`
}

type GameConstants_sub3 struct {
	Details string `json:"Details"`
	Icon    string `json:"Icon"`
	Name    string `json:"Name"`
}

type GameConstants_sub2 struct {
	Details string `json:"Details"`
	Name    string `json:"Name"`
}

type GameConstants_sub35 struct {
	Duration               int64 `json:"duration"`
	StackLimit             int64 `json:"stackLimit"`
	TicksOnActivations     bool  `json:"ticksOnActivations"`
	UseActivationsOfTarget bool  `json:"useActivationsOfTarget"`
}

type GameConstants_sub32 struct {
	Duration          int64 `json:"duration"`
	StackLimit        int64 `json:"stackLimit"`
	TicksOnEndOfRound bool  `json:"ticksOnEndOfRound"`
}

type GameConstants_sub16 struct {
	Duration   int64 `json:"duration"`
	StackLimit int64 `json:"stackLimit"`
}

type GameConstants_sub9 struct {
	Duration               int64 `json:"duration"`
	TicksOnActivations     bool  `json:"ticksOnActivations"`
	UseActivationsOfTarget bool  `json:"useActivationsOfTarget"`
}

type GameConstants_sub23 struct {
	Duration           int64 `json:"duration"`
	TicksOnActivations bool  `json:"ticksOnActivations"`
}

type GameConstants_sub24 struct {
	EffectTargetType    string `json:"effectTargetType"`
	EffectTriggerType   string `json:"effectTriggerType"`
	ShowInStatusPanel   bool   `json:"showInStatusPanel"`
	ShowInTargetPreview bool   `json:"showInTargetPreview"`
}

type GameConstants_sub10 struct {
	EffectTargetType  string `json:"effectTargetType"`
	EffectTriggerType string `json:"effectTriggerType"`
}

type GameConstants_sub36 struct {
	EffectTargetType    string `json:"effectTargetType"`
	ForceVisRebuild     bool   `json:"forceVisRebuild"`
	ShowInStatusPanel   bool   `json:"showInStatusPanel"`
	ShowInTargetPreview bool   `json:"showInTargetPreview"`
}

type GameConstants_sub33 struct {
	EffectTargetType string `json:"effectTargetType"`
	ForceVisRebuild  bool   `json:"forceVisRebuild"`
}

type GameConstants_sub14 struct {
	K string `json:"k"`
	V int64  `json:"v"`
}

type GameConstants_sub4 struct {
	K string `json:"k"`
	V string `json:"v"`
}

type GameConstants_sub27 struct {
	MechDefaultGunnery     int64 `json:"MechDefaultGunnery"`
	MechDefaultGuts        int64 `json:"MechDefaultGuts"`
	MechDefaultPiloting    int64 `json:"MechDefaultPiloting"`
	MechDefaultTactics     int64 `json:"MechDefaultTactics"`
	TurretDefaultGunnery   int64 `json:"TurretDefaultGunnery"`
	TurretDefaultGuts      int64 `json:"TurretDefaultGuts"`
	TurretDefaultPiloting  int64 `json:"TurretDefaultPiloting"`
	TurretDefaultTactics   int64 `json:"TurretDefaultTactics"`
	VehicleDefaultGunnery  int64 `json:"VehicleDefaultGunnery"`
	VehicleDefaultGuts     int64 `json:"VehicleDefaultGuts"`
	VehicleDefaultPiloting int64 `json:"VehicleDefaultPiloting"`
	VehicleDefaultTactics  int64 `json:"VehicleDefaultTactics"`
}

type GameConstants_sub17 struct {
	ModType   string `json:"modType"`
	ModValue  string `json:"modValue"`
	Operation string `json:"operation"`
	StatName  string `json:"statName"`
}

type GameConstants_sub21 struct {
	PhaseAssault        int64 `json:"PhaseAssault"`
	PhaseAssaultTurret  int64 `json:"PhaseAssaultTurret"`
	PhaseAssaultVehicle int64 `json:"PhaseAssaultVehicle"`
	PhaseHeavy          int64 `json:"PhaseHeavy"`
	PhaseHeavyTurret    int64 `json:"PhaseHeavyTurret"`
	PhaseHeavyVehicle   int64 `json:"PhaseHeavyVehicle"`
	PhaseLight          int64 `json:"PhaseLight"`
	PhaseLightTurret    int64 `json:"PhaseLightTurret"`
	PhaseLightVehicle   int64 `json:"PhaseLightVehicle"`
	PhaseMedium         int64 `json:"PhaseMedium"`
	PhaseMediumTurret   int64 `json:"PhaseMediumTurret"`
	PhaseMediumVehicle  int64 `json:"PhaseMediumVehicle"`
	PhaseSpecial        int64 `json:"PhaseSpecial"`
}
