package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	db *sql.DB
}

func (h *Handler) InitRoutes() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})
	app.Get("/product", h.GetProducts)

	app.Listen(":10001")
}

// GetProducts
// @Tags product
// @Summary      Получение товаров
// @Accept       json
// @Produce      json
// @Success 200 {object} Catalog
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router       /product [get]
func (h *Handler) GetProducts(c *fiber.Ctx) error {
	products, err := GetAllProducts(h.db)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"products": products})
}
