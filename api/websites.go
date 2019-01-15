package api

import (
	"net/http"
)

// GET /api/sites
func (api *API) getWebsites(w http.ResponseWriter, r *http.Request) error {
	result, err := api.database.GetWebsites()
	if err != nil {
		return err
	}
	return respond(w, http.StatusOK, envelope{Result: result})
}
