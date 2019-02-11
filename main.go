package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type Items map[string]map[string]interface{}

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
	//"chassisdef",
	files := getDefFiles("chassisdef")
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
		ChassisDefs[def.Description.UIName] = def
	}
	fmt.Printf("Loaded ChassisDefs: %d\n", len(ChassisDefs))

	//"hardpointdatadef",
	//"heraldrydef",
	//"lancedef",
	//"mechdef",
	files = getDefFiles("mechdef")
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
		MechDefs[def.Description.UIName] = def
	}
	fmt.Printf("Loaded MechDefs: %d\n", len(MechDefs))

	//"movedef",
	//"pathingdef",
	//"quirks",
	files = getDefFiles("quirks")
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
		Quirks[def.Description.UIName] = def
	}
	fmt.Printf("Loaded Quirks: %d\n", len(Quirks))

	//"shopdef",
	//"starsystemdef",
	//"turretchassisdef",
	//"turretdef",
	//"vehiclechassisdef",
	//"vehicledef",
	//"weapons",
	files = getDefFiles("weapons")
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
		Weapons[def.Description.UIName] = def
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
