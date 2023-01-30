package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(loadConfig),
)

type Configuration struct {
	Server struct {
		Port string `yaml:"port" env:"server_port"`
	} `yaml:"server"`
	Database struct {
		Uri      string `yaml:"uri" env:"mongo_uri"`
		Username string `yaml:"user" env:"mongo_user"`
		Password string `yaml:"pass" env:"mongo_pass"`
	} `yaml:"database"`
}

func loadConfig() *Configuration {
	var conf Configuration
	cleanenv.ReadConfig("..\\configs\\soulmates.yaml", &conf)

	return &conf
}
