package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerPort    string
	DSN           string
	AppSecret     string
	RedisAddr     string
	RedisPassword string
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

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env variables not loaded")
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if len(redisAddr) < 1 {
		return AppConfig{}, errors.New("redis address env variables not loaded")
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if len(redisPassword) < 0 {
		return AppConfig{}, errors.New("redis pass env variables not loaded")
	}

	return AppConfig{
		ServerPort:    httpPort,
		DSN:           DSN,
		AppSecret:     appSecret,
		RedisAddr:     redisAddr,
		RedisPassword: redisPassword,
	}, nil
}
