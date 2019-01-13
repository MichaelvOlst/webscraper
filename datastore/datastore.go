package datastore

import (
	"michaelvanolst.nl/scraper/datastore/sqlite"
	"michaelvanolst.nl/scraper/website"
)

// Database has the config for datastore
type Database struct {
	Driver   string `default:"sqlite3"`
	Host     string `default:""`
	User     string `default:""`
	Password string `default:""`
	Name     string `default:"scraper.db"`
	SSLMode  string `default:""`
}

// Datastore represents a database implementations
type Datastore interface {

	// websites
	GetWebsites() ([]*website.Website, error)
	GetWebsite(id int64) (*website.Website, error)
	SaveWebsite(w *website.Website) error
	DeleteWebsite(w *website.Website) error

	Close() error
}

// New returns a new Datastore
func New() (Datastore, error) {

	var store Datastore
	var err error

	store, err = sqlite.New()

	return store, err
}
