package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities/sql"
)

// Initialize the object that talks to the backend
var speaker = sql.SpeakerSQL{}
var speakerReader entities.SpeakerReader = speaker
var speakerWriter entities.SpeakerWriter = speaker
var speakerDeleter entities.SpeakerWriter = speaker

// GetAllSpeakers Gets all speakers from the db
// @Summary Get all speakers
// @Description Return a list of all speakers
// @Produce json
// @Success 200 {array} entities.Speaker
// @Failure 400 {} nil
// @Router /api/v1/speaker [get]
func GetAllSpeakers(w http.ResponseWriter, r *http.Request) {
	speakers, err := speakerReader.ReadAllSpeakers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	j, err := json.Marshal(speakers)
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to GetAllSpeakers")
	}
}
