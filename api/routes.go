package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes returns the routes for the api
func (api *API) Routes() *mux.Router {

	r := mux.NewRouter()

	r.Handle("/api/websites", HandlerFunc(api.getWebsitesHandler)).Methods(http.MethodGet)
	r.Handle("/api/websites", HandlerFunc(api.saveWebsitesHandler)).Methods(http.MethodPost)
	r.Handle("/api/websites/{id:[0-9]+}", HandlerFunc(api.deleteWebsitesHandler)).Methods(http.MethodDelete)
	r.Handle("/api/websites/{id:[0-9]+}", HandlerFunc(api.getWebsiteHandler)).Methods(http.MethodGet)

	r.Handle("/api/links", HandlerFunc(api.getLinksHandler)).Methods(http.MethodGet)
	r.Handle("/api/links/{id:[0-9]+}", HandlerFunc(api.getLinksHandler)).Methods(http.MethodGet)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	return r
}
