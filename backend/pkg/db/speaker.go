package db

type Speaker struct {
	Email     string
	FirstName string
	LastName  string
}

type SpeakerReaderWriterUpdaterDeleter interface {
	SpeakerReader
	SpeakerWriter
	SpeakerUpdater
	SpeakerDeleter
}

type SpeakerReader interface {
	ReadASpeaker(email string) (Speaker, error)
	ReadAllSpeakers() ([]Speaker, error)
}

type SpeakerWriter interface {
	WriteASpeaker(s Speaker) error
}

type SpeakerUpdater interface {
	UpdateASpeaker(email string, newSpeaker Speaker) error
}

type SpeakerDeleter interface {
	DeleteASpeaker(email string) error
}
