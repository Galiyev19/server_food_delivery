package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateSqlDB(dbDriver, dbPath, migrationPath string) (*sql.DB, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open(dbDriver, dbPath)
	if err != nil {
		fmt.Println("OPEN ERROR")
		return nil, fmt.Errorf("%w %s", err, op)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("PING ERROR")
		return nil, fmt.Errorf("%w %s", err, op)
	}

	stmt, err := os.ReadFile(migrationPath)
	if err != nil {
		fmt.Println("READ FILE ERROR")
		return nil, fmt.Errorf("%w %s", err, op)
	}

	_, err = db.Exec(string(stmt))
	if err != nil {
		return nil, fmt.Errorf("%w %s", err, op)
	}
	return db, nil
}

func InsertTestDataInUser(dbDriver, dbPath, initPath string) error {
	db, err := sql.Open(dbDriver, dbPath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	stmt, err := os.ReadFile(initPath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	_, err = db.Exec(string(stmt))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
