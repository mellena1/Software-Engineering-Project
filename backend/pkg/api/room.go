package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// roomAPI holds all of the api functions related to Rooms and all of the variables needed to access the backend
type roomAPI struct {
	roomReader  db.RoomReader
	roomWriter  db.RoomWriter
	roomUpdater db.RoomUpdater
	roomDeleter db.RoomDeleter
}

type writeARoomRequest struct {
	Name     *string
	Capacity *int
}

type updateARoomRequest struct {
	ID       int64
	Name     *string
	Capacity *int
}

// CreateRoomRoutes makes all of the routes for room related calls
func CreateRoomRoutes(roomDBFacade db.RoomReaderWriterUpdaterDeleter) []Route {
	roomAPI := roomAPI{
		roomReader:  roomDBFacade,
		roomWriter:  roomDBFacade,
		roomUpdater: roomDBFacade,
		roomDeleter: roomDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/rooms", roomAPI.getAllRooms, "GET"),
		NewRoute("/api/v1/room", roomAPI.getARoom, "GET"),
		NewRoute("/api/v1/room", roomAPI.writeARoom, "POST"),
		NewRoute("/api/v1/room", roomAPI.updateARoom, "PUT"),
		NewRoute("/api/v1/room", roomAPI.deleteARoom, "DELETE"),
	}

	return routes
}

// getAllRooms Gets all rooms from the db
// @Summary Get all rooms
// @Description Return a list of all rooms
// @Produce json
// @Success 200 {array} db.Room
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/rooms [get]
func (a roomAPI) getAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := a.roomReader.ReadAllRooms()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to read rooms from the db: %v", err)
		w.Write([]byte("Read from the backend failed"))
		return
	}
	j, _ := json.Marshal(rooms)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(j)
	if err != nil {
		log.Println("Failed to respond to Rooms")
	}
}

// getRoom Gets all rooms from the db
// @Summary Get a room
// @Description Returns a room
// @param id query int true "the room to retrieve"
// @Produce json
// @Success 200 {} db.Room
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/room [get]
func (a roomAPI) getARoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	room, err := a.roomReader.ReadARoom(requestedID)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to read room (%v) from the db: %v", requestedID, err)
		w.Write([]byte("Read from the backend failed"))
		return
	}
	j, _ := json.Marshal(room)
	_, err = w.Write(j)
	if err != nil {
		log.Println("Failed to respond to Room")
	}
}

// writeRoom Writes a room to the room table
// @Summary Write a room to the db
// @Description Write a room to the db
// @Accept json
// @Produce json
// @Param room body api.writeARoomRequest true "Room to write"
// @Success 200 {int} nil
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/room [post]
func (a roomAPI) writeARoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "unable to read body"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	roomRequest := writeARoomRequest{}
	err = json.Unmarshal(body, &roomRequest)
	if err != nil {
		msg := "failed to unmarshal json"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	id, err := a.roomWriter.WriteARoom(roomRequest.Name, roomRequest.Capacity)
	if err != nil {
		msg := "failed to write a room"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusServiceUnavailable)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	json, _ := json.Marshal(id)
	_, err = w.Write(json)
	if err != nil {
		msg := "GET Room failed to write back"
		log.Printf("%s: %v", msg, err)
		return
	}
}

// deleteARoom deletes a room from the room table
// @Summary Delete a room from the db
// @Description Delete a room from the db
// @Accept json
// @Produce json
// @param id query int true "the room to delete"
// @Success 200 "Deleted properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/room [delete]
func (a roomAPI) deleteARoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	err = a.roomDeleter.DeleteARoom(requestedID)
	if err != nil {
		msg := "failed to delete a room"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusServiceUnavailable)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	json, _ := json.Marshal(true)
	_, err = w.Write(json)
	if err != nil {
		msg := "DELETE Room failed to write back"
		log.Printf("%s: %v", msg, err)
		return
	}
}

// updateARoom update a room in the room table
// @Summary Update a room in the db
// @Description Update a room in the db
// @Accept json
// @Produce json
// @param id query int true "the room to delete"
// @Success 200 {boolean} nil
// @Failure 400 {boolean} nil
// @Router /api/v1/room [PUT]
// @Param room body api.updateARoomRequest true "Room to update"
func (a roomAPI) updateARoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "unable to read body"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	updateRequest := updateARoomRequest{}
	err = json.Unmarshal(body, &updateRequest)
	if err != nil {
		msg := "failed to unmarshal json"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	err = a.roomUpdater.UpdateARoom(updateRequest.ID, updateRequest.Name, updateRequest.Capacity)
	if err != nil {
		msg := "failed to update a room"
		log.Printf("%s: %v", msg, err)
		w.WriteHeader(http.StatusServiceUnavailable)
		response, _ := json.Marshal(msg)
		w.Write(response)
		return
	}

	json, _ := json.Marshal(true)
	_, err = w.Write(json)
	if err != nil {
		msg := "PUT Room update failed to write back"
		log.Printf("%s: %v", msg, err)
		return
	}
}
