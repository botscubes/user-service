package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type DBCongig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	DBname   string `yaml:"dbname"`
}

type ServerConfig struct {
}

type Config struct {
	Redis RedisConfig `yaml:"redis"`
	DB    DBCongig    `yaml:"db"`
	ServerConfig
}

func GetConfig(fileName string) Config {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	config := Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return config
}
