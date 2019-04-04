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
	myCountAPI := countAPI{
		countReader:  countDBFacade,
		countWriter:  countDBFacade,
		countUpdater: countDBFacade,
		countDeleter: countDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/counts", myCountAPI.getAllCounts, "GET"),
		NewRoute("/api/v1/count", myCountAPI.getCountsOfSession, "GET"),
		NewRoute("/api/v1/count", myCountAPI.writeACount, "POST"),
		NewRoute("/api/v1/count", myCountAPI.updateACount, "PUT"),
		NewRoute("/api/v1/count", myCountAPI.deleteACount, "DELETE"),
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
func (myCountAPI countAPI) getAllCounts(writer http.ResponseWriter, request *http.Request) {
	counts, err := myCountAPI.countReader.ReadAllCounts()
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	responseJSON, _ := json.Marshal(counts)
	writer.Write(responseJSON)
}

// getACount Get the beginning, middle, and end counts for a specified session from the db
// @Summary Get the beginning, middle, and end counts for a specified session from the db
// @Description Get the beginning, middle, and end counts for a specified session from the db
// @produce json
// @param id query int true "the session of the count to retrieve"
// @Success 200 {array} db.Count "the requested count (beginning/middle/end)"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/count [get]
func (myCountAPI countAPI) getCountsOfSession(writer http.ResponseWriter, request *http.Request) {
	requestedID, err := getIDFromQueries(request)
	switch err {
	case ErrQueryNotSet:
		ReportError(ErrQueryNotSet, "the \"id\" param was not set", http.StatusBadRequest, writer)
		return
	case ErrBadQuery:
		ReportError(ErrBadQuery, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, writer)
		return
	case ErrBadQueryType:
		ReportError(ErrBadQueryType, "the \"id\" param is not a number", http.StatusBadRequest, writer)
		return
	}

	count, err := myCountAPI.countReader.ReadCountsOfSession(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	responseJSON, _ := json.Marshal(count)
	writer.Write(responseJSON)
}

type writeACountRequest struct {
	Time      *string `json:"time" example:"beginning"`
	SessionID *int64  `json:"sessionID" example:"2"`
	UserName  *string `json:"userName" example:"Kenny Robinson"`
	Count     *int64  `json:"count" example:"30"`
}

func (request writeACountRequest) Validate() error {
	if request.SessionID == nil || request.Time == nil {
		return ErrInvalidRequest
	}
	return nil
}

// writeACount Add a count to the db
// @Summary Add a count to the db
// @Description Add a count to the db
// @accept json
// @produce json
// @param count body api.writeACountRequest true "the count to add"
// @Success 200 {integer} int64 "the id of the session count was added to"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/count [post]
func (myCountAPI countAPI) writeACount(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	countRequest := writeACountRequest{}
	err = json.Unmarshal(body, &countRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, writer)
		return
	}

	if err = countRequest.Validate(); err != nil {
		ReportError(err, "time and session must be set", http.StatusBadRequest, writer)
		return
	}

	id, err := myCountAPI.countWriter.WriteACount(countRequest.Time, countRequest.SessionID, countRequest.UserName, countRequest.Count)
	if err != nil {
		ReportError(err, "failed to write a count", http.StatusServiceUnavailable, writer)
		return
	}

	writeIDToClient(writer, id)
}

type updateACountRequest struct {
	Time      *string `json:"time" example:"beginning"`
	SessionID *int64  `json:"sessionID" example:"2"`
	UserName  *string `json:"userName" example:"Kenny Robinson"`
	Count     *int64  `json:"count" example:"30"`
}

func (request updateACountRequest) Validate() error {
	if request.SessionID == nil || request.Time == nil {
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
func (myCountAPI countAPI) updateACount(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	updateRequest := updateACountRequest{}
	err = json.Unmarshal(body, &updateRequest)
	if err != nil {
		ReportError(err, "failed to unmarshal json", http.StatusBadRequest, writer)
		return
	}

	if err = updateRequest.Validate(); err != nil {
		ReportError(err, "must set time and session for count you wish to update", http.StatusBadRequest, writer)
		return
	}

	err = myCountAPI.countUpdater.UpdateACount(updateRequest.Time, updateRequest.SessionID, updateRequest.UserName, updateRequest.Count)
	switch err {
	case nil:
		writer.Write(nil)
		return
	case db.ErrNothingChanged:
		ReportError(err, "nothing in the db was changed. id probably does not exist", http.StatusBadRequest, writer)
		return
	default:
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}
}

// deleteACount Delete a sessions existing counts from the db given a sessionID
// @Summary Delete a sessions existing counts from the db given a sessionID
// @Description Delete a sessions existing counts from the db given a sessionID
// @accept json
// @produce json
// @param id query int true "the session of counts to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/count [delete]
func (myCountAPI countAPI) deleteACount(writer http.ResponseWriter, request *http.Request) {
	requestedID, err := getIDFromQueries(request)
	switch err {
	case ErrQueryNotSet:
		ReportError(err, "the \"sessionID\" param was not set", http.StatusBadRequest, writer)
		return
	case ErrBadQuery:
		ReportError(err, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, writer)
		return
	case ErrBadQueryType:
		ReportError(err, "the \"sessionID\" param is not a number", http.StatusBadRequest, writer)
		return
	}

	err = myCountAPI.countDeleter.DeleteACount(requestedID)
	switch err {
	case nil:
		writer.Write(nil)
		return
	case db.ErrNothingChanged:
		ReportError(err, "nothing in the db was changed. id probably does not exist", http.StatusBadRequest, writer)
		return
	default:
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}
}
