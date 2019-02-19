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
func (s SessionMySQL) ReadASession(sessionID int) (db.Session, error) {
	if s.db == nil {
		return db.Session{}, ErrDBNotSet
	}

	q := "SELECT * FROM session WHERE sessionID = " + string(sessionID) + ";"

	rows, err := s.db.Query(q)
	if err != nil {
		return db.Session{}, err
	}
	defer rows.Close()

	session := db.Session{}
	hasNext := rows.Next()

	if hasNext == true {
		rows.Scan(&session.ID, &session.StartTime, &session.EndTime, &session.Title, &session.Speaker, &session.Room)
	} else {
		err = rows.Err()
		if err != nil {
			return session, err
		}
		return session, ErrNoRowsFound
	}

	hasNext = rows.Next()
	if hasNext == true {
		return session, ErrTooManyRows
	}

	return session, nil
}

// ReadAllSessions reads all sessions from the db
func (s SessionMySQL) ReadAllSessions() ([]db.Session, error) {
	if s.db == nil {
		return nil, ErrDBNotSet
	}

	q := "SELECT * FROM session;"

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sessions := []db.Session{}
	for rows.Next() {
		session := db.Session{}
		rows.Scan(&session.ID, &session.StartTime, &session.EndTime, &session.Title, &session.Speaker, &session.Room)
		sessions = append(sessions, session)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return sessions, nil
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
