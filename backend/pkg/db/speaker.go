package db

// Speaker holds all data about a speaker
type Speaker struct {
	ID        int64
	Email     *string `json:"email" example:"firstname.lastname@gmail.com"`
	FirstName *string `json:"firstName" example:"Bob"`
	LastName  *string `json:"lastName" example:"Smith"`
}

// NewSpeaker makes a new Speaker with default values
func NewSpeaker() Speaker {
	return Speaker{
		ID:        0,
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
	ReadASpeaker(speakerID int64) (Speaker, error)
	ReadAllSpeakers() ([]Speaker, error)
}

// SpeakerWriter implements all write related methods
type SpeakerWriter interface {
	WriteASpeaker(email string, firstName string, lastName string) error
}

// SpeakerUpdater implements all update related methods
type SpeakerUpdater interface {
	UpdateASpeaker(id int64, email string, firstName string, lastName string) error
}

// SpeakerDeleter implements all delete related methods
type SpeakerDeleter interface {
	DeleteASpeaker(id int64) error
}
