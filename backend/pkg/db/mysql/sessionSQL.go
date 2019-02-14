package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// SessionSQL implements SessionReaderWriterUpdaterDeleter
type SessionSQL struct {
	db *sql.DB
}

// NewSessionSQL makes a new SessionSQL object given a db
func NewSessionSQL(db *sql.DB) SessionSQL {
	return SessionSQL{db}
}

// ReadASession reads a session from the db given sessionID
func (SessionSQL) ReadASession(sessionID int) (db.Session, error) {
	return db.Session{}, nil
}

// ReadAllSessions reads all sessions from the db
func (SessionSQL) ReadAllSessions() ([]db.Session, error) {
	return []db.Session{}, nil
}

// WriteASession writes a session to the db
func (SessionSQL) WriteASession(s db.Session) error {
	return nil
}

// UpdateASession updates a session in the db given a sessionID and the updated session
func (SessionSQL) UpdateASession(sessionID int, newSession db.Session) error {
	return nil
}

// DeleteASession deletes a room given a sessionID
func (SessionSQL) DeleteASession(sessionID int) error {
	return nil
}
