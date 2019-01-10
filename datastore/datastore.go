package datastore

import (
	"github.com/spf13/viper"
	"michaelvanolst.nl/scraper/datastore/sqlite"
	"michaelvanolst.nl/scraper/website"
)

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

	if viper.Get("database.driver") == "sqlite3" {
		store, err = sqlite.New()
	}

	return store, err
}
