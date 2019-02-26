package api

import (
	"encoding/json"
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
		NewRoute("/api/v1/session", sessAPI.getAllSessions, "GET"),
	}

	return routes
}

// getAllSessions Gets all sessions from the db
// @Summary Get all sessions
// @Description Return a list of all sessions
// @Produce json
// @Success 200 {array} db.Session
// @Failure 400 {} nil
// @Router /api/v1/session [get]
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
