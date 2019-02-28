package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	}

	return routes
}

type getASessionRequest struct {
	ID int64 `json:"id" example:"1"`
}

// getAllSessions Gets all sessions from the db
// @Summary Get all sessions
// @Description Return a list of all sessions
// @Produce json
// @Success 200 {array} db.Session
// @Failure 400 {} nil
// @Router /api/v1/session [get]
func (s sessionAPI) getASession(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	sessionRequest := getASessionRequest{}
	json.Unmarshal(body, &sessionRequest)
	session, err := s.sessionReader.ReadASession(sessionRequest.ID)
	if err != nil {
		ReportError(err, "Failed to get a session", http.StatusBadRequest, w)
		return
	}

	j, err := json.Marshal(session)
	if err != nil {
		ReportError(err, "Failed to create session json", http.StatusBadRequest, w)
	}

	w.Write(j)
}

// getAllSessions Gets all sessions from the db
// @Summary Get all sessions
// @Description Return a list of all sessions
// @Produce json
// @Param roomID body api.getASessionRequest true "ID of the requested Session"
// @Success 200 {array} db.Session
// @Failure 400 {} nil
// @Router /api/v1/sessions [get]
func (s sessionAPI) getAllSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := s.sessionReader.ReadAllSessions()
	if err != nil {
		ReportError(err, "Failed to get all sessions", http.StatusBadRequest, w)
		return
	}

	j, err := json.Marshal(sessions)
	if err != nil {
		ReportError(err, "Failed to create sessions json", http.StatusBadRequest, w)
	}

	w.Write(j)
}

// writeASessionRequest request for writeASession
type writeASessionRequest struct {
	SpeakerID   *int    `json:"speakerID" example:"1"`
	RoomID      *int    `json:"roomID" example:"1"`
	TimeslotID  *int64  `json:"timeslotID" example:"1"`
	SessionName *string `json:"sessionName" example:"Microservices"`
}

// writeASession Add a session to the db
// @Summary Add a session
// @Description Add a session to the db
// @accept json
// @produce json
// @param session body api.writeASessionRequest true "the session to add"
// @Success 200 {} int "the id of the session added"
// @Failure 400 {} string "the request was bad"
// @Failure 503 {} string "failed to access the db"
// @Router /api/v1/session [post]
func (s sessionAPI) writeASession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	sessionRequest := writeASessionRequest{}
	json.Unmarshal(j, &sessionRequest)

	id, err := s.sessionWriter.WriteASession(sessionRequest.SpeakerID, sessionRequest.RoomID, sessionRequest.TimeslotID, sessionRequest.SessionName)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	response := map[string]int64{"id": id}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}

// updateASessionRequest request for updateASession
type updateASessionRequest struct {
	SessionID   int64   `json:"sessionID" example:"1"`
	SpeakerID   *int    `json:"speakerID" example:"1"`
	RoomID      *int    `json:"roomID" example:"1"`
	TimeslotID  *int64  `json:"timeslotID" example:"1"`
	SessionName *string `json:"sessionName" example:"Microservices"`
}

// updateASession Update an existing session in the db
// @Summary Update an existing session in the db
// @Description Update an existing session in the db
// @accept json
// @produce json
// @param session body api.updateASessionRequest true "the session to update with the new values"
// @Success 200 "Updated properly"
// @Failure 400 {} string "the request was bad"
// @Failure 503 {} string "failed to access the db"
// @Router /api/v1/session [put]
func (s sessionAPI) updateASession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	sessionRequest := updateASessionRequest{}
	json.Unmarshal(j, &sessionRequest)

	err = s.sessionUpdater.UpdateASession(sessionRequest.SessionID, sessionRequest.SpeakerID, sessionRequest.RoomID, sessionRequest.TimeslotID, sessionRequest.SessionName)
	if err != nil {
		var msg string
		var status int
		switch err {
		case db.ErrNothingChanged:
			msg = "nothing in the db was changed. id probably does not exist"
			status = http.StatusBadRequest
		default:
			msg = "failed to access the db"
			status = http.StatusServiceUnavailable
		}

		ReportError(err, msg, status, w)
		return
	}

	w.Write(nil)
}

// deleteASessionRequest request for deleteASession
type deleteASessionRequest struct {
	SessionID int64 `json:"sessionID" example:"1"`
}

// deleteASession Delete an existing session in the db
// @Summary Delete an existing session in the db
// @Description Delete an existing session in the db
// @accept json
// @produce json
// @param session body api.deleteASessionRequest true "the session to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} string "the request was bad"
// @Failure 503 {} string "failed to access the db"
// @Router /api/v1/session [delete]
func (s sessionAPI) deleteASession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	j, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	sessionRequest := deleteASessionRequest{}
	json.Unmarshal(j, &sessionRequest)

	err = s.sessionDeleter.DeleteASession(sessionRequest.SessionID)
	if err != nil {
		var msg string
		var status int
		switch err {
		case db.ErrNothingChanged:
			msg = "nothing in the db was changed. id probably does not exist"
			status = http.StatusBadRequest
		default:
			msg = "failed to access the db"
			status = http.StatusServiceUnavailable
		}

		ReportError(err, msg, status, w)
		return
	}

	w.Write(nil)
}
