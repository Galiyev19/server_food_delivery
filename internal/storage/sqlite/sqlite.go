package sqlite

import (
	"database/sql"
	"fmt"
	"food_delivery/internal/data"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(dbDriver, dbPath, migrationPath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open(dbDriver, dbPath)
	if err != nil {
		fmt.Println("ERROR")
		return nil, fmt.Errorf("%w %s", err, op)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("ERROR")
		return nil, fmt.Errorf("%w %s", err, op)
	}

	stmt, err := os.ReadFile(migrationPath)
	if err != nil {
		fmt.Println("ERROR")
		return nil, fmt.Errorf("%w %s", err, op)
	}

	_, err = db.Exec(string(stmt))
	if err != nil {
		return nil, fmt.Errorf("%w %s", err, op)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) SaveProduct(data data.Products) (int64, error) {
	const op = "storage.sqlite.SaveProduct"

	result, err := s.db.Exec(
		`INSERT INTO product (title,description,price,category,image,rating_rate,rating_count) VALUES (?, ?, ?, ?, ?, ?, ?);`,
		data.Title, data.Description, data.Price, data.Category, data.Image, data.Rating.Rate, data.Rating.Count)
	if err != nil {
		return -1, fmt.Errorf("%w %s", err, op)
	}
	return result.LastInsertId()
}
