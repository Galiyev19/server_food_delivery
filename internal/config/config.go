package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Address     string `json:"addr"`
	StorePath   string `json:"store_path"`
	StoreDriver string `json:"store_driver"`
}

func New() *Config {
	return &Config{}
}

func (cfg *Config) InitConfig(configPath string, config *Config) error {
	configJson, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	defer configJson.Close()
	body, err := ioutil.ReadAll(configJson)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	if err := json.Unmarshal(body, cfg); err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}
	return nil
}
