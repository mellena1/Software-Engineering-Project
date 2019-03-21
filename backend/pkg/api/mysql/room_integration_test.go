// +build integration

package mysql

import (
	"database/sql"
	"testing"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
	mysqldb "github.com/mellena1/Software-Engineering-Project/backend/pkg/db/mysql"
	"github.com/stretchr/testify/assert"
)

var validRoomTest = db.Room{
	ID:       db.Int64Ptr(1),
	Name:     db.StringPtr("beatty"),
	Capacity: db.Int64Ptr(100),
}

func insertValidRoom() error {
	_, err := apiObj.db.Exec("INSERT INTO room(`roomName`, `capacity`) VALUES (?, ?);", validRoomTest.Name, validRoomTest.Capacity)
	if err != nil {
		return err
	}
	return nil
}

func insertValidRoomWithNullCapacity() error {
	_, err := apiObj.db.Exec("INSERT INTO room(`roomName`, `capacity`) VALUES (\"beatty\", NULL);")
	if err != nil {
		return err
	}
	return nil
}

func getRoom1() (db.Room, error) {
	row := apiObj.db.QueryRow("SELECT roomID, roomName, capacity FROM room WHERE roomID = 1;")
	actual := db.NewRoom()

	capNull := sql.NullInt64{}
	err := row.Scan(actual.ID, actual.Name, &capNull)
	if err != nil {
		return db.Room{}, err
	}
	actual.Capacity = mysqldb.NullIntToInt(capNull)

	return actual, nil
}

func TestGetRoom(t *testing.T) {
	resetDB()

	err := insertValidRoom()
	if err != nil {
		t.Error(err)
	}

	expected := validRoomTest
	apiTester.Get("/api/v1/room").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestGetRoomNullCapacity(t *testing.T) {
	resetDB()

	err := insertValidRoomWithNullCapacity()
	if err != nil {
		t.Error(err)
	}

	expected := db.Room{
		ID:       db.Int64Ptr(1),
		Name:     db.StringPtr("beatty"),
		Capacity: nil,
	}
	apiTester.Get("/api/v1/room").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestGetInvalidRoomNotExist(t *testing.T) {
	resetDB()

	apiTester.Get("/api/v1/room").
		AddQuery("id", "2").
		Expect(t).
		Status(503).
		Done()
}

func TestGetInvalidRoomBadQuery(t *testing.T) {
	resetDB()

	apiTester.Get("/api/v1/room").
		AddQuery("id", "NaN").
		Expect(t).
		Status(400).
		Done()
}

func TestGetAllRooms(t *testing.T) {
	resetDB()

	err := insertValidRoom()
	if err != nil {
		t.Error(err)
	}
	err = insertValidRoomWithNullCapacity()
	if err != nil {
		t.Error(err)
	}

	expected := []db.Room{
		validRoomTest,
		db.Room{
			ID:       db.Int64Ptr(2),
			Name:     db.StringPtr("beatty"),
			Capacity: nil,
		},
	}
	apiTester.Get("/api/v1/rooms").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestAddRoom(t *testing.T) {
	resetDB()

	val := api.WriteARoomRequest{
		Name:     *validRoomTest.Name,
		Capacity: validRoomTest.Capacity,
	}
	apiTester.Post("/api/v1/room").
		JSON(val).
		Expect(t).
		Status(200).
		JSON(map[string]int{"id": 1}).
		Done()

	expected := validRoomTest
	actual, err := getRoom1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestAddRoomNullCapacity(t *testing.T) {
	resetDB()

	val := api.WriteARoomRequest{
		Name:     "beatty",
		Capacity: nil,
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
		Capacity: nil,
	}
	actual, err := getRoom1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestAddInvalidRoomEmptyName(t *testing.T) {
	resetDB()

	val := api.WriteARoomRequest{
		Name:     "",
		Capacity: db.Int64Ptr(100),
	}
	apiTester.Post("/api/v1/room").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateRoom(t *testing.T) {
	resetDB()

	err := insertValidRoom()
	if err != nil {
		t.Error(err)
	}

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
	actual, err := getRoom1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestUpdateRoomNullCapacity(t *testing.T) {
	resetDB()

	err := insertValidRoom()
	if err != nil {
		t.Error(err)
	}

	val := api.UpdateARoomRequest{
		ID:       db.Int64Ptr(1),
		Name:     "wentworth",
		Capacity: nil,
	}
	apiTester.Put("/api/v1/room").
		JSON(val).
		Expect(t).
		Status(200).
		Done()

	expected := db.Room{
		ID:       db.Int64Ptr(1),
		Name:     db.StringPtr("wentworth"),
		Capacity: nil,
	}
	actual, err := getRoom1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestUpdateInvalidRoomNullID(t *testing.T) {
	resetDB()

	val := api.UpdateARoomRequest{
		ID:       nil,
		Name:     "wentworth",
		Capacity: db.Int64Ptr(25),
	}
	apiTester.Put("/api/v1/room").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateInvalidRoomNotExist(t *testing.T) {
	resetDB()

	val := api.UpdateARoomRequest{
		ID:       db.Int64Ptr(2),
		Name:     "wentworth",
		Capacity: db.Int64Ptr(25),
	}
	apiTester.Put("/api/v1/room").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestDeleteRoom(t *testing.T) {
	resetDB()

	err := insertValidRoom()
	if err != nil {
		t.Error(err)
	}

	apiTester.Delete("/api/v1/room").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		Done()

	_, err = getRoom1()
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestDeleteInvalidRoomNotExist(t *testing.T) {
	resetDB()

	apiTester.Delete("/api/v1/room").
		AddQuery("id", "2").
		Expect(t).
		Status(400).
		Done()
}

func TestDeleteInvalidRoomBadQuery(t *testing.T) {
	resetDB()

	apiTester.Delete("/api/v1/room").
		AddQuery("id", "NaN").
		Expect(t).
		Status(400).
		Done()
}
