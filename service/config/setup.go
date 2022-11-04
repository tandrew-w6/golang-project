package config

import (
	"gopkg.in/gcfg.v1"
)

func Setup() (*Config, error) {
	cfgFile := "config.toml"
	cfg := &Config{}

	err := gcfg.ReadFileInto(cfg, cfgFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type Config struct {
	Application ApplicationConfig
	PostgresDB  DBConfig
}

type ApplicationConfig struct {
	ServiceName string
	ServerPort  string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}
