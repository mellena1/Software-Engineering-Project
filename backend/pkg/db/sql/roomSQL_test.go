package sql

import (
	"testing"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var columns = []string{"roomID", "capacity"}

func testReadAllRoomsValid(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occured when opening the stub db connection: %s", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT * FROM room").
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1,1\n2,2"))

	// Execute ReadAllRooms
	roomSQL := NewRoomSQL(db)
	actual, err := roomSQL.ReadAllRooms()
	if err != nil {
		t.Fatalf("an error occured when running ReadAllRooms(): %s", err)
	}

	// Make sure expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Make sure returned Rooms are correct
	expected := []entities.Room{
		entities.Room{ID: 1, Capacity: 1},
		entities.Room{ID: 2, Capacity: 2},
	}
	assert.Equal(t, actual, expected)
}

func testReadAllRoomsInvalid(t *testing.T) {
	roomSQL := RoomSQL{}
	_, err := roomSQL.ReadAllRooms()
	assert.Equal(t, err, ErrDBNotSet)
}
