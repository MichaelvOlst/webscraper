package datastore

import (
	"michaelvanolst.nl/scraper/datastore/sqlite"
	"michaelvanolst.nl/scraper/websites"
)

// Datastore represents a database implementations
type Datastore interface {

	// Websites
	GetWebsites() ([]*websites.Website, error)
	GetWebsite(id int64) (*websites.Website, error)
	SaveWebsite(w *websites.Website) error
	DeleteWebsite(w *websites.Website) error

	// Attributes
	GetAttributes(websiteID int64) ([]*websites.Attribute, error)
	GetAttribute(websiteID, id int64) (*websites.Attribute, error)
	SaveAttribute(a *websites.Attribute) error
	DeleteAttribute(a *websites.Attribute) error

	Close() error
}

// New returns a new Datastore
func New(c *Config) (Datastore, error) {

	var store Datastore
	var err error
	if c.Driver == "sqlite3" {
		store, err = sqlite.New(c.Name)
	}

	return store, err
}

// Config holds the database config
type Config struct {
	Driver   string `default:"sqlite3"`
	Host     string `default:""`
	User     string `default:""`
	Password string `default:""`
	Name     string `default:"scraper.db"`
	SSLMode  string `default:""`
}
