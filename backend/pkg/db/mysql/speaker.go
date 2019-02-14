package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// SpeakerMySQL implements SpeakerReaderWriterUpdaterDeleter
type SpeakerMySQL struct {
	db *sql.DB
}

// NewSpeakerMySQL makes a new SpeakerMySQL object given a db
func NewSpeakerMySQL(db *sql.DB) SpeakerMySQL {
	return SpeakerMySQL{db}
}

// ReadASpeaker reads a speaker from the db given email
func (SpeakerMySQL) ReadASpeaker(email string) (db.Speaker, error) {
	return db.Speaker{}, nil
}

// ReadAllSpeakers reads all speakers from the db
func (SpeakerMySQL) ReadAllSpeakers() ([]db.Speaker, error) {
	return []db.Speaker{}, nil
}

// WriteASpeaker writes a speaker to the db
func (SpeakerMySQL) WriteASpeaker(s db.Speaker) error {
	return nil
}

// UpdateASpeaker updates a speaker in the db given an email and the updated speaker
func (SpeakerMySQL) UpdateASpeaker(email string, newSpeaker db.Speaker) error {
	return nil
}

// DeleteASpeaker deletes a speaker given an email
func (SpeakerMySQL) DeleteASpeaker(email string) error {
	return nil
}
