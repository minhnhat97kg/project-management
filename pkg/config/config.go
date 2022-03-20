package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Port int `yaml:"port"`
}

func ReadConfig(path string) (Config, error) {
	var config Config
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err2 := yaml.Unmarshal(yamlFile, &config)
	if err2 != nil {
		return config, err2
	}
	return config, nil
}
