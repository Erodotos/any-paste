package routes

import (
	"frontend/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/post/:id", handlers.ReadPost)
}