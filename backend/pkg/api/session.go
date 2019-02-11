package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

var sessionReader db.SessionReader
var sessionWriter db.SessionWriter
var sessionUpdater db.SessionUpdater
var sessionDeleter db.SessionDeleter

func AddAllSessionRoutesToRouter(router *mux.Router, sessionDBFacade db.SessionReaderWriterUpdaterDeleter) {
	sessionReader = sessionDBFacade
	sessionWriter = sessionDBFacade
	sessionUpdater = sessionDBFacade
	sessionDeleter = sessionDBFacade

	router.HandleFunc("/api/v1/session", getAllSessions).Methods("GET")
}

// getAllSessions Gets all sessions from the db
// @Summary Get all sessions
// @Description Return a list of all sessions
// @Produce json
// @Success 200 {array} db.Session
// @Failure 400 {} nil
// @Router /api/v1/session [get]
func getAllSessions(w http.ResponseWriter, r *http.Request) {
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
