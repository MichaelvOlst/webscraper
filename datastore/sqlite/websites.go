package sqlite

import (
	"michaelvanolst.nl/scraper/websites"
)

// GetWebsites gets all the websites
func (db *SQLStore) GetWebsites() ([]*websites.Website, error) {
	results := []*websites.Website{}

	rows, err := db.Query("SELECT * FROM websites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		w := &websites.Website{}
		err := rows.Scan(&w.ID, &w.Name, &w.URL)
		if err != nil {
			return nil, err
		}

		results = append(results, w)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, err
}

// GetWebsite gets a single website
func (db *SQLStore) GetWebsite(id int64) (*websites.Website, error) {
	return nil, nil
}

// SaveWebsite saves a website
func (db *SQLStore) SaveWebsite(w *websites.Website) error {
	return nil
}

// DeleteWebsite deletes a website
func (db *SQLStore) DeleteWebsite(w *websites.Website) error {
	return nil
}
