package main

import (
	"frontend/config"
	"frontend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {
	htmlEngine := html.New("./views", ".html")

	config.LoadEnvironmentVariables()

    app := fiber.New(fiber.Config{
        Views: htmlEngine,
    })

	app.Use(logger.New())

	routes.SetUpRoutes(app)

	app.Listen(":3001")
}