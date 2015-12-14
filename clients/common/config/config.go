package config

import (
	"encoding/json"
	"github.com/karolhor/go-workshops-challange/common/config"
	"io/ioutil"
	"log"
)

type (
	LoggerConfig struct {
		config.WithRedisConfig
		LogFilePath string `json:"log_file_path"`
	}

	mongoDBConfig struct {
		URL string `json:"url"`
		DbName string `json:"db_name"`
	}

	MongoConfig struct {
		config.WithRedisConfig
		MongoDBConfig *mongoDBConfig `json:"mongo"`
	}

	EventStreamConfig struct {
		config.WithRedisConfig
		config.WithPortConfig

		StaticPath string `json:"static_path"`
	}
)

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

func NewMongoConfigFromJSONFile(configPath *string) *MongoConfig {
	configData, err := ioutil.ReadFile(*configPath)

	if err != nil {
		log.Fatalf("Could not open mongo config file. Path: '%v', Error: '%s", configPath, err)
	}

	config := &MongoConfig{}

	if err := json.Unmarshal(configData, config); err != nil {
		log.Fatalf("Mongo config is not JSON valid. %v", err)
	}

	return config
}

func NewEventStreamConfigFromJSONFile(configPath *string) *EventStreamConfig {
	configData, err := ioutil.ReadFile(*configPath)

	if err != nil {
		log.Fatalf("Could not open event stream config file. Path: '%v', Error: '%s", configPath, err)
	}

	config := &EventStreamConfig{}

	if err := json.Unmarshal(configData, config); err != nil {
		log.Fatalf("Event stream config is not JSON valid. %v", err)
	}

	return config
}