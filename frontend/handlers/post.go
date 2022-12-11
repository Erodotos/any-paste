package handlers

import (
	"encoding/json"
	"fmt"
	"frontend/models"
	"os"

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

func CreatePost(c *fiber.Ctx) error {

	client := fiber.AcquireClient()
	a := client.Post(os.Getenv("API_URL"))
	a.JSON(fiber.Map{"post": "123"})
	code, body, _ := a.Bytes()

	fmt.Println(code)

	post := models.Test{}
	json.Unmarshal(body, &post)

	fmt.Println(post)

	return c.JSON(post)
}
