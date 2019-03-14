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

// scanASpeaker takes in a speaker pointer and scans a row into it
func scanASpeaker(speaker *db.Speaker, row rowScanner) error {
	email, firstName, lastName := sql.NullString{}, sql.NullString{}, sql.NullString{}
	err := row.Scan(&speaker.ID, &email, &firstName, &lastName)
	speaker.Email, speaker.FirstName, speaker.LastName = nullStringToString(email), nullStringToString(firstName), nullStringToString(lastName)
	return err
}

// ReadASpeaker reads a speaker from the db given email
func (s SpeakerMySQL) ReadASpeaker(speakerID int64) (db.Speaker, error) {
	speaker := db.NewSpeaker()

	if s.db == nil {
		return speaker, ErrDBNotSet
	}

	stmt, err := s.db.Prepare("SELECT speakerID, email, firstName, lastName FROM speaker where speakerID = ?;")
	if err != nil {
		return speaker, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(speakerID)

	err = scanASpeaker(&speaker, row)

	return speaker, err
}

// ReadAllSpeakers reads all speakers from the db
func (s SpeakerMySQL) ReadAllSpeakers() ([]db.Speaker, error) {
	if s.db == nil {
		return nil, ErrDBNotSet
	}

	q := "SELECT speakerID, email, firstName, lastName FROM speaker;"

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	speakers := []db.Speaker{}
	for rows.Next() {
		newSpeaker := db.NewSpeaker()
		scanASpeaker(&newSpeaker, rows)
		speakers = append(speakers, newSpeaker)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return speakers, nil
}

// WriteASpeaker writes a speaker to the db
func (s SpeakerMySQL) WriteASpeaker(email *string, firstName *string, lastName *string) (int64, error) {
	if s.db == nil {
		return 0, ErrDBNotSet
	}

	stmt, err := s.db.Prepare("INSERT INTO speaker (`email`, `firstName`, `lastName`) VALUES (?, ?, ?);")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(stringToNullString(email), stringToNullString(firstName), stringToNullString(lastName))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateASpeaker updates a speaker in the db given an email and the updated speaker
func (s SpeakerMySQL) UpdateASpeaker(id int64, email *string, firstName *string, lastName *string) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	stmt, err := s.db.Prepare("UPDATE speaker SET email = ?, firstName = ?, lastName = ? WHERE speakerID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(stringToNullString(email), stringToNullString(firstName), stringToNullString(lastName), id)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return err
	} else if rows == 0 {
		return db.ErrNothingChanged
	}

	return nil
}

// DeleteASpeaker deletes a speaker given an email
func (s SpeakerMySQL) DeleteASpeaker(id int64) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	stmt, err := s.db.Prepare("DELETE FROM speaker WHERE speakerID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return err
	} else if rows == 0 {
		return db.ErrNothingChanged
	}

	return nil
}
