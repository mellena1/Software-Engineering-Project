package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// SessionMySQL implements SessionReaderWriterUpdaterDeleter
type SessionMySQL struct {
	db *sql.DB
}

// NewSessionMySQL makes a new SessionMySQL object given a db
func NewSessionMySQL(db *sql.DB) SessionMySQL {
	return SessionMySQL{db}
}

// ReadASession reads a session from the db given sessionID
func (SessionMySQL) ReadASession(sessionID int) (db.Session, error) {
	return db.Session{}, nil
}

// ReadAllSessions reads all sessions from the db
func (SessionMySQL) ReadAllSessions() ([]db.Session, error) {
	return []db.Session{}, nil
}

// WriteASession writes a session to the db
func (s SessionMySQL) WriteASession(speakerID *int, roomID *int, timeslotID *int, name *string) (int64, error) {
	if s.db == nil {
		return 0, ErrDBNotSet
	}

	statement, err := s.db.Prepare("INSERT INTO session (`speakerID`, `roomID`, `timeslotID`, `sessionName`) VALUES (?, ?, ?, ?);")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(speakerID, roomID, timeslotID, name)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateASession updates a session in the db given a sessionID and the updated session
func (s SessionMySQL) UpdateASession(sessionID int, speakerID *int, roomID *int, timeslotID *int, name *string) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	statement, err := s.db.Prepare("UPDATE session SET speakerID = ?, roomID = ?, timeslotID = ?, sessionName = ? WHERE sessionID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(speakerID, roomID, timeslotID, name, sessionID)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return err
	} else if rows == 0 {
		return ErrNothingChanged
	}

	return nil
}

// DeleteASession deletes a room given a sessionID
func (s SessionMySQL) DeleteASession(sessionID int64) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	statement, err := s.db.Prepare("DELETE FROM session WHERE sessionID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(sessionID)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return err
	} else if rows == 0 {
		return ErrNothingChanged
	}

	return nil
}
