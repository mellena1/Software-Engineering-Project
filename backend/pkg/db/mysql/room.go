package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// RoomMySQL implements RoomReaderWriterUpdaterDeleter
type RoomMySQL struct {
	db *sql.DB
}

// NewRoomMySQL makes a new RoomMySQL object given a db
func NewRoomMySQL(db *sql.DB) RoomMySQL {
	return RoomMySQL{db}
}

// ReadARoom reads a room from the db given roomName
func (r RoomMySQL) ReadARoom(roomName string) (db.Room, error) {
	return db.Room{}, nil
}

// ReadAllRooms reads all rooms from the db
func (r RoomMySQL) ReadAllRooms() ([]db.Room, error) {
	if r.db == nil {
		return nil, ErrDBNotSet
	}

	q := "SELECT * FROM room;"

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []db.Room{}
	for rows.Next() {
		newRoom := db.NewRoom()
		rows.Scan(newRoom.ID, newRoom.RoomName, newRoom.Capacity)
		rooms = append(rooms, newRoom)
	}
	return rooms, nil
}

// WriteARoom writes a room to the db
func (r RoomMySQL) WriteARoom(room db.Room) error {
	return nil
}

// UpdateARoom updates a room in the db given a roomName and the updated room
func (r RoomMySQL) UpdateARoom(roomName string, newRoom db.Room) error {
	return nil
}

// DeleteARoom deletes a room given a roomName
func (r RoomMySQL) DeleteARoom(roomName int) error {
	return nil
}
