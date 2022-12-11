package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Get app config from file.
func GetConfig(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
