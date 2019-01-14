package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes ...
func (api *API) Routes() *mux.Router {

	r := mux.NewRouter()

	r.Handle("/api/websites", HandlerFunc(api.getWebsites)).Methods(http.MethodGet)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	return r
}
