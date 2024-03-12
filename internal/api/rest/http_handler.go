package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go-blog-app/internal/helper"
	"gorm.io/gorm"
)

type RestHandler struct {
	App   *fiber.App
	DB    *gorm.DB
	Auth  helper.Auth
	Redis *redis.Client
}
