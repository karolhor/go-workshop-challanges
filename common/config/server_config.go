package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type RedisConfig struct {
	Address       string `json:"address"`
	PubSubChannel string `json:"pub_sub_channel"`
}

type ServerConfig struct {
	WithPortConfig
	RedisConfig *RedisConfig   `json:"redis"`
	Clients     *ClientsConfig `json:"clients"`
}

type ClientsConfig struct {
	JSONApiUrl string `json:"json_api_url"`
}

func NewServerConfigFromJSONFile(configPath *string) *ServerConfig {
	configData, err := ioutil.ReadFile(*configPath)

	if err != nil {
		log.Fatalf("Could not open server config file. Path: '%s', Error: '%s", configPath, err)
	}

	sc := &ServerConfig{}

	if err := json.Unmarshal(configData, sc); err != nil {
		log.Fatalf("Server config is not JSON valid. %v", err)
	}

	return sc
}

func NewJsonApiConfigFromJSONFile(configPath *string) *WithPortConfig {
	configData, err := ioutil.ReadFile(*configPath)

	if err != nil {
		log.Fatalf("Could not open json api config file. Path: '%s', Error: '%s", configPath, err)
	}

	config := &WithPortConfig{}

	if err := json.Unmarshal(configData, config); err != nil {
		log.Fatalf("Json API config is not JSON valid. %v", err)
	}

	return config
}
