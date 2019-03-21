// +build integration

package mysql

import (
	"database/sql"
	"testing"
	"time"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
	mysqldb "github.com/mellena1/Software-Engineering-Project/backend/pkg/db/mysql"
	"github.com/stretchr/testify/assert"
)

var validTimeslotTest = db.Timeslot{
	ID:        db.Int64Ptr(1),
	StartTime: parseTime(time.RFC3339, "2019-02-18T21:00:00Z"),
	EndTime:   parseTime(time.RFC3339, "2019-02-18T22:00:00Z"),
}

func insertValidTimeslot() error {
	_, err := apiObj.db.Exec("INSERT INTO timeslot(`startTime`, `endTime`) VALUES (?, ?);",
		validTimeslotTest.StartTime.Format(mysqldb.MySQLTimeFormat), validTimeslotTest.EndTime.Format(mysqldb.MySQLTimeFormat))
	if err != nil {
		return err
	}
	return nil
}

func getTimeslot1() (db.Timeslot, error) {
	row := apiObj.db.QueryRow("SELECT timeslotID, startTime, endTime FROM timeslot WHERE timeslotID = 1;")
	actual := db.NewTimeslot()

	err := row.Scan(actual.ID, &actual.StartTime, &actual.EndTime)
	if err != nil {
		return db.Timeslot{}, err
	}
	return actual, nil
}

func parseTime(format, value string) time.Time {
	t, err := time.Parse(format, value)
	if err != nil {
		panic(err)
	}
	return t
}

func TestGetTimeslot(t *testing.T) {
	resetDB()

	err := insertValidTimeslot()
	if err != nil {
		t.Error(err)
	}

	expected := validTimeslotTest
	apiTester.Get("/api/v1/timeslot").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestGetAllTimeslots(t *testing.T) {
	resetDB()

	err := insertValidTimeslot()
	if err != nil {
		t.Error(err)
	}
	err = insertValidTimeslot()
	if err != nil {
		t.Error(err)
	}

	expected := []db.Timeslot{
		validTimeslotTest,
		db.Timeslot{
			ID:        db.Int64Ptr(2),
			StartTime: parseTime(time.RFC3339, "2019-02-18T21:00:00Z"),
			EndTime:   parseTime(time.RFC3339, "2019-02-18T22:00:00Z"),
		},
	}
	apiTester.Get("/api/v1/timeslots").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestGetInvalidTimeslotNotExist(t *testing.T) {
	resetDB()

	apiTester.Get("/api/v1/timeslot").
		AddQuery("id", "2").
		Expect(t).
		Status(503).
		Done()
}

func TestGetInvalidTimeslotBadQuery(t *testing.T) {
	resetDB()

	apiTester.Get("/api/v1/timeslot").
		AddQuery("id", "NaN").
		Expect(t).
		Status(400).
		Done()
}

func TestAddTimeslot(t *testing.T) {
	resetDB()

	val := api.WriteATimeslotRequest{
		StartTime: db.StringPtr(validTimeslotTest.StartTime.Format(time.RFC3339)),
		EndTime:   db.StringPtr(validTimeslotTest.EndTime.Format(time.RFC3339)),
	}
	apiTester.Post("/api/v1/timeslot").
		JSON(val).
		Expect(t).
		Status(200).
		JSON(map[string]int{"id": 1}).
		Done()

	expected := validTimeslotTest
	actual, err := getTimeslot1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestAddTimeslotInvalidNulls(t *testing.T) {
	resetDB()

	val := api.WriteATimeslotRequest{
		StartTime: db.StringPtr("2019-02-18T21:00:00Z"),
		EndTime:   nil,
	}
	apiTester.Post("/api/v1/timeslot").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateTimeslot(t *testing.T) {
	resetDB()

	err := insertValidTimeslot()
	if err != nil {
		t.Error(err)
	}

	val := api.UpdateATimeslotRequest{
		ID:        db.Int64Ptr(1),
		StartTime: db.StringPtr("2019-02-18T21:00:00Z"),
		EndTime:   db.StringPtr("2019-02-18T23:00:00Z"),
	}
	apiTester.Put("/api/v1/timeslot").
		JSON(val).
		Expect(t).
		Status(200).
		Done()

	expected := db.Timeslot{
		ID:        db.Int64Ptr(1),
		StartTime: parseTime(time.RFC3339, "2019-02-18T21:00:00Z"),
		EndTime:   parseTime(time.RFC3339, "2019-02-18T23:00:00Z"),
	}
	actual, err := getTimeslot1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestUpdateInvalidTimeslotNull(t *testing.T) {
	resetDB()

	err := insertValidTimeslot()
	if err != nil {
		t.Error(err)
	}

	val := api.UpdateATimeslotRequest{
		ID:        db.Int64Ptr(1),
		StartTime: db.StringPtr("2019-02-18T21:00:00Z"),
		EndTime:   nil,
	}
	apiTester.Put("/api/v1/timeslot").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateInvalidTimeslotNullID(t *testing.T) {
	resetDB()

	val := api.UpdateATimeslotRequest{
		ID:        nil,
		StartTime: db.StringPtr("2019-02-18T21:00:00Z"),
		EndTime:   db.StringPtr("2019-02-18T23:00:00Z"),
	}
	apiTester.Put("/api/v1/timeslot").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateInvalidTimeslotNotExist(t *testing.T) {
	resetDB()

	val := api.UpdateATimeslotRequest{
		ID:        db.Int64Ptr(2),
		StartTime: db.StringPtr("2019-02-18T21:00:00Z"),
		EndTime:   db.StringPtr("2019-02-18T23:00:00Z"),
	}
	apiTester.Put("/api/v1/timeslot").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestDeleteTimeslot(t *testing.T) {
	resetDB()

	err := insertValidTimeslot()
	if err != nil {
		t.Error(err)
	}

	apiTester.Delete("/api/v1/timeslot").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		Done()

	_, err = getTimeslot1()
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestDeleteInvalidTimeslotNotExist(t *testing.T) {
	resetDB()

	apiTester.Delete("/api/v1/timeslot").
		AddQuery("id", "2").
		Expect(t).
		Status(400).
		Done()
}

func TestDeleteInvalidTimeslotBadQuery(t *testing.T) {
	resetDB()

	apiTester.Delete("/api/v1/timeslot").
		AddQuery("id", "NaN").
		Expect(t).
		Status(400).
		Done()
}
