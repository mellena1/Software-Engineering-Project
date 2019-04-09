package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

type roomAPI struct {
	roomReader  db.RoomReader
	roomWriter  db.RoomWriter
	roomUpdater db.RoomUpdater
	roomDeleter db.RoomDeleter
}

func CreateRoomRoutes(roomDBFacade db.RoomReaderWriterUpdaterDeleter) []Route {
	myRoomAPI := roomAPI{
		roomReader:  roomDBFacade,
		roomWriter:  roomDBFacade,
		roomUpdater: roomDBFacade,
		roomDeleter: roomDBFacade,
	}

	routes := []Route{
		NewRoute("/api/v1/rooms", myRoomAPI.getAllRooms, "GET"),
		NewRoute("/api/v1/room", myRoomAPI.getARoom, "GET"),
		NewRoute("/api/v1/room", myRoomAPI.writeARoom, "POST"),
		NewRoute("/api/v1/room", myRoomAPI.updateARoom, "PUT"),
		NewRoute("/api/v1/room", myRoomAPI.deleteARoom, "DELETE"),
	}

	return routes
}

// getAllRooms Gets all rooms from the db
// @Summary Get all rooms room from the db
// @Description Return a list of all rooms
// @Produce json
// @Success 200 {array} db.Room
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/rooms [get]
func (myRoomAPI roomAPI) getAllRooms(writer http.ResponseWriter, request *http.Request) {
	rooms, err := myRoomAPI.roomReader.ReadAllRooms()
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	responseJSON, _ := json.Marshal(rooms)
	writer.Write(responseJSON)
}

// getRoom Gets a room from the db given a specific roomID
// @Summary Gets a room from the db given a specific roomID
// @Description Returns a room
// @param id query int true "the room to retrieve"
// @Produce json
// @Success 200 {} db.Room
// @Failure 400 {} _ "the request was bad"
// @Failure 503 {} _ "failed to access the db"
// @Router /api/v1/room [get]
func (myRoomAPI roomAPI) getARoom(writer http.ResponseWriter, request *http.Request) {
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

	room, err := myRoomAPI.roomReader.ReadARoom(requestedID)
	if err != nil {
		ReportError(err, "failed to access the db", http.StatusServiceUnavailable, writer)
		return
	}

	responseJSON, _ := json.Marshal(room)
	writer.Write(responseJSON)
}

type WriteARoomRequest struct {
	Name     string `json:"name" example:"Beatty"`
	Capacity *int64 `json:"capacity" example:"50"`
}

func (request WriteARoomRequest) Validate() error {
	if request.Name == "" {
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
func (myRoomAPI roomAPI) writeARoom(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	roomRequest := WriteARoomRequest{}
	err = json.Unmarshal(body, &roomRequest)
	if err != nil {
		ReportError(err, "failed to unmarshal json", http.StatusBadRequest, writer)
		return
	}

	if err = roomRequest.Validate(); err != nil {
		ReportError(err, "must set name for a room", http.StatusBadRequest, writer)
		return
	}

	id, err := myRoomAPI.roomWriter.WriteARoom(roomRequest.Name, roomRequest.Capacity)
	if err != nil {
		ReportError(err, "failed to write a room", http.StatusServiceUnavailable, writer)
		return
	}

	writeIDToClient(writer, id)
}

type UpdateARoomRequest struct {
	ID       *int64 `json:"id" example:"1"`
	Name     string `json:"name" example:"Beatty"`
	Capacity *int64 `json:"capacity" example:"50"`
}

func (request UpdateARoomRequest) Validate() error {
	if request.Name == "" || request.ID == nil {
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
func (myRoomAPI roomAPI) updateARoom(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ReportError(err, "unable to read body", http.StatusBadRequest, writer)
		return
	}

	updateRequest := UpdateARoomRequest{}
	err = json.Unmarshal(body, &updateRequest)
	if err != nil {
		ReportError(err, "failed to unmarshal json", http.StatusBadRequest, writer)
		return
	}

	if err = updateRequest.Validate(); err != nil {
		ReportError(err, "must set name for a room and pass an ID", http.StatusBadRequest, writer)
		return
	}

	err = myRoomAPI.roomUpdater.UpdateARoom(*updateRequest.ID, updateRequest.Name, updateRequest.Capacity)
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
func (myRoomAPI roomAPI) deleteARoom(writer http.ResponseWriter, request *http.Request) {
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

	err = myRoomAPI.roomDeleter.DeleteARoom(requestedID)
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
