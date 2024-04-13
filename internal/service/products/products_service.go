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
	GetProductById(id int) (models.Products, error)
	UpdateProduct(id int, product models.Products) error
	DeleteProduct(id int) error
}

func (p *ProductsService) InsertProduct(product models.Products) error {
	productModel := models.Products{
		Title:       product.Title,
		Description: product.Description,
		Category:    product.Category,
		Price:       product.Price,
		Image:       product.Image,
		Rating:      product.Rating,
	}
	fmt.Println("Model: - :", productModel)

	if err := p.ProductRepository.InsertProduct(productModel); err != nil {
		return fmt.Errorf("Create product %v", err)
	}
	return nil
}

func (p *ProductsService) GetListProduct() error {
	return nil
}

func (p *ProductsService) GetProductById(id int) (models.Products, error) {
	product, err := p.ProductRepository.GetProductById(id)
	if err != nil {
		return models.Products{}, err
	}
	return product, nil
}

func (p *ProductsService) UpdateProduct(id int, product models.Products) error {
	findProduct, err := p.GetProductById(id)
	if err != nil {
		return err
	}

	findProduct.Title = product.Title               // ttile
	findProduct.Description = product.Description   // description
	findProduct.Price = product.Price               // price
	findProduct.Category = product.Category         // category
	findProduct.Image = product.Image               // image
	findProduct.Rating.Rate = product.Rating.Rate   // rate
	findProduct.Rating.Count = product.Rating.Count // count

	err = p.ProductRepository.UpdateProduct(findProduct)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (p *ProductsService) DeleteProduct(id int) error {
	_, err := p.GetProductById(id)
	if err != nil {
		return fmt.Errorf("product with id %d nout found", id)
	}

	err = p.ProductRepository.DeleteProduct(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
