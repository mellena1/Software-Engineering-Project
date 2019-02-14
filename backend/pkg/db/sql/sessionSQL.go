package sql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

type SessionSQL struct {
	db *sql.DB
}

func NewSessionSQL(db *sql.DB) SessionSQL {
	return SessionSQL{db}
}

func (SessionSQL) ReadASession(sessionID int) (db.Session, error) {
	return db.Session{}, nil
}

func (SessionSQL) ReadAllSessions() ([]db.Session, error) {
	return []db.Session{}, nil
}

func (SessionSQL) WriteASession(s db.Session) error {
	return nil
}

func (SessionSQL) UpdateASession(sessionID int, newSession db.Session) error {
	return nil
}

func (SessionSQL) DeleteASession(sessionID int) error {
	return nil
}
