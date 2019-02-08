package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities/sql"
)

// Initialize the object that talks to the backend
var session = sql.SessionSQL{}
var sessionReader entities.SessionReader = session
var sessionWriter entities.SessionWriter = session
var sessionDeleter entities.SessionWriter = session

// GetAllSessions Gets all sessions from the db
// @Summary Get all sessions
// @Description Return a list of all sessions
// @Produce json
// @Success 200 {array} entities.Session
// @Failure 400 {} nil
// @Router /api/v1/session [get]
func GetAllSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := sessionReader.ReadAllSessions()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	j, err := json.Marshal(sessions)
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to GetAllSpeakers")
	}
}
