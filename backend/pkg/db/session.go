package db

// Session holds all data about a session
type Session struct {
	ID       *int      `json:"id" example:"1"`
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
		ID:       IntPtr(0),
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
	ReadASession(sessionID int) (Session, error)
	ReadAllSessions() ([]Session, error)
}

// SessionWriter implements all write related methods
type SessionWriter interface {
	WriteASession(session Session) error
}

// SessionUpdater implements all update related methods
type SessionUpdater interface {
	UpdateASession(sessionID int, newSession Session) error
}

// SessionDeleter implements all delete related methods
type SessionDeleter interface {
	DeleteASession(sessionID int) error
}
