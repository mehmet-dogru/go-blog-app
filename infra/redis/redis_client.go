package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-blog-app/config"
	"log"
)

func ConnectRedis(config config.AppConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis connection error: %v", err)
	} else {
		log.Println("redis connection success")
	}

	return client
}
