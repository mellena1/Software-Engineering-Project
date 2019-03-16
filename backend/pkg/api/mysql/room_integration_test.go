// +build integration

package mysql

import (
	"database/sql"
	"testing"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
	"github.com/stretchr/testify/assert"
)

func insertValidRoom() error {
	_, err := apiObj.db.Exec("INSERT INTO room(`roomName`, `capacity`) VALUES (\"beatty\", 100);")
	if err != nil {
		return err
	}
	return nil
}

func TestGetRoom(t *testing.T) {
	resetDB(apiObj)

	insertValidRoom()

	expected := db.Room{
		ID:       db.Int64Ptr(1),
		Name:     db.StringPtr("beatty"),
		Capacity: db.Int64Ptr(100),
	}
	apiTester.Get("/api/v1/room").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestAddRoom(t *testing.T) {
	resetDB(apiObj)

	val := api.WriteARoomRequest{
		Name:     "beatty",
		Capacity: db.Int64Ptr(100),
	}
	apiTester.Post("/api/v1/room").
		JSON(val).
		Expect(t).
		Status(200).
		JSON(map[string]int{"id": 1}).
		Done()

	expected := db.Room{
		ID:       db.Int64Ptr(1),
		Name:     db.StringPtr("beatty"),
		Capacity: db.Int64Ptr(100),
	}
	row := apiObj.db.QueryRow("SELECT roomID, roomName, capacity FROM room WHERE roomID = 1;")
	actual := db.NewRoom()
	err := row.Scan(actual.ID, actual.Name, actual.Capacity)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestUpdateRoom(t *testing.T) {
	resetDB(apiObj)

	insertValidRoom()

	val := api.UpdateARoomRequest{
		ID:       db.Int64Ptr(1),
		Name:     "wentworth",
		Capacity: db.Int64Ptr(25),
	}
	apiTester.Put("/api/v1/room").
		JSON(val).
		Expect(t).
		Status(200).
		Done()

	expected := db.Room{
		ID:       db.Int64Ptr(1),
		Name:     db.StringPtr("wentworth"),
		Capacity: db.Int64Ptr(25),
	}
	row := apiObj.db.QueryRow("SELECT roomID, roomName, capacity FROM room WHERE roomID = 1;")
	actual := db.NewRoom()
	err := row.Scan(actual.ID, actual.Name, actual.Capacity)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestDeleteRoom(t *testing.T) {
	resetDB(apiObj)

	insertValidRoom()

	apiTester.Delete("/api/v1/room").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		Done()

	row := apiObj.db.QueryRow("SELECT roomID, roomName, capacity FROM room WHERE roomID = 1;")
	err := row.Scan()
	assert.Equal(t, sql.ErrNoRows, err)
}
