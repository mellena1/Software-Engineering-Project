package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// speakerAPI holds all of the api functions related to Speaker and all of the variables needed to access the backend
type speakerAPI struct {
	speakerReader  db.SpeakerReader
	speakerWriter  db.SpeakerWriter
	speakerUpdater db.SpeakerUpdater
	speakerDeleter db.SpeakerDeleter
}

// CreateSpeakerRoutes makes all of the routes for speaker related calls
func CreateSpeakerRoutes(speakerDBFacade db.SpeakerReaderWriterUpdaterDeleter) []Route {
	speakAPI := speakerAPI{
		speakerReader:  speakerDBFacade,
		speakerWriter:  speakerDBFacade,
		speakerUpdater: speakerDBFacade,
		speakerDeleter: speakerDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/speaker", speakAPI.getASpeaker, "GET"),
		NewRoute("/api/v1/speakers", speakAPI.getAllSpeakers, "GET"),
	}

	return routes
}

// getAllSpeakers Gets all speakers from the db
// @Summary Get all speakers
// @Description Return a list of all speakers
// @Produce json
// @Success 200 {array} db.Speaker
// @Failure 400 {} nil
// @Router /api/v1/speakers [get]
func (a speakerAPI) getAllSpeakers(w http.ResponseWriter, r *http.Request) {
	speakers, err := a.speakerReader.ReadAllSpeakers()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to read speakers from the db: %v", err)
		w.Write([]byte("Read from the backend failed"))
		return
	}
	j, err := json.Marshal(speakers)
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to getAllSpeakers")
	}
}

// getAllSpeakers Gets a speaker with the specified email from the db
// @Summary Get a speaker by email
// @Description Return a speaker with the specified email
// @Param email body string true "Email of the requested speaker"
// @Produce json
// @Success 200 {array} db.Speaker
// @Failure 400 {} nil
// @Router /api/v1/speaker [get]
func (a speakerAPI) getASpeaker(w http.ResponseWriter, r *http.Request) {
	email := getParamsFromRequest(r)
	speakers, err := a.speakerReader.ReadASpeaker(email)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to read speaker (%v) from the db: %v", email, err)
		w.Write([]byte("Read from the backend failed"))
		return
	}
	j, err := json.Marshal(speakers)
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to getASpeaker")
	}
}
