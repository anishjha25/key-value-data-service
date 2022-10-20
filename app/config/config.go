package config

import (
	"encoding/json"
	"log"
	"os"
)

var AppConfig Config

type Config struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	StorageHost string `json:"storageHost"`
}

func Load() {
	file, err := os.ReadFile("./config/config.json")
	if err != nil {
		log.Fatal("failed to open config file")
	}

	err = json.Unmarshal(file, &AppConfig)
	if err != nil {
		log.Fatal("failed to unmarshal config")
	}
}
