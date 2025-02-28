package repositories

import (
	"inventory/internal/app/domain/entities"
	"inventory/internal/app/domain"
	"inventory/internal/app/domain/ports"

	"errors"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ports.ProductRepository {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) FindById(id uint) (*entities.Product, error) {
	var product entities.Product
	if result := r.db.First(&product, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, result.Error
	}
	return &product, nil
}

func (r *GormProductRepository) FindAll() ([]*entities.Product, error) {
	var products []*entities.Product
	if result := r.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *GormProductRepository) Update(product *entities.Product) error {
	if result := r.db.Save(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormProductRepository) Save(product *entities.Product) error {
	if result := r.db.Create(&product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormProductRepository) Delete(id uint) error {
	if result := r.db.Delete(&entities.Product{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
