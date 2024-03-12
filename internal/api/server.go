package api

import (
	"github.com/gofiber/fiber/v2"
	"go-blog-app/config"
	"go-blog-app/internal/api/rest"
	"go-blog-app/internal/api/rest/handlers"
	"go-blog-app/internal/domain"
	"go-blog-app/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	} else {
		log.Println("database connection success ✅")
	}

	db.AutoMigrate(&domain.User{})

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
	}

	setupRoutes(rh)

	log.Fatal(app.Listen(config.ServerPort))
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
