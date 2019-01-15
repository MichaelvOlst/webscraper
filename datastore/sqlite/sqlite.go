package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // ...
)

// SQLStore ...
type SQLStore struct {
	*sql.DB
}

// Close the database
func (s *SQLStore) Close() error {
	return s.DB.Close()
}

// New returns a new sqllite connection
func New() (*SQLStore, error) {

	// databaseFile := viper.Get("database.filename")

	db, err := sql.Open("sqlite3", "./scraper.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	return &SQLStore{db}, nil
}
