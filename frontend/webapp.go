package main

import (
	
	"frontend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	htmlEngine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: htmlEngine,
    })

	routes.SetUpRoutes(app)

	app.Listen(":3001")
}