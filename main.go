package main

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/packr"
	"sync"
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
	var wg sync.WaitGroup

	// Load all the definition data
	wg.Add(7)
	go loadChassisDefs(&wg)
	go loadMechDefs(&wg)
	go loadQuirkDefs(&wg)
	go loadWeaponDefs(&wg)
	go loadAmmoDefs(&wg)
	go loadEngineDefs(&wg)
	go loadGearDefs(&wg)
	wg.Wait()

	//generateTestMech("mechdef_phoenixhawk_PXH-IIC")
	//generateTestMech("mechdef_catapult_CPLT-P")
	generateTestMech("mechdef_hatchetman_HCT-S7")
	//generateTestMech("mechdef_locust_LCT-2V")
	//generateTestMech("mechdef_hellspawn_HSN-8E")
	//generateTestMech("mechdef_osiris_OSR-3D")
	//generateTestMech("mechdef_raven_RVN-3L")
	//generateTestMech("mechdef_vindicator_VND-1R")
	//generateTestMech("mechdef_hunchback_HBK-4N")
	//generateTestMech("mechdef_javelin_JVN-10P")
	//generateTestMech("mechdef_centurion_CN9-YLW")
	//generateTestMech("mechdef_GyrFalcon_1")
}

type HardPoints struct {
	AntiPersonnel int
	Ballistic     int
	Energy        int
	Missile       int
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
	mech := NewMech(genmech)
	markdown := mech.generateWikiMarkdown()
	fmt.Println(markdown)
	//PrettyPrint(mech)
}
