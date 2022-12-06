package main

import (
	"log"
	"os"

	"github.com/fprasty/GoRestApiUDWijaya/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Konekdb()
	err := godotenv.Load("init.env") //Load file env
	if err != nil {
		log.Fatal("main > can't load init.env file")
	}

	port := os.Getenv("port") //Cek "port" from env file
	apk := fiber.New()
	apk.Listen(":" + port)

}
