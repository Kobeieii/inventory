package ports

import (
	"inventory/internal/app/domain/entities"
)

type ProductService interface {
	CreateProduct(product *entities.Product) error
	FindProductById(id uint) (*entities.Product, error)
	FindAllProducts() ([]*entities.Product, error)
	UpdateProduct(product *entities.Product) error
	DeleteProduct(id uint) error
}