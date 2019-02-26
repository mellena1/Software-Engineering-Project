package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// TimeslotMySQL implements TimeslotReaderWriterUpdaterDeleter
type TimeslotMySQL struct {
	db *sql.DB
}

// NewTimeslotMySQL makes a new TimeslotMySQL object given a db
func NewTimeslotMySQL(db *sql.DB) TimeslotMySQL {
	return TimeslotMySQL{db}
}

// ReadATimeslot reads a timeslot from the db given an id
func (t TimeslotMySQL) ReadATimeslot(id int64) (db.Timeslot, error) {
	return db.Timeslot{}, nil
}

// ReadAllTimeslots reads all timeslots from the db
func (t TimeslotMySQL) ReadAllTimeslots() ([]db.Timeslot, error) {
	return nil, nil
}

// WriteATimeslot writes a timeslot to the db
func (t TimeslotMySQL) WriteATimeslot(timeslot db.Timeslot) (int64, error) {
	if t.db == nil {
		return 0, ErrDBNotSet
	}

	statement, err := t.db.Prepare("INSERT INTO timeslot (`startTime`, `endTime`) VALUES (?, ?);")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(timeslot.StartTime.Format(db.TimeFormat), timeslot.EndTime.Format(db.TimeFormat))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateATimeslot updates a timeslot in the db given an id and the updated timeslot
func (t TimeslotMySQL) UpdateATimeslot(timeslot db.Timeslot) error {
	if t.db == nil {
		return ErrDBNotSet
	}

	statement, err := t.db.Prepare("UPDATE timeslot SET startTime = ?, endTime = ? WHERE timeslotID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(timeslot.StartTime.Format(db.TimeFormat), timeslot.EndTime.Format(db.TimeFormat), timeslot.ID)
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

// DeleteATimeslot deletes a timeslot given an id
func (t TimeslotMySQL) DeleteATimeslot(id int64) error {
	if t.db == nil {
		return ErrDBNotSet
	}

	statement, err := t.db.Prepare("DELETE FROM timeslot WHERE timeslotID = ?;")
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
