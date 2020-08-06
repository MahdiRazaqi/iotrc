package config

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Mongo struct {
		Host     string `json:"host"`
		DB       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mongo"`
}

// Config data
var Config config

// Load config file
func Load() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &Config); err != nil {
		return err
	}
	return nil
}
