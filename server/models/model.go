package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	dsn := "host=172.17.0.1 user=admin password=admin dbname=admin port=5432 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Job{})
	if err != nil {
		panic(err)
	}
}
