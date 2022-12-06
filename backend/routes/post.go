package routes

import (

	"github.com/gofiber/fiber/v2"
)

func PostRouter(api fiber.Router) {
	bookGroup := api.Group("/post")
	bookGroup.Get("/:id", handlers.ReadBook)
	bookGroup.Post("/", handlers.CreateBook)
}