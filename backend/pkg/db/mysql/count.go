package mysql

import (
	"database/sql"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// CountMySQL implements CountReaderWriterUpdaterDeleter
type CountMySQL struct {
	db *sql.DB
}

// NewCountMySQL makes a new CountMySQL object given a db
func NewCountMySQL(db *sql.DB) CountMySQL {
	return CountMySQL{db}
}

// scanACount takes in a count pointer and scans a row into it
func scanACount(count *db.Count, row rowScanner) error {
	return row.Scan(&count.Time, &count.SessionID, &count.UserName, &count.Count)
}

// scanACountBySpeaker takes in a session pointer and scans a row into it
// must be in order: speakerFirstName, speakerLastName
//					 sessionName,
//					time, sessionId, userName, count
func scanACountBySpeaker(session *db.CountBySpeakerResponse, row rowScanner) error {
	speakerFirstName, speakerLastName := sql.NullString{}, sql.NullString{}
	sessionName := sql.NullString{}
	time, sessionID, userName, count := sql.NullString{}, sql.NullInt64{}, sql.NullString{}, sql.NullInt64{}

	err := row.Scan(&speakerFirstName, &speakerLastName,
		&sessionName,
		&time, &sessionID, &userName, &count)

	session.SpeakerFirstName, session.SpeakerLastName = NullStringToString(speakerFirstName), NullStringToString(speakerLastName)
	session.SessionName = NullStringToString(sessionName)
	session.Time, session.SessionID, session.UserName, session.Count = NullStringToString(time), NullIntToInt(sessionID), NullStringToString(userName), NullIntToInt(count)

	return err
}

// ReadCountsOfSession reads a count from the db given a sessionID
func (myCountSQL CountMySQL) ReadCountsOfSession(sessionID int64) ([]db.Count, error) {
	if myCountSQL.db == nil {
		return nil, ErrDBNotSet
	}

	statement, err := myCountSQL.db.Prepare("SELECT time, sessionID, userName, count FROM count where sessionID = ?;")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := []db.Count{}
	for rows.Next() {
		newCount := db.NewCount()
		scanACount(&newCount, rows)
		counts = append(counts, newCount)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return counts, nil
}

// ReadAllCounts reads all counts from the db
func (myCountSQL CountMySQL) ReadAllCounts() ([]db.Count, error) {
	if myCountSQL.db == nil {
		return nil, ErrDBNotSet
	}

	query := "SELECT time, sessionID, userName, count FROM count;"

	rows, err := myCountSQL.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := []db.Count{}
	for rows.Next() {
		newCount := db.NewCount()
		scanACount(&newCount, rows)
		counts = append(counts, newCount)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return counts, nil
}

// ReadAllCountsBySpeaker reads all counts from the db and sorts them by speaker and session
func (myCountSQL CountMySQL) ReadAllCountsBySpeaker() ([]db.CountBySpeakerResponse, error) {
	if myCountSQL.db == nil {
		return nil, ErrDBNotSet
	}

	query := `SELECT speaker.firstName, speaker.lastName, session.sessionName,
				count.time, count.sessionID, count.userName, count.count 
			FROM session
			LEFT JOIN speaker ON session.speakerID = speaker.speakerID
			LEFT JOIN count ON session.sessionID = count.sessionID;`

	rows, err := myCountSQL.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	countsBySpeaker := []db.CountBySpeakerResponse{}
	for rows.Next() {
		newCountBySpeaker := db.NewCountBySpeaker()
		scanACountBySpeaker(&newCountBySpeaker, rows)
		countsBySpeaker = append(countsBySpeaker, newCountBySpeaker)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return countsBySpeaker, nil
}

// WriteACount writes a count to the db
func (myCountSQL CountMySQL) WriteACount(time *string, sessionID *int64, userName *string, count *int64) (int64, error) {
	if myCountSQL.db == nil {
		return 0, ErrDBNotSet
	}
	statement, err := myCountSQL.db.Prepare("INSERT INTO `count`(`time`, `sessionID`, `userName`, `count`) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(StringToNullString(time), IntToNullInt(sessionID), StringToNullString(userName), IntToNullInt(count))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateACount updates a count in the db given the session Id and the updated time of session
func (myCountSQL CountMySQL) UpdateACount(time *string, sessionID *int64, userName *string, count *int64) error {
	if myCountSQL.db == nil {
		return ErrDBNotSet
	}

	statement, err := myCountSQL.db.Prepare("UPDATE count SET time = ?, sessionID = ?, userName = ?, count = ? WHERE time = ? and sessionID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(StringToNullString(time), IntToNullInt(sessionID), StringToNullString(userName), IntToNullInt(count), StringToNullString(time), IntToNullInt(sessionID))
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

// DeleteACount deletes a count given an id
func (myCountSQL CountMySQL) DeleteACount(sessionID int64) error {
	if myCountSQL.db == nil {
		return ErrDBNotSet
	}

	statement, err := myCountSQL.db.Prepare("DELETE FROM count WHERE sessionID = ?;")
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(sessionID)
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
