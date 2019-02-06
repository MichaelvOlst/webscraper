package sqlite

import (
	"errors"

	"michaelvanolst.nl/scraper/models"
)

// GetLinks gets all the links
func (db *SQLStore) GetLinks() ([]*models.Link, error) {
	// results := []*models.Link{}

	// rows, err := db.Query("SELECT * FROM links")
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	l := &models.Link{}
	// 	err := rows.Scan(&w.ID, &w.Name, &w.URL, &w.Holder)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	attributes, err := db.GetAttributes(w.ID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	w.Attributes = attributes

	// 	results = append(results, w)
	// }

	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }

	// return results, err
	return nil, nil
}

// // GetWebsite gets a single website
// func (db *SQLStore) GetWebsite(id int64) (*models.Website, error) {

// 	stmt, err := db.Prepare("SELECT * FROM websites WHERE id=?")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	w := &models.Website{}
// 	err = stmt.QueryRow(id).Scan(&w.ID, &w.Name, &w.URL, &w.Holder)
// 	if err != nil {
// 		return nil, err
// 	}

// 	attributes, err := db.GetAttributes(w.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	w.Attributes = attributes

// 	return w, nil
// }

// SaveLink saves a website
func (db *SQLStore) SaveLink(l *models.Link) error {

	exists, err := db.CheckLinkExists(l.URL)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Link already exists")
	}

	query := `
		INSERT INTO links (website_id, url, price, address, status, image_url) 
		VALUES (?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(l.WebsiteID, l.URL, l.Price, l.Address, l.Status, l.ImageURL)
	if err != nil {
		return err
	}

	return nil
}

// CheckLinkExists deletes a website
func (db *SQLStore) CheckLinkExists(url string) (bool, error) {
	stmt, err := db.Prepare("SELECT COUNT(*) as count from links WHERE url=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var count int

	err = stmt.QueryRow(url).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, err
	}

	return true, nil
}

// // DeleteWebsite deletes a website
// func (db *SQLStore) DeleteWebsite(w *models.Website) error {
// 	stmt, err := db.Prepare("DELETE from websites WHERE id=?")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(w.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
