package db

// Speaker holds all data about a speaker
type Speaker struct {
	ID        *string
	Email     *string
	FirstName *string
	LastName  *string
}

// NewSpeaker makes a new Speaker with default values
func NewSpeaker() Speaker {
	return Speaker{
		ID:        StringPtr(""),
		Email:     StringPtr(""),
		FirstName: StringPtr(""),
		LastName:  StringPtr(""),
	}
}

// SpeakerReaderWriterUpdaterDeleter implements everything that a facade for a Speaker would need
type SpeakerReaderWriterUpdaterDeleter interface {
	SpeakerReader
	SpeakerWriter
	SpeakerUpdater
	SpeakerDeleter
}

// SpeakerReader implements all read related methods
type SpeakerReader interface {
	ReadASpeaker(email string) (Speaker, error)
	ReadAllSpeakers() ([]Speaker, error)
}

// SpeakerWriter implements all write related methods
type SpeakerWriter interface {
	WriteASpeaker(s Speaker) error
}

// SpeakerUpdater implements all update related methods
type SpeakerUpdater interface {
	UpdateASpeaker(email string, newSpeaker Speaker) error
}

// SpeakerDeleter implements all delete related methods
type SpeakerDeleter interface {
	DeleteASpeaker(email string) error
}
