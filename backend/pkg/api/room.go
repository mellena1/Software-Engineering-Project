package api

import (
	"encoding/json"
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
		NewRoute("/api/v1/room", roomAPI.getRoom, "GET"),
	}

	return routes
}

// getAllRooms Gets all rooms from the db
// @Summary Get all rooms
// @Description Return a list of all rooms
// @Produce json
// @Success 200 {array} db.Room
// @Failure 400 {} nil
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
// @Produce json
// @Success 200 {} db.Room
// @Failure 400 {} nil
// @Router /api/v1/room [get]
func (a roomAPI) getRoom(w http.ResponseWriter, r *http.Request) {
	room, err := a.roomReader.ReadARoom() //This wants a string
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("Failed to read room from the db: %v", err)
		w.Write([]byte("Read from the backend failed"))
		return
	}
	j, _ := json.Marshal(room)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err = w.Write(j)
	if err != nil {
		log.Println("Failed to respond to Room")
	}
}
