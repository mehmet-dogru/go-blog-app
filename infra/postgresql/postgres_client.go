package postgresql

import (
	"go-blog-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectPostgres(config config.AppConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v/n", err)
	}
	log.Println("database connection success")
	return db, nil
}
