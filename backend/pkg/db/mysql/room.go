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

// scanARoom takes in a room pointer and scans a row into it
func scanARoom(room *db.Room, row rowScanner) error {
	capacity := sql.NullInt64{}
	err := row.Scan(&room.ID, room.Name, &capacity)
	room.Capacity = NullIntToInt(capacity)
	return err
}

// ReadARoom reads a room from the db given roomID
func (myRoomSQL RoomMySQL) ReadARoom(roomID int64) (db.Room, error) {
	if myRoomSQL.db == nil {
		return db.Room{}, ErrDBNotSet
	}

	statement, err := myRoomSQL.db.Prepare("SELECT roomID, roomName, capacity FROM room WHERE roomID = ?;")
	if err != nil {
		return db.Room{}, err
	}
	defer statement.Close()

	room := db.NewRoom()
	row := statement.QueryRow(roomID)

	err = scanARoom(&room, row)

	return room, err
}

// ReadAllRooms reads all rooms from the db
func (myRoomSQL RoomMySQL) ReadAllRooms() ([]db.Room, error) {
	if myRoomSQL.db == nil {
		return nil, ErrDBNotSet
	}

	query := "SELECT roomID, roomName, capacity FROM room;"

	rows, err := myRoomSQL.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []db.Room{}
	for rows.Next() {
		newRoom := db.NewRoom()
		scanARoom(&newRoom, rows)
		rooms = append(rooms, newRoom)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

// WriteARoom writes a room to the db
func (myRoomSQL RoomMySQL) WriteARoom(name string, capacity *int64) (int64, error) {
	if myRoomSQL.db == nil {
		return 0, ErrDBNotSet
	}
	statement, err := myRoomSQL.db.Prepare("INSERT INTO room(`roomName`, `capacity`) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(name, IntToNullInt(capacity))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateARoom updates a room in the db given a roomName and the updated room
func (myRoomSQL RoomMySQL) UpdateARoom(id int64, name string, capacity *int64) error {
	if myRoomSQL.db == nil {
		return ErrDBNotSet
	}

	statement, err := myRoomSQL.db.Prepare("UPDATE room SET roomName = ?, capacity = ? WHERE roomID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(name, IntToNullInt(capacity), id)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return err
	} else if rows == 0 {
		return db.ErrNothingChanged
	}

	return nil
}

// DeleteARoom deletes a room given a roomName
func (myRoomSQL RoomMySQL) DeleteARoom(id int64) error {
	if myRoomSQL.db == nil {
		return ErrDBNotSet
	}

	statement, err := myRoomSQL.db.Prepare("DELETE FROM room WHERE roomID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return err
	} else if rows == 0 {
		return db.ErrNothingChanged
	}

	return nil
}
