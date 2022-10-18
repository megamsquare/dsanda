package config

import (
	"gorm.io/gorm"
	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	// dsn := "host=db user=dsanda password=postgrespass dbname=postgresdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "postgres://dsanda:postgrespass@localhost:5432/postgresdb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database! \n Error: " + err.Error())
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
