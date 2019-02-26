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
func (s SpeakerMySQL) ReadASpeaker(emailID string) (db.Speaker, error) {
	if s.db == nil {
		return db.Speaker{}, ErrDBNotSet
	}

	query := `SELECT * FROM speaker where email = ?;`

	rows, err := s.db.Query(query, emailID)
	if err != nil {
		return db.Speaker{}, err
	}

	defer rows.Close()

	speaker := db.NewSpeaker()
	for rows.Next() {
		rows.Scan(speaker.Email, speaker.FirstName, speaker.LastName)
	}

	return speaker, nil
}

// ReadAllSpeakers reads all speakers from the db
func (s SpeakerMySQL) ReadAllSpeakers() ([]db.Speaker, error) {
	if s.db == nil {
		return nil, ErrDBNotSet
	}

	query := "SELECT * FROM speaker;"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	speakers := []db.Speaker{}
	for rows.Next() {
		newSpeaker := db.NewSpeaker()
		rows.Scan(newSpeaker.Email, newSpeaker.FirstName, newSpeaker.LastName)
		speakers = append(speakers, newSpeaker)
	}
	return speakers, nil
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
