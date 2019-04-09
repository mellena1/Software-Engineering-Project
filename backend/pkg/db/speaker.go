package db

type Speaker struct {
	ID        *int64  `json:"id" example:"1"`
	Email     *string `json:"email" example:"firstname.lastname@gmail.com"`
	FirstName *string `json:"firstName" example:"Bob"`
	LastName  *string `json:"lastName" example:"Smith"`
}

func NewSpeaker() Speaker {
	return Speaker{
		ID:        Int64Ptr(0),
		Email:     StringPtr(""),
		FirstName: StringPtr(""),
		LastName:  StringPtr(""),
	}
}

type SpeakerReaderWriterUpdaterDeleter interface {
	SpeakerReader
	SpeakerWriter
	SpeakerUpdater
	SpeakerDeleter
}

type SpeakerReader interface {
	ReadASpeaker(speakerID int64) (Speaker, error)
	ReadAllSpeakers() ([]Speaker, error)
}

type SpeakerWriter interface {
	WriteASpeaker(email *string, firstName *string, lastName *string) (int64, error)
}

type SpeakerUpdater interface {
	UpdateASpeaker(id int64, email *string, firstName *string, lastName *string) error
}

type SpeakerDeleter interface {
	DeleteASpeaker(id int64) error
}
