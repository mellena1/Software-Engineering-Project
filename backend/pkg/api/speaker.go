package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

var speakerReader db.SpeakerReader
var speakerWriter db.SpeakerWriter
var speakerUpdater db.SpeakerUpdater
var speakerDeleter db.SpeakerDeleter

func CreateSpeakerRoutes(apiObj API, speakerDBFacade db.SpeakerReaderWriterUpdaterDeleter) {
	speakerReader = speakerDBFacade
	speakerWriter = speakerDBFacade
	speakerUpdater = speakerDBFacade
	speakerDeleter = speakerDBFacade

	apiObj.CreateRouteWithMethods("/api/v1/speaker", getAllSpeakers, "GET")
}

// getAllSpeakers Gets all speakers from the db
// @Summary Get all speakers
// @Description Return a list of all speakers
// @Produce json
// @Success 200 {array} db.Speaker
// @Failure 400 {} nil
// @Router /api/v1/speaker [get]
func getAllSpeakers(w http.ResponseWriter, r *http.Request) {
	speakers, err := speakerReader.ReadAllSpeakers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	j, err := json.Marshal(speakers)
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to getAllSpeakers")
	}
}
