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
	speaker.Email, speaker.FirstName, speaker.LastName = NullStringToString(email), NullStringToString(firstName), NullStringToString(lastName)
	return err
}

// ReadASpeaker reads a speaker from the db given email
func (mySpeakerSQL SpeakerMySQL) ReadASpeaker(speakerID int64) (db.Speaker, error) {
	if mySpeakerSQL.db == nil {
		return db.Speaker{}, ErrDBNotSet
	}

	statement, err := mySpeakerSQL.db.Prepare("SELECT speakerID, email, firstName, lastName FROM speaker where speakerID = ?;")
	if err != nil {
		return db.Speaker{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(speakerID)

	speaker := db.NewSpeaker()
	err = scanASpeaker(&speaker, row)

	return speaker, err
}

// ReadAllSpeakers reads all speakers from the db
func (mySpeakerSQL SpeakerMySQL) ReadAllSpeakers() ([]db.Speaker, error) {
	if mySpeakerSQL.db == nil {
		return nil, ErrDBNotSet
	}

	query := "SELECT speakerID, email, firstName, lastName FROM speaker;"

	rows, err := mySpeakerSQL.db.Query(query)
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
func (mySpeakerSQL SpeakerMySQL) WriteASpeaker(email *string, firstName *string, lastName *string) (int64, error) {
	if mySpeakerSQL.db == nil {
		return 0, ErrDBNotSet
	}

	statement, err := mySpeakerSQL.db.Prepare("INSERT INTO speaker (`email`, `firstName`, `lastName`) VALUES (?, ?, ?);")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(StringToNullString(email), StringToNullString(firstName), StringToNullString(lastName))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateASpeaker updates a speaker in the db given an email and the updated speaker
func (mySpeakerSQL SpeakerMySQL) UpdateASpeaker(id int64, email *string, firstName *string, lastName *string) error {
	if mySpeakerSQL.db == nil {
		return ErrDBNotSet
	}

	statement, err := mySpeakerSQL.db.Prepare("UPDATE speaker SET email = ?, firstName = ?, lastName = ? WHERE speakerID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(StringToNullString(email), StringToNullString(firstName), StringToNullString(lastName), id)
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
func (mySpeakerSQL SpeakerMySQL) DeleteASpeaker(id int64) error {
	if mySpeakerSQL.db == nil {
		return ErrDBNotSet
	}

	statement, err := mySpeakerSQL.db.Prepare("DELETE FROM speaker WHERE speakerID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(id)
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
