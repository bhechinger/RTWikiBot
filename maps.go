package main

//go:generate gojson -name BuildingDef -input json/buildingdef.json -o buildingdef.go
var BuildingDefs = map[string]BuildingDef{}

//go:generate gojson -name ChassisDef -input json/chassisdef.json -o chassisdef.go
var ChassisDefs = map[string]ChassisDef{}

//go:generate gojson -name HardPointDataDef -input json/hardpointdatadef.json -o hardpointdatadef.go
var HardPointDataDefs = map[string]HardPointDataDef{}

//go:generate gojson -name HeraldryDef -input json/heraldrydef.json -o heraldrydef.go
var HeraldryDefs = map[string]HeraldryDef{}

//go:generate gojson -name LanceDef -input json/lancedef.json -o lancedef.go
var LanceDefs = map[string]LanceDef{}

//go:generate gojson -name MechDef -input json/mechdef.json -o mechdef.go
var MechDefs = map[string]MechDef{}

//go:generate gojson -name MoveDef -input json/movedef.json -o movedef.go
var MoveDefs = map[string]MoveDef{}

//go:generate gojson -name PathingDef -input json/pathingdef.json -o pathingdef.go
var PathingDefs = map[string]PathingDef{}

//go:generate gojson -name Quirk -input json/quirks.json -o quirks.go
var Quirks = map[string]Quirk{}

//go:generate gojson -name ShopDef -input json/shopdef.json -o shopdef.go
var ShopDefs = map[string]ShopDef{}

//go:generate gojson -name StarSystemDef -input json/starsystemdef.json -o starsystemdef.go
var StarSystemDefs = map[string]StarSystemDef{}

//go:generate gojson -name TurretChassisDef -input json/turretchassisdef.json -o turretchassisdef.go
var TurretChassisDefs = map[string]TurretChassisDef{}

//go:generate gojson -name TurretDef -input json/turretdef.json -o turretdef.go
var TurretDefs = map[string]TurretDef{}

//go:generate gojson -name VehicleChassisDef -input json/vehiclechassisdef.json -o vehiclechassisdef.go
var VehicleChassisDefs = map[string]VehicleChassisDef{}

//go:generate gojson -name VehicleDef -input json/vehicledef.json -o vehicledef.go
var VehicleDefs = map[string]VehicleDef{}

//go:generate gojson -name Weapon -input json/weapons.json -o weapons.go
var Weapons = map[string]Weapon{}
