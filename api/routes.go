package api

import (
	"net/http"
	"text/template"

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

	r.HandleFunc("/email", func(w http.ResponseWriter, r *http.Request) {

		td := struct {
			Address  string
			URL      string
			ImageURL string
			Price    string
		}{
			Address:  "Gooierserf 292",
			URL:      "https://www.floberg.nl/aanbod/woningaanbod/huizen/koop/huis-4449832-Gooierserf-292/",
			ImageURL: "https://images.realworks.nl/servlets/images/media.objectmedia/74094533.jpg?height=280&check=sha256%3A8c253f568761741af0bd5bd19d5a5d898bbfbc9221cc87e607eba5377319a7ee&width=420&resize=5",
			Price:    "€ 200.000,- k.k.",
		}

		tmpl := template.Must(template.ParseFiles("public/email.html"))
		tmpl.Execute(w, td)
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	return r
}
