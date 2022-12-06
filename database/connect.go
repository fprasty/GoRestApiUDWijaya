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
	err := godotenv.Load("init.env") //Load file env
	if err != nil {
		log.Fatal("Konekdb > error load init.env")
	}
	dsn := os.Getenv("dsn") //Cek "dsn" from env file

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //Open Database from MySql
	if err != nil {
		panic("Konekdb > error open database from MySql")
	}

	DB = database
}
