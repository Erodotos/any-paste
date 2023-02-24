package main

import (
	"backend/config"
	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	config.LoadEnvironmentVariables()
	database.InitializeDatabase()

	routes.SetUpRoutes(app)
	app.Use(logger.New())
	app.Listen(":3000")
}
