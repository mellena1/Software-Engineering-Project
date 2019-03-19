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

var validSpeakerTest = db.Speaker{
	ID:        db.Int64Ptr(1),
	Email:     db.StringPtr("test@test.com"),
	FirstName: db.StringPtr("Bob"),
	LastName:  db.StringPtr("Smith"),
}

func insertValidSpeaker() error {
	_, err := apiObj.db.Exec("INSERT INTO speaker(`email`, `firstName`, `lastName`) VALUES (?, ?, ?);",
		validSpeakerTest.Email, validSpeakerTest.FirstName, validSpeakerTest.LastName)
	if err != nil {
		return err
	}
	return nil
}

func insertValidSpeakerWithNulls() error {
	_, err := apiObj.db.Exec("INSERT INTO speaker(`email`, `firstName`, `lastName`) VALUES (\"test@test.com\", NULL, NULL);")
	if err != nil {
		return err
	}
	return nil
}

func getSpeaker1() (db.Speaker, error) {
	row := apiObj.db.QueryRow("SELECT speakerID, email, firstName, lastName FROM speaker WHERE speakerID = 1;")
	actual := db.NewSpeaker()

	emailNull, fNameNull, lNameNull := sql.NullString{}, sql.NullString{}, sql.NullString{}
	err := row.Scan(actual.ID, &emailNull, &fNameNull, &lNameNull)
	if err != nil {
		return db.Speaker{}, err
	}
	actual.Email, actual.FirstName, actual.LastName = mysqldb.NullStringToString(emailNull), mysqldb.NullStringToString(fNameNull), mysqldb.NullStringToString(lNameNull)
	return actual, nil
}

func TestGetSpeaker(t *testing.T) {
	resetDB()

	err := insertValidSpeaker()
	if err != nil {
		t.Error(err)
	}

	expected := validSpeakerTest
	apiTester.Get("/api/v1/speaker").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestGetSpeakerWithNulls(t *testing.T) {
	resetDB()

	err := insertValidSpeakerWithNulls()
	if err != nil {
		t.Error(err)
	}

	expected := db.Speaker{
		ID:        db.Int64Ptr(1),
		Email:     db.StringPtr("test@test.com"),
		FirstName: nil,
		LastName:  nil,
	}
	apiTester.Get("/api/v1/speaker").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestGetInvalidSpeakerNotExist(t *testing.T) {
	resetDB()

	apiTester.Get("/api/v1/speaker").
		AddQuery("id", "2").
		Expect(t).
		Status(503).
		Done()
}

func TestGetInvalidSpeakerBadQuery(t *testing.T) {
	resetDB()

	apiTester.Get("/api/v1/speaker").
		AddQuery("id", "NaN").
		Expect(t).
		Status(400).
		Done()
}

func TestAddSpeaker(t *testing.T) {
	resetDB()

	val := api.WriteASpeakerRequest{
		Email:     validSpeakerTest.Email,
		FirstName: validSpeakerTest.FirstName,
		LastName:  validSpeakerTest.LastName,
	}
	apiTester.Post("/api/v1/speaker").
		JSON(val).
		Expect(t).
		Status(200).
		JSON(map[string]int{"id": 1}).
		Done()

	expected := validSpeakerTest
	actual, err := getSpeaker1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestAddSpeakerNulls(t *testing.T) {
	resetDB()

	val := api.WriteASpeakerRequest{
		Email:     db.StringPtr("test@test.com"),
		FirstName: nil,
		LastName:  nil,
	}
	apiTester.Post("/api/v1/speaker").
		JSON(val).
		Expect(t).
		Status(200).
		JSON(map[string]int{"id": 1}).
		Done()

	expected := db.Speaker{
		ID:        db.Int64Ptr(1),
		Email:     db.StringPtr("test@test.com"),
		FirstName: nil,
		LastName:  nil,
	}
	actual, err := getSpeaker1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestAddSpeakerInvalidAllNulls(t *testing.T) {
	resetDB()

	val := api.WriteASpeakerRequest{
		Email:     nil,
		FirstName: nil,
		LastName:  nil,
	}
	apiTester.Post("/api/v1/speaker").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateSpeaker(t *testing.T) {
	resetDB()

	err := insertValidSpeaker()
	if err != nil {
		t.Error(err)
	}

	val := api.UpdateASpeakerRequest{
		ID:        db.Int64Ptr(1),
		Email:     db.StringPtr("test123@test.com"),
		FirstName: db.StringPtr("Joe"),
		LastName:  db.StringPtr("Adams"),
	}
	apiTester.Put("/api/v1/speaker").
		JSON(val).
		Expect(t).
		Status(200).
		Done()

	expected := db.Speaker{
		ID:        db.Int64Ptr(1),
		Email:     db.StringPtr("test123@test.com"),
		FirstName: db.StringPtr("Joe"),
		LastName:  db.StringPtr("Adams"),
	}
	actual, err := getSpeaker1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestUpdateSpeakerNulls(t *testing.T) {
	resetDB()

	err := insertValidSpeaker()
	if err != nil {
		t.Error(err)
	}

	val := api.UpdateASpeakerRequest{
		ID:        db.Int64Ptr(1),
		Email:     db.StringPtr("test123@test.com"),
		FirstName: nil,
		LastName:  nil,
	}
	apiTester.Put("/api/v1/speaker").
		JSON(val).
		Expect(t).
		Status(200).
		Done()

	expected := db.Speaker{
		ID:        db.Int64Ptr(1),
		Email:     db.StringPtr("test123@test.com"),
		FirstName: nil,
		LastName:  nil,
	}
	actual, err := getSpeaker1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestUpdateInvalidSpeakerNullID(t *testing.T) {
	resetDB()

	val := api.UpdateASpeakerRequest{
		ID:        nil,
		Email:     db.StringPtr("test123@test.com"),
		FirstName: db.StringPtr("Joe"),
		LastName:  db.StringPtr("Adams"),
	}
	apiTester.Put("/api/v1/speaker").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateInvalidSpeakerNotExist(t *testing.T) {
	resetDB()

	val := api.UpdateASpeakerRequest{
		ID:        db.Int64Ptr(2),
		Email:     db.StringPtr("test123@test.com"),
		FirstName: db.StringPtr("Joe"),
		LastName:  db.StringPtr("Adams"),
	}
	apiTester.Put("/api/v1/speaker").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestDeleteSpeaker(t *testing.T) {
	resetDB()

	err := insertValidSpeaker()
	if err != nil {
		t.Error(err)
	}

	apiTester.Delete("/api/v1/speaker").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		Done()

	_, err = getSpeaker1()
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestDeleteInvalidSpeakerNotExist(t *testing.T) {
	resetDB()

	apiTester.Delete("/api/v1/speaker").
		AddQuery("id", "2").
		Expect(t).
		Status(400).
		Done()
}

func TestDeleteInvalidSpeakerBadQuery(t *testing.T) {
	resetDB()

	apiTester.Delete("/api/v1/speaker").
		AddQuery("id", "NaN").
		Expect(t).
		Status(400).
		Done()
}
