package db

// Room holds all data about a room
type Room struct {
	ID       *int
	Name     *string
	Capacity *int
}

// NewRoom makes a new Room with default values
func NewRoom() Room {
	return Room{
		ID:       IntPtr(0),
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
	ReadARoom(roomName string) (Room, error)
	ReadAllRooms() ([]Room, error)
}

// RoomWriter implements all write related methods
type RoomWriter interface {
	WriteARoom(name *string, capacity *int) (int64, error)
}

// RoomUpdater implements all update related methods
type RoomUpdater interface {
	UpdateARoom(name *string, newRoom Room) error
}

// RoomDeleter implements all delete related methods
type RoomDeleter interface {
	DeleteARoom(id *int) error
}
