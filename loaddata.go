package main

import (
	"bytes"
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

func loadData() {
	//"buildingdef",
	//files := getDefFiles("buildingdef")
	//fmt.Printf("Found '%d' BuildingDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := BuildingDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	BuildingDefs[def.Description.ID] = def
	//}
	//fmt.Printf("Loaded BuildingDefs: %d\n", len(BuildingDefs))

	//"chassisdef",
	files := getDefFiles("chassisdef")
	fmt.Printf("Found '%d' ChassisDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := ChassisDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		ChassisDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded ChassisDefs: %d\n", len(ChassisDefs))

	//"hardpointdatadef",
	//files = getDefFiles("hardpointdatadef")
	//fmt.Printf("Found '%d' HardPointDataDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := HardPointDataDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	HardPointDataDefs[def.ID] = def
	//}
	//fmt.Printf("Loaded HardPointDataDefs: %d\n", len(HardPointDataDefs))

	//"heraldrydef",
	//files = getDefFiles("heraldrydef")
	//fmt.Printf("Found '%d' HeraldryDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := HeraldryDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	HeraldryDefs[def.Description.ID] = def
	//}
	//fmt.Printf("Loaded HeraldryDefs: %d\n", len(HeraldryDefs))

	//"lancedef",
	//files = getDefFiles("lancedef")
	//fmt.Printf("Found '%d' LanceDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := LanceDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	LanceDefs[def.Description.Name] = def
	//}
	//fmt.Printf("Loaded LanceDefs: %d\n", len(LanceDefs))

	//"mechdef",
	files = getDefFiles("mechdef")
	fmt.Printf("Found '%d' MechDef files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := MechDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		MechDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded MechDefs: %d\n", len(MechDefs))

	//"movedef",
	//files = getDefFiles("movedef")
	//fmt.Printf("Found '%d' MoveDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := MoveDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	MoveDefs[def.Description.Name] = def
	//}
	//fmt.Printf("Loaded MoveDefs: %d\n", len(MoveDefs))

	//"pathingdef",
	//files = getDefFiles("pathingdef")
	//fmt.Printf("Found '%d' PathingDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := PathingDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	PathingDefs[def.Description.Name] = def
	//}
	//fmt.Printf("Loaded PathingDefs: %d\n", len(PathingDefs))

	//"quirks",
	files = getDefFiles("quirks")
	fmt.Printf("Found '%d' Quirk files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := Quirk{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		Quirks[def.Description.ID] = def
	}
	fmt.Printf("Loaded Quirks: %d\n", len(Quirks))

	//"shopdef",
	//files = getDefFiles("shopdef")
	//fmt.Printf("Found '%d' ShopDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := ShopDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	ShopDefs[def.ID] = def
	//}
	//fmt.Printf("Loaded ShopDefs: %d\n", len(ShopDefs))

	//"starsystemdef",
	//files = getDefFiles("starsystemdef")
	//fmt.Printf("Found '%d' StarSystemDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := StarSystemDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	StarSystemDefs[def.Description.Name] = def
	//}
	//fmt.Printf("Loaded StarSystemDefs: %d\n", len(StarSystemDefs))

	//"turretchassisdef",
	//files = getDefFiles("turretchassisdef")
	//fmt.Printf("Found '%d' TurretChassisDefs files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := TurretChassisDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	TurretChassisDefs[def.Description.Name] = def
	//}
	//fmt.Printf("Loaded TurretChassisDefs: %d\n", len(TurretChassisDefs))

	//"turretdef",
	//files = getDefFiles("turretdef")
	//fmt.Printf("Found '%d' TurretDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := TurretDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	TurretDefs[def.Description.Name] = def
	//}
	//fmt.Printf("Loaded TurretDefs: %d\n", len(TurretDefs))

	//"vehiclechassisdef",
	//files = getDefFiles("vehiclechassisdef")
	//fmt.Printf("Found '%d' VehicleChassisDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := VehicleChassisDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	VehicleChassisDefs[def.Description.Name] = def
	//}
	//fmt.Printf("Loaded VehicleChassisDefs: %d\n", len(VehicleChassisDefs))

	//"vehicledef",
	//files = getDefFiles("vehicledef")
	//fmt.Printf("Found '%d' VehicleDef files... ", len(files))
	//for f := range files {
	//	fp, err := os.Open(files[f])
	//	if err != nil {
	//		panic(err)
	//	}
	//	fileByte, _ := ioutil.ReadAll(fp)
	//	def := VehicleDef{}
	//	err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
	//	if err != nil {
	//		fmt.Println(files[f])
	//		fmt.Println(err)
	//	}
	//	fp.Close()
	//	VehicleDefs[def.Description.Name] = def
	//}
	//fmt.Printf("Loaded VehicleDefs: %d\n", len(VehicleDefs))

	//"weapons",
	files = getDefFiles("weapons")
	fmt.Printf("Found '%d' Weapon files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := Weapon{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		Weapons[def.Description.ID] = def
	}
	fmt.Printf("Loaded Weapons: %d\n", len(Weapons))

	// engines
	files = getDefFiles("emod")
	fmt.Printf("Found '%d' Engine files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := EngineDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		EngineDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded Engines: %d\n", len(EngineDefs))

	// HeatSinks
	files = getDefFiles("Gear_HeatSink")
	fmt.Printf("Found '%d' HeatSink files... ", len(files))
	for f := range files {
		fp, err := os.Open(files[f])
		if err != nil {
			panic(err)
		}
		fileByte, _ := ioutil.ReadAll(fp)
		def := HeatSinkDef{}
		err = json.Unmarshal(bytes.Trim(fileByte, "\xef\xbb\xbf"), &def)
		if err != nil {
			fmt.Println(files[f])
			fmt.Println(err)
		}
		fp.Close()
		HeatSinkDefs[def.Description.ID] = def
	}
	fmt.Printf("Loaded HeatSink: %d\n", len(HeatSinkDefs))
}
