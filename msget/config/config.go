package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	HostAddr string
	ConnStr  string
}

func Init() (*Config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
