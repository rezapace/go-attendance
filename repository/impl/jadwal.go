package database

import (
	"presensee_project/config"
	"presensee_project/model"
)

func CreateJadwal(jadwal *model.Jadwal) error {
	if err := config.DB.Create(jadwal).Error; err != nil {
		return err
	}
	return nil
}

func GetJadwals() ([]model.Jadwal, error) {
	var jadwals []model.Jadwal
	if err := config.DB.Find(&jadwals).Error; err != nil {
		return nil, err
	}
	return jadwals, nil
}

func GetJadwalByID(id uint) (model.Jadwal, error) {
	var jadwal model.Jadwal
	if err := config.DB.First(&jadwal, id).Error; err != nil {
		return model.Jadwal{}, err
	}
	return jadwal, nil
}

func UpdateJadwal(jadwal *model.Jadwal) error {
	if err := config.DB.Save(jadwal).Error; err != nil {
		return err
	}
	return nil
}

func DeleteJadwal(id uint) error {
	if err := config.DB.Delete(&model.Jadwal{}, id).Error; err != nil {
		return err
	}
	return nil
}
