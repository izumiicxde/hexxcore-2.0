package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL         string `env:"DATABASE_URL"`
	PORT           string `env:"PORT"`
	API_ENDPOINT   string `env:"API_ENDPOINT"`
	JWT_SECRET     string `env:"JWT_SECRET"`
	RESEND_API_KEY string `env:"RESEND_API_KEY"`
	URLS           string `env:"FRONTEND"`
}

var Envs Config
var Validator = validator.New(validator.WithRequiredStructEnabled())

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	if err := env.Parse(&Envs); err != nil {
		panic(err)
	}
}
