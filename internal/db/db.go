package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB = nil

func Connect() {
	connectionString := "postgres://postgres:1234@localhost:5432/test"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Connection = db
}
