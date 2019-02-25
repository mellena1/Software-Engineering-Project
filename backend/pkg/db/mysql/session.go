package mysql

import (
	"database/sql"

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

// ReadASession reads a session from the db given sessionID
func (s SessionMySQL) ReadASession(sessionID int) (db.Session, error) {
	if s.db == nil {
		return db.Session{}, ErrDBNotSet
	}

	q := `SELECT session.sessionID, speaker.*, room.*, timeslot.*, session.sessionName FROM session
		INNER JOIN speaker ON session.speakerID = speaker.speakerID
		INNER JOIN room ON session.roomID = room.roomID
		INNER JOIN timeslot ON session.timeslotID = timeslot.timeslotID
		WHERE session.sessionID = ` + string(sessionID) + ";"

	rows, err := s.db.Query(q)
	if err != nil {
		return db.Session{}, err
	}
	defer rows.Close()

	session := db.Session{}
	hasNext := rows.Next()

	if hasNext == true {
		rows.Scan(session.ID, session.Speaker.Email, session.Speaker.FirstName,
			session.Speaker.LastName, session.Room.ID, session.Room.RoomName,
			session.Room.Capacity, session.Timeslot.ID, session.Timeslot.StartTime,
			session.Timeslot.EndTime, session.Name)
	} else {
		err = rows.Err()
		if err != nil {
			return session, err
		}
		return session, ErrNoRowsFound
	}

	// check if the session ID returned multiple rows, if so error
	hasNext = rows.Next()
	if hasNext == true {
		return session, ErrTooManyRows
	}

	return session, nil
}

// ReadAllSessions reads all sessions from the db
func (s SessionMySQL) ReadAllSessions() ([]db.Session, error) {
	if s.db == nil {
		return nil, ErrDBNotSet
	}

	q := `SELECT session.sessionID, speaker.*, room.*, timeslot.*, session.sessionName FROM session
		INNER JOIN speaker ON session.speakerID = speaker.speakerID
		INNER JOIN room ON session.roomID = room.roomID
		INNER JOIN timeslot ON session.timeslotID = timeslot.timeslotID;`

	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sessions := []db.Session{}
	for rows.Next() {
		session := db.Session{}
		rows.Scan(session.ID, session.Speaker.Email, session.Speaker.FirstName,
			session.Speaker.LastName, session.Room.ID, session.Room.RoomName,
			session.Room.Capacity, session.Timeslot.ID, session.Timeslot.StartTime,
			session.Timeslot.EndTime, session.Name)
		sessions = append(sessions, session)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

// WriteASession writes a session to the db
func (s SessionMySQL) WriteASession(session db.Session) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	stmt, err := s.db.Prepare("INSERT INTO session VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(session.StartTime.Format("2019-02-18T21:00:00"), session.EndTime.Format("2019-02-18T21:00:00"),
		session.Title, session.Speaker.Email, session.Room.RoomName)
	if err != nil {
		return err
	}

	return nil
}

// UpdateASession updates a session in the db given a sessionID and the updated session
func (SessionMySQL) UpdateASession(sessionID int, newSession db.Session) error {
	return nil
}

// DeleteASession deletes a room given a sessionID
func (SessionMySQL) DeleteASession(sessionID int) error {
	return nil
}
