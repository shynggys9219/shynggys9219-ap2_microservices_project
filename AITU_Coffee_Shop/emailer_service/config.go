package main

import "github.com/caarlos0/env/v10"

type Config struct {
	SendGridMail struct {
		APIKey string `env:"SEND_GRID_API_KEY,notEmpty"`
	}
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
