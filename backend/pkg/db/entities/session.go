package entities

type Session struct {
	StartTime int
	EndTime   int
	Title     string
	Speaker   *Speaker
	Room      *Room
}

type SessionReader interface {
	ReadASession(startTime int, roomID int) (Session, error)
	ReadAllSessions() ([]Session, error)
}

type SessionWriter interface {
	WriteASession(s Session) error
}

type SessionDeleter interface {
	DeleteASession(startTime int, roomID int) error
}
