package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=dsanda password=postgrespass dbname=postgresdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database! \n Error: " + err.Error())
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}