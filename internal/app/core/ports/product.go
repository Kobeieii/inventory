package ports

import (
	"inventory/internal/app/core/domain/model"
)

type ProductRepository interface {
	FindById(id uint) (*model.Product, error)
	FindAll() ([]*model.Product, error)
	Update(product *model.Product) error
	Save(product *model.Product) error
	Delete(id uint) error
}
