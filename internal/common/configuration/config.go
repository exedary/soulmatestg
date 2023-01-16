package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfiguration struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}

type ConfigurationBase struct {
	Port string                 `default:"5000" yaml:"port"`
	Db   *DatabaseConfiguration `yaml:"db"`
}

func NewConfig(configPath string) (*ConfigurationBase, error) {
	config := &ConfigurationBase{}

	configurationFile, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}

	defer configurationFile.Close()

	yamlDecoder := yaml.NewDecoder(configurationFile)

	if err := yamlDecoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
