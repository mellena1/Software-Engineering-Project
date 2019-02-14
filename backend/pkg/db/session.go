package db

import "time"

type Session struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
	Title     string
	Speaker   *Speaker
	Room      *Room
}

type SessionReaderWriterUpdaterDeleter interface {
	SessionReader
	SessionWriter
	SessionUpdater
	SessionDeleter
}

type SessionReader interface {
	ReadASession(sessionID int) (Session, error)
	ReadAllSessions() ([]Session, error)
}

type SessionWriter interface {
	WriteASession(s Session) error
}

type SessionUpdater interface {
	UpdateASession(sessionID int, newSession Session) error
}

type SessionDeleter interface {
	DeleteASession(sessionID int) error
}
