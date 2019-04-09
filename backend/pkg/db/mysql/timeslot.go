package mysql

import (
	"database/sql"
	"time"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

const (
	MySQLTimeFormat = "2006-01-02 15:04:05"
)

type TimeslotMySQL struct {
	db *sql.DB
}

func NewTimeslotMySQL(db *sql.DB) TimeslotMySQL {
	return TimeslotMySQL{db}
}

func scanATimeslot(timeslot *db.Timeslot, row rowScanner) error {
	return row.Scan(&timeslot.ID, &timeslot.StartTime, &timeslot.EndTime)
}

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
