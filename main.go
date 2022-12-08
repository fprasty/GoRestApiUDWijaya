package main

import (
	"log"
	"os"

	"github.com/fprasty/GoRestApiUDWijaya/database"
	"github.com/fprasty/GoRestApiUDWijaya/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.KonekDB()
	err := godotenv.Load("init.env")
	if err != nil {
		log.Fatal("main>Can't load init.env")
	}
	port := os.Getenv("port")
	apk := fiber.New()
	routes.Setup(apk)
	apk.Listen(":" + port)
}
