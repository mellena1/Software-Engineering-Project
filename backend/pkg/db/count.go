package db

// Count holds all data about a count
type Count struct {
	Time      *string
	UserID    *int64
	SessionID *int64
	Count     *int64
}

// NewCount makes a new Count with default values
func NewCount() Count {
	return Count{
		UserID:    Int64Ptr(0),
		SessionID: Int64Ptr(0),
		Time:      StringPtr(""),
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
	WriteACount(sessionID int64, time string, count int64) (int64, error)
}

// CountUpdater implements all update related methods
type CountUpdater interface {
	UpdateACount(sessionID int64, userID, int64, time string, updatedCount int64) error
}

// CountDeleter implements all delete related methods
type CountDeleter interface {
	DeleteACount(sessionID int64) error
}
