package main

import (
	"backend/routes"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetUpRoutes(app)
	app.Use(logger.New())
	app.Listen(":3000")
}