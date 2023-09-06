package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func App() {
	app := fiber.New()

	app.Static("/assets", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK")
	})

	log.Fatal(app.Listen(":4000"))
}
