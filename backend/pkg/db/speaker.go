package db

type Speaker struct {
	ID   int
	Name string
}

type SpeakerReaderWriterUpdaterDeleter interface {
	SpeakerReader
	SpeakerWriter
	SpeakerUpdater
	SpeakerDeleter
}

type SpeakerReader interface {
	ReadASpeaker(id int) (Speaker, error)
	ReadAllSpeakers() ([]Speaker, error)
}

type SpeakerWriter interface {
	WriteASpeaker(s Speaker) error
}

type SpeakerUpdater interface {
	UpdateASpeaker(id int, newSpeaker Speaker) error
}

type SpeakerDeleter interface {
	DeleteASpeaker(id int) error
}
