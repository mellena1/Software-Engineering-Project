package api

import (
	"encoding/json"
	"io/ioutil"
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

type getASpeakerRequest struct {
	ID int
}
type writeASpeakerRequest struct {
	Email     string
	FirstName string
	LastName  string
}
type updateASpeakerRequest struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
}
type deleteASpeakerRequest struct {
	ID int
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
		NewRoute("/api/v1/speaker", speakAPI.writeASpeaker, "POST"),
		NewRoute("/api/v1/speaker", speakAPI.updateASpeaker, "PUT"),
		NewRoute("/api/v1/speaker", speakAPI.deleteASpeaker, "DELETE"),
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to getAllSpeakers")
	}
}

// getAllSpeakers Gets a speaker with the specified email from the db
// @Summary Get a speaker by email
// @Description Return a speaker with the specified email
// @Param speakerID body api.getASpeakerRequest true "ID of the requested speaker"
// @Produce json
// @Success 200 {array} db.Speaker
// @Failure 400 {} nil
// @Router /api/v1/speaker [get]
func (a speakerAPI) getASpeaker(w http.ResponseWriter, r *http.Request) {

	var data getASpeakerRequest
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)

	if err != nil {
		return
	}

	speakerID := data.ID

	speakers, err := a.speakerReader.ReadASpeaker(speakerID)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to read speaker (%v) from the db: %v", speakerID, err)
		w.Write([]byte("Read from the backend failed"))
		return
	}
	j, err := json.Marshal(speakers)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to getASpeaker")
	}
}

// writeASpeaker Inserts a speaker into the database
// @Summary Write a speaker
// @Description Inserts a speaker with the specified email, firstName and lastName
// @Param Speaker body api.writeASpeakerRequest true "Speaker that wants to be added to the db (no ID)"
// @Produce json
// @Success 200
// @Failure 400 {} nil
// @Router /api/v1/speaker [post]
func (a speakerAPI) writeASpeaker(w http.ResponseWriter, r *http.Request) {

	var data writeASpeakerRequest
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)

	if err != nil {
		return
	}

	err = a.speakerWriter.WriteASpeaker(data.Email, data.FirstName, data.LastName)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to write speaker (%v) to the db: %v", data.Email, err)
		w.Write([]byte("Write failed"))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		log.Fatal("Failed to respond to writeASpeaker")
	}
}

// updateASpeaker Edits a speaker already in the database
// @Summary Edit a speaker
// @Description Return a speaker with the specified email
// @Param Speaker body api.updateASpeakerRequest true "Speaker struct that wants to be added to the db"
// @Produce json
// @Success 200
// @Failure 400 {} nil
// @Router /api/v1/speaker [put]
func (a speakerAPI) updateASpeaker(w http.ResponseWriter, r *http.Request) {

	var data updateASpeakerRequest
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)

	if err != nil {
		return
	}

	err = a.speakerUpdater.UpdateASpeaker(data.ID, data.Email, data.FirstName, data.LastName)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to update speaker (%v) to the db: %v", data.ID, err)
		w.Write([]byte("Write failed"))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		log.Fatal("Failed to respond to writeASpeaker")
	}
}

// deleteASpeaker Delete a speaker from the database
// @Summary Delete a speaker
// @Description Delete a speaker with the specified id
// @Param speakerID body api.deleteASpeakerRequest true "SpeakerID of the speaker to be deleted from the db"
// @Produce json
// @Success 200
// @Failure 400 {} nil
// @Router /api/v1/speaker [delete]
func (a speakerAPI) deleteASpeaker(w http.ResponseWriter, r *http.Request) {

	var data deleteASpeakerRequest
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)

	if err != nil {
		return
	}

	err = a.speakerDeleter.DeleteASpeaker(data.ID)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to delete speaker (%v) from the db: %v", data.ID, err)
		w.Write([]byte("Write failed"))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		log.Fatal("Failed to respond to writeASpeaker")
	}
}
