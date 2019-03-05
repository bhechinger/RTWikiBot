package main

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr"
	"go.4amlunch.net/RTWikiBot/defs"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"text/template"
)

// TODO: See why these are wrong

// HSN-8E
// stab is 48 should be 44
// jump heat should be 13 and shutdown should be 100
// hex movement should be 11/7 and jump should be 210
// Armor should be 520 stability should be 100%
// Damage should be 45 base -22 quirk DFA should be 90

// OSR-3D
// Stab should be 35
// Alpha should be 90 and jump should be 8
// movement should be 15/9 jump should be 150/8
// Armor should be 355
// Damage should me 30-15 DFA should be 60

// RCN-3L
// Alpha should be 37 jump should be 0 shutdown should be 100
// movement should be 11/7 - 6
// Armor should be 625
// dmg should be 35-17 70 DFA

// VND-1R
// Stab should be 25
// Alpha should be 64 jump should be 8 and shutdown 100
// movement should be 8/5 and jump should be 150
// Armor should be 720
// dmg should be 45-22 90 DFA

// HBK-4N
// Alpha should be 66 jump 0 100 shutdown
// movement 8/5 - 4
// Armor should be 800
// dmg should be 50-25 DFA should be 100

// JVN-10P
// Stab should be 21
// Alpha should be 32 , 13 jump 100 shutdown
// movement 11/7
// Armor should be 320
// dmg should be 30-15 60 DFA

// CN9-YLW
// Firepower should be 205
// Alpha should be 43 jump should be 6 and shutdown 100
// Movement should be 170/100 7/4 max jump 120
// dmg should be 55-27 DFA should be 110
// really does 48

func main() {
	loadChassisDefs()
	loadMechDefs()
	loadQuirkDefs()
	loadWeaponDefs()
	loadAmmoDefs()
	loadEngineDefs()
	loadGearDefs()

	generateTestMech("mechdef_phoenixhawk_PXH-IIC")
	generateTestMech("mechdef_catapult_CPLT-P")
	generateTestMech("mechdef_hatchetman_HCT-S7")
	generateTestMech("mechdef_locust_LCT-2V")
	generateTestMech("mechdef_hellspawn_HSN-8E")
	generateTestMech("mechdef_osiris_OSR-3D")
	generateTestMech("mechdef_raven_RVN-3L")
	generateTestMech("mechdef_vindicator_VND-1R")
	generateTestMech("mechdef_hunchback_HBK-4N")
	generateTestMech("mechdef_javelin_JVN-10P")
	generateTestMech("mechdef_centurion_CN9-YLW")
}

type HardPoints struct {
	AntiPersonnel int
	Ballistic     int
	Energy        int
	Missile       int
}

type Mech struct {
	ChassisDef    defs.ChassisDef
	MechDef       defs.MechDef
	QuirkText     string
	HardPoints    HardPoints
	Damage        int
	DFADamage     int
	DFASelfDamage int
	Melee         struct {
		Bonus int
		Total int
	}
	Stability int
	Heat      int
	HeatDmg   int
	Structure int
	Armor     int
	HeatSink  int
	JumpJets  int
	Distance  struct {
		Walk   int
		Sprint int
		Jump   int
	}
	Hex struct {
		Walk   int
		Sprint int
		Jump   int
	}
	JumpHeat   int
	Shutdown   int
	CAPName    string
	StockRoles string
}

