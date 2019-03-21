package db

// Count holds all data about a count
type Count struct {
	Time      *string
	SessionID *int64
	UserID    *int64
	Count     *int64
}

// NewCount makes a new Count with default values
func NewCount() Count {
	return Count{
		Time:      StringPtr(""),
		SessionID: Int64Ptr(0),
		UserID:    Int64Ptr(0),
		Count:     Int64Ptr(0),
	}
}

// CountReaderWriterUpdaterDeleter implements everything that a facade for a Count would need
type CountReaderWriterUpdaterDeleter interface {
	CountReader
	CountWriter
	CountUpdater
	CountDeleter
}

// CountReader implements all read related methods
type CountReader interface {
	ReadACount(sessionID int64) ([]Count, error)
	ReadAllCounts() ([]Count, error)
}

// CountWriter implements all write related methods
type CountWriter interface {
	WriteACount(time *string, sessionID *int64, userID *int64, count *int64) (int64, error)
}

// CountUpdater implements all update related methods
type CountUpdater interface {
	UpdateACount(time *string, sessionID *int64, userID *int64, count *int64) error
}

// CountDeleter implements all delete related methods
type CountDeleter interface {
	DeleteACount(sessionID int64) error
}
