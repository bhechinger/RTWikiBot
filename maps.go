package main

//go:generate gojson -name BuildingDef -input "../RogueTech/RogueTech Core/StreamingAssets/data/buildings/buildingdef_Civilian_Large_Defend.json" -o buildingdef.go
var BuildingDefs = map[string]BuildingDef{}

//go:generate gojson -name ChassisDef -input ../RogueTech/RogueMechs/chassis/chassisdef_catapult_CPLT-C21.json -o chassisdef.go
var ChassisDefs = map[string]ChassisDef{}

//go:generate gojson -name MechDef -input ../RogueTech/RogueMechs/mech/mechdef_quickdraw_QKD-7X.json -o mechdef.go
var MechDefs = map[string]MechDef{}

//go:generate gojson -name Quirk -input ../RogueTech/RogueModuleTech/quirks/upgrade/Gear_FCS_ARTV.json -o quirks.go
var Quirks = map[string]Quirk{}

//go:generate gojson -name Weapon -input ../RogueTech/RogueModuleTech/Primitive/weapons/Weapon_Autocannon_AC10_Quicsell.json -o weapons.go
var Weapons = map[string]Weapon{}
