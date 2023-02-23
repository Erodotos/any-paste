package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend/models"
	"io/ioutil"
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

	resp, err := http.Get(fmt.Sprintf("%s%s%s", os.Getenv("API_URL"), "/api/post/", postId))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Body)

	var res models.PostResponse
	json.NewDecoder(resp.Body).Decode(&res)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	fmt.Println(body)

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
