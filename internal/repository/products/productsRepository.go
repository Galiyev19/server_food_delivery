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
	GetProductById(id int) (models.Products, error)
	UpdateProduct(product models.Products) error
	DeleteProduct(id int) error
}

func (p *ProductRepository) InsertProduct(product models.Products) error {
	stmt := `INSERT INTO product (created_at, title, description, category, price, image, rating_rate, rating_count)
	VALUES (datetime('now', 'localtime'), ?, ?, ?, ?, ?, ?, ?);`
	if _, err := p.db.Exec(stmt, product.Title, product.Description, product.Category, product.Price, product.Image, product.Rating.Rate, product.Rating.Count); err != nil {
		return fmt.Errorf("p.db.Exec %v", err)
	}
	return nil
}

func (p *ProductRepository) GetListProduct() error {
	return nil
}

func (p *ProductRepository) GetProductById(id int) (models.Products, error) {
	var product models.Products

	stmt := `SELECT * FROM product where id = ?`

	if err := p.db.QueryRow(stmt, id).Scan(&product.ID,
		&product.CreadAt, &product.Title, &product.Description,
		&product.Price, &product.Category, &product.Image,
		&product.Rating.Rate, &product.Rating.Count); err != nil {
		return models.Products{}, fmt.Errorf("Not found product")
	}

	return product, nil
}

func (p *ProductRepository) UpdateProduct(product models.Products) error {
	stmt := `UPDATE product 
	SET title = ?, description = ?,price = ?,category = ?,image = ?,rating_rate = ?,rating_count = ?
	WHERE id = ?;`

	if _, err := p.db.Exec(stmt, product.Title, product.Description, product.Price, product.Category, product.Image, product.Rating.Rate, product.Rating.Count, product.ID); err != nil {
		return fmt.Errorf("p.db.Exec %v", err)
	}
	return nil
}

// Delete product

func (a *ProductRepository) DeleteProduct(id int) error {
	stmt := `DELETE FROM product WHERE id = ?`
	if _, err := a.db.Exec(stmt, id); err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	return nil
}
