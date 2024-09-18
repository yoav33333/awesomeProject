package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello ENV! " + os.Getenv("TEST_ENV"))
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen("0.0.0.0" + port))
}
