package main

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("server1/index.html")
	})

	app.Listen(":3000")
}
