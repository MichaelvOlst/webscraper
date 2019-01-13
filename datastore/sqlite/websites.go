package sqlite

import "michaelvanolst.nl/scraper/website"

// GetWebsites gets all the websites
func (db *SQLStore) GetWebsites() ([]*website.Website, error) {
	return nil, nil
}

// GetWebsite gets a single website
func (db *SQLStore) GetWebsite(id int64) (*website.Website, error) {
	return nil, nil
}

// SaveWebsite saves a website
func (db *SQLStore) SaveWebsite(w *website.Website) error {
	return nil
}

// DeleteWebsite deletes a website
func (db *SQLStore) DeleteWebsite(w *website.Website) error {
	return nil
}
