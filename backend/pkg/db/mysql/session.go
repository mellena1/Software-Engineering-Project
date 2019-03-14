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
func (s SessionMySQL) ReadASession(sessionID int64) (db.Session, error) {
	session := db.NewSession()

	if s.db == nil {
		return session, ErrDBNotSet
	}

	stmt, err := s.db.Prepare(`
		SELECT session.sessionID, speaker.speakerID, speaker.email, 
			speaker.firstName, speaker.lastName, room.roomID, room.roomName, room.capacity, 
			timeslot.timeslotID, timeslot.startTime, timeslot.endTime, session.sessionName 
		FROM session
		INNER JOIN speaker ON session.speakerID = speaker.speakerID
		INNER JOIN room ON session.roomID = room.roomID
		INNER JOIN timeslot ON session.timeslotID = timeslot.timeslotID
		WHERE session.sessionID = ?;`)
	if err != nil {
		return session, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(sessionID)

	err = row.Scan(&session.ID, &session.Speaker.ID, session.Speaker.Email, session.Speaker.FirstName,
		session.Speaker.LastName, &session.Room.ID, session.Room.Name,
		session.Room.Capacity, &session.Timeslot.ID, &session.Timeslot.StartTime,
		&session.Timeslot.EndTime, session.Name)

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
		newSession := db.NewSession()
		rows.Scan(&newSession.ID, &newSession.Speaker.ID, newSession.Speaker.Email, newSession.Speaker.FirstName,
			newSession.Speaker.LastName, &newSession.Room.ID, newSession.Room.Name,
			newSession.Room.Capacity, &newSession.Timeslot.ID, &newSession.Timeslot.StartTime,
			&newSession.Timeslot.EndTime, newSession.Name)
		sessions = append(sessions, newSession)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

// WriteASession writes a session to the db
func (s SessionMySQL) WriteASession(speakerID *int, roomID *int, timeslotID *int64, name *string) (int64, error) {
	if s.db == nil {
		return 0, ErrDBNotSet
	}

	stmt, err := s.db.Prepare("INSERT INTO session (`speakerID`, `roomID`, `timeslotID`, `sessionName`) VALUES (?, ?, ?, ?);")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(speakerID, roomID, timeslotID, name)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UpdateASession updates a session in the db given a sessionID and the updated session
func (s SessionMySQL) UpdateASession(sessionID int64, speakerID *int, roomID *int, timeslotID *int64, name *string) error {
	if s.db == nil {
		return ErrDBNotSet
	}

	stmt, err := s.db.Prepare("UPDATE session SET speakerID = ?, roomID = ?, timeslotID = ?, sessionName = ? WHERE sessionID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(speakerID, roomID, timeslotID, name, sessionID)
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
