package repository

import (
	"inventory/internal/app/core/domain/model"
	"inventory/internal/app/core/domain"
	"inventory/internal/app/core/ports"

	"errors"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ports.ProductRepository {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) FindById(id uint) (*model.Product, error) {
	var product model.Product
	if result := r.db.First(&product, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, result.Error
	}
	return &product, nil
}

func (r *GormProductRepository) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	if result := r.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *GormProductRepository) Update(product *model.Product) error {
	if result := r.db.Save(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormProductRepository) Save(product *model.Product) error {
	if result := r.db.Create(&product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormProductRepository) Delete(id uint) error {
	if result := r.db.Delete(&model.Product{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
