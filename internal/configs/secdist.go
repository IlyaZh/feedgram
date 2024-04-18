package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type SecDist struct {
	Telegram struct {
		Token string `yaml:"token"`
	} `yaml:"telegram""`
	Mysql struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
}

func NewSecDist(path string) SecDist {
	var s SecDist
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &s)
	if err != nil {
		log.Fatalf("Error while loading secdist file: %s", path)
	}
	return s
}
