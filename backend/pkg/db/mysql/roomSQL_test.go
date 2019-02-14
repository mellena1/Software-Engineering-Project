package mysql

import (
	"testing"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var columns = []string{"roomID", "capacity"}

func TestReadAllRoomsValid(t *testing.T) {
	mockdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error occured when opening the stub db connection: %s", err)
	}
	defer mockdb.Close()

	mock.ExpectQuery("SELECT * FROM room;").
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("Room1,1\nRoom2,2"))

	// Execute ReadAllRooms
	roomSQL := NewRoomMySQL(mockdb)
	actual, err := roomSQL.ReadAllRooms()
	if err != nil {
		t.Fatalf("an error occured when running ReadAllRooms(): %s", err)
	}

	// Make sure expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Make sure returned Rooms are correct
	expected := []db.Room{
		db.Room{RoomName: "Room1", Capacity: 1},
		db.Room{RoomName: "Room2", Capacity: 2},
	}
	assert.Equal(t, expected, actual)
}

func TestReadAllRoomsInvalid(t *testing.T) {
	roomSQL := RoomMySQL{}
	_, err := roomSQL.ReadAllRooms()
	assert.Equal(t, ErrDBNotSet, err)
}
