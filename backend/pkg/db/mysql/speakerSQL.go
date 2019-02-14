package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// SpeakerSQL implements SpeakerReaderWriterUpdaterDeleter
type SpeakerSQL struct {
	db *sql.DB
}

// NewSpeakerSQL makes a new SpeakerSQL object given a db
func NewSpeakerSQL(db *sql.DB) SpeakerSQL {
	return SpeakerSQL{db}
}

// ReadASpeaker reads a speaker from the db given email
func (SpeakerSQL) ReadASpeaker(email string) (db.Speaker, error) {
	return db.Speaker{}, nil
}

// ReadAllSpeakers reads all speakers from the db
func (SpeakerSQL) ReadAllSpeakers() ([]db.Speaker, error) {
	return []db.Speaker{}, nil
}

// WriteASpeaker writes a speaker to the db
func (SpeakerSQL) WriteASpeaker(s db.Speaker) error {
	return nil
}

// UpdateASpeaker updates a speaker in the db given an email and the updated speaker
func (SpeakerSQL) UpdateASpeaker(email string, newSpeaker db.Speaker) error {
	return nil
}

// DeleteASpeaker deletes a speaker given an email
func (SpeakerSQL) DeleteASpeaker(email string) error {
	return nil
}
