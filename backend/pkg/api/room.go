package api

import (
	"encoding/json"
	"io/ioutil"
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
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	j, _ := json.Marshal(rooms)
	w.Write(j)
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

	room, err := a.roomReader.ReadARoom(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, w)
		return
	}

	j, _ := json.Marshal(room)
	w.Write(j)
}

// WriteARoomRequest request for writeARoom
type WriteARoomRequest struct {
	Name     string `json:"name" example:"Beatty"`
	Capacity *int64 `json:"capacity" example:"50"`
}

// Validate validates a WriteARoomRequest
func (r WriteARoomRequest) Validate() error {
	if r.Name == "" {
		return ErrInvalidRequest
	}
	return nil
}

// writeRoom Writes a room to the room table
// @Summary Write a room to the db
// @Description Write a room to the db
// @Accept json
// @Produce json
// @Param room body api.WriteARoomRequest true "Room to write"
// @Success 200 {int} nil
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/room [post]
func (a roomAPI) writeARoom(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	roomRequest := WriteARoomRequest{}
	err = json.Unmarshal(body, &roomRequest)
	if err != nil {
		ReportError(err, "failed to unmarshal json", http.StatusBadRequest, w)
		return
	}

	if err = roomRequest.Validate(); err != nil {
		ReportError(err, "must set name for a room", http.StatusBadRequest, w)
		return
	}

	id, err := a.roomWriter.WriteARoom(roomRequest.Name, roomRequest.Capacity)
	if err != nil {
		ReportError(err, "failed to write a room", http.StatusServiceUnavailable, w)
		return
	}

	writeIDToClient(w, id)
}

// UpdateARoomRequest request for updateARoom
type UpdateARoomRequest struct {
	ID       *int64 `json:"id" example:"1"`
	Name     string `json:"name" example:"Beatty"`
	Capacity *int64 `json:"capacity" example:"50"`
}

// Validate validates a UpdateARoomRequest
func (r UpdateARoomRequest) Validate() error {
	if r.Name == "" || r.ID == nil {
		return ErrInvalidRequest
	}
	return nil
}

// updateARoom update a room in the room table
// @Summary Update a room in the db
// @Description Update a room in the db
// @Accept json
// @Produce json
// @param id query int true "the room to delete"
// @Success 200 "Updated properly"
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/room [PUT]
// @Param room body api.UpdateARoomRequest true "Room to update"
func (a roomAPI) updateARoom(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, w)
		return
	}

	updateRequest := UpdateARoomRequest{}
	err = json.Unmarshal(body, &updateRequest)
	if err != nil {
		ReportError(err, "failed to unmarshal json", http.StatusBadRequest, w)
		return
	}

	if err = updateRequest.Validate(); err != nil {
		ReportError(err, "must set name for a room and pass an ID", http.StatusBadRequest, w)
		return
	}

	err = a.roomUpdater.UpdateARoom(*updateRequest.ID, updateRequest.Name, updateRequest.Capacity)
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

	err = a.roomDeleter.DeleteARoom(requestedID)
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
