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

func GetJadwals() (jadwals []model.Jadwal, err error) {
	if err = config.DB.Model(&model.Jadwal{}).Find(&jadwals).Error; err != nil {
		return
	}
	return
}

func GetJadwalByID(id uint) (jadwal model.Jadwal, err error) {
	if err = config.DB.First(&jadwal, id).Error; err != nil {
		return
	}
	return
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
