package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(configFile string) (Config, error) {
	log.Println("loading config file: ", configFile)

	var config = Config{}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	if len(config.Issues) == 0 {
		return config, nil
	}

	log.Println("loaded config file: ", &config)
	return config, nil
}
