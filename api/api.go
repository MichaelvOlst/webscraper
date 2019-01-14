package api

import (
	"michaelvanolst.nl/scraper/datastore"
)

// API ...
type API struct {
	database datastore.Datastore
}

// Config holds the config for the api
type Config struct {
	Host string `default:"127.0.0.1"`
	Port string `default:"8080"`
}

// New instantiates a new API object
func New(db datastore.Datastore) *API {
	return &API{
		database: db,
	}
}
