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

type writeASpeakerRequest struct {
	Email     string
	FirstName string
	LastName  string
}
type updateASpeakerRequest struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
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
// @Failure 503 {} _ "failed to access the db"
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
// @param id query int true "the speaker to retrieve"
// @Produce json
// @Success 200 {object} db.Speaker
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/speaker [get]
func (a speakerAPI) getASpeaker(w http.ResponseWriter, r *http.Request) {
	requestedID, err := getIDFromQueries(r)
	switch err {
	case ErrQueryNotSet:
		ReportError(ErrQueryNotSet, "the \"id\" param was not set", http.StatusBadRequest, w)
		return
	case ErrBadQuery:
		ReportError(ErrBadQuery, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, w)
		return
	case ErrBadQueryType:
		ReportError(ErrBadQueryType, "the \"id\" param is not a number", http.StatusBadRequest, w)
		return
	}

	speaker, err := a.speakerReader.ReadASpeaker(requestedID)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to read speaker (%v) from the db: %v", requestedID, err)
		w.Write([]byte("Read from the backend failed"))
		return
	}
	j, err := json.Marshal(speaker)
	w.Write(j)
}

// writeASpeaker Inserts a speaker into the database
// @Summary Write a speaker
// @Description Inserts a speaker with the specified email, firstName and lastName
// @Param Speaker body api.writeASpeakerRequest true "Speaker that wants to be added to the db (no ID)"
// @Produce json
// @Success 200
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
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
	json, _ := json.Marshal(data.Email)
	_, err = w.Write(json)
	if err != nil {
		log.Fatal("Failed to respond to writeASpeaker")
	}
}

// updateASpeaker Edits a speaker already in the database
// @Summary Edit a speaker
// @Description Return a speaker with the specified email
// @Param Speaker body api.updateASpeakerRequest true "Speaker struct that wants to be updated in the db"
// @Produce json
// @Success 200
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
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
	json, _ := json.Marshal(data.ID)
	_, err = w.Write(json)
	if err != nil {
		log.Fatal("Failed to respond to writeASpeaker")
	}
}

// deleteASpeaker Delete a speaker from the database
// @Summary Delete a speaker
// @Description Delete a speaker with the specified id
// @param id query int true "the speaker to delete"
// @Produce json
// @Success 200
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/speaker [delete]
func (a speakerAPI) deleteASpeaker(w http.ResponseWriter, r *http.Request) {
	requestedID, err := getIDFromQueries(r)
	switch err {
	case ErrQueryNotSet:
		ReportError(ErrQueryNotSet, "the \"id\" param was not set", http.StatusBadRequest, w)
		return
	case ErrBadQuery:
		ReportError(ErrBadQuery, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, w)
		return
	case ErrBadQueryType:
		ReportError(ErrBadQueryType, "the \"id\" param is not a number", http.StatusBadRequest, w)
		return
	}
	err = a.speakerDeleter.DeleteASpeaker(requestedID)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to delete speaker (%v) from the db: %v", requestedID, err)
		w.Write([]byte("delete failed"))
		return
	}
}
