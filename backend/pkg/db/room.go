package db

// Room holds all data about a room
type Room struct {
	ID       int64   `json:"id" example:"1"`
	Name     *string `json:"name" example:"My Room Name"`
	Capacity *int    `json:"capacity" example:"50"`
}

// NewRoom makes a new Room with default values
func NewRoom() Room {
	return Room{
		ID:       0,
		Name:     StringPtr(""),
		Capacity: IntPtr(0),
	}
}

// RoomReaderWriterUpdaterDeleter implements everything that a facade for a Room would need
type RoomReaderWriterUpdaterDeleter interface {
	RoomReader
	RoomWriter
	RoomUpdater
	RoomDeleter
}

// RoomReader implements all read related methods
type RoomReader interface {
	ReadARoom(roomID int64) (Room, error)
	ReadAllRooms() ([]Room, error)
}

// RoomWriter implements all write related methods
type RoomWriter interface {
	WriteARoom(name string, capacity int) (int64, error)
}

// RoomUpdater implements all update related methods
type RoomUpdater interface {
	UpdateARoom(id int64, name string, capacity int) error
}

// RoomDeleter implements all delete related methods
type RoomDeleter interface {
	DeleteARoom(id int64) error
}
