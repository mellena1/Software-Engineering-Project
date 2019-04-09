package db

type Room struct {
	ID       *int64  `json:"id" example:"1"`
	Name     *string `json:"name" example:"My Room Name"`
	Capacity *int64  `json:"capacity" example:"50"`
}

func NewRoom() Room {
	return Room{
		ID:       Int64Ptr(0),
		Name:     StringPtr(""),
		Capacity: Int64Ptr(0),
	}
}

type RoomReaderWriterUpdaterDeleter interface {
	RoomReader
	RoomWriter
	RoomUpdater
	RoomDeleter
}

type RoomReader interface {
	ReadARoom(roomID int64) (Room, error)
	ReadAllRooms() ([]Room, error)
}

type RoomWriter interface {
	WriteARoom(name string, capacity *int64) (int64, error)
}

type RoomUpdater interface {
	UpdateARoom(id int64, name string, capacity *int64) error
}

type RoomDeleter interface {
	DeleteARoom(id int64) error
}
