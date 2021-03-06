package db

import "time"

const (
	// TimeFormat to use to parse time strings with
	TimeFormat = time.RFC3339
)

// Timeslot holds all data about a timeslot
type Timeslot struct {
	ID        *int64    `json:"id" example:"1"`
	StartTime time.Time `json:"startTime" example:"2019-02-18T21:00:00Z"`
	EndTime   time.Time `json:"endTime" example:"2019-10-01T23:00:00Z"`
}

// NewTimeslot makes a new Timeslot with default values
func NewTimeslot() Timeslot {
	return Timeslot{
		ID:        Int64Ptr(0),
		StartTime: time.Time{},
		EndTime:   time.Time{},
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
	ReadATimeslot(id int64) (Timeslot, error)
	ReadAllTimeslots() ([]Timeslot, error)
}

// TimeslotWriter implements all write related methods
type TimeslotWriter interface {
	WriteATimeslot(startTime, endTime time.Time) (int64, error)
}

// TimeslotUpdater implements all update related methods
type TimeslotUpdater interface {
	UpdateATimeslot(id int64, startTime, endTime time.Time) error
}

// TimeslotDeleter implements all delete related methods
type TimeslotDeleter interface {
	DeleteATimeslot(id int64) error
}
