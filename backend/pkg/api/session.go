package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// sessionAPI holds all of the api functions related to Sessions and all of the variables needed to access the backend
type sessionAPI struct {
	sessionReader  db.SessionReader
	sessionWriter  db.SessionWriter
	sessionUpdater db.SessionUpdater
	sessionDeleter db.SessionDeleter
}

// CreateSessionRoutes makes all of the routes for session related calls
func CreateSessionRoutes(sessionDBFacade db.SessionReaderWriterUpdaterDeleter) []Route {
	mySessionAPI := sessionAPI{
		sessionReader:  sessionDBFacade,
		sessionWriter:  sessionDBFacade,
		sessionUpdater: sessionDBFacade,
		sessionDeleter: sessionDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/session", mySessionAPI.getASession, "GET"),
		NewRoute("/api/v1/sessions", mySessionAPI.getAllSessions, "GET"),
		NewRoute("/api/v1/session", mySessionAPI.writeASession, "POST"),
		NewRoute("/api/v1/session", mySessionAPI.updateASession, "PUT"),
		NewRoute("/api/v1/session", mySessionAPI.deleteASession, "DELETE"),
		NewRoute("/api/v1/sessionsBySpeaker", mySessionAPI.getAllSessionsBySpeaker, "GET"),
		NewRoute("/api/v1/sessionsByTimeslot", mySessionAPI.getAllSessionsByTimeslot, "GET"),
	}

	return routes
}

// getASession Gets a session from the db given a specified sessionID
// @Summary Gets a session from the db given a specified sessionID
// @Description Gets a session from the db given a specified sessionID
// @Produce json
// @param id query int true "the session to retrieve"
// @Success 200 {object} db.Session
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/session [get]
func (mySessionAPI sessionAPI) getASession(writer http.ResponseWriter, request *http.Request) {
	requestedID, err := getIDFromQueries(request)
	switch err {
	case ErrQueryNotSet:
		ReportError(err, "the \"id\" param was not set", http.StatusBadRequest, writer)
		return
	case ErrBadQuery:
		ReportError(err, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, writer)
		return
	case ErrBadQueryType:
		ReportError(err, "the \"id\" param is not a number", http.StatusBadRequest, writer)
		return
	}

	session, err := mySessionAPI.sessionReader.ReadASession(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	responseJSON, _ := json.Marshal(session)
	writer.Write(responseJSON)
}

// getAllSessions Gets all sessions from the db
// @Summary Gets all sessions from the db
// @Description Return a list of all sessions
// @Produce json
// @Success 200 {array} db.Session
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/sessions [get]
func (mySessionAPI sessionAPI) getAllSessions(writer http.ResponseWriter, request *http.Request) {
	sessions, err := mySessionAPI.sessionReader.ReadAllSessions()
	if err != nil {
		ReportError(err, "Failed to get all sessions", http.StatusBadRequest, writer)
		return
	}

	responseJSON, _ := json.Marshal(sessions)
	writer.Write(responseJSON)
}

// WriteASessionRequest request for writeASession
type WriteASessionRequest struct {
	SpeakerID   *int64  `json:"speakerID" example:"1"`
	RoomID      *int64  `json:"roomID" example:"1"`
	TimeslotID  *int64  `json:"timeslotID" example:"1"`
	SessionName *string `json:"sessionName" example:"Microservices"`
}

// Validate validates a WriteASessionRequest
func (request WriteASessionRequest) Validate() error {
	if request.SpeakerID == nil && request.RoomID == nil && request.TimeslotID == nil && (request.SessionName == nil || *request.SessionName == "") {
		return ErrInvalidRequest
	}
	return nil
}

// @Summary Get all sessions, sorted by speaker
// @Description Return a list of all sessions sorted by speaker
// @Produce json
// @Success 200 {array} string "An array of speakerName: Sessions map entries"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/sessionsBySpeaker [get]
func (mySessionAPI sessionAPI) getAllSessionsBySpeaker(writer http.ResponseWriter, request *http.Request) {
	sessions, err := mySessionAPI.sessionReader.ReadAllSessions()
	if err != nil {
		ReportError(err, "Failed to get all session", http.StatusBadRequest, writer)
		return
	}
	sessionsBySpeaker := make(map[string][]db.Session)
	for _, session := range sessions {
		key := *session.Speaker.FirstName + " " + *session.Speaker.LastName
		sessionsBySpeaker[key] = append(sessionsBySpeaker[key], session)
	}

	responseJSON, err := json.Marshal(sessionsBySpeaker)
	if err != nil {
		ReportError(err, "Failed to marshal data", http.StatusBadRequest, writer)
		return
	}
	writer.Write(responseJSON)
}

// getAllSessions Gets all sessions from the db, sorted by timeslot
// @Summary Get all sessions, sorted by timeslot
// @Description Return a list of all sessions, sorted by timeslot
// @Produce json
// @Success 200 {array} string "An array of timeslotString: Sessions map entries"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/sessionsByTimeslot [get]
func (mySessionAPI sessionAPI) getAllSessionsByTimeslot(writer http.ResponseWriter, request *http.Request) {
	sessions, err := mySessionAPI.sessionReader.ReadAllSessions()
	if err != nil {
		ReportError(err, "Failed to get all session", http.StatusBadRequest, writer)
		return
	}

	// Making a map of sessions to timeslots, where the key is just the time of the session (i.e 12:00 to 1:00)
	sessionsByTimeslot := make(map[string][]db.Session)
	for _, session := range sessions {
		key := session.Timeslot.StartTime.Format(time.Kitchen) + "-" + session.Timeslot.EndTime.Format(time.Kitchen)
		sessionsByTimeslot[key] = append(sessionsByTimeslot[key], session)
	}

	responseJSON, err := json.Marshal(sessionsByTimeslot)
	if err != nil {
		ReportError(err, "Failed to marshal data", http.StatusBadRequest, writer)
		return
	}
	writer.Write(responseJSON)
}

// writeASession Add a session to the db
// @Summary Add a session
// @Description Add a session to the db
// @accept json
// @produce json
// @param session body api.WriteASessionRequest true "the session to add"
// @Success 200 {} int "the id of the session added"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/session [post]
func (mySessionAPI sessionAPI) writeASession(writer http.ResponseWriter, request *http.Request) {
	requestJSON, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	sessionRequest := WriteASessionRequest{}
	err = json.Unmarshal(requestJSON, &sessionRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, writer)
		return
	}

	if err = sessionRequest.Validate(); err != nil {
		ReportError(err, "must set one of speakerID, roomID, timeslotID, sessionName", http.StatusBadRequest, writer)
		return
	}

	id, err := mySessionAPI.sessionWriter.WriteASession(sessionRequest.SpeakerID, sessionRequest.RoomID, sessionRequest.TimeslotID, sessionRequest.SessionName)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	writeIDToClient(writer, id)
}

