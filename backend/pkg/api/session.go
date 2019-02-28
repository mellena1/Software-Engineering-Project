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
	}

	return routes
}

type getASessionRequest struct {
	ID int `json:"id" example:"1"`
}

// getAllSessions Gets all sessions from the db
// @Summary Get all sessions
// @Description Return a list of all sessions
// @Produce json
// @Success 200 {array} db.Session
// @Failure 400 {} nil
// @Router /api/v1/session [get]
func (a sessionAPI) getASession(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	sessionRequest := getASessionRequest{}
	json.Unmarshal(body, &sessionRequest)
	session, err := a.sessionReader.ReadASession(sessionRequest.ID)
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
func (a sessionAPI) getAllSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := a.sessionReader.ReadAllSessions()
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
