package config

import (
	"errors"
	"log"
	"os"

	"github.com/akhilmk/GoRESTAPI/model"
	"github.com/akhilmk/GoRESTAPI/util"
)

const CONFIG_FILE = "config/config.json"
const CONFIG_FILE_ERROR = "ERROR: config.json file not present in current directory"

var AppConfig model.Config

func init() {
	loadAppConfig()
}

// Loading json config file.
func loadAppConfig() {

	if _, err := os.Stat(CONFIG_FILE); errors.Is(err, os.ErrNotExist) {
		log.Fatal(CONFIG_FILE_ERROR)
	}

	jsonData, _ := os.ReadFile(CONFIG_FILE)
	err := util.ByteToStruct(jsonData, &AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
