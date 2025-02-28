package services

import (
	"inventory/internal/app/domain"
	"inventory/internal/app/domain/entities"
	"inventory/internal/app/domain/ports"
)

type ProductServiceImp struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &ProductServiceImp{repo: repo}
}

func (s *ProductServiceImp) CreateProduct(product *entities.Product) error {
	if product.Price < 0 {
		return domain.ErrInvalidPrice
	}

	if err := s.repo.Save(product); err != nil {
		return err
	}

	return nil
}

func (s *ProductServiceImp) FindProductById(id uint) (*entities.Product, error) {
	return s.repo.FindById(id)
}

func (s *ProductServiceImp) FindAllProducts() ([]*entities.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductServiceImp) UpdateProduct(product *entities.Product) error {
	if err := s.repo.Update(product); err != nil {
		return err
	}
	return nil
}

func (s *ProductServiceImp) DeleteProduct(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
