package services

import (
	"inventory/internal/app/domain"
	"inventory/internal/app/domain/entities"
	"inventory/internal/app/domain/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

var _ ports.ProductService = &ProductService{}

func NewProductService(repo ports.ProductRepository) ProductService {
	return ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *entities.Product) error {
	if product.Price < 0 {
		return domain.ErrInvalidPrice
	}

	if err := s.repo.Save(product); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) FindProductById(id uint) (*entities.Product, error) {
	return s.repo.FindById(id)
}

func (s *ProductService) FindAllProducts() ([]*entities.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) UpdateProduct(product *entities.Product) error {
	if err := s.repo.Update(product); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
