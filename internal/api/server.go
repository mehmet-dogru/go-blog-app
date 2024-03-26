package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"go-blog-app/config"
	_ "go-blog-app/docs"
	"go-blog-app/infra/postgresql"
	"go-blog-app/infra/redis"
	"go-blog-app/internal/api/rest"
	"go-blog-app/internal/api/rest/handlers"
	"go-blog-app/internal/domain"
	"go-blog-app/internal/helper"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := postgresql.ConnectPostgres(config)

	err = db.AutoMigrate(&domain.User{}, &domain.Article{})
	if err != nil {
		return
	}

	rdsDB := redis.ConnectRedis(config)

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:   app,
		DB:    db,
		Auth:  auth,
		Redis: rdsDB,
	}

	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault)
	setupRoutes(rh)

	log.Fatal(app.Listen(config.ServerPort))
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
	handlers.SetupArticleRoutes(rh)
}
