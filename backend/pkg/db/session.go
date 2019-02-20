package db

import "time"

// Session holds all data about a session
type Session struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
	Title     string
	Speaker   *Speaker
	Room      *Room
}

// SessionReaderWriterUpdaterDeleter implements everything that a facade for a Session would need
type SessionReaderWriterUpdaterDeleter interface {
	SessionReader
	SessionWriter
	SessionUpdater
	SessionDeleter
}

// SessionReader implements all read related methods
type SessionReader interface {
	ReadASession(sessionID int) (Session, error)
	ReadAllSessions() ([]Session, error)
}

// SessionWriter implements all write related methods
type SessionWriter interface {
	WriteASession(session Session) error
}

// SessionUpdater implements all update related methods
type SessionUpdater interface {
	UpdateASession(sessionID int, newSession Session) error
}

// SessionDeleter implements all delete related methods
type SessionDeleter interface {
	DeleteASession(sessionID int) error
}
