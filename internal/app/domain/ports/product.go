package ports

import (
	"inventory/internal/app/domain/entities"
)

type ProductRepository interface {
	FindById(id uint) (*entities.Product, error)
	FindAll() ([]*entities.Product, error)
	Update(product *entities.Product) error
	Save(product *entities.Product) error
	Delete(id uint) error
}
