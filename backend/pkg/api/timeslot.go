package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// timeslotAPI holds all of the api functions related to Timeslots and all of the variables needed to access the backend
type timeslotAPI struct {
	timeslotReader  db.TimeslotReader
	timeslotWriter  db.TimeslotWriter
	timeslotUpdater db.TimeslotUpdater
	timeslotDeleter db.TimeslotDeleter
}

// CreateTimeslotRoutes makes all of the routes for room related calls
func CreateTimeslotRoutes(timeslotDBFacade db.TimeslotReaderWriterUpdaterDeleter) []Route {
	tAPI := timeslotAPI{
		timeslotReader:  timeslotDBFacade,
		timeslotWriter:  timeslotDBFacade,
		timeslotUpdater: timeslotDBFacade,
		timeslotDeleter: timeslotDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/timeslots", tAPI.getAllTimeslots, "GET"),
		NewRoute("/api/v1/timeslot", tAPI.getATimeslot, "GET"),
		NewRoute("/api/v1/timeslot", tAPI.writeATimeslot, "POST"),
		NewRoute("/api/v1/timeslot", tAPI.updateATimeslot, "PUT"),
		NewRoute("/api/v1/timeslot", tAPI.deleteATimeslot, "DELETE"),
	}

	return routes
}

// getAllTimeslots Get all timeslots from the db
// @Summary Get all timeslots from the db
// @Description Get all timeslots from the db
// @produce json
// @Success 200 {array} db.Timeslot "all timeslots in the db"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/timeslots [get]
func (t timeslotAPI) getAllTimeslots(w http.ResponseWriter, r *http.Request) {
	timeslots, err := t.timeslotReader.ReadAllTimeslots()
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	responseJSON, _ := json.Marshal(timeslots)
	w.Write(responseJSON)
}

// getATimeslot Get a timeslot from the db
// @Summary Get a timeslot from the db
// @Description Get a timeslot from the db
// @produce json
// @param id query int true "the timeslot to retrieve"
// @Success 200 {object} db.Timeslot "the requested timeslot"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/timeslot [get]
func (t timeslotAPI) getATimeslot(w http.ResponseWriter, r *http.Request) {
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

	timeslot, err := t.timeslotReader.ReadATimeslot(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	responseJSON, _ := json.Marshal(timeslot)
	w.Write(responseJSON)
}

// WriteATimeslotRequest request for writeATimeslot
type WriteATimeslotRequest struct {
	StartTime *string `json:"startTime" example:"2019-02-18 21:00:00"`
	EndTime   *string `json:"endTime" example:"2019-10-01 23:00:00"`
}

// Validate validates a WriteATimeslotRequest
func (r WriteATimeslotRequest) Validate() error {
	if r.StartTime == nil || r.EndTime == nil {
		return ErrInvalidRequest
	}
	return nil
}

// writeATimeslot Add a timeslot to the db
// @Summary Add a timeslot
// @Description Add a timeslot to the db
// @accept json
// @produce json
// @param timeslot body api.WriteATimeslotRequest true "the timeslot to add"
// @Success 200 {integer} int64 "the id of the timeslot added"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/timeslot [post]
func (t timeslotAPI) writeATimeslot(w http.ResponseWriter, r *http.Request) {
	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	timeslotRequest := WriteATimeslotRequest{}
	err = json.Unmarshal(j, &timeslotRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, w)
		return
	}

	if err = timeslotRequest.Validate(); err != nil {
		ReportError(err, "must set both startTime and endTime", http.StatusBadRequest, w)
		return
	}

	startTime, err := time.Parse(db.TimeFormat, *timeslotRequest.StartTime)
	if err != nil {
		msg := fmt.Sprintf("start time invalid. please use the format %s", db.TimeFormat)
		ReportError(err, msg, http.StatusBadRequest, w)
		return
	}

	endTime, err := time.Parse(db.TimeFormat, *timeslotRequest.EndTime)
	if err != nil {
		msg := fmt.Sprintf("end time invalid. please use the format %s", db.TimeFormat)
		ReportError(err, msg, http.StatusBadRequest, w)
		return
	}

	id, err := t.timeslotWriter.WriteATimeslot(startTime, endTime)
	if err != nil {
		ReportError(err, "failed to write a room", http.StatusServiceUnavailable, w)
		return
	}

	writeIDToClient(w, id)
}

// UpdateATimeslotRequest request for updateATimeslot
type UpdateATimeslotRequest struct {
	ID        *int64  `json:"id" example:"1"`
	StartTime *string `json:"startTime" example:"2019-02-18 21:00:00"`
	EndTime   *string `json:"endTime" example:"2019-10-01 23:00:00"`
}

// Validate validates a UpdateATimeslotRequest
func (r UpdateATimeslotRequest) Validate() error {
	if r.ID == nil {
		return ErrInvalidRequest
	}
	if r.StartTime == nil || r.EndTime == nil {
		return ErrInvalidRequest
	}
	return nil
}

// updateATimeslot Update an existing timeslot in the db
// @Summary Update an existing timeslot in the db
// @Description Update an existing timeslot in the db
// @accept json
// @produce json
// @param timeslot body api.UpdateATimeslotRequest true "the timeslot to update with the new values"
// @Success 200 "Updated properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/timeslot [put]
func (t timeslotAPI) updateATimeslot(w http.ResponseWriter, r *http.Request) {
	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	timeslotRequest := UpdateATimeslotRequest{}
	err = json.Unmarshal(j, &timeslotRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, w)
		return
	}

	if err = timeslotRequest.Validate(); err != nil {
		ReportError(err, "must set both startTime and endTime", http.StatusBadRequest, w)
		return
	}

	startTime, err := time.Parse(db.TimeFormat, *timeslotRequest.StartTime)
	if err != nil {
		msg := fmt.Sprintf("start time invalid. please use the format %s", db.TimeFormat)
		ReportError(err, msg, http.StatusBadRequest, w)
		return
	}

	endTime, err := time.Parse(db.TimeFormat, *timeslotRequest.EndTime)
	if err != nil {
		msg := fmt.Sprintf("end time invalid. please use the format %s", db.TimeFormat)
		ReportError(err, msg, http.StatusBadRequest, w)
		return
	}

	err = t.timeslotUpdater.UpdateATimeslot(*timeslotRequest.ID, startTime, endTime)
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

// deleteATimeslotRequest request for deleteATimeslot
type deleteATimeslotRequest struct {
	ID int64 `json:"id" example:"1"`
}

// deleteATimeslot Delete an existing timeslot in the db
// @Summary Delete an existing timeslot in the db
// @Description Delete an existing timeslot in the db
// @accept json
// @produce json
// @param id query int true "the timeslot to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/timeslot [delete]
func (t timeslotAPI) deleteATimeslot(w http.ResponseWriter, r *http.Request) {
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

	err = t.timeslotDeleter.DeleteATimeslot(requestedID)
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
