package api

import (
	"net/http"
	"strconv"

	"michaelvanolst.nl/scraper/models"

	"github.com/gorilla/mux"
)

// GET /api/websites
func (api *API) getLinksHandler(w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	sid, ok := vars["id"]
	var result []*models.Link
	var err error
	if ok {
		id, err := strconv.ParseInt(sid, 10, 64)
		if err != nil {
			return err
		}

		result, err = api.database.GetLinksByID(id)
		if err != nil {
			return err
		}
	} else {
		result, err = api.database.GetAllLinks()
		if err != nil {
			return err
		}
	}

	return respond(w, http.StatusOK, envelope{Result: result})
}
