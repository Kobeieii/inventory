package service

import (
	"inventory/internal/app/core/domain"
	"inventory/internal/app/core/domain/model"
	"inventory/internal/app/core/ports"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	FindProductById(id uint) (*model.Product, error)
	FindAllProducts() ([]*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id uint) error
}

type ProductServiceImpl struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) CreateProduct(product *model.Product) error {
	if product.Price < 0 {
		return domain.ErrInvalidPrice
	}

	if err := s.repo.Save(product); err != nil {
		return err
	}

	return nil
}

func (s *ProductServiceImpl) FindProductById(id uint) (*model.Product, error) {
	return s.repo.FindById(id)
}

func (s *ProductServiceImpl) FindAllProducts() ([]*model.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductServiceImpl) UpdateProduct(product *model.Product) error {
	if err := s.repo.Update(product); err != nil {
		return err
	}
	return nil
}

func (s *ProductServiceImpl) DeleteProduct(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
