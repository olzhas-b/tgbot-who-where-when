package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config Config

type Config struct {
	HTTP struct {
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"HTTP"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		User     string `yaml:"userName"`
		Password string `yaml:"password"`
	} `yaml:"DATABASE"`
	Telegram struct {
		Token string `yaml:"token"`
		Debug bool   `yaml:"debug"`
	} `yaml:"TELEGRAM"`
}

func InitConfig(fileName string) (Config, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Config{}, fmt.Errorf("Couldn't read config file got error %v", err)
	}

	if err = yaml.Unmarshal(file, &config); err != nil {
		return Config{}, fmt.Errorf("Cound't parse config file got error %v", err)
	}

	return config, nil
}

func GetConfig() Config {
	return config
}
