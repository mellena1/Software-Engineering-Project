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
func (SessionMySQL) WriteASession(s db.Session) error {
	return nil
}

// UpdateASession updates a session in the db given a sessionID and the updated session
func (SessionMySQL) UpdateASession(sessionID int, newSession db.Session) error {
	return nil
}

// DeleteASession deletes a room given a sessionID
func (SessionMySQL) DeleteASession(sessionID int) error {
	return nil
}
