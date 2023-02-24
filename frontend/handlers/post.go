package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend/models"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func ReadPost(c *fiber.Ctx) error {

	postId := c.Params("id")

	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(fmt.Sprintf("%s%s%s", os.Getenv("API_URL"), "/api/post/", postId))

	if err := a.Parse(); err != nil {
		panic(err)
	}

	code, body, errs := a.Bytes()
	if code != 200 && len(errs) > 0 {
		return c.Render("post", fiber.Map{
			"Post": errs,
		})
	}

	var res models.PostResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return c.Render("post", fiber.Map{
			"Post": err,
		})
	}

	return c.Render("post", fiber.Map{
		"Post": res.Data,
	})
}

func CreatePost(c *fiber.Ctx) error {

	newPost := new(models.Post)

	if err := c.BodyParser(newPost); err != nil {
		return c.Status(400).JSON(fiber.Map{"data": nil, "error": err.Error()})
	}

	json_data, err := json.Marshal(newPost)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s%s", os.Getenv("API_URL"), "/api/post"), "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	var res models.PostResponse
	json.NewDecoder(resp.Body).Decode(&res)

	return c.JSON(fiber.Map{"data": res.Data, "error": res.Error})
}
