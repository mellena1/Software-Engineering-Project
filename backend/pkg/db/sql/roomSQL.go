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

func (r RoomSQL) ReadARoom(id int) (db.Room, error) {
	return db.Room{}, nil
}

func (r RoomSQL) ReadAllRooms() ([]db.Room, error) {
	if r.db == nil {
		return nil, ErrDBNotSet
	}

	q := `
	SELECT * FROM room;
	`

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}

	rooms := []db.Room{}
	for rows.Next() {
		newRoom := db.Room{}
		rows.Scan(&newRoom)
		rooms = append(rooms, newRoom)
	}
	return rooms, nil
}

func (r RoomSQL) WriteARoom(room db.Room) error {
	return nil
}

func (r RoomSQL) UpdateARoom(oldRoomID int, newRoom db.Room) error {
	return nil
}

func (r RoomSQL) DeleteARoom(id int) error {
	return nil
}
