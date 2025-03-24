package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
    ReposPath string `yaml:"repos_path"`
    Port      string `yaml:"port"`
}

func LoadConfig(path string) (Config, error) {
	// Default configuration
	config := Config{
		Port:      "8080",
		ReposPath: "./repos",
	}

	// If config file exists, override defaults
	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		if err != nil {
			return config, err
		}

		if err := yaml.Unmarshal(data, &config); err != nil {
			return config, err
		}
	} else {
		return config, errors.New("config file not found, using defaults")
	}

	return config, nil
}