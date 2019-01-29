package sqlite

import "database/sql"

// Migrate creates the tables for the app
func Migrate(db *sql.DB) error {
	var err error
	err = createWebsitesTable(db)
	if err != nil {
		return err
	}

	return nil
}

func createWebsitesTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS websites (
		id INTEGER PRIMARY KEY, 
		name VARCHAR(255), 
		url VARCHAR(255),
		holder VARCHAR(255)
	)
	`
	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}
