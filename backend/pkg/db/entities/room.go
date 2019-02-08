package entities

type Room struct {
	ID       int
	Capacity int
}

type RoomReader interface {
	ReadARoom(id int) (Room, error)
	ReadAllRooms() ([]Room, error)
}

type RoomWriter interface {
	WriteARoom(r Room) error
}

type RoomDeleter interface {
	DeleteARoom(id int) error
}