// UpdateASessionRequest request for updateASession
type UpdateASessionRequest struct {
	SessionID   *int64  `json:"sessionID" example:"1"`
	SpeakerID   *int64  `json:"speakerID" example:"1"`
	RoomID      *int64  `json:"roomID" example:"1"`
	TimeslotID  *int64  `json:"timeslotID" example:"1"`
	SessionName *string `json:"sessionName" example:"Microservices"`
}

// Validate validates a UpdateASessionRequest
func (request UpdateASessionRequest) Validate() error {
	if request.SessionID == nil {
		return ErrInvalidRequest
	}
	if request.SpeakerID == nil && request.RoomID == nil && request.TimeslotID == nil && request.SessionName == nil {
		return ErrInvalidRequest
	}
	return nil
}

// updateASession Update an existing session in the db
// @Summary Update an existing session in the db
// @Description Update an existing session in the db
// @accept json
// @produce json
// @param session body api.UpdateASessionRequest true "the session to update with the new values"
// @Success 200 "Updated properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/session [put]
func (mySessionAPI sessionAPI) updateASession(writer http.ResponseWriter, request *http.Request) {
	requestJSON, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	sessionRequest := UpdateASessionRequest{}
	err = json.Unmarshal(requestJSON, &sessionRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, writer)
		return
	}

	if err = sessionRequest.Validate(); err != nil {
		ReportError(err, "must set sessionID and must set one of speakerID, roomID, timeslotID, sessionName", http.StatusBadRequest, writer)
		return
	}

	err = mySessionAPI.sessionUpdater.UpdateASession(*sessionRequest.SessionID, sessionRequest.SpeakerID, sessionRequest.RoomID, sessionRequest.TimeslotID, sessionRequest.SessionName)
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

// deleteASession Delete an existing session in the db
// @Summary Delete an existing session in the db
// @Description Delete an existing session in the db
// @produce json
// @param id query int true "the session to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/session [delete]
func (mySessionAPI sessionAPI) deleteASession(writer http.ResponseWriter, request *http.Request) {
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

	err = mySessionAPI.sessionDeleter.DeleteASession(requestedID)
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
