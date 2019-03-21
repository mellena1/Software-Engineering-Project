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
	return row.Scan(&count.Time, &count.UserID, &count.SessionID, &count.Count)
}

// ReadACount reads a count from the db given a sessionID
func (c CountMySQL) ReadACount(sessionID int64) ([]db.Count, error) {
	if c.db == nil {
		return nil, ErrDBNotSet
	}

	stmt, err := c.db.Prepare("SELECT time, userID, sessionID, count FROM count where sessionID = ?;")
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

	q := "SELECT time, sessionID, userID, count FROM count;"

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
