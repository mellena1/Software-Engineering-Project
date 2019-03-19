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

var validSessionTest = db.Session{
	ID:       db.Int64Ptr(1),
	Timeslot: &validTimeslotTest,
	Speaker:  &validSpeakerTest,
	Room:     &validRoomTest,
	Name:     db.StringPtr("session 1"),
}

func insertValidSessionRequirements() error {
	err := insertValidRoom()
	if err != nil {
		return err
	}
	err = insertValidSpeaker()
	if err != nil {
		return err
	}
	err = insertValidTimeslot()
	if err != nil {
		return err
	}
	return nil
}

func insertValidSession() error {
	err := insertValidSessionRequirements()
	if err != nil {
		return err
	}

	_, err = apiObj.db.Exec("INSERT INTO session(`speakerID`, `roomID`, `timeslotID`, `sessionName`) VALUES (?, ?, ?, ?);",
		validSessionTest.Speaker.ID, validSessionTest.Room.ID, validSessionTest.Timeslot.ID, validSessionTest.Name)
	if err != nil {
		return err
	}
	return nil
}

func insertValidSessionNulls() error {
	_, err := apiObj.db.Exec("INSERT INTO session(`speakerID`, `roomID`, `timeslotID`, `sessionName`) VALUES (NULL, NULL, NULL, \"session 1\");")
	if err != nil {
		return err
	}
	return nil
}

// scanASession takes in a session pointer and scans a row into it
func scanASession(session *db.Session, row *sql.Row) error {
	speakerID, roomID, timeslotID := sql.NullInt64{}, sql.NullInt64{}, sql.NullInt64{}
	speakerEmail, speakerFirstName, speakerLastName := sql.NullString{}, sql.NullString{}, sql.NullString{}
	roomCapacity := sql.NullInt64{}

	err := row.Scan(&session.ID, &speakerID, &speakerEmail, &speakerFirstName,
		&speakerLastName, &roomID, session.Room.Name,
		&roomCapacity, &timeslotID, &session.Timeslot.StartTime,
		&session.Timeslot.EndTime, session.Name)

	if speakerID.Valid {
		session.Speaker.ID = mysqldb.NullIntToInt(speakerID)
		session.Speaker.Email = mysqldb.NullStringToString(speakerEmail)
		session.Speaker.FirstName = mysqldb.NullStringToString(speakerFirstName)
		session.Speaker.LastName = mysqldb.NullStringToString(speakerLastName)
	} else {
		session.Speaker = nil
	}

	if roomID.Valid {
		session.Room.ID = mysqldb.NullIntToInt(roomID)
		session.Room.Capacity = mysqldb.NullIntToInt(roomCapacity)
	} else {
		session.Room = nil
	}

	if timeslotID.Valid {
		session.Timeslot.ID = mysqldb.NullIntToInt(timeslotID)
	} else {
		session.Timeslot = nil
	}

	return err
}

func getSession1() (db.Session, error) {
	row := apiObj.db.QueryRow(`
		SELECT session.sessionID, speaker.speakerID, speaker.email, 
			speaker.firstName, speaker.lastName, room.roomID, room.roomName, room.capacity, 
			timeslot.timeslotID, timeslot.startTime, timeslot.endTime, session.sessionName 
		FROM session
		LEFT JOIN speaker ON session.speakerID = speaker.speakerID
		LEFT JOIN room ON session.roomID = room.roomID
		LEFT JOIN timeslot ON session.timeslotID = timeslot.timeslotID
		WHERE session.sessionID = 1;`)
	actual := db.NewSession()

	err := scanASession(&actual, row)
	if err != nil {
		return db.Session{}, err
	}
	return actual, nil
}

