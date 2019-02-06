package sqlite

import "michaelvanolst.nl/scraper/models"

// GetWebsites gets all the websites
func (db *SQLStore) GetWebsites() ([]*models.Website, error) {
	results := []*models.Website{}

	rows, err := db.Query("SELECT * FROM websites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		w := &models.Website{}
		err := rows.Scan(&w.ID, &w.Name, &w.URL, &w.Holder)
		if err != nil {
			return nil, err
		}

		attributes, err := db.GetAttributes(w.ID)
		if err != nil {
			return nil, err
		}
		w.Attributes = attributes

		results = append(results, w)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, err
}

// GetWebsite gets a single website
func (db *SQLStore) GetWebsite(id int64) (*models.Website, error) {

	stmt, err := db.Prepare("SELECT * FROM websites WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	w := &models.Website{}
	err = stmt.QueryRow(id).Scan(&w.ID, &w.Name, &w.URL, &w.Holder)
	if err != nil {
		return nil, err
	}

	attributes, err := db.GetAttributes(w.ID)
	if err != nil {
		return nil, err
	}
	w.Attributes = attributes

	return w, nil
}

// SaveWebsite saves a website
func (db *SQLStore) SaveWebsite(w *models.Website) error {

	stmt, err := db.Prepare("INSERT INTO websites (name, url, holder) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(w.Name, w.URL, w.Holder)
	if err != nil {
		return err
	}

	return nil
}

// DeleteWebsite deletes a website
func (db *SQLStore) DeleteWebsite(w *models.Website) error {
	stmt, err := db.Prepare("DELETE from websites WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(w.ID)
	if err != nil {
		return err
	}

	return nil
}
