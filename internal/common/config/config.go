package config

import "github.com/vakhia/artilight/pkg/env"

type Config struct {
	Environment string
	Port        string
	Database    *Database
	Token       *Token
	GCS         *GCS
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

type GCS struct {
	Bucket          string // Google Cloud Storage bucket name
	ProjectId       string
	CredentialsPath string // Path to the Google Cloud credentials file
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
		GCS: &GCS{
			ProjectId:       env.MustGet("GCS_PROJECT_ID"),
			Bucket:          env.MustGet("GCS_BUCKET"),
			CredentialsPath: env.MustGet("GOOGLE_APPLICATION_CREDENTIALS"),
		},
	}, nil
}
