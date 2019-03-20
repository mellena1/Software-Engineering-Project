package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

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

	j, _ := json.Marshal(speakers)
	w.Write(j)
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
		ReportError(err, "the \"id\" param was not set", http.StatusBadRequest, w)
		return
	case ErrBadQuery:
		ReportError(err, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, w)
		return
	case ErrBadQueryType:
		ReportError(err, "the \"id\" param is not a number", http.StatusBadRequest, w)
		return
	}

	speaker, err := a.speakerReader.ReadASpeaker(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	j, _ := json.Marshal(speaker)
	w.Write(j)
}

// writeASpeakerRequest request for writeASpeaker
type writeASpeakerRequest struct {
	Email     *string `json:"email" example:"person@gmail.com"`
	FirstName *string `json:"firstName" example:"Bob"`
	LastName  *string `json:"lastName" example:"Smith"`
}

var validateEmail, _ = regexp.Compile(`[a-zA-Z0-9\.\-\_]+@[a-zA-Z0-9\.\-\_]+`)
var validateName, _ = regexp.Compile(`[a-zA-Z\.\-]+`)

// Validate validates a writeASpeakerRequest
func (r writeASpeakerRequest) Validate() error {
	atLeastOneField := false

	if r.Email != nil {
		if *r.Email != validateEmail.FindString(*r.Email) {
			return ErrInvalidEmail
		}
		atLeastOneField = true
	}

	if r.FirstName != nil {
		if *r.FirstName != validateName.FindString(*r.FirstName) {
			return ErrInvalidName
		}
		atLeastOneField = true
	}

	if r.LastName != nil {
		if *r.LastName != validateName.FindString(*r.LastName) {
			return ErrInvalidName
		}
		atLeastOneField = true
	}

	if !atLeastOneField {
		return ErrInvalidRequest
	}

	return nil
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	speakerRequest := writeASpeakerRequest{}
	err = json.Unmarshal(body, &speakerRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, w)
		return
	}

	if err = speakerRequest.Validate(); err != nil {
		ReportError(err, "Failed to validate speaker request", http.StatusBadRequest, w)
		return
	}

	id, err := a.speakerWriter.WriteASpeaker(speakerRequest.Email, speakerRequest.FirstName, speakerRequest.LastName)
	if err != nil {
		ReportError(err, "failed to write a room", http.StatusServiceUnavailable, w)
		return
	}

	writeIDToClient(w, id)
}

// updateASpeakerRequest request for updateASpeaker
type updateASpeakerRequest struct {
	ID        *int64  `json:"id" example:"1"`
	Email     *string `json:"email" example:"person@gmail.com"`
	FirstName *string `json:"firstName" example:"Bob"`
	LastName  *string `json:"lastName" example:"Smith"`
}

// Validate validates a updateASpeakerRequest
func (r updateASpeakerRequest) Validate() error {
	if r.ID == nil {
		return ErrInvalidRequest
	}

	atLeastOneField := false

	if r.Email != nil {
		if !validateEmail.MatchString(*r.Email) {
			return ErrInvalidEmail
		}
		atLeastOneField = true
	}

	if r.FirstName != nil {
		if !validateName.MatchString(*r.FirstName) {
			return ErrInvalidName
		}
		atLeastOneField = true
	}

	if r.LastName != nil {
		if !validateName.MatchString(*r.LastName) {
			return ErrInvalidName
		}
		atLeastOneField = true
	}

	if !atLeastOneField {
		return ErrInvalidRequest
	}

	return nil
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	speakerRequest := updateASpeakerRequest{}
	err = json.Unmarshal(body, &speakerRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, w)
		return
	}

	if err = speakerRequest.Validate(); err != nil {
		ReportError(err, "must pass an id and one of email, firstName, and lastName must be set", http.StatusBadRequest, w)
		return
	}

	err = a.speakerUpdater.UpdateASpeaker(*speakerRequest.ID, speakerRequest.Email, speakerRequest.FirstName, speakerRequest.LastName)
	switch err {
	case nil:
		w.Write(nil)
		return
	case db.ErrNothingChanged:
		ReportError(err, "nothing in the db was changed. id probably does not exist", http.StatusBadRequest, w)
		return
	default:
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
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
		ReportError(err, "the \"id\" param was not set", http.StatusBadRequest, w)
		return
	case ErrBadQuery:
		ReportError(err, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, w)
		return
	case ErrBadQueryType:
		ReportError(err, "the \"id\" param is not a number", http.StatusBadRequest, w)
		return
	}

	err = a.speakerDeleter.DeleteASpeaker(requestedID)
	switch err {
	case nil:
		w.Write(nil)
		return
	case db.ErrNothingChanged:
		ReportError(err, "nothing in the db was changed. id probably does not exist", http.StatusBadRequest, w)
		return
	default:
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}
}
