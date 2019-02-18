package main

import "go.4amlunch.net/RTWikiBot/defs"

//go:generate gojson -name BuildingDef -pkg defs -input json/buildingdef.json -o defs/buildingdef.go
var BuildingDefs = map[string]defs.BuildingDef{}

//go:generate gojson -name ChassisDef -pkg defs -input json/chassisdef.json -o defs/chassisdef.go
var ChassisDefs = map[string]defs.ChassisDef{}

//go:generate gojson -name HardPointDataDef -pkg defs -input json/hardpointdatadef.json -o defs/hardpointdatadef.go
var HardPointDataDefs = map[string]defs.HardPointDataDef{}

//go:generate gojson -name HeraldryDef -pkg defs -input json/heraldrydef.json -o defs/heraldrydef.go
var HeraldryDefs = map[string]defs.HeraldryDef{}

//go:generate gojson -name LanceDef -pkg defs -input json/lancedef.json -o defs/lancedef.go
var LanceDefs = map[string]defs.LanceDef{}

//go:generate gojson -name MechDef -pkg defs -input json/mechdef.json -o defs/mechdef.go
var MechDefs = map[string]defs.MechDef{}

//go:generate gojson -name MoveDef -pkg defs -input json/movedef.json -o defs/movedef.go
var MoveDefs = map[string]defs.MoveDef{}

//go:generate gojson -name PathingDef -pkg defs -input json/pathingdef.json -o defs/pathingdef.go
var PathingDefs = map[string]defs.PathingDef{}

//go:generate gojson -name Quirk -pkg defs -input json/quirks.json -o defs/quirks.go
var Quirks = map[string]defs.Quirk{}

//go:generate gojson -name ShopDef -pkg defs -input json/shopdef.json -o defs/shopdef.go
var ShopDefs = map[string]defs.ShopDef{}

//go:generate gojson -name StarSystemDef -pkg defs -input json/starsystemdef.json -o defs/starsystemdef.go
var StarSystemDefs = map[string]defs.StarSystemDef{}

//go:generate gojson -name TurretChassisDef -pkg defs -input json/turretchassisdef.json -o defs/turretchassisdef.go
var TurretChassisDefs = map[string]defs.TurretChassisDef{}

//go:generate gojson -name TurretDef -pkg defs -input json/turretdef.json -o defs/turretdef.go
var TurretDefs = map[string]defs.TurretDef{}

//go:generate gojson -name VehicleChassisDef -pkg defs -input json/vehiclechassisdef.json -o defs/vehiclechassisdef.go
var VehicleChassisDefs = map[string]defs.VehicleChassisDef{}

//go:generate gojson -name VehicleDef -pkg defs -input json/vehicledef.json -o defs/vehicledef.go
var VehicleDefs = map[string]defs.VehicleDef{}

//go:generate gojson -name Weapon -pkg defs -input json/weapons.json -o defs/weapons.go
var Weapons = map[string]defs.Weapon{}

//go:generate gojson -name EngineDef -pkg defs -input json/emod.json -o defs/emod.go
var EngineDefs = map[string]defs.EngineDef{}

//go:generate gojson -name GearDef -pkg defs -input json/geardef.json -o defs/geardef.go
var GearDefs = map[string]defs.GearDef{}

//go:generate gojson -name Settings -pkg defs -input json/Settings.json -o defs/settings.go
