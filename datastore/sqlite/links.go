package sqlite

import (
	"errors"
	"time"

	"michaelvanolst.nl/scraper/models"
)

// GetAllLinks gets all the links
func (db *SQLStore) GetAllLinks() ([]*models.Link, error) {
	results := []*models.Link{}

	rows, err := db.Query("SELECT id, url, image_url, address, price, status FROM links ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		l := &models.Link{}
		err := rows.Scan(&l.ID, &l.URL, &l.ImageURL, &l.Address, &l.Price, &l.Status)
		if err != nil {
			return nil, err
		}

		results = append(results, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, err
}

// GetLinksByID gets all the links
func (db *SQLStore) GetLinksByID(ID int64) ([]*models.Link, error) {
	results := []*models.Link{}

	query := `
		SELECT id, url, image_url, address, price, status FROM links 
		WHERE website_id = ? ORDER BY id DESC
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		l := &models.Link{}
		err := rows.Scan(&l.ID, &l.URL, &l.ImageURL, &l.Address, &l.Price, &l.Status)
		if err != nil {
			return nil, err
		}

		results = append(results, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, err
}

// SaveLink saves a website
func (db *SQLStore) SaveLink(l *models.Link) error {

	exists, err := db.CheckLinkExists(l.URL)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Link already exists")
	}

	t := time.Now()

	l.CreatedDate = t.Format("2006-01-02 15:04:05")
	l.ModifiedDate = t.Format("2006-01-02 15:04:05")

	query := `
		INSERT INTO links (website_id, url, price, address, status, image_url, modified_date, created_date) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(l.WebsiteID, l.URL, l.Price, l.Address, l.Status, l.ImageURL, l.ModifiedDate, l.CreatedDate)
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
