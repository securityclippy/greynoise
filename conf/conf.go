package conf

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	APIKey string `json:"api_key"`
	Endpoint string  `json:"endpoint"`
}

func ReadConfig(configPath string) (Config, error) {
	config := Config{}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}