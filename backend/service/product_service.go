package service

import (
	"product_manage/model"
	"product_manage/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.repo.UpdateProduct(product)
}
func (s *ProductService) DeleteProduct(productID string) error {
	return s.repo.DeleteProduct(productID)
}
func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.repo.GetAllProducts()
}
func (s *ProductService) GetProductById(productID string) (model.Product, error) {
	return s.repo.GetProductById(productID)
}
