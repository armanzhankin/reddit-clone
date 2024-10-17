package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUrl string
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	return Config{
		DbUrl: os.Getenv("DB_URL"),
	}, nil
}
