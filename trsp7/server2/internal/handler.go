package internal

import (
	"database/sql"
	_ "frontend3_server2/docs"
	"math/rand"
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

	app.Get("/statham/quotes", h.GetQuotes)
	app.Get("/statham/img_url", h.GetImgURL)

	app.Listen(":10002")
}

// GetQuotes
// @Tags statham
// @Summary      Получение цитат
// @Accept       json
// @Produce      json
// @Success 200 {object} Message
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router       /statham/quotes [get]
func (h *Handler) GetQuotes(c *fiber.Ctx) error {
	quotes := []string{
		"Настоящий мужчина, как ковер тети Зины — с каждым годом лысеет.",
		"Мама учила не ругаться матом, но жизнь научила не ругаться матом при маме.",
		"Если тебе где-то не рады в рваных носках, то и в целых туда идти не стоит.",
		"«Жи-ши» пиши от души.",
		"Никогда не доверяйте таксистам, ведь они могут вас подвезти.",
		"Однажды городской тип купил поселок. Теперь это поселок городского типа.",
	}

	return c.Status(200).JSON(Message{Message: quotes[rand.Intn(len(quotes))]})
}

// GetImgURL
// @Tags statham
// @Summary      Получение картинок
// @Accept       json
// @Produce      json
// @Success 200 {object} Message
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router       /statham/img_url [get]
func (h *Handler) GetImgURL(c *fiber.Ctx) error {
	imgs := []string{
		"https://s.newslab.ru/photoalbum/1728/m/18064.jpg",
		"https://i.pinimg.com/originals/47/e9/35/47e935b8050380243a7cf5a1110b13ab.jpg",
		"https://sun9-79.userapi.com/impg/-TITt0dqVICrsPv2OPrd2M2Q5D2RoaZF7Xuwsg/tnADmZI60Vw.jpg?size=800x950&quality=95&sign=bc4359416dc23fbcec7302c3c35aea02&c_uniq_tag=kwXDWfElSQdnKNJU5v5WMp-maXAi3CG20yTQDvGG1L4&type=album",
		"https://avatars.mds.yandex.net/get-marketpic/8916051/pic648e303318e4df573d8fc62218c0387f/orig",
		"https://cdn-m-net.dstv.com/images/SearchEngineOptimization/2024/06/05/1013145/16/1717582137-34_GettyImages_1496879697.jpeg",
	}

	return c.Status(200).JSON(Message{Message: imgs[rand.Intn(len(imgs))]})
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
