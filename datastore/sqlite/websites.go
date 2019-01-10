package sqlite

import "michaelvanolst.nl/scraper/website"

// GetWebsites gets all the websites
func (db *SqlStore) GetWebsites() ([]*website.Website, error) {
	return nil, nil
}

// GetWebsite gets a single website
func (db *SqlStore) GetWebsite(id int64) (*website.Website, error) {
	return nil, nil
}

// SaveWebsite saves a website
func (db *SqlStore) SaveWebsite(w *website.Website) error {
	return nil
}

// DeleteWebsite deletes a website
func (db *SqlStore) DeleteWebsite(w *website.Website) error {
	return nil
}
