package app

import (
	"log"

	"github.com/a-h/templ"
	hellopage "github.com/edwincarlflores/go-htmx/templates/hello"

	"github.com/gofiber/fiber/v2"
)

func HTML(c *fiber.Ctx, comp templ.Component) error {
	c.Set("Content-Type", fiber.MIMETextHTML)
	return comp.Render(c.Context(), c.Response().BodyWriter())
}

func App() {
	app := fiber.New()

	app.Static("/assets", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return HTML(c, hellopage.HelloPage())
	})

	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return HTML(c, hellopage.HelloName(name))
	})

	log.Fatal(app.Listen(":4000"))
}
