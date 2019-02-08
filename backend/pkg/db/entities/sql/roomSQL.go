package sql

import (
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db/entities"
)

type RoomSQL struct{}

func (RoomSQL) ReadARoom(id int) (entities.Room, error) {
	return entities.Room{}, nil
}

func (RoomSQL) ReadAllRooms() ([]entities.Room, error) {
	return []entities.Room{}, nil
}

func (RoomSQL) WriteARoom(r entities.Room) error {
	return nil
}

func (RoomSQL) DeleteARoom(id int) error {
	return nil
}
