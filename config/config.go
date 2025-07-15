package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	SMTPServer   string `env:"SMTP_SERVER"`
	SMTPUser     string `env:"SMTP_USER"`
	SMTPPassword string `env:"SMTP_PASSWORD"`
	JWTSecret string `env:JWT_SECRET`
	APIPort string `env:"API_PORT"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
		return nil, err
	}
	return &cfg, nil
}