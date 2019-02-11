package sql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

type SpeakerSQL struct {
	db *sql.DB
}

func NewSpeakerSQL(db *sql.DB) SpeakerSQL {
	return SpeakerSQL{db}
}

func (SpeakerSQL) ReadASpeaker(id int) (db.Speaker, error) {
	return db.Speaker{}, nil
}

func (SpeakerSQL) ReadAllSpeakers() ([]db.Speaker, error) {
	return []db.Speaker{}, nil
}

func (SpeakerSQL) WriteASpeaker(s db.Speaker) error {
	return nil
}

func (SpeakerSQL) UpdateASpeaker(id int, newSpeaker db.Speaker) error {
	return nil
}

func (SpeakerSQL) DeleteASpeaker(id int) error {
	return nil
}
