package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Server().MaxConnsPerIP = 10

	routes.SetUpRoutes(app)
	middleware.MiddlewareInit(app)

	app.Listen(":3000")
}