func TestGetSession(t *testing.T) {
	resetDB()

	err := insertValidSession()
	if err != nil {
		t.Error(err)
	}

	expected := validSessionTest
	apiTester.Get("/api/v1/session").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestGetSessionNulls(t *testing.T) {
	resetDB()

	err := insertValidSessionNulls()
	if err != nil {
		t.Error(err)
	}

	expected := db.Session{
		ID:       db.Int64Ptr(1),
		Timeslot: nil,
		Speaker:  nil,
		Room:     nil,
		Name:     db.StringPtr("session 1"),
	}
	apiTester.Get("/api/v1/session").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		JSON(expected).
		Done()
}

func TestGetInvalidSessionNotExist(t *testing.T) {
	resetDB()

	apiTester.Get("/api/v1/session").
		AddQuery("id", "2").
		Expect(t).
		Status(503).
		Done()
}

func TestGetInvalidSessionBadQuery(t *testing.T) {
	resetDB()

	apiTester.Get("/api/v1/session").
		AddQuery("id", "NaN").
		Expect(t).
		Status(400).
		Done()
}

func TestAddSession(t *testing.T) {
	resetDB()

	insertValidSessionRequirements()

	val := api.WriteASessionRequest{
		SpeakerID:   validSessionTest.Speaker.ID,
		RoomID:      validSessionTest.Room.ID,
		TimeslotID:  validSessionTest.Timeslot.ID,
		SessionName: validSessionTest.Name,
	}
	apiTester.Post("/api/v1/session").
		JSON(val).
		Expect(t).
		Status(200).
		JSON(map[string]int{"id": 1}).
		Done()

	expected := validSessionTest
	actual, err := getSession1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestAddSessionNulls(t *testing.T) {
	resetDB()

	insertValidSessionRequirements()

	val := api.WriteASessionRequest{
		SpeakerID:   nil,
		RoomID:      nil,
		TimeslotID:  nil,
		SessionName: validSessionTest.Name,
	}
	apiTester.Post("/api/v1/session").
		JSON(val).
		Expect(t).
		Status(200).
		JSON(map[string]int{"id": 1}).
		Done()

	expected := db.Session{
		ID:       db.Int64Ptr(1),
		Speaker:  nil,
		Room:     nil,
		Timeslot: nil,
		Name:     validSessionTest.Name,
	}
	actual, err := getSession1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestAddSessionInvalidNulls(t *testing.T) {
	resetDB()

	val := api.WriteASessionRequest{
		SpeakerID:   nil,
		RoomID:      nil,
		TimeslotID:  nil,
		SessionName: nil,
	}
	apiTester.Post("/api/v1/session").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateSession(t *testing.T) {
	resetDB()

	insertValidSession()

	val := api.UpdateASessionRequest{
		SessionID:   db.Int64Ptr(1),
		SpeakerID:   validSessionTest.Speaker.ID,
		RoomID:      validSessionTest.Room.ID,
		TimeslotID:  validSessionTest.Timeslot.ID,
		SessionName: db.StringPtr("updated"),
	}
	apiTester.Put("/api/v1/session").
		JSON(val).
		Expect(t).
		Status(200).
		Done()

	expected := db.Session{
		ID:       db.Int64Ptr(1),
		Timeslot: &validTimeslotTest,
		Speaker:  &validSpeakerTest,
		Room:     &validRoomTest,
		Name:     db.StringPtr("updated"),
	}
	actual, err := getSession1()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

func TestUpdateInvalidSessionNull(t *testing.T) {
	resetDB()

	insertValidSession()

	val := api.UpdateASessionRequest{
		SessionID:   db.Int64Ptr(1),
		SpeakerID:   nil,
		RoomID:      nil,
		TimeslotID:  nil,
		SessionName: nil,
	}
	apiTester.Put("/api/v1/session").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateInvalidSessionNullID(t *testing.T) {
	resetDB()

	val := api.UpdateASessionRequest{
		SessionID:   nil,
		SpeakerID:   validSessionTest.Speaker.ID,
		RoomID:      validSessionTest.Room.ID,
		TimeslotID:  validSessionTest.Timeslot.ID,
		SessionName: db.StringPtr("updated"),
	}
	apiTester.Put("/api/v1/session").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestUpdateInvalidSessionNotExist(t *testing.T) {
	resetDB()

	val := api.UpdateASessionRequest{
		SessionID:   db.Int64Ptr(2),
		SpeakerID:   validSessionTest.Speaker.ID,
		RoomID:      validSessionTest.Room.ID,
		TimeslotID:  validSessionTest.Timeslot.ID,
		SessionName: db.StringPtr("updated"),
	}
	apiTester.Put("/api/v1/session").
		JSON(val).
		Expect(t).
		Status(400).
		Done()
}

func TestDeleteSession(t *testing.T) {
	resetDB()

	insertValidSession()

	apiTester.Delete("/api/v1/session").
		AddQuery("id", "1").
		Expect(t).
		Status(200).
		Done()

	_, err := getSession1()
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestDeleteInvalidSessionNotExist(t *testing.T) {
	resetDB()

	apiTester.Delete("/api/v1/session").
		AddQuery("id", "2").
		Expect(t).
		Status(400).
		Done()
}

func TestDeleteInvalidSessionBadQuery(t *testing.T) {
	resetDB()

	apiTester.Delete("/api/v1/session").
		AddQuery("id", "NaN").
		Expect(t).
		Status(400).
		Done()
}
