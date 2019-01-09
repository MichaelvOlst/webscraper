package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // ...
)

// Storage holds db
type Storage struct {
	db *sql.DB
}

// New returns a new storage
func New() (*Storage, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}
