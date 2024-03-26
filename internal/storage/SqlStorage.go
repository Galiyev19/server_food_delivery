package storage

import (
	"database/sql"
	"fmt"
	"os"
)

type SqlStorage struct {
	db *sql.DB
}

func New(dbDriver, dbPath, migrationPath string) (*SqlStorage, error) {
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
	return &SqlStorage{db: db}, nil
}
