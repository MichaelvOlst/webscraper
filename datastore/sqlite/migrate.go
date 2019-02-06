package sqlite

import "database/sql"

// Migrate creates the tables for the app
func Migrate(db *sql.DB) error {
	var err error
	err = createWebsitesTable(db)
	if err != nil {
		return err
	}

	err = createAttributesTable(db)
	if err != nil {
		return err
	}

	err = createLinksTable(db)
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

func createAttributesTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS attributes (
		id INTEGER PRIMARY KEY, 
		website_id INTEGER, 
		search VARCHAR(255),
		type VARCHAR(255),
		value VARCHAR(255),
		FOREIGN KEY(website_id) REFERENCES websites(id)
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

func createLinksTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS links (
		id INTEGER PRIMARY KEY, 
		website_id INTEGER,
		url VARCHAR(255),
		price VARCHAR(255), 
		address VARCHAR(255),
		status VARCHAR(255),
		image_url VARCHAR(255),
		modified_date TEXT,
		created_date TEXT,
		FOREIGN KEY(website_id) REFERENCES websites(id)
		
	);

	CREATE UNIQUE INDEX unique_link ON links (url);

	
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
