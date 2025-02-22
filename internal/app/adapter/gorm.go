package adapter

import (
	"inventory/internal/app/domain/model"
	"inventory/internal/app/domain/repository"

	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) repository.ProductRepository {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) Save(product *model.Product) error {
	if result := r.db.Create(&product); result.Error != nil {
		return result.Error
	}
	return nil
}
