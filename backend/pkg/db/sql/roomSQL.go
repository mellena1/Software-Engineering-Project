package sql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

type RoomSQL struct {
	db *sql.DB
}

func NewRoomSQL(db *sql.DB) RoomSQL {
	return RoomSQL{db}
}

func (r RoomSQL) ReadARoom(roomName string) (db.Room, error) {
	return db.Room{}, nil
}

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

func (r RoomSQL) WriteARoom(room db.Room) error {
	return nil
}

func (r RoomSQL) UpdateARoom(roomName string, newRoom db.Room) error {
	return nil
}

func (r RoomSQL) DeleteARoom(roomName int) error {
	return nil
}
