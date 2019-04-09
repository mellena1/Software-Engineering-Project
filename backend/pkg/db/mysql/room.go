package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

type RoomMySQL struct {
	db *sql.DB
}

func NewRoomMySQL(db *sql.DB) RoomMySQL {
	return RoomMySQL{db}
}

func scanARoom(room *db.Room, row rowScanner) error {
	capacity := sql.NullInt64{}
	err := row.Scan(&room.ID, room.Name, &capacity)
	room.Capacity = NullIntToInt(capacity)
	return err
}

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
