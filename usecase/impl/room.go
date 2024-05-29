package usecase

import (
	"presensee_project/model"
	database "presensee_project/repository/impl"
)

func CreateRoom(room *model.Room) error {
	err := database.CreateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func GetRoom(id uint) (model.Room, error) {
	room, err := database.GetRoomByID(id)
	if err != nil {
		return model.Room{}, err
	}
	return room, nil
}

func GetListRooms() ([]model.Room, error) {
	rooms, err := database.GetRooms()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func UpdateRoom(room *model.Room) error {
	err := database.UpdateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRoom(id uint) error {
	err := database.DeleteRoom(id)
	if err != nil {
		return err
	}
	return nil
}
