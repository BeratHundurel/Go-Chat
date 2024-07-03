package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

func Get() *gorm.DB {
	return dbInstance
}

func init() {
	dsn := "host=localhost user=postgres password=Berat9730 dbname=go_chat port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	dbInstance = db
}
