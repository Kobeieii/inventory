package services

import (
	"inventory/internal/app/domain"
	"inventory/internal/app/domain/entities"
	"inventory/internal/app/domain/ports"
)

type ProductService interface {
	CreateProduct(product *entities.Product) error
	FindProductById(id uint) (*entities.Product, error)
	FindAllProducts() ([]*entities.Product, error)
	UpdateProduct(product *entities.Product) error
	DeleteProduct(id uint) error
}

type ProductServiceImpl struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) CreateProduct(product *entities.Product) error {
	if product.Price < 0 {
		return domain.ErrInvalidPrice
	}

	if err := s.repo.Save(product); err != nil {
		return err
	}

	return nil
}

func (s *ProductServiceImpl) FindProductById(id uint) (*entities.Product, error) {
	return s.repo.FindById(id)
}

func (s *ProductServiceImpl) FindAllProducts() ([]*entities.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductServiceImpl) UpdateProduct(product *entities.Product) error {
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
