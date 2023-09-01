package config

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig(configFile string) (*Config, error) {
	config := &Config{}

	// Find the full path to the config file
	_, currentFilePath, _, _ := runtime.Caller(0)
	configFilePath := filepath.Join(filepath.Dir(currentFilePath), configFile)

	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
