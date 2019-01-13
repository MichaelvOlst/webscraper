package api

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"michaelvanolst.nl/scraper/datastore"
	"michaelvanolst.nl/scraper/website"
)

// API ...
type API struct {
	database datastore.Datastore
}

// New instantiates a new API object
func New(db datastore.Datastore) *API {
	return &API{
		database: db,
	}
}

// Start the api
func (a *API) Start() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("public", true)))
	api := r.Group("/api")
	website.RegisterRoutes(api.Group("/websites"))

	r.Run()
}
