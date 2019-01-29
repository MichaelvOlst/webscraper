package sqlite

import "michaelvanolst.nl/scraper/websites"

// GetAttributes gets all the attributes for a website
func (db *SQLStore) GetAttributes(websiteID int64) ([]*websites.Attribute, error) {
	return nil, nil
}

// GetAttribute gets a single attribute for a website
func (db *SQLStore) GetAttribute(websiteID, id int64) (*websites.Attribute, error) {
	return nil, nil
}

// SaveAttribute saves a single attribute for a website
func (db *SQLStore) SaveAttribute(a *websites.Attribute) error {
	return nil
}

// DeleteAttribute deletes a single attribute for a website
func (db *SQLStore) DeleteAttribute(a *websites.Attribute) error {
	return nil
}
