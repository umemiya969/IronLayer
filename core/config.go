package core

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Listen  string `yaml:"listen"`
		Backend string `yaml:"backend"`
	} `yaml:"server"`

	Rules struct {
		RateLimit map[string]string `yaml:"rate_limit"`
		BlockPath []string          `yaml:"block_paths"`
	} `yaml:"rules"`
}

func LoadConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
