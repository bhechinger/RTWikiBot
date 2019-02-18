package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.4amlunch.net/RTWikiBot/defs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func getDefFiles(defType string, quirk bool) []string {
	var regexpstr string
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if quirk {
		regexpstr = fmt.Sprintf("%s.*json$", defType)
	} else {
		regexpstr = fmt.Sprintf("%s[^/]*json$", defType)
	}
	var files []string
	err = filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(regexpstr, path)
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

func loadSettings() defs.Settings {
	fp, err := os.Open("../RTWikiBot/json/Settings.json")
	if err != nil {
		panic(err)
	}
	fileByte, _ := ioutil.ReadAll(fp)
	def := defs.Settings{}
	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	if err != nil {
		fmt.Println("json/Settings.json")
		fmt.Println(err)
	}
	fp.Close()

	return def
}

func loadBuildingDefs() {
	files := getDefFiles("buildingdef", false)
	fmt.Printf("Found '%d' BuildingDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.BuildingDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		BuildingDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded BuildingDefs: %d\n", len(BuildingDefs))
}

func loadChassisDefs() {
	files := getDefFiles("chassisdef", false)
	fmt.Printf("Found '%d' ChassisDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.ChassisDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		ChassisDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded ChassisDefs: %d\n", len(ChassisDefs))
}

func loadHardPointDataDefs() {
	files := getDefFiles("hardpointdatadef", false)
	fmt.Printf("Found '%d' HardPointDataDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.HardPointDataDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		HardPointDataDefs[def.ID] = def
	}
	fmt.Printf("Loaded HardPointDataDefs: %d\n", len(HardPointDataDefs))
}

func loadHeraldryDefs() {
	files := getDefFiles("heraldrydef", false)
	fmt.Printf("Found '%d' HeraldryDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.HeraldryDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		HeraldryDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded HeraldryDefs: %d\n", len(HeraldryDefs))
}

func loadLanceDefs() {
	files := getDefFiles("lancedef", false)
	fmt.Printf("Found '%d' LanceDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.LanceDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		LanceDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded LanceDefs: %d\n", len(LanceDefs))
}

func loadMechDefs() {
	files := getDefFiles("mechdef", false)
	fmt.Printf("Found '%d' MechDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.MechDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		MechDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded MechDefs: %d\n", len(MechDefs))
}

func loadMoveDefs() {
	files := getDefFiles("movedef", false)
	fmt.Printf("Found '%d' MoveDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.MoveDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		MoveDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded MoveDefs: %d\n", len(MoveDefs))
}

func loadPathingDefs() {
	files := getDefFiles("pathingdef", false)
	fmt.Printf("Found '%d' PathingDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.PathingDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		PathingDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded PathingDefs: %d\n", len(PathingDefs))
}

func loadQuirkDefs() {
	files := getDefFiles("quirks", true)
	fmt.Printf("Found '%d' Quirk files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.Quirk{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		Quirks[def.Description.ID] = def
	}
	fmt.Printf("Loaded Quirks: %d\n", len(Quirks))
}

func loadShopDefs() {
	files := getDefFiles("shopdef", false)
	fmt.Printf("Found '%d' ShopDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.ShopDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		ShopDefs[def.ID] = def
	}
	fmt.Printf("Loaded ShopDefs: %d\n", len(ShopDefs))
}

func loadStarSystemDefs() {
	files := getDefFiles("starsystemdef", false)
	fmt.Printf("Found '%d' StarSystemDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.StarSystemDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		StarSystemDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded StarSystemDefs: %d\n", len(StarSystemDefs))
}

func loadTurretChassisDefs() {
	files := getDefFiles("turretchassisdef", false)
	fmt.Printf("Found '%d' TurretChassisDefs files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.TurretChassisDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		TurretChassisDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded TurretChassisDefs: %d\n", len(TurretChassisDefs))
}

func loadTurretDefs() {
	files := getDefFiles("turretdef", false)
	fmt.Printf("Found '%d' TurretDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.TurretDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		TurretDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded TurretDefs: %d\n", len(TurretDefs))
}

func loadVehicleChassisDefs() {
	files := getDefFiles("vehiclechassisdef", false)
	fmt.Printf("Found '%d' VehicleChassisDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.VehicleChassisDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		VehicleChassisDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded VehicleChassisDefs: %d\n", len(VehicleChassisDefs))
}

func loadVehicleDefs() {
	files := getDefFiles("vehicledef", false)
	fmt.Printf("Found '%d' VehicleDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.VehicleDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		VehicleDefs[def.Description.Name] = def
	}
	fmt.Printf("Loaded VehicleDefs: %d\n", len(VehicleDefs))
}

func loadWeaponDefs() {
	files := getDefFiles("Weapon_", false)
	fmt.Printf("Found '%d' Weapon files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.Weapon{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		Weapons[def.Description.ID] = def
	}
	fmt.Printf("Loaded Weapons: %d\n", len(Weapons))
}

func loadAmmoDefs() {
	files := getDefFiles("Ammo_", false)
	fmt.Printf("Found '%d' Ammo files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.AmmoDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		AmmoDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded Ammo: %d\n", len(AmmoDefs))
}

func loadEngineDefs() {
	files := getDefFiles("emod", false)
	fmt.Printf("Found '%d' Engine files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.EngineDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		EngineDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded Engines: %d\n", len(EngineDefs))
}

func loadGearDefs() {
	files := getDefFiles("Gear_", false)
	fmt.Printf("Found '%d' Gear files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := defs.GearDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		GearDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded Gear: %d\n", len(GearDefs))
}
