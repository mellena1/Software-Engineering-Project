package entities

type Speaker struct {
	ID   int
	Name string
}

type SpeakerReader interface {
	ReadASpeaker(id int) (Speaker, error)
	ReadAllSpeakers() ([]Speaker, error)
}

type SpeakerWriter interface {
	WriteASpeaker(s Speaker) error
}

type SpeakerDeleter interface {
	DeleteASpeaker(id int) error
}
