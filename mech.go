package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.4amlunch.net/RTWikiBot/defs"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var hexSize = 24

type Mech struct {
	ChassisDef defs.ChassisDef
	MechDef    defs.MechDef
	CCMod      defs.CCMod
	Constants  defs.GameConstants

	Movement struct {
		Rating   int
		Distance struct {
			Walk   int
			Sprint int
			Jump   int
		}
		Hex struct {
			Walk   int
			Sprint int
			Jump   int
		}
	}

	Heat struct {
		Self     int
		Sink     int
		Jump     int
		Shutdown int
	}

	QuirkText     string
	HardPoints    HardPoints
	Damage        int
	HeatDmg       int
	DFADamage     int
	DFASelfDamage int
	Melee         struct {
		Bonus     int
		Total     int
		Stability int
	}
	Stability  int
	Structure  int
	Armor      int
	JumpJets   int
	CAPName    string
	StockRoles string
	AmmoType   []string
}

func loadCCMod() defs.CCMod {
	file := "CustomComponents/mod.json"

	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	fileByte, _ := ioutil.ReadAll(fp)
	def := defs.CCMod{}

	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	if err != nil {
		fmt.Println(file)
		fmt.Println(err)
	}

	err = fp.Close()
	if err != nil {
		panic(err)
	}

	return def
}

func loadConstants() defs.GameConstants {
	file := "RogueTech Core/StreamingAssets/data/constants/CombatGameConstants.json"

	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	fileByte, _ := ioutil.ReadAll(fp)
	def := defs.GameConstants{}

	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	if err != nil {
		fmt.Println(file)
		fmt.Println(err)
	}

	err = fp.Close()
	if err != nil {
		panic(err)
	}

	return def
}

