package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host     string `json:"host"`
	RatePath string `json:"rate_path"`
}

func GetConfig(file string) *Config {
	var config Config
	filePath := os.Getenv("CONFIG")
	if filePath == "" {
		filePath = file
	}

	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return &config
}
