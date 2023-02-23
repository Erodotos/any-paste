package handlers

import (
	"backend/dal"
	"backend/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	if err := dal.CreatePost(post); err != nil {
		return c.Status(500).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"data": post.ID, "error": nil})
}

func ReadPost(c *fiber.Ctx) error {
	postId := c.Params("id")

	fmt.Println(postId)

	result, err := dal.ReadPost(postId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	fmt.Println(result.Post)
	fmt.Printf("%T", result.Post)

	return c.Status(200).JSON("aloou")
}
