package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func PostRouter(api fiber.Router) {
	bookGroup := api.Group("/post")
	bookGroup.Get("/:id", handlers.ReadPost)
	bookGroup.Post("/", handlers.CreatePost)
}