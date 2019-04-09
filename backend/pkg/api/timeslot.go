package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

type timeslotAPI struct {
	timeslotReader  db.TimeslotReader
	timeslotWriter  db.TimeslotWriter
	timeslotUpdater db.TimeslotUpdater
	timeslotDeleter db.TimeslotDeleter
}

func CreateTimeslotRoutes(timeslotDBFacade db.TimeslotReaderWriterUpdaterDeleter) []Route {
	myTimeslotAPI := timeslotAPI{
		timeslotReader:  timeslotDBFacade,
		timeslotWriter:  timeslotDBFacade,
		timeslotUpdater: timeslotDBFacade,
		timeslotDeleter: timeslotDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/timeslots", myTimeslotAPI.getAllTimeslots, "GET"),
		NewRoute("/api/v1/timeslot", myTimeslotAPI.getATimeslot, "GET"),
		NewRoute("/api/v1/timeslot", myTimeslotAPI.writeATimeslot, "POST"),
		NewRoute("/api/v1/timeslot", myTimeslotAPI.updateATimeslot, "PUT"),
		NewRoute("/api/v1/timeslot", myTimeslotAPI.deleteATimeslot, "DELETE"),
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
func (myTimeslotAPI timeslotAPI) getAllTimeslots(writer http.ResponseWriter, request *http.Request) {
	timeslots, err := myTimeslotAPI.timeslotReader.ReadAllTimeslots()
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	responseJSON, _ := json.Marshal(timeslots)
	writer.Write(responseJSON)
}

// getATimeslot Get a timeslot with a specific timeslotID from the db
// @Summary Get a timeslot with a specific timeslotID from the db
// @Description Get a timeslot from the db
// @produce json
// @param id query int true "the timeslot to retrieve"
// @Success 200 {object} db.Timeslot "the requested timeslot"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/timeslot [get]
func (myTimeslotAPI timeslotAPI) getATimeslot(writer http.ResponseWriter, request *http.Request) {
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

	timeslot, err := myTimeslotAPI.timeslotReader.ReadATimeslot(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	responseJSON, _ := json.Marshal(timeslot)
	writer.Write(responseJSON)
}

type WriteATimeslotRequest struct {
	StartTime *string `json:"startTime" example:"2019-02-18T21:00:00Z"`
	EndTime   *string `json:"endTime" example:"2019-10-01T23:00:00Z"`
}

func (request WriteATimeslotRequest) Validate() error {
	if request.StartTime == nil || request.EndTime == nil {
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
func (myTimeslotAPI timeslotAPI) writeATimeslot(writer http.ResponseWriter, request *http.Request) {
	requestJSON, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	timeslotRequest := WriteATimeslotRequest{}
	err = json.Unmarshal(requestJSON, &timeslotRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, writer)
		return
	}

	if err = timeslotRequest.Validate(); err != nil {
		ReportError(err, "must set both startTime and endTime", http.StatusBadRequest, writer)
		return
	}

	startTime, err := time.Parse(db.TimeFormat, *timeslotRequest.StartTime)
	if err != nil {
		msg := fmt.Sprintf("start time invalid. please use the format %s", db.TimeFormat)
		ReportError(err, msg, http.StatusBadRequest, writer)
		return
	}

	endTime, err := time.Parse(db.TimeFormat, *timeslotRequest.EndTime)
	if err != nil {
		msg := fmt.Sprintf("end time invalid. please use the format %s", db.TimeFormat)
		ReportError(err, msg, http.StatusBadRequest, writer)
		return
	}

	id, err := myTimeslotAPI.timeslotWriter.WriteATimeslot(startTime, endTime)
	if err != nil {
		ReportError(err, "failed to write a room", http.StatusServiceUnavailable, writer)
		return
	}

	writeIDToClient(writer, id)
}

type UpdateATimeslotRequest struct {
	ID        *int64  `json:"id" example:"1"`
	StartTime *string `json:"startTime" example:"2019-02-18T21:00:00Z"`
	EndTime   *string `json:"endTime" example:"2019-10-01T23:00:00Z"`
}

func (request UpdateATimeslotRequest) Validate() error {
	if request.ID == nil {
		return ErrInvalidRequest
	}
	if request.StartTime == nil || request.EndTime == nil {
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
func (myTimeslotAPI timeslotAPI) updateATimeslot(writer http.ResponseWriter, request *http.Request) {
	requestJSON, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	timeslotRequest := UpdateATimeslotRequest{}
	err = json.Unmarshal(requestJSON, &timeslotRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, writer)
		return
	}

	if err = timeslotRequest.Validate(); err != nil {
		ReportError(err, "must set both startTime and endTime", http.StatusBadRequest, writer)
		return
	}

	startTime, err := time.Parse(db.TimeFormat, *timeslotRequest.StartTime)
	if err != nil {
		msg := fmt.Sprintf("start time invalid. please use the format %s", db.TimeFormat)
		ReportError(err, msg, http.StatusBadRequest, writer)
		return
	}

	endTime, err := time.Parse(db.TimeFormat, *timeslotRequest.EndTime)
	if err != nil {
		msg := fmt.Sprintf("end time invalid. please use the format %s", db.TimeFormat)
		ReportError(err, msg, http.StatusBadRequest, writer)
		return
	}

	err = myTimeslotAPI.timeslotUpdater.UpdateATimeslot(*timeslotRequest.ID, startTime, endTime)
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
func (myTimeslotAPI timeslotAPI) deleteATimeslot(writer http.ResponseWriter, request *http.Request) {
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

	err = myTimeslotAPI.timeslotDeleter.DeleteATimeslot(requestedID)
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
