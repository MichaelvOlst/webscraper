package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

// SqlStore ...
type SqlStore struct {
	*sql.DB
}

// Close the database
func (db *SqlStore) Close() error {
	return db.Close()
}

// New returns a new sqllite connection
func New() (*SqlStore, error) {

	databaseFile := viper.Get("database.filename")

	db, err := sql.Open("sqlite3", databaseFile.(string))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &SqlStore{db}, nil
}
