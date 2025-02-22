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
