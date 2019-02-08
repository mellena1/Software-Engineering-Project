package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities/sql"
)

// Initialize the object that talks to the backend
var room = sql.RoomSQL{}
var roomReader entities.RoomReader = room
var roomWriter entities.RoomWriter = room
var roomDeleter entities.RoomWriter = room

// GetAllRooms Gets all rooms from the db
// @Summary Get all rooms
// @Description Return a list of all rooms
// @Produce json
// @Success 200 {array} entities.Room
// @Failure 400 {} nil
// @Router /api/v1/room [get]
func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := roomReader.ReadAllRooms()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	j, _ := json.Marshal(rooms)
	_, err = w.Write(j)
	if err != nil {
		log.Fatal("Failed to respond to GetAllRooms")
	}
}
