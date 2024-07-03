package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Get() *gorm.DB {
	return dbInstance
}

func init() {
	dsn := "host=localhost user=postgres password=Berat9730 dbname=go_templ_mia port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbInstance = db
}