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

func main() {
	loadChassisDefs()
	loadMechDefs()
	loadQuirkDefs()
	loadWeaponDefs()
	loadEngineDefs()
	loadGearDefs()

	//generateTestMech("mechdef_phoenixhawk_PXH-IIC")
	//generateTestMech("mechdef_catapult_CPLT-P")
	generateTestMech("mechdef_hatchetman_HCT-S7")
	//generateTestMech("mechdef_locust_LCT-2V")
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
	Stability float64
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
	mechdef := MechDefs[genmech]
	chassisdef := ChassisDefs[mechdef.ChassisID]
	mech := &Mech{
		MechDef:    mechdef,
		ChassisDef: chassisdef,
	}
	mech.CAPName = strings.ToUpper(chassisdef.PrefabBase)
	mech.StockRoles = strings.Replace(chassisdef.StockRole, " & ", "]] [[", -1)

	var movement = &Movement{Tonnage: chassisdef.Tonnage}

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

	bList := NewBonuses()
	var engineMultiplier = 1
	var engine int
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
			mech.Damage += int(Weapons[item.ComponentDefID].Damage)
			mech.Stability += Weapons[item.ComponentDefID].Instability
			weaponHeat += int(Weapons[item.ComponentDefID].HeatGenerated)
			mech.HeatDmg += int(Weapons[item.ComponentDefID].HeatDamage)
		}

		if item.ComponentDefType == "HeatSink" {
			if strings.Contains(item.ComponentDefID, "emod_engine_") {
				if strings.Contains(item.ComponentDefID, "emod_engine_cooling") {
					mech.HeatSink += int(EngineDefs[item.ComponentDefID].Custom.EngineHeatBlock.HeatSinkCount)
				} else {
					engineRating, err := strconv.Atoi(EngineDefs[item.ComponentDefID].Custom.EngineCore.Rating)
					movement.Rating = int64(engineRating)
					if err == nil {
						engine = int(movement.Rating/25) * 3
					} else {
						log.Fatal(err)
					}
				}
			}

			if strings.Contains(item.ComponentDefID, "emod_kit") {
				if strings.Contains(EngineDefs[item.ComponentDefID].Custom.Cooling.HeatSinkDefID, "Double") {
					engineMultiplier = 2
				}

			}

			if strings.Contains(item.ComponentDefID, "Gear_HeatSink_") {
				mech.HeatSink += int(GearDefs[item.ComponentDefID].DissipationCapacity)
			}
		}
	}

	//PrettyPrint(bList.ApplyBonus("dfa", 1))
	//PrettyPrint(bList.ApplyBonus("melee", 1))
	//PrettyPrint(bList.ApplyBonus("meleeStab", 1))
	//PrettyPrint(bList.ApplyBonus("targetHeat", 1))
	mech.Heat = bList.ApplyBonus("selfHeat", bList.ApplyBonus("weaponHeat", weaponHeat))
	mech.HeatSink += engineMultiplier * engine
	mech.Distance.Walk = bList.ApplyBonus("walk", int(movement.CalcWalkDistance()))
	mech.Distance.Sprint = bList.ApplyBonus("run", int(movement.CalcSprintDistance()))
	mech.Hex.Walk = int(math.Round(float64(mech.Distance.Walk / 30)))
	mech.Hex.Sprint = int(math.Round(float64(mech.Distance.Sprint / 30)))
	mech.Distance.Jump = bList.ApplyBonus("jump", mech.JumpJets*30)
	//	13/6,5 * 0,9 +3
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
