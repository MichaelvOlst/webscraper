package api

import (
	"net/http"
)

// GET /api/sites
func (api *API) getWebsites(w http.ResponseWriter, r *http.Request) error {
	// result, err := api.database.GetSites()
	// if err != nil {
	// 	return err
	// }
	result := "websites"
	return respond(w, http.StatusOK, envelope{Data: result})
}
