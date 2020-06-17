package main

import (
	"github.com/ryanrolds/club/pkg/signaling"

	"github.com/ilyakaznacheev/cleanenv"
)

type EnvironmentValue string

const (
	EnvironmentTests      = EnvironmentValue("tests")
	EnvironmentProduction = EnvironmentValue("production")
)

type Config struct {
	Environment       EnvironmentValue `yaml:"env" env:"ENV" env-default:"local"`
	DefaultGroupLimit int              `yaml:"default_group_limit" env:"DEFAULT_GROUP_LIMIT" env-default:"12"`
	Groups            []GroupConfig    `yaml:"groups"`
	Port              int              `yaml:"port" env:"PORT" env-default:"3001"`
}

type GroupConfig struct {
	ID    signaling.NodeID `yaml:"id"`
	Limit int              `yaml:"limit"`
}

func GetConfig(filename string) (*Config, error) {
	config := &Config{}
	err := cleanenv.ReadConfig(filename, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
