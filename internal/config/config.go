package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(loadConfigFromFile),
	fx.Provide(loadAuthConfigFromEnv),
)

type Oauth struct {
	ClientID     string   `env:"g_services_client_id"`
	ClientSecret string   `env:"g_services_client_secret"`
	RedirectURL  string   `yaml:"redirect"`
	Scopes       []string `yaml:"scopes"`
}

type Configuration struct {
	Server struct {
		Port string `yaml:"port" env:"server_port"`
	} `yaml:"server"`
	Database struct {
		Uri      string `yaml:"uri" env:"mongo_uri"`
		Username string `yaml:"user" env:"mongo_user"`
		Password string `yaml:"pass" env:"mongo_pass"`
	} `yaml:"database"`
	GoogleOauth Oauth `yaml:"oauth"`
}

func loadConfigFromFile() *Configuration {
	var conf Configuration
	cleanenv.ReadConfig("..\\configs\\soulmates.yaml", &conf)

	return &conf
}

func loadAuthConfigFromEnv() *Oauth {
	var authConfig Oauth
	cleanenv.ReadEnv(authConfig)

	return &authConfig
}
