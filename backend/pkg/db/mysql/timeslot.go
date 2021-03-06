package mysql

import (
	"database/sql"
	"time"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

const (
	// MySQLTimeFormat format used to write to the mysql db
	MySQLTimeFormat = "2006-01-02 15:04:05"
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
func (myTimeslotSQL TimeslotMySQL) ReadATimeslot(id int64) (db.Timeslot, error) {
	if myTimeslotSQL.db == nil {
		return db.Timeslot{}, ErrDBNotSet
	}

	statement, err := myTimeslotSQL.db.Prepare("SELECT timeslotID, startTime, endTime FROM timeslot where timeslotID = ?;")
	if err != nil {
		return db.Timeslot{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)

	timeslot := db.NewTimeslot()
	err = scanATimeslot(&timeslot, row)

	return timeslot, err
}

// ReadAllTimeslots reads all timeslots from the db
func (myTimeslotSQL TimeslotMySQL) ReadAllTimeslots() ([]db.Timeslot, error) {
	if myTimeslotSQL.db == nil {
		return nil, ErrDBNotSet
	}

	query := "SELECT timeslotID, startTime, endTime FROM timeslot;"

	rows, err := myTimeslotSQL.db.Query(query)
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
func (myTimeslotSQL TimeslotMySQL) WriteATimeslot(startTime, endTime time.Time) (int64, error) {
	if myTimeslotSQL.db == nil {
		return 0, ErrDBNotSet
	}

	statement, err := myTimeslotSQL.db.Prepare("INSERT INTO timeslot (`startTime`, `endTime`) VALUES (?, ?);")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(startTime.Format(MySQLTimeFormat), endTime.Format(MySQLTimeFormat))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateATimeslot updates a timeslot in the db given an id and the updated timeslot
func (myTimeslotSQL TimeslotMySQL) UpdateATimeslot(id int64, startTime, endTime time.Time) error {
	if myTimeslotSQL.db == nil {
		return ErrDBNotSet
	}

	statement, err := myTimeslotSQL.db.Prepare("UPDATE timeslot SET startTime = ?, endTime = ? WHERE timeslotID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(startTime.Format(MySQLTimeFormat), endTime.Format(MySQLTimeFormat), id)
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
func (myTimeslotSQL TimeslotMySQL) DeleteATimeslot(id int64) error {
	if myTimeslotSQL.db == nil {
		return ErrDBNotSet
	}

	statement, err := myTimeslotSQL.db.Prepare("DELETE FROM timeslot WHERE timeslotID = ?;")
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
