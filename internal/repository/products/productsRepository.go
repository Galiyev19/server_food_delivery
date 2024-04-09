package products

import (
	"database/sql"
	"fmt"
	"food_delivery/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductsRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

type IProductRepository interface {
	InsertProduct(product models.Products) error
	GetListProduct() error
	GetProductById(id int) error
}

func (p *ProductRepository) InsertProduct(product models.Products) error {
	stmt := `INSERT INTO product (title, description, category, price, image, rating)
		VALUES (?, ?, ?, ?, ?, ?);`
	if _, err := p.db.Exec(stmt, product.Title, product.Description, product.Category, product.Price, product.Image, product.Rating); err != nil {
		return fmt.Errorf("p.db.Exec %v", err)
	}
	return nil
}

func (p *ProductRepository) GetListProduct() error {
	return nil
}

func (p *ProductRepository) GetProductById(id int) error {
	return nil
}
