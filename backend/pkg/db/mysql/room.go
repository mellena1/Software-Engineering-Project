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

// ReadARoom reads a room from the db given roomID
func (r RoomMySQL) ReadARoom(roomID int) (db.Room, error) {
	if r.db == nil {
		return db.Room{}, ErrDBNotSet
	}

	newRoom := db.NewRoom()

	q := ("SELECT * FROM room WHERE roomID = ?;")

	row := r.db.QueryRow(q, roomID)
	switch err := row.Scan(newRoom.ID, newRoom.Name, newRoom.Capacity); err {
	case sql.ErrNoRows:
		return db.Room{}, err
	case nil:
		return newRoom, nil
	default:
		return db.Room{}, err
	}
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
		rows.Scan(newRoom.ID, newRoom.Name, newRoom.Capacity)
		rooms = append(rooms, newRoom)
	}
	return rooms, nil
}

// WriteARoom writes a room to the db
func (r RoomMySQL) WriteARoom(name *string, capacity *int) (int64, error) {
	if r.db == nil {
		return 0, ErrDBNotSet
	}
	insert, err := r.db.Prepare("INSERT INTO room(`roomName`, `capacity`) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer insert.Close()

	result, err := insert.Exec(*name, *capacity)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateARoom updates a room in the db given a roomName and the updated room
func (r RoomMySQL) UpdateARoom(id *int, name *string, capacity *int) error {
	if r.db == nil {
		return ErrDBNotSet
	}

	statement, err := r.db.Prepare("UPDATE room SET roomName = ?, capacity = ? WHERE roomID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(*name, *capacity, *id)
	if err != nil {
		return err
	}

	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

// DeleteARoom deletes a room given a roomName
func (r RoomMySQL) DeleteARoom(id *int) error {
	if r.db == nil {
		return ErrDBNotSet
	}

	statement, err := r.db.Prepare("DELETE FROM room WHERE roomID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(*id)
	if err != nil {
		return err
	}

	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}
