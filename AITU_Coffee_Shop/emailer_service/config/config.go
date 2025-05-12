package config

import (
	"github.com/caarlos0/env/v10"
)

type (
	Config struct {
		MailerKey string `env:"MAILER_API_KEY,notEmpty"`
		Nats      Nats

		Version string `env:"VERSION"`
	}

	// Nats configuration for main application
	Nats struct {
		Hosts        []string `env:"NATS_HOSTS,notEmpty" envSeparator:","`
		NKey         string   `env:"NATS_NKEY,notEmpty"`
		IsTest       bool     `env:"NATS_IS_TEST,notEmpty" envDefault:"true"`
		NatsSubjects NatsSubjects
	}

	// NatsSubjects for main application
	NatsSubjects struct {
		CustomerEventSubject string `env:"NATS_CUSTOMER_EVENT_SUBJECT,notEmpty"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
