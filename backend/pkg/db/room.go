package db

type Room struct {
	ID       int
	Capacity int
}

type RoomReaderWriterUpdaterDeleter interface {
	RoomReader
	RoomWriter
	RoomUpdater
	RoomDeleter
}

type RoomReader interface {
	ReadARoom(id int) (Room, error)
	ReadAllRooms() ([]Room, error)
}

type RoomWriter interface {
	WriteARoom(r Room) error
}

type RoomUpdater interface {
	UpdateARoom(oldRoomID int, newRoom Room) error
}

type RoomDeleter interface {
	DeleteARoom(id int) error
}
