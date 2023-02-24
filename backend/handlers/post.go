package handlers

import (
	"backend/dal"
	"backend/models"
	"strconv"

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

	return c.Status(201).JSON(fiber.Map{"data": strconv.FormatUint(uint64(post.ID), 10), "error": nil})
}

func ReadPost(c *fiber.Ctx) error {
	postId := c.Params("id")

	result, err := dal.ReadPost(postId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": result.Post, "error": nil})
}
