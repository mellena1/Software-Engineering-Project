package sql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// RoomSQL implements RoomReaderWriterUpdaterDeleter
type RoomSQL struct {
	db *sql.DB
}

// NewRoomSQL makes a new RoomSQL object given a db
func NewRoomSQL(db *sql.DB) RoomSQL {
	return RoomSQL{db}
}

// ReadARoom reads a room from the db given roomName
func (r RoomSQL) ReadARoom(roomName string) (db.Room, error) {
	return db.Room{}, nil
}

// ReadAllRooms reads all rooms from the db
func (r RoomSQL) ReadAllRooms() ([]db.Room, error) {
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
		newRoom := db.Room{}
		rows.Scan(&newRoom.RoomName, &newRoom.Capacity)
		rooms = append(rooms, newRoom)
	}
	return rooms, nil
}

// WriteARoom writes a room to the db
func (r RoomSQL) WriteARoom(room db.Room) error {
	return nil
}

// UpdateARoom updates a room in the db given a roomName and the updated room
func (r RoomSQL) UpdateARoom(roomName string, newRoom db.Room) error {
	return nil
}

// DeleteARoom deletes a room given a roomName
func (r RoomSQL) DeleteARoom(roomName int) error {
	return nil
}
