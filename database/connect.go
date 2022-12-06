package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Konekdb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Tidak bisa load file env")
	}
	dsn := os.Getenv("dsn")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Tidak bisa terkoneksi ke database")
	}
	DB = database
}
