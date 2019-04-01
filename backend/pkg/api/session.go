package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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
	sessAPI := sessionAPI{
		sessionReader:  sessionDBFacade,
		sessionWriter:  sessionDBFacade,
		sessionUpdater: sessionDBFacade,
		sessionDeleter: sessionDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/session", sessAPI.getASession, "GET"),
		NewRoute("/api/v1/sessions", sessAPI.getAllSessions, "GET"),
		NewRoute("/api/v1/session", sessAPI.writeASession, "POST"),
		NewRoute("/api/v1/session", sessAPI.updateASession, "PUT"),
		NewRoute("/api/v1/session", sessAPI.deleteASession, "DELETE"),
		NewRoute("/api/v1/sessionsByTimeslot", sessAPI.getAllSessionsByTimeslot, "GET"),
	}

	return routes
}

// getASession Gets a session from the db
// @Summary Get a session
// @Description Return a session
// @Produce json
// @param id query int true "the session to retrieve"
// @Success 200 {object} db.Session
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/session [get]
func (s sessionAPI) getASession(w http.ResponseWriter, r *http.Request) {
	requestedID, err := getIDFromQueries(r)
	switch err {
	case ErrQueryNotSet:
		ReportError(err, "the \"id\" param was not set", http.StatusBadRequest, w)
		return
	case ErrBadQuery:
		ReportError(err, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, w)
		return
	case ErrBadQueryType:
		ReportError(err, "the \"id\" param is not a number", http.StatusBadRequest, w)
		return
	}

	session, err := s.sessionReader.ReadASession(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	j, _ := json.Marshal(session)
	w.Write(j)
}

// getAllSessions Gets all sessions from the db
// @Summary Get all sessions
// @Description Return a list of all sessions
// @Produce json
// @Success 200 {array} db.Session
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/sessions [get]
func (s sessionAPI) getAllSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := s.sessionReader.ReadAllSessions()
	if err != nil {
		ReportError(err, "Failed to get all sessions", http.StatusBadRequest, w)
		return
	}

	j, _ := json.Marshal(sessions)
	w.Write(j)
}

// WriteASessionRequest request for writeASession
type WriteASessionRequest struct {
	SpeakerID   *int64  `json:"speakerID" example:"1"`
	RoomID      *int64  `json:"roomID" example:"1"`
	TimeslotID  *int64  `json:"timeslotID" example:"1"`
	SessionName *string `json:"sessionName" example:"Microservices"`
}

// Validate validates a WriteASessionRequest
func (r WriteASessionRequest) Validate() error {
	if r.SpeakerID == nil && r.RoomID == nil && r.TimeslotID == nil && r.SessionName == nil {
		return ErrInvalidRequest
	}
	return nil
}

// getAllSessions Gets all sessions from the db
// @Summary Get all sessions
// @Description Return a list of all sessions
// @Produce json
// @Success 200 {array} map[string][]*string
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/sessionsByTimeslot [get]
func (s sessionAPI) getAllSessionsByTimeslot(w http.ResponseWriter, r *http.Request) {
	sessions, err := s.sessionReader.ReadAllSessions()
	if err != nil {
		ReportError(err, "Failed to get ", http.StatusBadRequest, w)
		return
	}

	// Making a map of sessions to timeslots, where the key is just the time of the session (i.e 12:00 to 1:00)
	sessionsByTimeslot := make(map[string][]*string)
	for _, session := range sessions {
		key := strconv.Itoa(session.Timeslot.StartTime.Hour()) + "-" + strconv.Itoa(session.Timeslot.EndTime.Hour())
		sessionsByTimeslot[key] = append(sessionsByTimeslot[key], session.Name)
	}

	j, _ := json.Marshal(sessionsByTimeslot)
	w.Write(j)
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
func (s sessionAPI) writeASession(w http.ResponseWriter, r *http.Request) {
	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	sessionRequest := WriteASessionRequest{}
	err = json.Unmarshal(j, &sessionRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, w)
		return
	}

	if err = sessionRequest.Validate(); err != nil {
		ReportError(err, "must set one of speakerID, roomID, timeslotID, sessionName", http.StatusBadRequest, w)
		return
	}

	id, err := s.sessionWriter.WriteASession(sessionRequest.SpeakerID, sessionRequest.RoomID, sessionRequest.TimeslotID, sessionRequest.SessionName)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	writeIDToClient(w, id)
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
func (r UpdateASessionRequest) Validate() error {
	if r.SessionID == nil {
		return ErrInvalidRequest
	}
	if r.SpeakerID == nil && r.RoomID == nil && r.TimeslotID == nil && r.SessionName == nil {
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
func (s sessionAPI) updateASession(w http.ResponseWriter, r *http.Request) {
	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	sessionRequest := UpdateASessionRequest{}
	err = json.Unmarshal(j, &sessionRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, w)
		return
	}

	if err = sessionRequest.Validate(); err != nil {
		ReportError(err, "must set sessionID and must set one of speakerID, roomID, timeslotID, sessionName", http.StatusBadRequest, w)
		return
	}

	err = s.sessionUpdater.UpdateASession(*sessionRequest.SessionID, sessionRequest.SpeakerID, sessionRequest.RoomID, sessionRequest.TimeslotID, sessionRequest.SessionName)
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

// deleteASession Delete an existing session in the db
// @Summary Delete an existing session in the db
// @Description Delete an existing session in the db
// @produce json
// @param id query int true "the session to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/session [delete]
func (s sessionAPI) deleteASession(w http.ResponseWriter, r *http.Request) {
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

	err = s.sessionDeleter.DeleteASession(requestedID)
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
