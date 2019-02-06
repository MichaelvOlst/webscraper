package sqlite

import "michaelvanolst.nl/scraper/models"

// GetAttributes gets all the attributes for a website
func (db *SQLStore) GetAttributes(websiteID int64) ([]*models.Attribute, error) {

	results := []*models.Attribute{}

	stmt, err := db.Prepare("SELECT id, website_id, search, type, value FROM attributes WHERE website_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(websiteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		a := &models.Attribute{}
		err := rows.Scan(&a.ID, &a.WebsiteID, &a.Search, &a.Type, &a.Value)
		if err != nil {
			return nil, err
		}

		results = append(results, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, err
}

// SaveAttribute saves a single attribute for a website
func (db *SQLStore) SaveAttribute(a *models.Attribute) error {
	return nil
}

// DeleteAttribute deletes a single attribute for a website
func (db *SQLStore) DeleteAttribute(a *models.Attribute) error {
	return nil
}
