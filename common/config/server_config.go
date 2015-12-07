package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ServerConfig struct {
	WithPortConfig
	WithRedisConfig
	Clients *ClientsConfig `json:"clients"`
}

type ClientsConfig struct {
	JSONApiUrl string `json:"json_api_url"`
}

func NewServerConfigFromJSONFile(configPath *string) *ServerConfig {
	configData, err := ioutil.ReadFile(*configPath)

	if err != nil {
		log.Fatalf("Could not open server config file. Path: '%v', Error: '%s", configPath, err)
	}

	sc := &ServerConfig{}

	if err := json.Unmarshal(configData, sc); err != nil {
		log.Fatalf("Server config is not JSON valid. %v", err)
	}

	return sc
}