func NewMech(genmech string) Mech {
	mech := Mech{}
	mech.CCMod = loadCCMod()
	mech.Constants = loadConstants()
	mech.MechDef = MechDefs[genmech]
	mech.ChassisDef = ChassisDefs[mech.MechDef.ChassisID]
	mech.CAPName = strings.ToUpper(mech.ChassisDef.PrefabBase)
	mech.StockRoles = strings.Replace(mech.ChassisDef.StockRole, " & ", "]] [[", -1)

	// I hate having to do this twice but I really need Ammo before I do the rest
	for i := range mech.MechDef.Inventory {
		item := mech.MechDef.Inventory[i]

		if item.ComponentDefType == "AmmunitionBox" {
			mech.AmmoType = append(mech.AmmoType, item.ComponentDefID)
		}
	}

	feList := make(map[string]int)
	for e := range mech.ChassisDef.FixedEquipment {
		feList[mech.ChassisDef.FixedEquipment[e].ComponentDefID] += 1
		newItem := defs.MechDef_sub4{
			ComponentDefID:   mech.ChassisDef.FixedEquipment[e].ComponentDefID,
			ComponentDefType: mech.ChassisDef.FixedEquipment[e].ComponentDefType,
			DamageLevel:      mech.ChassisDef.FixedEquipment[e].DamageLevel,
			GUID:             mech.ChassisDef.FixedEquipment[e].GUID,
			HardpointSlot:    mech.ChassisDef.FixedEquipment[e].HardpointSlot,
			MountedLocation:  mech.ChassisDef.FixedEquipment[e].MountedLocation,
			SimGameUID:       mech.ChassisDef.FixedEquipment[e].SimGameUID,
			HasPrefabName:    mech.ChassisDef.FixedEquipment[e].HasPrefabName,
			PrefabName:       mech.ChassisDef.FixedEquipment[e].PrefabName,
		}
		mech.MechDef.Inventory = append(mech.MechDef.Inventory, newItem)
	}

	var qt strings.Builder
	for e := range feList {
		qt.WriteString("* [[")
		qt.WriteString(Quirks[e].Description.Name)
		if Quirks[e].ComponentType == "AmmunitionBox" {
			mech.AmmoType = append(mech.AmmoType, e)
		}
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

	for l := range mech.MechDef.Locations {
		mech.Armor += int(mech.MechDef.Locations[l].AssignedArmor)
	}

	hardPoints := map[string]int{
		"AntiPersonnel": 0,
		"Ballistic":     0,
		"Energy":        0,
		"Missile":       0,
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

	mech.addDefaults()

	heatSinkKitIDs := map[string]bool{
		"emod_kit_shs":            true,
		"emod_kit_cdhs":           true,
		"emod_kit_dhs_proto":      true,
		"emod_kit_chs":            true,
		"emod_kit_dhs":            true,
		"emod_kit_dhs_widowmaker": true,
		"emod_kit_dhs_royal":      true,
	}

	var (
		hsValue      int
		coolingValue int
	)
	for i := range mech.MechDef.Inventory {
		item := mech.MechDef.Inventory[i]

		if item.ComponentDefType == "HeatSink" {
			engineSlotsMatched, err := regexp.Match(`^emod_engineslots.*$`, []byte(item.ComponentDefID))
			if err != nil {
				log.Fatal(err)
			}

			engineMatched, err := regexp.Match(`^emod_engine_\d{3}$`, []byte(item.ComponentDefID))
			if err != nil {
				log.Fatal(err)
			}

			coolingMatched, err := regexp.Match(`^emod_engine_cooling_\d{1,2}$`, []byte(item.ComponentDefID))
			if err != nil {
				log.Fatal(err)
			}

			if engineSlotsMatched {
				continue
			} else if engineMatched {
				mech.Movement.Rating, err = strconv.Atoi(
					EngineDefs[item.ComponentDefID].Custom.EngineCore.Rating)
				if err != nil {
					log.Fatal(err)
				}
			} else if coolingMatched {
				coolingValue = int(EngineDefs[item.ComponentDefID].Custom.EngineHeatBlock.HeatSinkCount)
			} else if heatSinkKitIDs[item.ComponentDefID] {
				hs := GearDefs[EngineDefs[item.ComponentDefID].Custom.Cooling.HeatSinkDefID].Custom.BonusDescriptions.Bonuses
				for bonus := range hs {
					hptMatched, err := regexp.Match(`^HeatPerTurn: .*$`, []byte(hs[bonus]))
					if err != nil {
						log.Fatal(err)
					}
					if hptMatched {
						s := strings.Split(hs[bonus], ": -")
						fmt.Printf("HPT Matched: %v\n", s[1])
						hsValue, err = strconv.Atoi(s[1])
						if err != nil {
							log.Fatal(err)
						}
						mech.Heat.Sink += hsValue * 10
					}
				}
			} else {
				fmt.Printf("Not Engine: ")
				PrettyPrint(item)
				PrettyPrint(EngineDefs[item.ComponentDefID])
			}
		} else if item.ComponentDefType == "JumpJet" {
			jj := GearDefs[item.ComponentDefID]
			if jj.Custom.Category == nil {
				mech.JumpJets += 1
			}
		} else if item.ComponentDefType == "Weapon" {
			//fmt.Printf("Weapon: ")
			weapon := Weapons[item.ComponentDefID]
			//PrettyPrint(weapon)
			mech.Damage += int(weapon.Damage * weapon.ShotsWhenFired)
			mech.HeatDmg += int(weapon.HeatDamage)
			mech.Heat.Self += int(weapon.HeatGenerated)
			mech.Stability += int(weapon.Instability)
		}
	}
	mech.Heat.Sink += hsValue * coolingValue
	// TODO: Verify this is true
	mech.Heat.Jump = int(math.Round(float64(mech.JumpJets) * mech.Constants.Heat.JumpHeatPerUnit * float64(mech.Movement.Hex.Jump)))
	if mech.Heat.Jump < 3 {
		mech.Heat.Jump = 3
	}
	mech.Heat.Shutdown = int(mech.Constants.Heat.MaxHeat)

	mech.Movement.Distance.Walk = mech.CalcWalkDistance()
	mech.Movement.Distance.Sprint = mech.CalcSprintDistance()
	mech.Movement.Distance.Jump = mech.JumpJets * hexSize

	mech.Movement.Hex.Walk = mech.Movement.Distance.Walk / hexSize
	mech.Movement.Hex.Sprint = mech.Movement.Distance.Sprint / hexSize

	mech.getBonuses()

	mech.Melee.Total = int(mech.ChassisDef.MeleeDamage) + mech.Melee.Bonus
	return mech
}

func (m *Mech) addDefaults() {
	taggedDefaults := m.CCMod.Settings.TaggedDefaults
	defaults := m.CCMod.Settings.Defaults
	var item defs.GearDef
	var itemId string

cdOuter:
	for c := range m.ChassisDef.Custom.ChassisDefaults {
		for i := range m.MechDef.Inventory {
			itemId = m.MechDef.Inventory[i].ComponentDefID
			item = GearDefs[itemId]
			if len(item.Custom.Category) != 0 {
				if m.ChassisDef.Custom.ChassisDefaults[c].CategoryID == item.Custom.Category[c].CategoryID {
					continue cdOuter
				}
			}
		}
		m.addItem(itemId, item)
		itemId = ""
	}

tdOuter:
	for td := range taggedDefaults {
		if !taggedDefaults[td].AddIfNotPresent {
			continue tdOuter
		}

		for t := range m.ChassisDef.ChassisTags.Items {
			if taggedDefaults[td].Tag == m.ChassisDef.ChassisTags.Items[t] {
				for i := range m.MechDef.Inventory {
					itemId = m.MechDef.Inventory[i].ComponentDefID
					item = GearDefs[itemId]
					for cat := range item.Custom.Category {
						if item.Custom.Category[cat].CategoryID == taggedDefaults[td].CategoryID {
							if taggedDefaults[td].DefID == itemId {
								continue tdOuter
							}
						}
					}
				}
			}
		}
		m.addItem(itemId, item)
		itemId = ""
	}

dOuter:
	for d := range defaults {
		if !defaults[d].AddIfNotPresent {
			continue dOuter
		}

		for i := range m.MechDef.Inventory {
			itemId = m.MechDef.Inventory[i].ComponentDefID
			item = GearDefs[itemId]
			for cat := range item.Custom.Category {
				if item.Custom.Category[cat].CategoryID == defaults[d].CategoryID {
					if defaults[d].DefID == itemId {
						continue dOuter
					}
				}
			}
		}
		m.addItem(itemId, item)
		itemId = ""
	}

}

func (m *Mech) addItem(itemId string, item defs.GearDef) {
	if itemId == "" {
		return
	}

	newItem := defs.MechDef_sub4{}
	newItem.ComponentDefID = itemId
	newItem.ComponentDefType = item.ComponentType
	newItem.PrefabName = item.PrefabIdentifier

	m.MechDef.Inventory = append(m.MechDef.Inventory, newItem)
}

func (m *Mech) getMovementPoints() float64 {
	//fmt.Printf("getting MP: %v %v\n", m.Movement.Rating, m.ChassisDef.Tonnage)
	return float64(int64(m.Movement.Rating) / m.ChassisDef.Tonnage)
}

func RoundBy5(value float64) int64 {
	return int64(value) / 5 * 5
}

func (m *Mech) CalcWalkDistance() int {
	// numbers the result of the best fit line for the game movement
	var walkSpeedFixed = 26.05
	var walkSpeedMultiplier = 23.14

	return int(RoundBy5(walkSpeedFixed + m.getMovementPoints()*walkSpeedMultiplier))
}

func (m *Mech) CalcSprintDistance() int {
	// numbers the result of the best fit line for the game movement
	var runSpeedFixed = 52.43
	var runSpeedMultiplier = 37.29

	return int(RoundBy5(runSpeedFixed + m.getMovementPoints()*runSpeedMultiplier))
}

func (m *Mech) getBonuses() {
	bonuses := NewBonuses()
	for i := range m.MechDef.Inventory {
		comp := GearDefs[m.MechDef.Inventory[i].ComponentDefID]
		for b := range comp.Custom.BonusDescriptions.Bonuses {
			bonusType, bonus := bonuses.getBonus(comp.Custom.BonusDescriptions.Bonuses[b])
			//fmt.Printf("bonusType: %v, bonus: ", bonusType)
			//PrettyPrint(bonus)
			if bonusType == "run" {
				m.Movement.Distance.Sprint = bonuses.ApplyBonus(bonus, m.Movement.Distance.Sprint)
			} else if bonusType == "walk" {
				m.Movement.Distance.Walk = bonuses.ApplyBonus(bonus, m.Movement.Distance.Walk)
			} else if bonusType == "jump" {
				m.Movement.Distance.Jump = bonuses.ApplyBonus(bonus, m.Movement.Distance.Jump)
			} else if bonusType == "dfa" {
				m.DFADamage = bonuses.ApplyBonus(bonus, m.DFADamage)
			} else if bonusType == "dfaSelf" {
				m.DFASelfDamage = bonuses.ApplyBonus(bonus, m.DFASelfDamage)
			} else if bonusType == "melee" {
				m.Melee.Bonus = bonuses.ApplyBonus(bonus, m.Melee.Bonus)
			} else if bonusType == "meleeStab" {
				m.Melee.Stability = bonuses.ApplyBonus(bonus, m.Melee.Stability)
			} else if bonusType == "targetDamage" {
				m.Damage = bonuses.ApplyBonus(bonus, m.Damage)
			} else if bonusType == "targetHeat" {
				m.HeatDmg = bonuses.ApplyBonus(bonus, m.HeatDmg)
			} else if bonusType == "targetStab" {
				m.Stability = bonuses.ApplyBonus(bonus, m.Stability)
			} else if bonusType == "selfHeat" {
				m.Heat.Self = bonuses.ApplyBonus(bonus, m.Heat.Self)
			} else if bonusType == "weaponHeat" {
				m.Heat.Self = bonuses.ApplyBonus(bonus, m.Heat.Self)
			}
		}
	}
}

func (m Mech) generateWikiMarkdown() string {
	var tpl bytes.Buffer

	templateText, err := box.FindString("mech.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("test").Delims("{&", "&}").Parse(templateText)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(&tpl, m)
	if err != nil {
		log.Fatal(err)
	}

	return tpl.String()
}
