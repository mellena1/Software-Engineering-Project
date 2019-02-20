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
	query := `SELECT * FROM speaker where email = ?;`
	speaker := db.Speaker{}
	var (
		email string
		firstName string
		lastName string
	)

	rows, err := s.db.Query(query, emailID)
	if err != nil {
		return speaker, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&email, &firstName, &lastName)
		if err != nil {
			return speaker, err
		}
		speaker = db.Speaker{
			Email: email,
			FirstName: firstName,
			LastName: lastName,
		}	
		
	}
	err = rows.Err()
	if err != nil {
		return speaker, err
	}

	return speaker, nil
}


// ReadAllSpeakers reads all speakers from the db
func (s SpeakerMySQL) ReadAllSpeakers() ([]db.Speaker, error) {
	query := `SELECT * FROM speaker;`
	var speakers = []db.Speaker{}

	var (
		email string
		firstName string
		lastName string
	)

	rows, err := s.db.Query(query)
	if err != nil {
		return []db.Speaker{}, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&email, &firstName, &lastName)
		if err != nil {
			return []db.Speaker{}, err
		}
		speaker := db.Speaker{
			Email: email,
			FirstName: firstName,
			LastName: lastName,
		}	
		
		speakers = append(speakers, speaker)
	}
	err = rows.Err()
	if err != nil {
		return []db.Speaker{}, err
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
