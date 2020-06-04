package main

import (
	"time"

	"github.com/ryanrolds/club/pkg/signaling"

	"github.com/ilyakaznacheev/cleanenv"
)

type EnvironmentValue string

const (
	EnvironmentLocal      = EnvironmentValue("local")
	EnvironmentTest       = EnvironmentValue("test")
	EnvironmentStaging    = EnvironmentValue("staging")
	EnvironmentProduction = EnvironmentValue("production")
)

type GroupConfig struct {
	ID    signaling.GroupID `yaml:"id"`
	Limit int               `yaml:"limit"`
}

type Config struct {
	Environment       EnvironmentValue `yaml:"env" env:"ENV" env-default:"local"`
	DefaultGroupLimit int              `yaml:"default_group_limit" env:"DEFAULT_GROUP_LIMIT" env-default:"12"`
	Groups            []GroupConfig    `yaml:"groups"`
	ReaperInterval    time.Duration    `yaml:"reaper_interval" env:"REAPER_INTERVAL" env-default:"15s"`
	Port              int              `yaml:"port" env:"PORT" env-default:"3001"`
}

func GetConfig(filename string) (*Config, error) {
	config := &Config{}
	err := cleanenv.ReadConfig(filename, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
