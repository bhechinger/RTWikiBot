package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func getDefFiles(defType string) []string {
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	err = filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(fmt.Sprintf("%s.*json$", defType), path)
			if err == nil && r {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func main() {
	//"buildingdef",
	files := getDefFiles("buildingdef")
	fmt.Printf("Found '%d' BuildingDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := BuildingDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		BuildingDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded BuildingDefs: %d\n", len(BuildingDefs))

	//"chassisdef",
	files = getDefFiles("chassisdef")
	fmt.Printf("Found '%d' ChassisDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := ChassisDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		ChassisDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded ChassisDefs: %d\n", len(ChassisDefs))

	//"hardpointdatadef",
	files = getDefFiles("hardpointdatadef")
	fmt.Printf("Found '%d' HardPointDataDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := HardPointDataDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		HardPointDataDefs[def.ID] = def
	}
	fmt.Printf("Loaded HardPointDataDefs: %d\n", len(HardPointDataDefs))

	//"heraldrydef",
	files = getDefFiles("heraldrydef")
	fmt.Printf("Found '%d' HeraldryDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := HeraldryDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		HeraldryDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded HeraldryDefs: %d\n", len(HeraldryDefs))

	//"lancedef",
	files = getDefFiles("lancedef")
	fmt.Printf("Found '%d' LanceDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := LanceDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		LanceDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded LanceDefs: %d\n", len(LanceDefs))

	//"mechdef",
	files = getDefFiles("mechdef")
	fmt.Printf("Found '%d' MechDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := MechDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		MechDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded MechDefs: %d\n", len(MechDefs))

	//"movedef",
	files = getDefFiles("movedef")
	fmt.Printf("Found '%d' MoveDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := MoveDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		MoveDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded MoveDefs: %d\n", len(MoveDefs))

	//"pathingdef",
	files = getDefFiles("pathingdef")
	fmt.Printf("Found '%d' PathingDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := PathingDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		PathingDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded PathingDefs: %d\n", len(PathingDefs))

	//"quirks",
	files = getDefFiles("quirks")
	fmt.Printf("Found '%d' Quirk files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := Quirk{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		Quirks[def.Description.ID] = def
	}
	fmt.Printf("Loaded Quirks: %d\n", len(Quirks))

	//"shopdef",
	files = getDefFiles("shopdef")
	fmt.Printf("Found '%d' ShopDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := ShopDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		ShopDefs[def.ID] = def
	}
	fmt.Printf("Loaded ShopDefs: %d\n", len(ShopDefs))

	//"starsystemdef",
	files = getDefFiles("starsystemdef")
	fmt.Printf("Found '%d' StarSystemDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := StarSystemDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		StarSystemDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded StarSystemDefs: %d\n", len(StarSystemDefs))

	//"turretchassisdef",
	files = getDefFiles("turretchassisdef")
	fmt.Printf("Found '%d' TurretChassisDefs files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := TurretChassisDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		TurretChassisDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded TurretChassisDefs: %d\n", len(TurretChassisDefs))

	//"turretdef",
	files = getDefFiles("turretdef")
	fmt.Printf("Found '%d' TurretDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := TurretDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		TurretDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded TurretDefs: %d\n", len(TurretDefs))

	//"vehiclechassisdef",
	files = getDefFiles("vehiclechassisdef")
	fmt.Printf("Found '%d' VehicleChassisDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := VehicleChassisDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		VehicleChassisDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded VehicleChassisDefs: %d\n", len(VehicleChassisDefs))

	//"vehicledef",
	files = getDefFiles("vehicledef")
	fmt.Printf("Found '%d' VehicleDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := VehicleDef{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		VehicleDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded VehicleDefs: %d\n", len(VehicleDefs))

	//"weapons",
	files = getDefFiles("weapons")
	fmt.Printf("Found '%d' Weapon files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		fileByte, _ := ioutil.ReadAll(fp)
		def := Weapon{}
		json.Unmarshal(fileByte, &def)
		fp.Close()
		Weapons[def.Description.ID] = def
	}
	fmt.Printf("Loaded Weapons: %d\n", len(Weapons))
}

//func (m * Mech) WikiOutput() (string, error) {
//	template_text := `{&.MechDef.Description.UIName&}
//{| class="wikitable sortable mw-collapsible" style="background: black"
//|-
//!Name
//!Signature
//|'''FACTIONS (3039+)'''
//!Weight!!Class!!Hardpoints!!Current Quirks
//|-
//!{{FP icon|Catapult.jpg|CATAPULT}}
//|{&.ChassisDef.VariantName&}
//|
//*
//*
//*
//*
//*
//|{&.ChassisDef.Tonnage&}
//|{&.ChassisDef.WeightClass&}
//|2B 4E 4M 0S 6JJ
//|[[{&.ChassisDef.StockRole&}]]
//LRM Ammo is "Pirate Special"
//
//+10% Jump Distance
//
//-Reducing ranges
//
//-90 min range
//
//dealing extra damage and have +1 heat damage per missile
//
//Cannot equip Special Missile Ammo
//|}
//{| class="wikitable" style="background: black"
//!style="color: grey" | Firepower:
//|style="color: grey" |58 DMG, 5 Stab.
//|-
//!style="color: silver" |Heat:
//|49 Heatsinking, 61 Alpha strike, 8 Jump, 110 Shutdown
//|-
//!style="color: grey" |Movement:
//|style="color: grey" |200 / 8 hex Sprint, 115 / 5 hex Walk, 150 Jump, 4 TT
//|-
//!style="color: silver" |Range:
//|580 max, 200 Opt
//|-
//!style="color: grey" |Durability:
//|style="color: grey" |521 Structure, 1040 armour, 100% stab def, 45 dfa self
//|-
//!style="color: silver" |Melee:
//|{&.ChassisDef.MeleeDamage&} Base DMG, -0 Quirk, 33 Total Dmg, {&.ChassisDef.MeleeInstability&} Stab., {&.ChassisDef.DFADamage&} DFA
//|}
//[[File:Catapult-CPLT-P.png|frameless|1316x1316px]]`
//
//	tmpl, err := template.New("test").Delims("{&", "&}").Parse(template_text)
//	if err != nil { panic(err) }
//	err = tmpl.Execute(os.Stdout, m)
//	if err != nil { panic(err) }
//
//	return "", nil
//}
