package api

import (
	"encoding/json"
	"io/ioutil"
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
		NewRoute("/api/v1/count", cAPI.getCountsOfSession, "GET"),
		NewRoute("/api/v1/count", cAPI.writeACount, "POST"),
		NewRoute("/api/v1/count", cAPI.updateACount, "PUT"),
		NewRoute("/api/v1/count", cAPI.deleteACount, "DELETE"),
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
func (c countAPI) getCountsOfSession(w http.ResponseWriter, r *http.Request) {
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

	count, err := c.countReader.ReadCountsOfSession(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	responseJSON, _ := json.Marshal(count)
	w.Write(responseJSON)
}

// writeACountRequest request for writeACount
type writeACountRequest struct {
	Time      *string `json:"time" example:"beginning"`
	SessionID *int64  `json:"sessionID" example:"2"`
	UserID    *int64  `json:"userID" example:"1"`
	Count     *int64  `json:"count" example:"30"`
}

// Validate validates a writeACountRequest
func (r writeACountRequest) Validate() error {
	if r.SessionID == nil || r.Time == nil {
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	countRequest := writeACountRequest{}
	err = json.Unmarshal(body, &countRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, w)
		return
	}

	if err = countRequest.Validate(); err != nil {
		ReportError(err, "time and session must be set", http.StatusBadRequest, w)
		return
	}

	id, err := c.countWriter.WriteACount(countRequest.Time, countRequest.SessionID, countRequest.UserID, countRequest.Count)
	if err != nil {
		ReportError(err, "failed to write a count", http.StatusServiceUnavailable, w)
		return
	}

	writeIDToClient(w, id)
}

// updateACountRequest request for updateACount
type updateACountRequest struct {
	Time      *string `json:"time" example:"beginning"`
	SessionID *int64  `json:"sessionID" example:"2"`
	UserID    *int64  `json:"userID" example:"1"`
	Count     *int64  `json:"count" example:"30"`
}

// Validate validates a updateACountRequest
func (r updateACountRequest) Validate() error {
	if r.SessionID == nil || r.Time == nil {
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	updateRequest := updateACountRequest{}
	err = json.Unmarshal(body, &updateRequest)
	if err != nil {
		ReportError(err, "failed to unmarshal json", http.StatusBadRequest, w)
		return
	}

	if err = updateRequest.Validate(); err != nil {
		ReportError(err, "must set time and session for count you wish to update", http.StatusBadRequest, w)
		return
	}

	err = c.countUpdater.UpdateACount(updateRequest.Time, updateRequest.SessionID, updateRequest.UserID, updateRequest.Count)
	switch err {
	case nil:
		w.Write(nil)
		return
	case db.ErrNothingChanged:
		ReportError(err, "nothing in the db was changed. id probably does not exist", http.StatusBadRequest, w)
		return
	default:
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}
}

// deleteACount Delete an existing count in the db
// @Summary Delete an existing count in the db
// @Description Delete an existing session's counts in the db
// @accept json
// @produce json
// @param id query int true "the session of counts to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/count [delete]
func (c countAPI) deleteACount(w http.ResponseWriter, r *http.Request) {
	requestedID, err := getIDFromQueries(r)
	switch err {
	case ErrQueryNotSet:
		ReportError(err, "the \"sessionID\" param was not set", http.StatusBadRequest, w)
		return
	case ErrBadQuery:
		ReportError(err, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, w)
		return
	case ErrBadQueryType:
		ReportError(err, "the \"sessionID\" param is not a number", http.StatusBadRequest, w)
		return
	}

	err = c.countDeleter.DeleteACount(requestedID)
	switch err {
	case nil:
		w.Write(nil)
		return
	case db.ErrNothingChanged:
		ReportError(err, "nothing in the db was changed. id probably does not exist", http.StatusBadRequest, w)
		return
	default:
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}
}
