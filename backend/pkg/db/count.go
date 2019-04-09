package db

// Count holds all data about a count
type Count struct {
	Time      *string
	SessionID *int64
	UserName  *string
	Count     *int64
}

// NewCount makes a new Count with default values
func NewCount() Count {
	return Count{
		Time:      StringPtr(""),
		SessionID: Int64Ptr(0),
		UserName:  StringPtr(""),
		Count:     Int64Ptr(0),
	}
}

// CountBySpeakerResponse response struct for phase 3
type CountBySpeakerResponse struct {
	SessionName      *string `json:"time" example:"Microservices"`
	SpeakerFirstName *string `json:"speakerFirstName" example:"Kenny"`
	SpeakerLastName  *string `json:"speakerLastName" example:"Robinson"`
	Time             *string `json:"countTime" example:"beginning/middle/end"`
	SessionID        *int64  `json:"SessionID" example:"1"`
	UserName         *string `json:"userName" example:"Kenny Robinson"`
	Count            *int64  `json:"count" example:"30"`
}

// NewCountBySpeaker makes a new Count with default values
func NewCountBySpeaker() CountBySpeakerResponse {
	return CountBySpeakerResponse{
		SpeakerFirstName: StringPtr(""),
		SpeakerLastName:  StringPtr(""),
		SessionName:      StringPtr(""),
		Time:             StringPtr(""),
		SessionID:        Int64Ptr(0),
		UserName:         StringPtr(""),
		Count:            Int64Ptr(0),
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
	ReadCountsOfSession(sessionID int64) ([]Count, error)
	ReadAllCountsBySpeaker() ([]CountBySpeakerResponse, error)
	ReadAllCounts() ([]Count, error)
}

// CountWriter implements all write related methods
type CountWriter interface {
	WriteACount(time *string, sessionID *int64, userName *string, count *int64) (int64, error)
}

// CountUpdater implements all update related methods
type CountUpdater interface {
	UpdateACount(time *string, sessionID *int64, userName *string, count *int64) error
}

// CountDeleter implements all delete related methods
type CountDeleter interface {
	DeleteACount(sessionID int64) error
}
