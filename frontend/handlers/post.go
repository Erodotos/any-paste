package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func ReadPost(c *fiber.Ctx) error {
	return c.Render("post", fiber.Map{
		"Post": "Hello, World!",
	})
}
