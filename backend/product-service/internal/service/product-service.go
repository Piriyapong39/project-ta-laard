package service

import (
	"errors"
	"fmt"
	"product-service/internal/models"
	"product-service/internal/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product models.Product) error {
	uuid := uuid.New()
	product.ProductID = uuid.String()
	fmt.Println(product.MainCategory)
	// if product.ProductName == "" ||
	// 	product.Price == 0 ||
	// 	product.Stock == 0 ||
	// 	product.MainCategory == 0 ||
	// 	product.SubCategory == 0 ||
	// 	product.UserID == 0 {
	// 	return errors.New("all fields must be filled")
	// }
	fmt.Println(product)

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

func (r *ProductService) DeleteProductById(ProductID string) error {
	if ProductID == "" {
		return errors.New("product id must be filled")
	}
	if err := r.repo.DeleteProductById(ProductID); err != nil {
		return err
	}
	return nil
}

func (r *ProductService) InactivateProductById(ProductID string) error {
	if ProductID == "" {
		return errors.New("product id must be filled")
	}
	if err := r.repo.InactivateProductById(ProductID); err != nil {
		return err
	}
	return nil
}
