package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func LoadENV(envFile string) error {
	if err := godotenv.Load(envFile); err != nil {
		return err
	}

	return nil
}

func LoadAllconfig() {
	if err := LoadENV(".env"); err != nil {
		log.Warn("Failed to load .env file, falling back to OS environment variables")
	}

	if os.Getenv("APP_HOST") == "" {
		log.Error("Environment variable APP_HOST is not set.")
		panic("Environment variable APP_HOST is not set.")
	}

	LoadAppConfig()
	LoadDBConfig()

	log.Info("Successfully loaded configurations.")
}
