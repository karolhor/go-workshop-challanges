package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/karolhor/go-workshops-challange/common/config"
)

type LoggerConfig struct {
	config.WithRedisConfig
	LogFilePath string `json:"log_file_path"`
}

func NewLoggerConfigFromJSONFile(configPath *string) *LoggerConfig {
	configData, err := ioutil.ReadFile(*configPath)

	if err != nil {
		log.Fatalf("Could not open logger config file. Path: '%v', Error: '%s", configPath, err)
	}

	config := &LoggerConfig{}

	if err := json.Unmarshal(configData, config); err != nil {
		log.Fatalf("Logger config is not JSON valid. %v", err)
	}

	return config
}

func NewJsonApiConfigFromJSONFile(configPath *string) *config.WithPortConfig {
	configData, err := ioutil.ReadFile(*configPath)

	if err != nil {
		log.Fatalf("Could not open json api config file. Path: '%v', Error: '%s", configPath, err)
	}

	config := &config.WithPortConfig{}

	if err := json.Unmarshal(configData, config); err != nil {
		log.Fatalf("Json API config is not JSON valid. %v", err)
	}

	return config
}
