package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

var roomReader db.RoomReader
var roomWriter db.RoomWriter
var roomUpdater db.RoomUpdater
var roomDeleter db.RoomDeleter

// CreateRoomRoutes makes all of the routes for room related calls
func CreateRoomRoutes(apiObj API, roomDBFacade db.RoomReaderWriterUpdaterDeleter) {
	roomReader = roomDBFacade
	roomWriter = roomDBFacade
	roomUpdater = roomDBFacade
	roomDeleter = roomDBFacade

	apiObj.CreateRouteWithMethods("/api/v1/room", getAllRooms, "GET")
}

// getAllRooms Gets all rooms from the db
// @Summary Get all rooms
// @Description Return a list of all rooms
// @Produce json
// @Success 200 {array} db.Room
// @Failure 400 {} nil
// @Router /api/v1/room [get]
func getAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := roomReader.ReadAllRooms()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	j, _ := json.Marshal(rooms)
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to getAllRooms")
	}
}
