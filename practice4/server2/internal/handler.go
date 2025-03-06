package internal

import (
	"database/sql"
	_ "frontend3_server2/docs"
	"strconv"

	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Db *sql.DB
}

func (h *Handler) InitRoutes() {
	app := fiber.New()
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Post("/product", h.CreateProducts)
	app.Put("/product", h.UpdateProduct)
	app.Delete("/product/:id", h.DeleteProduct)
	app.Get("/messenger", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	app.Listen(":10002")
}

// CreateProducts
// @Tags product
// @Summary      Создание товаров
// @Accept       json
// @Produce      json
// @Param data body CreateProductRequest true "Данные товаров"
// @Success 200 {object} Message
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router       /product [post]
func (h *Handler) CreateProducts(c *fiber.Ctx) error {
	var catalog Catalog
	err := c.BodyParser(&catalog)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	var ids []int
	for _, product := range catalog.Products {
		id, err := CreateProduct(h.Db, product)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		ids = append(ids, id)
	}

	return c.Status(200).JSON(Ids{Ids: ids})
}

// UpdateProduct
// @Tags product
// @Summary      Обновление товара
// @Accept       json
// @Produce      json
// @Param data body Product true "Данные товара"
// @Success 200 {object} Message
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router       /product [put]
func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	var product Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err = UpdateProduct(h.Db, product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "product updated"})
}

// DeleteProduct
// @Tags product
// @Summary      Удаление товара
// @Accept       json
// @Produce      json
// @Param id path int true "id товара"
// @Success 200 {object} Message
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router       /product/{id} [delete]
func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err = DeleteProduct(h.Db, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "product deleted"})
}
