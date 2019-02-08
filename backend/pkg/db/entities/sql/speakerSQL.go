package sql

import (
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities"
)

type SpeakerSQL struct{}

func (SpeakerSQL) ReadASpeaker(id int) (entities.Speaker, error) {
	return entities.Speaker{}, nil
}

func (SpeakerSQL) ReadAllSpeakers() ([]entities.Speaker, error) {
	return []entities.Speaker{}, nil
}

func (SpeakerSQL) WriteASpeaker(s entities.Speaker) error {
	return nil
}

func (SpeakerSQL) DeleteASpeaker(id int) error {
	return nil
}
