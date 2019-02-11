package main

//go:generate gojson -name BuildingDef -input "../RogueTech/RogueTech Core/StreamingAssets/data/buildings/buildingdef_Civilian_Large_Defend.json" -o buildingdef.go
var BuildingDefs = map[string]BuildingDef{}

//go:generate gojson -name ChassisDef -input ../RogueTech/RogueMechs/chassis/chassisdef_catapult_CPLT-C21.json -o chassisdef.go
var ChassisDefs = map[string]ChassisDef{}

//go:generate gojson -name HardPointDataDef -input "../RogueTech/RogueTech Core/StreamingAssets/data/hardpoints/hardpointdatadef_jagermech.json" -o hardpointdatadef.go
var HardPointDataDefs = map[string]HardPointDataDef{}

//go:generate gojson -name HeraldryDef -input ../RogueTech/InnerSphereMap/Heraldry/heraldrydef_ClanGhostBear.json -o heraldrydef.go
var HeraldryDefs = map[string]HeraldryDef{}

//go:generate gojson -name LanceDef -input ../RogueTech/RogueContracts/lance/Convoy/lancedef_vehicle_d7_dynamic_convoy.json -o lancedef.go
var LanceDefs = map[string]LanceDef{}

//go:generate gojson -name MechDef -input ../RogueTech/RogueMechs/mech/mechdef_quickdraw_QKD-7X.json -o mechdef.go
var MechDefs = map[string]MechDef{}

//go:generate gojson -name MoveDef -input ../RogueTech/RogueMechs/move/movedef_catapult_CPLT-S7.json -o movedef.go
var MoveDefs = map[string]MoveDef{}

//go:generate gojson -name PathingDef -input ../RogueTech/RogueTanks/pathing/pathingdef_heavy_tracked.json -o pathingdef.go
var PathingDefs = map[string]PathingDef{}

//go:generate gojson -name Quirk -input ../RogueTech/RogueModuleTech/quirks/upgrade/Gear_FCS_ARTV.json -o quirks.go
var Quirks = map[string]Quirk{}

//go:generate gojson -name ShopDef -input ../RogueTech/JK_VariantsCampaign/shops/shopdef_JK_Mech_blackmarket.json -o shopdef.go
var ShopDefs = map[string]ShopDef{}

//go:generate gojson -name StarSystemDef -input ../RogueTech/InnerSphereMap/StarSystems/starsystemdef_Tangipahoa.json -o starsystemdef.go
var StarSystemDefs = map[string]StarSystemDef{}

//go:generate gojson -name TurretChassisDef -input ../RogueTech/RogueTurrets/TurretChassis/turretchassisdef_Heavy_lrm2.json -o turretchassisdef.go
var TurretChassisDefs = map[string]TurretChassisDef{}

//go:generate gojson -name TurretDef -input ../RogueTech/RogueTurrets/Turret/turretdef_Light_LRM3.json -o turretdef.go
var TurretDefs = map[string]TurretDef{}

//go:generate gojson -name VehicleChassisDef -input ../RogueTech/RogueTanks/CLANK/vehiclechassis/vehiclechassisdef_MARS__HAG_CLAN.json -o vehiclechassisdef.go
var VehicleChassisDefs = map[string]VehicleChassisDef{}

//go:generate gojson -name VehicleDef -input ../RogueTech/RogueTanks/CLANK/vehicle/vehicledef_EPONA_CLAN_RT.json -o vehicledef.go
var VehicleDefs = map[string]VehicleDef{}

//go:generate gojson -name Weapon -input ../RogueTech/RogueModuleTech/Primitive/weapons/Weapon_Autocannon_AC10_Quicsell.json -o weapons.go
var Weapons = map[string]Weapon{}
