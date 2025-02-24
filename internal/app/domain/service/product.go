package service

import (
	"errors"
	"inventory/internal/app/domain/model"
	"inventory/internal/app/domain/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	FindProductById(id uint) (*model.Product, error)
}

type ProductServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) CreateProduct(product *model.Product) error {
	if product.Price < 0 {
		return errors.New("price must be positive")
	}

	if err := s.repo.Save(product); err != nil {
		return err
	}

	return nil
}

func (s *ProductServiceImpl) FindProductById(id uint) (*model.Product, error) {
	return s.repo.FindById(id)
}
