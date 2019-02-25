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
	if r.db == nil {
		return nil, ErrDBNotSet
	}

	q := "SELECT * FROM room WHERE roomName = ?;"
	selectedRooms := db.Room{}

	rows, err := r.db.Query(q, roomName)
	if err != nil {
		return selectedRooms, err
	}

	defer rows.Close()

	for rows.next() {
		selectedRooms = db.Room{
			roomID: roomID;
			roomName: roomName;
			Capacity: Capacity;
		}
	}
	return selectedRooms, nil
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
		rows.Scan(newRoom.RoomName, newRoom.Capacity)
		rooms = append(rooms, newRoom)
	}
	return rooms, nil
}

// WriteARoom writes a room to the db
func (r RoomMySQL) WriteARoom(room db.Room) error {
	if r.db == nil {
		return nil, ErrDBNotSet
	}

	q := "" //I am not sure how to insert here based on what is being passed

	return nil
}

// UpdateARoom updates a room in the db given a roomName and the updated room
func (r RoomMySQL) UpdateARoom(roomName string, newRoom db.Room) error {
	//Are we adding more parameters here as well?
	return nil
}

// DeleteARoom deletes a room given a roomName
func (r RoomMySQL) DeleteARoom(roomName int) error {
	if r.db == nil {
		return nil, ErrDBNotSet
	}

	q := "DELETE FROM room WHERE roomName = ?"
	rows, err := r.db.Query(q, roomName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return nil
}
