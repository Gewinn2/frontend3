package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	f := InitRoutes()
	f.Listen(":8080")
}

func InitRoutes() *fiber.App {
	f := fiber.New()

	f.Use(cors.New())
	f.Post("/signup", SignUp)
	f.Post("/login", Login)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return WithJWTAuth(c, "12345678")
	})
	authGroup.Get("/check", CheckLogin)

	return f
}
