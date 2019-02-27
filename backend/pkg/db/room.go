package db

// Room holds all data about a room
type Room struct {
	ID       *int    `json:"id" example:"1"`
	Name     *string `json:"name" example:"My Room Name"`
	Capacity *int    `json:"capacity" example:"50"`
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
	WriteARoom(r Room) error
}

// RoomUpdater implements all update related methods
type RoomUpdater interface {
	UpdateARoom(roomName string, newRoom Room) error
}

// RoomDeleter implements all delete related methods
type RoomDeleter interface {
	DeleteARoom(roomName int) error
}
