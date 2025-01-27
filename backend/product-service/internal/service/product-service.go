package service

import (
	"errors"
	"product-service/internal/models"
	"product-service/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product models.Product) error {

	if product.ProductName == "" ||
		product.Price == 0 ||
		product.Stock == 0 ||
		product.MainCategory == 0 ||
		product.SubCategory == 0 ||
		product.UserID == 0 {
		return errors.New("all fields must be filled")
	}
	if product.Price < 0 {
		return errors.New("price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("stock must be greater than 0")
	}
	if err := s.repo.CreateProduct(product); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetProduct(productFilter models.ProductFilter, page int, userId uint) ([]models.ResponseProduct, error) {

	results, err := s.repo.GetProducts(productFilter, page, userId)
	if err != nil {
		return []models.ResponseProduct{}, err
	}

	return results, nil
}

func (r *ProductService) DeleteProductById(ProductID string, userId uint) error {
	if ProductID == "" {
		return errors.New("product id must be filled")
	}
	if err := r.repo.DeleteProductById(ProductID, userId); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) InactivateProductById(productId string, userId uint) error {
	if productId == "" {
		return errors.New("product id must be filled")
	}
	if err := s.repo.InactivateProductById(productId, userId); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) ActivateProduct(productId string, userId uint) error {
	if productId == "" {
		return errors.New("product id is required")
	}
	if err := s.repo.ActivateProduct(productId, userId); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProduct(product models.Product, userId uint) error {

	return nil
}
