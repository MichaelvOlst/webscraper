package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"michaelvanolst.nl/scraper/websites"
)

// GET /api/websites
func (api *API) getWebsitesHandler(w http.ResponseWriter, r *http.Request) error {
	result, err := api.database.GetWebsites()
	if err != nil {
		return err
	}
	return respond(w, http.StatusOK, envelope{Result: result})
}

// GET /api/websites/{id:[0-9]+}
func (api *API) getWebsiteHandler(w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return err
	}

	result, err := api.database.GetWebsite(id)
	if err != nil {
		return err
	}

	return respond(w, http.StatusOK, envelope{Result: result})
}

// POST /api/websites
func (api *API) saveWebsitesHandler(w http.ResponseWriter, r *http.Request) error {
	var s websites.Website

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		return err
	}

	if err := api.database.SaveWebsite(&s); err != nil {
		return err
	}

	return respond(w, http.StatusOK, envelope{Result: s})
}

// DELETE /api/websites
func (api *API) deleteWebsitesHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return err
	}

	if err := api.database.DeleteWebsite(&websites.Website{ID: id}); err != nil {
		return err
	}

	return respond(w, http.StatusOK, envelope{Result: true})
}
