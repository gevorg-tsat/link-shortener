package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		GRCP `yaml:"grcp"`
		DB   `yaml:"db"`
	}
	GRCP struct {
		Host string `env-required:"true" yaml:"host" env:"GRCP_HOST"`
		Port int    `env-required:"true" yaml:"port" env:"GRCP_PORT"`
	}
	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port int    `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}
	DB struct {
		Port     int    `yaml:"port" env:"PORT" env-default:"5432"`
		Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
		Name     string `yaml:"name" env:"NAME" env-default:"postgres"`
		User     string `yaml:"user" env:"USER" env-default:"postgres"`
		Password string `yaml:"password" env:"PASSWORD" env-default:"postgres"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
