package sql

import (
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities"
)

type SessionSQL struct{}

func (SessionSQL) ReadASession(startTime int, roomID int) (entities.Session, error) {
	return entities.Session{}, nil
}

func (SessionSQL) ReadAllSessions() ([]entities.Session, error) {
	return []entities.Session{}, nil
}

func (SessionSQL) WriteASession(s entities.Session) error {
	return nil
}

func (SessionSQL) DeleteASession(startTime int, roomID int) error {
	return nil
}
