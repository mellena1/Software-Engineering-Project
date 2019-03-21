package mysql

import (
	"database/sql"
	"time"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"
)

// SessionMySQL implements SessionReaderWriterUpdaterDeleter
type SessionMySQL struct {
	db *sql.DB
}

// NewSessionMySQL makes a new SessionMySQL object given a db
func NewSessionMySQL(db *sql.DB) SessionMySQL {
	return SessionMySQL{db}
}

// scanASession takes in a session pointer and scans a row into it
// must be in order: session.ID, speakerID, speakerEmail, speakerFirstName,
//					 speakerLastName, roomID, roomName,
//					 roomCapacity, timeslotID, timeslotStartTime,
//					 timeslotEndTime, sessionName
func scanASession(session *db.Session, row rowScanner) error {
	speakerID, roomID, timeslotID := sql.NullInt64{}, sql.NullInt64{}, sql.NullInt64{}
	speakerEmail, speakerFirstName, speakerLastName := sql.NullString{}, sql.NullString{}, sql.NullString{}
	roomName, roomCapacity := sql.NullString{}, sql.NullInt64{}
	timeslotStartTime, timeslotEndTime := sql.NullString{}, sql.NullString{}
	sessionName := sql.NullString{}

	err := row.Scan(&session.ID, &speakerID, &speakerEmail, &speakerFirstName,
		&speakerLastName, &roomID, &roomName,
		&roomCapacity, &timeslotID, &timeslotStartTime,
		&timeslotEndTime, &sessionName)

	if speakerID.Valid {
		session.Speaker.ID = NullIntToInt(speakerID)
		session.Speaker.Email, session.Speaker.FirstName, session.Speaker.LastName = NullStringToString(speakerEmail), NullStringToString(speakerFirstName), NullStringToString(speakerLastName)
	} else {
		session.Speaker = nil
	}

	if roomID.Valid {
		session.Room.ID = NullIntToInt(roomID)
		session.Room.Name = NullStringToString(roomName)
		session.Room.Capacity = NullIntToInt(roomCapacity)
	} else {
		session.Room = nil
	}

	if timeslotID.Valid {
		session.Timeslot.ID = NullIntToInt(timeslotID)

		// I think the time vals are in RFC3339 because we set ParseTime = true on the mysql-driver
		sTime := NullStringToString(timeslotStartTime)                   // should never be null
		session.Timeslot.StartTime, _ = time.Parse(time.RFC3339, *sTime) // mysql time will always be in this format

		eTime := NullStringToString(timeslotEndTime)                   // should never be null
		session.Timeslot.EndTime, _ = time.Parse(time.RFC3339, *eTime) // mysql time will always be in this format
	} else {
		session.Timeslot = nil
	}

	session.Name = NullStringToString(sessionName)

	return err
}

// ReadASession reads a session from the db given sessionID
func (s SessionMySQL) ReadASession(sessionID int64) (db.Session, error) {
	if s.db == nil {
		return db.Session{}, ErrDBNotSet
	}

	stmt, err := s.db.Prepare(`
		SELECT session.sessionID, speaker.speakerID, speaker.email, 
			speaker.firstName, speaker.lastName, room.roomID, room.roomName, room.capacity, 
			timeslot.timeslotID, timeslot.startTime, timeslot.endTime, session.sessionName 
		FROM session
		LEFT JOIN speaker ON session.speakerID = speaker.speakerID
		LEFT JOIN room ON session.roomID = room.roomID
		LEFT JOIN timeslot ON session.timeslotID = timeslot.timeslotID
		WHERE session.sessionID = ?;`)
	if err != nil {
		return db.Session{}, err
	}
	defer stmt.Close()

	session := db.NewSession()
	row := stmt.QueryRow(sessionID)
	err = scanASession(&session, row)

	return session, err
}

// ReadAllSessions reads all sessions from the db
func (s SessionMySQL) ReadAllSessions() ([]db.Session, error) {
	if s.db == nil {
		return nil, ErrDBNotSet
	}

	q := `
		SELECT session.sessionID, speaker.speakerID, speaker.email, 
			speaker.firstName, speaker.lastName, room.roomID, room.roomName, room.capacity, 
			timeslot.timeslotID, timeslot.startTime, timeslot.endTime, session.sessionName 
		FROM session
		LEFT JOIN speaker ON session.speakerID = speaker.speakerID
		LEFT JOIN room ON session.roomID = room.roomID
		LEFT JOIN timeslot ON session.timeslotID = timeslot.timeslotID;`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sessions := []db.Session{}
	for rows.Next() {
		newSession := db.NewSession()
		scanASession(&newSession, rows)
		sessions = append(sessions, newSession)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

// WriteASession writes a session to the db
func (s SessionMySQL) WriteASession(speakerID *int64, roomID *int64, timeslotID *int64, name *string) (int64, error) {
	if s.db == nil {
		return 0, ErrDBNotSet
	}

	stmt, err := s.db.Prepare("INSERT INTO session (`speakerID`, `roomID`, `timeslotID`, `sessionName`) VALUES (?, ?, ?, ?);")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(IntToNullInt(speakerID), IntToNullInt(roomID), IntToNullInt(timeslotID), StringToNullString(name))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateASession updates a session in the db given a sessionID and the updated session
func (s SessionMySQL) UpdateASession(sessionID int64, speakerID *int64, roomID *int64, timeslotID *int64, name *string) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	stmt, err := s.db.Prepare("UPDATE session SET speakerID = ?, roomID = ?, timeslotID = ?, sessionName = ? WHERE sessionID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(IntToNullInt(speakerID), IntToNullInt(roomID), IntToNullInt(timeslotID), StringToNullString(name), sessionID)
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

// DeleteASession deletes a room given a sessionID
func (s SessionMySQL) DeleteASession(sessionID int64) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	stmt, err := s.db.Prepare("DELETE FROM session WHERE sessionID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(sessionID)
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
