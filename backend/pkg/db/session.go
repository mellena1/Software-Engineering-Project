package db

// Session holds all data about a session
type Session struct {
	ID       *int64    `json:"id" example:"1"`
	Timeslot *Timeslot `json:"timeslot"`
	Name     *string   `json:"name" example:"Session Name"`
	Speaker  *Speaker  `json:"speaker"`
	Room     *Room     `json:"room"`
}

// NewSession makes a new Session with default values
func NewSession() Session {
	room := NewRoom()
	speaker := NewSpeaker()
	timeslot := NewTimeslot()
	return Session{
		ID:       Int64Ptr(0),
		Timeslot: &timeslot,
		Name:     StringPtr(""),
		Speaker:  &speaker,
		Room:     &room,
	}
}

// SessionReaderWriterUpdaterDeleter implements everything that a facade for a Session would need
type SessionReaderWriterUpdaterDeleter interface {
	SessionReader
	SessionWriter
	SessionUpdater
	SessionDeleter
}

// SessionReader implements all read related methods
type SessionReader interface {
	ReadASession(sessionID int64) (Session, error)
	ReadAllSessions() ([]Session, error)
}

// SessionWriter implements all write related methods
type SessionWriter interface {
	WriteASession(speakerID *int64, roomID *int64, timeslotID *int64, name *string) (int64, error)
}

// SessionUpdater implements all update related methods
type SessionUpdater interface {
	UpdateASession(sessionID int64, speakerID *int64, roomID *int64, timeslotID *int64, name *string) error
}

// SessionDeleter implements all delete related methods
type SessionDeleter interface {
	DeleteASession(sessionID int64) error
}
