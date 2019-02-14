package main

import (
	"encoding/json"
	"fmt"
	"go.4amlunch.net/RTWikiBot/defs"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	loadData()
	generateTestMech("mechdef_catapult_CPLT-P")
	generateTestMech("mechdef_hatchetman_HCT-S7")
}

type HardPoints struct {
	AntiPersonnel int
	Ballistic     int
	Energy        int
	Missile       int
}

type Mech struct {
	ChassisDef defs.ChassisDef
	MechDef    defs.MechDef
	QuirkText  string
	HardPoints HardPoints
	Damage     int64
	Stability  float64
	Heat       int64
	HeatDmg    int64
	Structure  int64
	Armor      int64
	HeatSink   int64
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func generateTestMech(genmech string) {
	mechdef := MechDefs[genmech]
	chassisdef := ChassisDefs[mechdef.ChassisID]
	mech := &Mech{
		MechDef:    mechdef,
		ChassisDef: chassisdef,
	}

	feList := make(map[string]int)
	for e := range mech.ChassisDef.FixedEquipment {
		feList[mech.ChassisDef.FixedEquipment[e].ComponentDefID] += 1
	}

	var qt strings.Builder
	for e := range feList {
		qt.WriteString("* [[")
		qt.WriteString(Quirks[e].Description.Name)
		qt.WriteString("]] x")
		qt.WriteString(strconv.Itoa(feList[e]))
		qt.WriteString("\n")

		bonuses := Quirks[e].Custom.BonusDescriptions.Bonuses
		for q := range bonuses {
			qt.WriteString("** ")
			qt.WriteString(bonuses[q])
			qt.WriteString("\n")
		}
	}

	mech.QuirkText = qt.String()

	hardPoints := map[string]int{
		"AntiPersonnel": 0,
		"Ballistic":     0,
		"Energy":        0,
		"Missile":       0,
	}

	for l := range mech.MechDef.Locations {
		mech.Armor += mech.MechDef.Locations[l].AssignedArmor
	}

	for l := range mech.ChassisDef.Locations {
		location := mech.ChassisDef.Locations[l]
		for hp := range location.Hardpoints {
			hardPoints[location.Hardpoints[hp].WeaponMount] += 1
		}

		mech.Structure += location.InternalStructure
	}

	mech.HardPoints.AntiPersonnel = hardPoints["AntiPersonnel"]
	mech.HardPoints.Ballistic = hardPoints["Ballistic"]
	mech.HardPoints.Energy = hardPoints["Energy"]
	mech.HardPoints.Missile = hardPoints["Missile"]

	var engineMultiplier int64 = 1
	var engine int64
	for i := range mech.MechDef.Inventory {
		item := mech.MechDef.Inventory[i]
		if item.ComponentDefType == "Weapon" {
			mech.Damage += Weapons[item.ComponentDefID].Damage
			mech.Stability += Weapons[item.ComponentDefID].Instability
			mech.Heat += Weapons[item.ComponentDefID].HeatGenerated
			mech.HeatDmg += Weapons[item.ComponentDefID].HeatDamage
		}

		if item.ComponentDefType == "HeatSink" {
			if strings.Contains(item.ComponentDefID, "emod_engine_") {
				if strings.Contains(item.ComponentDefID, "emod_engine_cooling") {
					mech.HeatSink += EngineDefs[item.ComponentDefID].Custom.EngineHeatBlock.HeatSinkCount
				} else {
					engineRating, err := strconv.Atoi(EngineDefs[item.ComponentDefID].Custom.EngineCore.Rating)
					if err == nil {
						engine = int64(engineRating/25) * 3
					} else {
						fmt.Printf("Error converting to int: %s\n", err.Error())
					}
				}
			}

			if strings.Contains(item.ComponentDefID, "emod_kit") {
				if strings.Contains(EngineDefs[item.ComponentDefID].Custom.Cooling.HeatSinkDefID, "Double") {
					engineMultiplier = 2
				}

			}

			if strings.Contains(item.ComponentDefID, "Gear_HeatSink_") {
				mech.HeatSink += HeatSinkDefs[item.ComponentDefID].DissipationCapacity
			}
		}
	}

	mech.HeatSink += engineMultiplier * engine

	template_text := `{&.MechDef.Description.UIName&}
{| class="wikitable sortable mw-collapsible" style="background: black"
|-
!Name
!Signature
|'''FACTIONS (3039+)'''
!Weight!!Class!!Hardpoints!!Current Quirks
|-
!{{FP icon|Catapult.jpg|CATAPULT}}
|{&.ChassisDef.VariantName&}
|
*
*
*
*
*
|{&.ChassisDef.Tonnage&}
|{&.ChassisDef.WeightClass&}
|{&.HardPoints.Ballistic&}B {&.HardPoints.Energy&}E {&.HardPoints.Missile&}M {&.HardPoints.AntiPersonnel&}S {&.ChassisDef.MaxJumpjets&}JJ
|[[{&.ChassisDef.StockRole&}]]
{&.QuirkText&}

|}
{| class="wikitable" style="background: black"
!style="color: grey" | Firepower:
|style="color: grey" |{&.Damage&} DMG, {&.HeatDmg&} Heat DMG, {&.Stability&} Stab.
|-
!style="color: silver" |Heat:
|{&.HeatSink&} Heatsinking, {&.Heat&} Alpha strike, 8 Jump, 110 Shutdown
|-
!style="color: grey" |Movement:
|style="color: grey" |200 / 8 hex Sprint, 115 / 5 hex Walk, 150 Jump, 4 TT
|-
!style="color: silver" |Range:
|580 max, 200 Opt (To be removed?)
|-
!style="color: grey" |Durability:
|style="color: grey" |{&.Structure&} Structure, {&.Armor&} Armour, {&.ChassisDef.Stability&}% Stab Def, {&.ChassisDef.DFASelfDamage&} DFA Self DMG
|-
!style="color: silver" |Melee:
|{&.ChassisDef.MeleeDamage&} Base DMG, -0 Quirk, 33 Total Dmg, {&.ChassisDef.MeleeInstability&} Stab., {&.ChassisDef.DFADamage&} DFA
|}
[[File:Catapult-CPLT-P.png|frameless|1316x1316px]]
`

	tmpl, err := template.New("test").Delims("{&", "&}").Parse(template_text)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, mech)
	if err != nil {
		panic(err)
	}
}
