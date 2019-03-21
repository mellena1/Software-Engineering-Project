package api

import (
	"encoding/json"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// countAPI holds all of the api functions related to Counts and all of the variables needed to access the backend
type countAPI struct {
	countReader  db.CountReader
	countWriter  db.CountWriter
	countUpdater db.CountUpdater
	countDeleter db.CountDeleter
}

// CreateCountRoutes makes all of the routes for room related calls
func CreateCountRoutes(countDBFacade db.CountReaderWriterUpdaterDeleter) []Route {
	cAPI := countAPI{
		countReader:  countDBFacade,
		countWriter:  countDBFacade,
		countUpdater: countDBFacade,
		countDeleter: countDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/counts", cAPI.getAllCounts, "GET"),
		NewRoute("/api/v1/count", cAPI.getACount, "GET"),
		// NewRoute("/api/v1/count", cAPI.writeACount, "POST"),
		// NewRoute("/api/v1/count", cAPI.updateACount, "PUT"),
		// NewRoute("/api/v1/count", cAPI.deleteACount, "DELETE"),
	}

	return routes
}

// getAllCounts Get all counts from the db
// @Summary Get all counts from the db
// @Description Get all counts from the db
// @produce json
// @Success 200 {array} db.Count "all counts in the db"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/counts [get]
func (c countAPI) getAllCounts(w http.ResponseWriter, r *http.Request) {
	counts, err := c.countReader.ReadAllCounts()
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	responseJSON, _ := json.Marshal(counts)
	w.Write(responseJSON)
}

// getACount Get a count from the db
// @Summary Get a count from the db
// @Description Get a count from the db
// @produce json
// @param id query int true "the session of the count to retrieve"
// @Success 200 {array} db.Count "the requested count (beginning/middle/end)"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/count [get]
func (c countAPI) getACount(w http.ResponseWriter, r *http.Request) {
	requestedID, err := getIDFromQueries(r)
	switch err {
	case ErrQueryNotSet:
		ReportError(ErrQueryNotSet, "the \"id\" param was not set", http.StatusBadRequest, w)
		return
	case ErrBadQuery:
		ReportError(ErrBadQuery, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, w)
		return
	case ErrBadQueryType:
		ReportError(ErrBadQueryType, "the \"id\" param is not a number", http.StatusBadRequest, w)
		return
	}

	count, err := c.countReader.ReadACount(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	responseJSON, _ := json.Marshal(count)
	w.Write(responseJSON)
}
