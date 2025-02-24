package adapter

import (
	"inventory/internal/app/domain/model"
	"inventory/internal/app/domain/service"

	"github.com/gofiber/fiber/v2"
)

type HttpProductHandler struct {
	service service.ProductService
}

func (h *HttpProductHandler) RegisterRoutes(api fiber.Router) {
	productApi := api.Group("/products")
	productApi.Post("/", h.CreateProduct)
	productApi.Get("/:id", h.FindProductById)
	productApi.Get("/", h.FindAllProducts)
	productApi.Patch("/:id", h.UpdateProduct)
	productApi.Delete("/:id", h.DeleteProduct)
}

func NewHttpProductHandler(service service.ProductService) HttpProductHandler {
	return HttpProductHandler{service: service}
}

func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *HttpProductHandler) FindProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	product, err := h.service.FindProductById(uint(id))
	if product == nil && err == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *HttpProductHandler) FindAllProducts(c *fiber.Ctx) error {
	products, err := h.service.FindAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (h *HttpProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	product, err := h.service.FindProductById(uint(id))
	if product == nil && err == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var updateProduct model.Product
	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	updateProduct.ID = uint(id)
	if err := h.service.UpdateProduct(&updateProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(updateProduct)

}

func (h *HttpProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"details": "Deleted Success"})
}