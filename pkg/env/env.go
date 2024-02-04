package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	dev        = "dev"
	production = "production"
	env        = "env"
)

func MustGet(key string) string {
	val := os.Getenv(key)
	if val == "" && key != "PORT" {
		panic("Env key missing " + key)
	}
	return val
}

func CheckDotEnv() {
	err := godotenv.Load()
	if err != nil && os.Getenv(env) == dev {
		log.Println("Error loading .env file")
	}
}
