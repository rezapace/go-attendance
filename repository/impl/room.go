package database

import (
	"presensee_project/config"
	"presensee_project/model"
)

func CreateRoom(room *model.Room) error {
	if err := config.DB.Create(room).Error; err != nil {
		return err
	}
	return nil
}

func GetRooms() (rooms []model.Room, err error) {
	if err = config.DB.Model(&model.Room{}).Find(&rooms).Error; err != nil {
		return
	}
	return
}

func GetRoomByID(id uint) (room model.Room, err error) {
	if err = config.DB.First(&room, id).Error; err != nil {
		return
	}
	return
}

func UpdateRoom(room *model.Room) error {
	if err := config.DB.Save(room).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRoom(id uint) error {
	if err := config.DB.Delete(&model.Room{}, id).Error; err != nil {
		return err
	}
	return nil
}