// Globals
//var settings = loadSettings()
var box = packr.NewBox("./templates")
var heat = struct {
	OverheatLevel    float64
	MaxHeat          int
	WalkHeat         int
	SprintHeat       int
	JumpHeatUnitSize float64
	JumpHeatPerUnit  float64
	JumpHeatMin      int
	EngineDamageHeat int
}{
	0.41,
	110,
	3,
	8,
	6.5,
	0.9,
	3,
	5,
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func generateTestMech(genmech string) {
	bList := NewBonuses()
	mechdef := MechDefs[genmech]
	chassisdef := ChassisDefs[mechdef.ChassisID]
	mech := &Mech{
		MechDef:    mechdef,
		ChassisDef: chassisdef,
	}
	mech.CAPName = strings.ToUpper(chassisdef.PrefabBase)
	mech.StockRoles = strings.Replace(chassisdef.StockRole, " & ", "]] [[", -1)

	var movement = &Movement{Tonnage: chassisdef.Tonnage}

	var ammoType []string
	// I hate having to do this twice but I really need Ammo before I do the rest
	for i := range mech.MechDef.Inventory {
		item := mech.MechDef.Inventory[i]

		if item.ComponentDefType == "AmmunitionBox" {
			ammoType = append(ammoType, item.ComponentDefID)
		}
	}

	feList := make(map[string]int)
	for e := range mech.ChassisDef.FixedEquipment {
		feList[mech.ChassisDef.FixedEquipment[e].ComponentDefID] += 1
	}

	var qt strings.Builder
	for e := range feList {
		qt.WriteString("* [[")
		qt.WriteString(Quirks[e].Description.Name)
		if Quirks[e].ComponentType == "AmmunitionBox" {
			ammoType = append(ammoType, e)
		}
		qt.WriteString("]] x")
		qt.WriteString(strconv.Itoa(feList[e]))
		qt.WriteString("\n")

		bonuses := Quirks[e].Custom.BonusDescriptions.Bonuses
		for q := range bonuses {
			bList.AddBonus(bonuses[q])
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
		mech.Armor += int(mech.MechDef.Locations[l].AssignedArmor)
	}

	for l := range mech.ChassisDef.Locations {
		location := mech.ChassisDef.Locations[l]
		for hp := range location.Hardpoints {
			hardPoints[location.Hardpoints[hp].WeaponMount] += 1
		}

		mech.Structure += int(location.InternalStructure)
	}

	mech.HardPoints.AntiPersonnel = hardPoints["AntiPersonnel"]
	mech.HardPoints.Ballistic = hardPoints["Ballistic"]
	mech.HardPoints.Energy = hardPoints["Energy"]
	mech.HardPoints.Missile = hardPoints["Missile"]

	var engineHeatSinkDC int
	var weaponHeat int
	for i := range mech.MechDef.Inventory {
		item := mech.MechDef.Inventory[i]

		bonuses := GearDefs[item.ComponentDefID].Custom.BonusDescriptions.Bonuses
		if bonuses != nil {
			for b := range bonuses {
				bList.AddBonus(bonuses[b])
			}
		}

		if item.ComponentDefType == "JumpJet" {
			mech.JumpJets += 1
		}

		if item.ComponentDefType == "Weapon" {
			// select Ammo (we pick the first because how else do we know?)
			ammoBonuses := NewBonuses()
			for at := range ammoType {
				for ci := range AmmoDefs[ammoType[at]].Custom.Category {
					ammo := fmt.Sprintf("%sAmmo", Weapons[item.ComponentDefID].AmmoCategory)
					if AmmoDefs[ammoType[at]].Custom.Category[ci].CategoryID == ammo {
						for ab := range AmmoDefs[ammoType[at]].Custom.BonusDescriptions.Bonuses {
							ammoBonuses.AddBonus(AmmoDefs[ammoType[at]].Custom.BonusDescriptions.Bonuses[ab])
						}
						break
					}
				}
			}

			mech.Damage += ammoBonuses.ApplyBonus("targetDamage", int(Weapons[item.ComponentDefID].Damage)) * int(Weapons[item.ComponentDefID].ShotsWhenFired)
			mech.HeatDmg += ammoBonuses.ApplyBonus("targetHeat", int(Weapons[item.ComponentDefID].HeatDamage)) * int(Weapons[item.ComponentDefID].ShotsWhenFired)
			mech.Stability += ammoBonuses.ApplyBonus("targetStab", int(math.Round(Weapons[item.ComponentDefID].Instability))) * int(Weapons[item.ComponentDefID].ShotsWhenFired)
			weaponHeat += int(Weapons[item.ComponentDefID].HeatGenerated)
		}

		if item.ComponentDefType == "HeatSink" {
			if strings.Contains(item.ComponentDefID, "emod_engine_") {
				if strings.Contains(item.ComponentDefID, "emod_engine_cooling") {
					mech.HeatSink += int(EngineDefs[item.ComponentDefID].Custom.EngineHeatBlock.HeatSinkCount)
				} else {
					engineRating, err := strconv.Atoi(EngineDefs[item.ComponentDefID].Custom.EngineCore.Rating)
					if err != nil {
						log.Fatal(err)
					}
					movement.Rating = int64(engineRating)
				}
			}

			if strings.Contains(item.ComponentDefID, "emod_kit") {
				engineHeatSinkDC = int(GearDefs[EngineDefs[item.ComponentDefID].Custom.Cooling.HeatSinkDefID].DissipationCapacity)
			}

			if strings.Contains(item.ComponentDefID, "Gear_HeatSink_") {
				mech.HeatSink += int(GearDefs[item.ComponentDefID].DissipationCapacity)
			}
		}
	}

	mech.Heat = bList.ApplyBonus("selfHeat", bList.ApplyBonus("weaponHeat", weaponHeat))
	mech.HeatSink += int(movement.Rating/25) * engineHeatSinkDC
	mech.Distance.Walk = bList.ApplyBonus("walk", int(movement.CalcWalkDistance()))
	mech.Distance.Sprint = bList.ApplyBonus("run", int(movement.CalcSprintDistance()))
	mech.Hex.Walk = int(math.Round(float64(mech.Distance.Walk / 30)))
	mech.Hex.Sprint = int(math.Round(float64(mech.Distance.Sprint / 30)))
	mech.Distance.Jump = bList.ApplyBonus("jump", mech.JumpJets*30)
	mech.JumpHeat = int(math.Round(float64(mech.JumpJets)/heat.JumpHeatUnitSize*heat.JumpHeatPerUnit + float64(heat.JumpHeatMin)))
	mech.Shutdown = heat.MaxHeat
	mech.DFADamage = bList.ApplyBonus("dfa", int(chassisdef.DFADamage))
	mech.DFASelfDamage = bList.ApplyBonus("dfaSelf", int(chassisdef.DFASelfDamage))
	mech.Melee.Bonus = bList.ApplyBonus("melee", int(chassisdef.MeleeDamage))
	mech.Melee.Total = int(chassisdef.MeleeDamage) + mech.Melee.Bonus

	templateText, err := box.FindString("mech.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("test").Delims("{&", "&}").Parse(templateText)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, mech)
	if err != nil {
		log.Fatal(err)
	}
}
