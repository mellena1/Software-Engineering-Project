package mysql

import (
	"database/sql"
	"time"

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

// scanATimeslot takes in a timeslot pointer and scans a row into it
func scanATimeslot(timeslot *db.Timeslot, row rowScanner) error {
	return row.Scan(&timeslot.ID, &timeslot.StartTime, &timeslot.EndTime)
}

// ReadATimeslot reads a timeslot from the db given an id
func (t TimeslotMySQL) ReadATimeslot(id int64) (db.Timeslot, error) {
	if t.db == nil {
		return db.Timeslot{}, ErrDBNotSet
	}

	stmt, err := t.db.Prepare("SELECT timeslotID, startTime, endTime FROM timeslot where timeslotID = ?;")
	if err != nil {
		return db.Timeslot{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	timeslot := db.NewTimeslot()
	err = scanATimeslot(&timeslot, row)

	return timeslot, err
}

// ReadAllTimeslots reads all timeslots from the db
func (t TimeslotMySQL) ReadAllTimeslots() ([]db.Timeslot, error) {
	if t.db == nil {
		return nil, ErrDBNotSet
	}

	q := "SELECT timeslotID, startTime, endTime FROM timeslot;"

	rows, err := t.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	timeslots := []db.Timeslot{}
	for rows.Next() {
		newTimeslot := db.NewTimeslot()
		scanATimeslot(&newTimeslot, rows)
		timeslots = append(timeslots, newTimeslot)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return timeslots, nil
}

// WriteATimeslot writes a timeslot to the db
func (t TimeslotMySQL) WriteATimeslot(startTime, endTime time.Time) (int64, error) {
	if t.db == nil {
		return 0, ErrDBNotSet
	}

	stmt, err := t.db.Prepare("INSERT INTO timeslot (`startTime`, `endTime`) VALUES (?, ?);")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(startTime.Format(db.TimeFormat), endTime.Format(db.TimeFormat))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateATimeslot updates a timeslot in the db given an id and the updated timeslot
func (t TimeslotMySQL) UpdateATimeslot(id int64, startTime, endTime time.Time) error {
	if t.db == nil {
		return ErrDBNotSet
	}

	stmt, err := t.db.Prepare("UPDATE timeslot SET startTime = ?, endTime = ? WHERE timeslotID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(startTime.Format(db.TimeFormat), endTime.Format(db.TimeFormat), id)
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

	stmt, err := t.db.Prepare("DELETE FROM timeslot WHERE timeslotID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
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
