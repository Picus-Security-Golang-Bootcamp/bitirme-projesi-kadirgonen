package service

import (
	"errors"

	model "HW/app/models"
	"HW/app/repo"
)

type ProductService struct {
	productRepo repo.ProductRepository
}

func NewProductService(pr repo.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: pr,
	}
}
func (s *ProductService) GetAllProducts(pageIndex, pageSize int) ([]model.Product, int) {
	items, count := s.productRepo.GetAllProducts(pageIndex, pageSize)

	return items, count
}

func (s *ProductService) GetProduct(productId int) *model.Product {
	return s.productRepo.GetProduct(productId)
}

func (s *ProductService) GetProductStockCode(StockCode string) *model.Product {
	return s.productRepo.GetProductStockCode(StockCode)
}

func (s *ProductService) SearchProducts(query string) []model.Product {
	return s.productRepo.SearchProduct(query)
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	productExists := s.productRepo.GetProductStockCode(product.StockCode)
	if productExists != nil {
		return errors.New("product with same sku already exist in database")
	}

	err := s.productRepo.CreateProduct(product)
	if err != nil {
		return errors.New("an unknown error occurred during operation")
	}

	return nil
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	err := s.productRepo.UpdateProduct(product)
	if err != nil {
		return errors.New("an unknown error occurred during operation")
	}

	return nil
}

func (s *ProductService) DeleteProduct(productId int) error {
	err := s.productRepo.DeleteProduct(productId)
	if err != nil {
		return errors.New("an unknown error occurred during operation")
	}

	return nil
}
