package main

import (
	
	"frontend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	htmlEngine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: htmlEngine,
    })

	app.Use(logger.New())

	routes.SetUpRoutes(app)

	app.Listen(":3001")
}