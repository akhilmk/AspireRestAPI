package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

const CONFIG_FILE = "config/config.json"
const CONFIG_FILE_ERROR = "ERROR: config.json file not present in current directory"

var AppConfig config

func init() {
	loadAppConfig()
}

// Loading json config file.
func loadAppConfig() {

	if _, err := os.Stat(CONFIG_FILE); errors.Is(err, os.ErrNotExist) {
		log.Fatal(CONFIG_FILE_ERROR)
	}

	jsonData, _ := os.ReadFile(CONFIG_FILE)
	err := json.Unmarshal(jsonData, &AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}

// Struct to hold application configurations.
type config struct {
	DbInt  bool   `json:"db_init"`
	DbHost string `json:"db_host"`
	DbUser string `json:"db_user"`
	DbName string `json:"db_name"`
}
