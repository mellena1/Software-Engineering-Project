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
func (s SpeakerMySQL) ReadASpeaker(speakerID int) (db.Speaker, error) {
	if s.db == nil {
		return db.Speaker{}, ErrDBNotSet
	}

	query := `SELECT * FROM speaker where speakerID = ?;`

	rows, err := s.db.Query(query, speakerID)
	if err != nil {
		return db.Speaker{}, err
	}

	defer rows.Close()

	speaker := db.NewSpeaker()
	for rows.Next() {
		rows.Scan(speaker.ID, speaker.Email, speaker.FirstName, speaker.LastName)
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
		rows.Scan(newSpeaker.ID, newSpeaker.Email, newSpeaker.FirstName, newSpeaker.LastName)
		speakers = append(speakers, newSpeaker)
	}
	return speakers, nil
}

// WriteASpeaker writes a speaker to the db
func (s SpeakerMySQL) WriteASpeaker(speaker db.Speaker) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	stmt, err := s.db.Prepare("INSERT INTO speaker (`email`, `firstName`, `lastName`) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(speaker.Email, speaker.FirstName, speaker.LastName)

	if err != nil {
		return err
	}

	return nil
}

// UpdateASpeaker updates a speaker in the db given an email and the updated speaker
func (s SpeakerMySQL) UpdateASpeaker(id int, newSpeaker db.Speaker) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	oldSpeaker, err := s.ReadASpeaker(id)

	if err != nil {
		return err
	}

	speakerID := oldSpeaker.ID
	_, err = s.db.Exec(`UPDATE speaker SET email = ?, firstName = ?, lastName = ? WHERE id = ?`, newSpeaker.Email, newSpeaker.FirstName, newSpeaker.LastName, speakerID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteASpeaker deletes a speaker given an email
func (s SpeakerMySQL) DeleteASpeaker(email string) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	_, err := s.db.Exec(`DELETE FROM speaker WHERE email = ?`, email)

	if err != nil {
		return err
	}

	return nil
}
