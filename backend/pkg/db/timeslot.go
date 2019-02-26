package db

import "time"

const (
	// TimeFormat to use to parse time strings with
	TimeFormat = "2006-01-02 15:04:05"
)

// Timeslot holds all data about a timeslot
type Timeslot struct {
	ID        *int64
	StartTime *time.Time
	EndTime   *time.Time
}

// NewTimeslot makes a new Timeslot with default values
func NewTimeslot() Timeslot {
	return Timeslot{
		ID:        Int64Ptr(0),
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
	ReadATimeslot(id int64) (Timeslot, error)
	ReadAllTimeslots() ([]Timeslot, error)
}

// TimeslotWriter implements all write related methods
type TimeslotWriter interface {
	WriteATimeslot(timeslot Timeslot) (int64, error)
}

// TimeslotUpdater implements all update related methods
type TimeslotUpdater interface {
	UpdateATimeslot(timeslot Timeslot) error
}

// TimeslotDeleter implements all delete related methods
type TimeslotDeleter interface {
	DeleteATimeslot(id int64) error
}
