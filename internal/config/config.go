package config

import (
	"github.com/vakhia/artilight/pkg/env"
)

type Config struct {
	Environment string
	Port        string
	Database    *Database
	Token       *Token
}

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

type Token struct {
	Secret string
	Issuer string
}

func NewConfig() (*Config, error) {
	env.CheckDotEnv()
	port := env.MustGet("SERVER_PORT")

	if port == "" {
		port = "3000"
	}

	return &Config{
		Environment: env.MustGet("ENV"),
		Port:        port,
		Database: &Database{
			Host:     env.MustGet("DATABASE_HOST"),
			Port:     env.MustGet("DATABASE_PORT"),
			User:     env.MustGet("DATABASE_USER"),
			DB:       env.MustGet("DATABASE_NAME"),
			Password: env.MustGet("DATABASE_PASSWORD"),
		},
		Token: &Token{
			Secret: env.MustGet("TOKEN_SECRET"),
			Issuer: env.MustGet("TOKEN_ISSUER"),
		},
	}, nil
}
