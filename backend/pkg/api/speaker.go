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
		NewRoute("/api/v1/speaker/{email}", speakAPI.getASpeaker, "GET"),
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
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
// @Produce json
// @Success 200 {array} db.Speaker
// @Failure 400 {} nil
// @Router /api/v1/speaker/{email} [get]
func (a speakerAPI) getASpeaker(w http.ResponseWriter, r *http.Request) {
	email := "audrey.kirlin@example.org" //placeholder email to test functionalilty until we figure out retrieving the {email} from the path
	speakers, err := a.speakerReader.ReadASpeaker(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	j, err := json.Marshal(speakers)
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to getAllSpeakers")
	}
}