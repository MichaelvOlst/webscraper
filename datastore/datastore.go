package datastore

import (
	"michaelvanolst.nl/scraper/datastore/sqlite"
	"michaelvanolst.nl/scraper/models"
)

// Datastore represents a database implementations
type Datastore interface {

	// Websites
	GetWebsites() ([]*models.Website, error)
	GetWebsite(id int64) (*models.Website, error)
	SaveWebsite(w *models.Website) error
	DeleteWebsite(w *models.Website) error

	// Attributes
	GetAttributes(websiteID int64) ([]*models.Attribute, error)
	SaveAttribute(a *models.Attribute) error
	DeleteAttribute(a *models.Attribute) error

	// Links
	GetAllLinks() ([]*models.Link, error)
	GetLinksByID(id int64) ([]*models.Link, error)
	SaveLink(l *models.Link) error
	CheckLinkExists(url string) (bool, error)

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
