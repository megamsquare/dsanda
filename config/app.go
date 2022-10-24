package config

import (
	"gorm.io/gorm"
	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"github.com/megamsquare/dsanda/db"
)

var DB *gorm.DB

func ConnectDB() {
	// dsn := "host=db user=dsanda password=postgrespass dbname=postgresdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "postgres://dsanda:postgrespass@db:5432/postgresdb"
	dbase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database! \n Error: " + err.Error())
	}
	DB = dbase

	
	// Migrate the schema
	// golang migrate working with gorm
	db.Migrate(dsn)

}

func GetDB() *gorm.DB {
	return DB
}
