package db

import "time"

const (
	TimeFormat = time.RFC3339
)

type Timeslot struct {
	ID        *int64    `json:"id" example:"1"`
	StartTime time.Time `json:"startTime" example:"2019-02-18T21:00:00Z"`
	EndTime   time.Time `json:"endTime" example:"2019-10-01T23:00:00Z"`
}

func NewTimeslot() Timeslot {
	return Timeslot{
		ID:        Int64Ptr(0),
		StartTime: time.Time{},
		EndTime:   time.Time{},
	}
}

type TimeslotReaderWriterUpdaterDeleter interface {
	TimeslotReader
	TimeslotWriter
	TimeslotUpdater
	TimeslotDeleter
}

type TimeslotReader interface {
	ReadATimeslot(id int64) (Timeslot, error)
	ReadAllTimeslots() ([]Timeslot, error)
}

type TimeslotWriter interface {
	WriteATimeslot(startTime, endTime time.Time) (int64, error)
}

type TimeslotUpdater interface {
	UpdateATimeslot(id int64, startTime, endTime time.Time) error
}

type TimeslotDeleter interface {
	DeleteATimeslot(id int64) error
}
