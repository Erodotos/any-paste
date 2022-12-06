package main

import (
	"backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetUpRoutes(app)

	app.Listen(":3000")
}