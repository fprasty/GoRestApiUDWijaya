package routes

import (
	"github.com/fprasty/GoRestApiUDWijaya/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
