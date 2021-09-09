package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var CurrentPath string
var SetupParams SetupParameters

type SetupParameters struct {
	Port          int
	CacheCapacity int
}

func init() {
	setCurrentPath()
	loadParams()
}

func setCurrentPath() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	CurrentPath = filepath.Dir(ex)
	log.Println("Current directory path set as [" + CurrentPath + "]")
}

func loadParams() {
	params, _ := ioutil.ReadFile(CurrentPath + "/config/config.json")
	log.Println("Log file path")
	SetupParams = SetupParameters{}
	_ = json.Unmarshal([]byte(params), &SetupParams)
	log.Println("Config parameters loaded...")
}
