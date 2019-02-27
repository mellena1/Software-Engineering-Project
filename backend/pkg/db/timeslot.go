package db

import "time"

// Timeslot holds all data about a timeslot
type Timeslot struct {
	ID        *int       `json:"id" example:"1"`
	StartTime *time.Time `json:"startTime" example:"2019-02-18 21:00:00"`
	EndTime   *time.Time `json:"endTime" example:"2019-10-01 23:00:00"`
}

// NewTimeslot makes a new Timeslot with default values
func NewTimeslot() Timeslot {
	return Timeslot{
		ID:        IntPtr(0),
		StartTime: &time.Time{},
		EndTime:   &time.Time{},
	}
}

// TimeslotReaderWriterUpdaterDeleter implements everything that a facade for a Room would need
type TimeslotReaderWriterUpdaterDeleter interface {
	TimeslotReader
	TimeslotWriter
	TimeslotUpdater
	TimeslotDeleter
}

// TimeslotReader implements all read related methods
type TimeslotReader interface {
	ReadATimeslot(id int) (Timeslot, error)
	ReadAllTimeslots() ([]Timeslot, error)
}

// TimeslotWriter implements all write related methods
type TimeslotWriter interface {
	WriteATimeslot(t Timeslot) error
}

// TimeslotUpdater implements all update related methods
type TimeslotUpdater interface {
	UpdateATimeslot(id int, newTimeslot Timeslot) error
}

// TimeslotDeleter implements all delete related methods
type TimeslotDeleter interface {
	DeleteATimeslot(id int) error
}
