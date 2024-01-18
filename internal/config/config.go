package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

// Get app config from file.
func GetConfig() (*Config, error) {
	var c Config
	if err := envconfig.Process(context.Background(), &c); err != nil {
		return nil, err
	}
	return &c, nil

	// data, err := os.ReadFile(fileName)
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//
	// config := Config{}
	// err = yaml.Unmarshal([]byte(data), &config)
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//
	// return &config, err
}
