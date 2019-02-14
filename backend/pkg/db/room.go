package db

type Room struct {
	RoomName string
	Capacity int
}

type RoomReaderWriterUpdaterDeleter interface {
	RoomReader
	RoomWriter
	RoomUpdater
	RoomDeleter
}

type RoomReader interface {
	ReadARoom(roomName string) (Room, error)
	ReadAllRooms() ([]Room, error)
}

type RoomWriter interface {
	WriteARoom(r Room) error
}

type RoomUpdater interface {
	UpdateARoom(roomName string, newRoom Room) error
}

type RoomDeleter interface {
	DeleteARoom(roomName int) error
}
