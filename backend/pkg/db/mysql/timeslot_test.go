package mysql

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
	"github.com/stretchr/testify/assert"
)

var (
	timeslotColumns       = []string{"timeslotID", "startTime", "endTime"}
	timeslotID      int64 = 15
	startTime             = "2019-02-18 21:00:00"
	endTime               = "2019-10-01 23:00:00"
	testTimeslot          = db.Timeslot{
		ID:        timeslotID,
		StartTime: parseTime(db.TimeFormat, startTime),
		EndTime:   parseTime(db.TimeFormat, endTime),
	}
)

func parseTime(format, value string) time.Time {
	t, err := time.Parse(format, value)
	if err != nil {
		panic(err)
	}
	return t
}

func TestWriteATimeslotValid(t *testing.T) {
	mockdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error occured when opening the stub db connection: %s", err)
	}
	defer mockdb.Close()

	mock.ExpectPrepare("INSERT INTO timeslot (`startTime`, `endTime`) VALUES (?, ?);").ExpectExec().
		WithArgs(startTime, endTime).
		WillReturnResult(sqlmock.NewResult(int64(timeslotID), 1))

	// Execute WriteATimeslot
	timeslotSQL := NewTimeslotMySQL(mockdb)
	id, err := timeslotSQL.WriteATimeslot(testTimeslot.StartTime, testTimeslot.EndTime)
	if err != nil {
		t.Fatalf("an error occured when running WriteATimeslot: %s", err)
	}

	// Make sure expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.Equal(t, timeslotID, id)
}

func TestWriteATimeslotInvalid(t *testing.T) {
	timeslotSQL := TimeslotMySQL{}
	_, err := timeslotSQL.WriteATimeslot(testTimeslot.StartTime, testTimeslot.EndTime)
	assert.Equal(t, ErrDBNotSet, err)
}

func TestUpdateATimeslotValid(t *testing.T) {
	mockdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error occured when opening the stub db connection: %s", err)
	}
	defer mockdb.Close()

	mock.ExpectPrepare("UPDATE timeslot SET startTime = ?, endTime = ? WHERE timeslotID = ?;").ExpectExec().
		WithArgs(startTime, endTime, timeslotID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute UpdateATimeslot
	timeslotSQL := NewTimeslotMySQL(mockdb)
	err = timeslotSQL.UpdateATimeslot(testTimeslot.ID, testTimeslot.StartTime, testTimeslot.EndTime)
	if err != nil {
		t.Fatalf("an error occured when running UpdateATimeslot: %s", err)
	}

	// Make sure expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateATimeslotInvalid(t *testing.T) {
	timeslotSQL := TimeslotMySQL{}
	err := timeslotSQL.UpdateATimeslot(testTimeslot.ID, testTimeslot.StartTime, testTimeslot.EndTime)
	assert.Equal(t, ErrDBNotSet, err)
}

func TestDeleteATimeslotValid(t *testing.T) {
	mockdb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error occured when opening the stub db connection: %s", err)
	}
	defer mockdb.Close()

	mock.ExpectPrepare("DELETE FROM timeslot WHERE timeslotID = ?;").ExpectExec().
		WithArgs(timeslotID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute UpdateATimeslot
	timeslotSQL := NewTimeslotMySQL(mockdb)
	err = timeslotSQL.DeleteATimeslot(timeslotID)
	if err != nil {
		t.Fatalf("an error occured when running DeleteATimeslot: %s", err)
	}

	// Make sure expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteATimeslotInvalid(t *testing.T) {
	timeslotSQL := TimeslotMySQL{}
	err := timeslotSQL.DeleteATimeslot(timeslotID)
	assert.Equal(t, ErrDBNotSet, err)
}
