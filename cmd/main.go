package main

import (
	"go-blog-app/config"
	"go-blog-app/internal/api"
	"log"
)

// @title Go Blog API
// @version 1.0
// @description This is a blog api
// @contact.email mdogru685@gmail.com
// @BasePath /
func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}
	api.StartServer(cfg)
}
