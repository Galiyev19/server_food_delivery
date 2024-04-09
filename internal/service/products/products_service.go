package products

import "food_delivery/internal/repository/products"

type ProductsService struct {
	ProductRepository products.IProductRepository
}

func NewProductService(productRepo products.IProductRepository) *ProductsService {
	return &ProductsService{
		ProductRepository: productRepo,
	}
}

type IProductsService interface {
	InsertProduct() error
	GetListProduct() error
	GetProductById(id int) error
}

func (p *ProductsService) InsertProduct() error {
	return nil
}

func (p *ProductsService) GetListProduct() error {
	return nil
}

func (p *ProductsService) GetProductById(id int) error {
	return nil
}
