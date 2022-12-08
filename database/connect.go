package database

import (
	"log"
	"os"

	"github.com/fprasty/GoRestApiUDWijaya/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func KonekDB() {
	err := godotenv.Load("init.env")
	if err != nil {
		log.Fatal("KonekDB>Can't open init.env")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("KonekDB>Database can't loaded")
	}
	DB = database

	database.AutoMigrate(
		&models.User{},
	)
}
