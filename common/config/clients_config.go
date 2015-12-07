package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type LoggerConfig struct {
	WithRedisConfig
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

func NewJsonApiConfigFromJSONFile(configPath *string) *WithPortConfig {
	configData, err := ioutil.ReadFile(*configPath)

	if err != nil {
		log.Fatalf("Could not open json api config file. Path: '%v', Error: '%s", configPath, err)
	}

	config := &WithPortConfig{}

	if err := json.Unmarshal(configData, config); err != nil {
		log.Fatalf("Json API config is not JSON valid. %v", err)
	}

	return config
}
