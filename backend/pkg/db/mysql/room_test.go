package mysql

import (
	"testing"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var roomColumns = []string{"ID", "roomName", "capacity"}

func TestReadAllRoomsValid(t *testing.T) {
	mockdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error occured when opening the stub db connection: %s", err)
	}
	defer mockdb.Close()

	mock.ExpectQuery("SELECT roomID, roomName, capacity FROM room;").
		WillReturnRows(sqlmock.NewRows(roomColumns).FromCSVString("10,Room1,1\n20,Room2,2"))

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
		db.Room{ID: db.Int64Ptr(10), Name: db.StringPtr("Room1"), Capacity: db.Int64Ptr(1)},
		db.Room{ID: db.Int64Ptr(20), Name: db.StringPtr("Room2"), Capacity: db.Int64Ptr(2)},
	}
	assert.Equal(t, expected, actual)
}

func TestReadAllRoomsInvalid(t *testing.T) {
	roomSQL := RoomMySQL{}
	_, err := roomSQL.ReadAllRooms()
	assert.Equal(t, ErrDBNotSet, err)
}
