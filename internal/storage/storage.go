package storage

import "database/sql"

type Store interface {
	// Products
	CreaeProducts() error
}

type Storage struct {
	DB *sql.DB
}

func NewStorage(sqlStorage *SqlStorage) *Storage {
	return &Storage{
		DB: sqlStorage.db,
	}
}

func (s *Storage) CreaeProducts() error {
	return nil
}
