package db

type Session struct {
	ID       *int64    `json:"id" example:"1"`
	Timeslot *Timeslot `json:"timeslot"`
	Name     *string   `json:"name" example:"Session Name"`
	Speaker  *Speaker  `json:"speaker"`
	Room     *Room     `json:"room"`
}

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

type SessionReaderWriterUpdaterDeleter interface {
	SessionReader
	SessionWriter
	SessionUpdater
	SessionDeleter
}

type SessionReader interface {
	ReadASession(sessionID int64) (Session, error)
	ReadAllSessions() ([]Session, error)
}

type SessionWriter interface {
	WriteASession(speakerID *int64, roomID *int64, timeslotID *int64, name *string) (int64, error)
}

type SessionUpdater interface {
	UpdateASession(sessionID int64, speakerID *int64, roomID *int64, timeslotID *int64, name *string) error
}

type SessionDeleter interface {
	DeleteASession(sessionID int64) error
}
