package handlers

import (
	"inventory/internal/app/domain"
	"inventory/internal/app/domain/ports"
	"inventory/internal/app/domain/entities"
	"inventory/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func toProductEntity(p *ProductDTO) *entities.Product {
	return &entities.Product{
		ID:     p.ID,
		Name:   p.Name,
		Price:  p.Price,
	}
}

func toProductDTO(p *entities.Product) *ProductDTO {
	return &ProductDTO{
		ID:     p.ID,
		Name:   p.Name,
		Price:  p.Price,
	}
}

type HttpProductHandler struct {
	service ports.ProductService
}

func (h *HttpProductHandler) RegisterRoutes(api fiber.Router) {
	productApi := api.Group("/products")
	productApi.Post("/", h.CreateProduct)
	productApi.Get("/:id", h.FindProductById)
	productApi.Get("/", h.FindAllProducts)
	productApi.Patch("/:id", h.UpdateProduct)
	productApi.Delete("/:id", h.DeleteProduct)
}

func NewHttpProductHandler(service ports.ProductService) HttpProductHandler {
	return HttpProductHandler{service: service}
}

func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	var productDTO ProductDTO
	if err := c.BodyParser(&productDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	product := toProductEntity(&productDTO)
	if err := h.service.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(toProductDTO(product))
}

func (h *HttpProductHandler) FindProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	product, err := h.service.FindProductById(uint(id))
	if err != nil {
		if err == domain.ErrProductNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product Not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(toProductDTO(product))
}

func (h *HttpProductHandler) FindAllProducts(c *fiber.Ctx) error {
	products, err := h.service.FindAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(utils.Map(products, toProductDTO))
}

func (h *HttpProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	product, err := h.service.FindProductById(uint(id))
	if err != nil {
		if err == domain.ErrProductNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product Not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var updateProduct entities.Product
	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	updateProduct.ID = product.ID
	if err := h.service.UpdateProduct(&updateProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(toProductDTO(&updateProduct))

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
