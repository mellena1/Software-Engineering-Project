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

// writeACountRequest request for writeACount
type writeACountRequest struct {
	StartTime *string `json:"startTime" example:"2019-02-18T21:00:00Z"`
	EndTime   *string `json:"endTime" example:"2019-10-01T23:00:00Z"`
}

// Validate validates a writeACountRequest
func (r writeACountRequest) Validate() error {
	if r.StartTime == nil || r.EndTime == nil {
		return ErrInvalidRequest
	}
	return nil
}

// writeACount Add a count to the db
// @Summary Add a count
// @Description Add a count to the db
// @accept json
// @produce json
// @param count body api.writeACountRequest true "the count to add"
// @Success 200 {integer} int64 "the id of the session count was added to"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/count [post]
func (c countAPI) writeACount(w http.ResponseWriter, r *http.Request) {

}

// updateACountRequest request for updateACount
type updateACountRequest struct {
	ID        *int64  `json:"id" example:"1"`
	StartTime *string `json:"startTime" example:"2019-02-18T21:00:00Z"`
	EndTime   *string `json:"endTime" example:"2019-10-01T23:00:00Z"`
}

// Validate validates a updateACountRequest
func (r updateACountRequest) Validate() error {
	if r.ID == nil {
		return ErrInvalidRequest
	}
	if r.StartTime == nil || r.EndTime == nil {
		return ErrInvalidRequest
	}
	return nil
}

// updateACount Update an existing count in the db
// @Summary Update an existing count in the db
// @Description Update an existing count in the db
// @accept json
// @produce json
// @param count body api.updateACountRequest true "the count to update with the new values"
// @Success 200 "Updated properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/count [put]
func (c countAPI) updateACount(w http.ResponseWriter, r *http.Request) {

}

// deleteACountRequest request for deleteATimeslot
type deleteACountRequest struct {
	ID int64 `json:"id" example:"1"`
}

// deleteACount Delete an existing count in the db
// @Summary Delete an existing count in the db
// @Description Delete an existing count in the db
// @accept json
// @produce json
// @param id query int true "the count to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/count [delete]
func (t timeslotAPI) deleteACount(w http.ResponseWriter, r *http.Request) {

}
