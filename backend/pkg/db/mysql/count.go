package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// CountMySQL implements CountReaderWriterUpdaterDeleter
type CountMySQL struct {
	db *sql.DB
}

// NewCountMySQL makes a new CountMySQL object given a db
func NewCountMySQL(db *sql.DB) CountMySQL {
	return CountMySQL{db}
}

// scanACount takes in a count pointer and scans a row into it
func scanACount(count *db.Count, row rowScanner) error {
	return row.Scan(&count.Time, &count.SessionID, &count.UserName, &count.Count)
}

// ReadCountsOfSession reads a count from the db given a sessionID
func (c CountMySQL) ReadCountsOfSession(sessionID int64) ([]db.Count, error) {
	if c.db == nil {
		return nil, ErrDBNotSet
	}

	stmt, err := c.db.Prepare("SELECT time, sessionID, userName, count FROM count where sessionID = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := []db.Count{}
	for rows.Next() {
		newCount := db.NewCount()
		scanACount(&newCount, rows)
		counts = append(counts, newCount)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return counts, nil
}

// ReadAllCounts reads all counts from the db
func (c CountMySQL) ReadAllCounts() ([]db.Count, error) {
	if c.db == nil {
		return nil, ErrDBNotSet
	}

	q := "SELECT time, sessionID, userName, count FROM count;"

	rows, err := c.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := []db.Count{}
	for rows.Next() {
		newCount := db.NewCount()
		scanACount(&newCount, rows)
		counts = append(counts, newCount)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return counts, nil
}

// WriteACount writes a count to the db
func (c CountMySQL) WriteACount(time *string, sessionID *int64, userName *string, count *int64) (int64, error) {
	if c.db == nil {
		return 0, ErrDBNotSet
	}
	stmt, err := c.db.Prepare("INSERT INTO `count`(`time`, `sessionID`, `userName`, `count`) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(StringToNullString(time), IntToNullInt(sessionID), StringToNullString(userName), IntToNullInt(count))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateACount updates a count in the db given the session Id and the updated time of session
func (c CountMySQL) UpdateACount(time *string, sessionID *int64, userName *string, count *int64) error {
	if c.db == nil {
		return ErrDBNotSet
	}

	stmt, err := c.db.Prepare("UPDATE count SET time = ?, sessionID = ?, userName = ?, count = ? WHERE time = ? and sessionID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(StringToNullString(time), IntToNullInt(sessionID), StringToNullString(userName), IntToNullInt(count), StringToNullString(time), IntToNullInt(sessionID))
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

// DeleteACount deletes a count given an id
func (c CountMySQL) DeleteACount(sessionID int64) error {
	if c.db == nil {
		return ErrDBNotSet
	}

	stmt, err := c.db.Prepare("DELETE FROM count WHERE sessionID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(sessionID)
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