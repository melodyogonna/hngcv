package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type ContactDetails struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Phone   string `form:"phone"`
	Message string `form:"message"`
}

func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"name":    "Anyaegbulam Melody",
			"email":   "melodyogonna@gmail.com",
			"phone":   "+2348022522386",
			"address": "Enugu, Nigeria",
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		ct := new(ContactDetails)

		if err := c.BodyParser(ct); err != nil {
			return err
		}

		return c.Render("greet", fiber.Map{
			"name": ct.Name,
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
