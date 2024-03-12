package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerPort string
	DSN        string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load()

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not loaded")
	}

	DSN := os.Getenv("DSN")
	if len(DSN) < 1 {
		return AppConfig{}, errors.New("env variables not loaded")
	}

	return AppConfig{
		ServerPort: httpPort,
		DSN:        DSN,
	}, nil
}
