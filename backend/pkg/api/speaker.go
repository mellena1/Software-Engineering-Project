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
	mySpeakerAPI := speakerAPI{
		speakerReader:  speakerDBFacade,
		speakerWriter:  speakerDBFacade,
		speakerUpdater: speakerDBFacade,
		speakerDeleter: speakerDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/speaker", mySpeakerAPI.getASpeaker, "GET"),
		NewRoute("/api/v1/speakers", mySpeakerAPI.getAllSpeakers, "GET"),
		NewRoute("/api/v1/speaker", mySpeakerAPI.writeASpeaker, "POST"),
		NewRoute("/api/v1/speaker", mySpeakerAPI.updateASpeaker, "PUT"),
		NewRoute("/api/v1/speaker", mySpeakerAPI.deleteASpeaker, "DELETE"),
	}

	return routes
}

// getAllSpeakers Gets all speakers from the db
// @Summary Gets all speakers from the db
// @Description Return a list of all speakers
// @Produce json
// @Success 200 {array} db.Speaker
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/speakers [get]
func (mySpeakerAPI speakerAPI) getAllSpeakers(writer http.ResponseWriter, request *http.Request) {
	speakers, err := mySpeakerAPI.speakerReader.ReadAllSpeakers()
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to read speakers from the db: %v", err)
		writer.Write([]byte("Read from the backend failed"))
		return
	}

	responseJSON, _ := json.Marshal(speakers)
	writer.Write(responseJSON)
}

// getAllSpeakers Gets a speaker with the specified speakerID from the db
// @Summary Gets a speaker with the specified speakerID from the db
// @Description Return a speaker with the specified speakerID from the db
// @param id query int true "the speaker to retrieve"
// @Produce json
// @Success 200 {object} db.Speaker
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/speaker [get]
func (mySpeakerAPI speakerAPI) getASpeaker(writer http.ResponseWriter, request *http.Request) {
	requestedID, err := getIDFromQueries(request)
	switch err {
	case ErrQueryNotSet:
		ReportError(err, "the \"id\" param was not set", http.StatusBadRequest, writer)
		return
	case ErrBadQuery:
		ReportError(err, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, writer)
		return
	case ErrBadQueryType:
		ReportError(err, "the \"id\" param is not a number", http.StatusBadRequest, writer)
		return
	}

	speaker, err := mySpeakerAPI.speakerReader.ReadASpeaker(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	responseJSON, _ := json.Marshal(speaker)
	writer.Write(responseJSON)
}

// WriteASpeakerRequest request for writeASpeaker
type WriteASpeakerRequest struct {
	Email     *string `json:"email" example:"person@gmail.com"`
	FirstName *string `json:"firstName" example:"Bob"`
	LastName  *string `json:"lastName" example:"Smith"`
}

var validateEmail, _ = regexp.Compile(`[a-zA-Z0-9\.\-\_]+@[a-zA-Z0-9\.\-\_]+`)
var validateName, _ = regexp.Compile(`[a-zA-Z\.\-]+`)

// Validate validates a writeASpeakerRequest
func (request WriteASpeakerRequest) Validate() error {
	atLeastOneField := false

	if request.Email != nil && *request.Email != "" {
		if *request.Email != validateEmail.FindString(*request.Email) {
			return ErrInvalidEmail
		}
		atLeastOneField = true
	}

	if request.FirstName != nil && *request.FirstName != "" {
		if *request.FirstName != validateName.FindString(*request.FirstName) {
			return ErrInvalidName
		}
		atLeastOneField = true
	}

	if request.LastName != nil && *request.LastName != "" {
		if *request.LastName != validateName.FindString(*request.LastName) {
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
// @Param Speaker body api.WriteASpeakerRequest true "Speaker that wants to be added to the db (no ID)"
// @Produce json
// @Success 200
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/speaker [post]
func (mySpeakerAPI speakerAPI) writeASpeaker(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	speakerRequest := WriteASpeakerRequest{}
	err = json.Unmarshal(body, &speakerRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, writer)
		return
	}

	if err = speakerRequest.Validate(); err != nil {
		ReportError(err, "Failed to validate speaker request", http.StatusBadRequest, writer)
		return
	}

	id, err := mySpeakerAPI.speakerWriter.WriteASpeaker(speakerRequest.Email, speakerRequest.FirstName, speakerRequest.LastName)
	if err != nil {
		ReportError(err, "failed to write a room", http.StatusServiceUnavailable, writer)
		return
	}

	writeIDToClient(writer, id)
}

// UpdateASpeakerRequest request for updateASpeaker
type UpdateASpeakerRequest struct {
	ID        *int64  `json:"id" example:"1"`
	Email     *string `json:"email" example:"person@gmail.com"`
	FirstName *string `json:"firstName" example:"Bob"`
	LastName  *string `json:"lastName" example:"Smith"`
}

// Validate validates a UpdateASpeakerRequest
func (request UpdateASpeakerRequest) Validate() error {
	if request.ID == nil {
		return ErrInvalidRequest
	}

	atLeastOneField := false

	if request.Email != nil && *request.Email != "" {
		if *request.Email != validateEmail.FindString(*request.Email) {
			return ErrInvalidEmail
		}
		atLeastOneField = true
	}

	if request.FirstName != nil && *request.FirstName != "" {
		if *request.FirstName != validateName.FindString(*request.FirstName) {
			return ErrInvalidName
		}
		atLeastOneField = true
	}

	if request.LastName != nil && *request.LastName != "" {
		if *request.LastName != validateName.FindString(*request.LastName) {
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
// @Param Speaker body api.UpdateASpeakerRequest true "Speaker struct that wants to be updated in the db"
// @Produce json
// @Success 200
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/speaker [put]
func (mySpeakerAPI speakerAPI) updateASpeaker(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	speakerRequest := UpdateASpeakerRequest{}
	err = json.Unmarshal(body, &speakerRequest)
	if err != nil {
		ReportError(err, "json is unable to be unmarshaled", http.StatusBadRequest, writer)
		return
	}

	if err = speakerRequest.Validate(); err != nil {
		ReportError(err, "must pass an id and one of email, firstName, and lastName must be set", http.StatusBadRequest, writer)
		return
	}

	err = mySpeakerAPI.speakerUpdater.UpdateASpeaker(*speakerRequest.ID, speakerRequest.Email, speakerRequest.FirstName, speakerRequest.LastName)
	switch err {
	case nil:
		writer.Write(nil)
		return
	case db.ErrNothingChanged:
		ReportError(err, "nothing in the db was changed. id probably does not exist", http.StatusBadRequest, writer)
		return
	default:
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
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
func (mySpeakerAPI speakerAPI) deleteASpeaker(writer http.ResponseWriter, request *http.Request) {
	requestedID, err := getIDFromQueries(request)
	switch err {
	case ErrQueryNotSet:
		ReportError(err, "the \"id\" param was not set", http.StatusBadRequest, writer)
		return
	case ErrBadQuery:
		ReportError(err, "you are only allowed to specify 1 id at a time", http.StatusBadRequest, writer)
		return
	case ErrBadQueryType:
		ReportError(err, "the \"id\" param is not a number", http.StatusBadRequest, writer)
		return
	}

	err = mySpeakerAPI.speakerDeleter.DeleteASpeaker(requestedID)
	switch err {
	case nil:
		writer.Write(nil)
		return
	case db.ErrNothingChanged:
		ReportError(err, "nothing in the db was changed. id probably does not exist", http.StatusBadRequest, writer)
		return
	default:
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}
}
