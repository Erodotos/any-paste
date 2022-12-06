package handlers

import (
	"github.com/erodotos/any-paste/dal"
	"github.com/erodotos/any-paste/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	if err := dal.CreateBook(book); err != nil {
		return c.Status(500).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"data": "success", "error": nil})
}

func ReadPost(c *fiber.Ctx) error {
	book_id := c.Params("id")

	err, result := dal.ReadBook(book_id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": result, "error": nil})
}
