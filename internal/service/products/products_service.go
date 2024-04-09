package products

import (
	"fmt"
	"food_delivery/internal/models"
	"food_delivery/internal/repository/products"
)

type ProductsService struct {
	ProductRepository products.IProductRepository
}

func NewProductService(productRepo products.IProductRepository) *ProductsService {
	return &ProductsService{
		ProductRepository: productRepo,
	}
}

type IProductsService interface {
	InsertProduct(product models.Products) error
	GetListProduct() error
	GetProductById(id int) error
}

func (p *ProductsService) InsertProduct(product models.Products) error {
	producModel := models.Products{
		Title:       product.Title,
		Description: product.Description,
		Category:    product.Category,
		Price:       product.Price,
		Image:       product.Image,
		Rating:      product.Rating,
	}

	if err := p.ProductRepository.InsertProduct(producModel); err != nil {
		return fmt.Errorf("Create product %v", err)
	}
	return nil
}

func (p *ProductsService) GetListProduct() error {
	return nil
}

func (p *ProductsService) GetProductById(id int) error {
	return nil
}
