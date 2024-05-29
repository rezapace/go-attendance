package database

import (
	"presensee_project/config"
	"presensee_project/model"
)

func CreateJurusan(jurusan *model.Jurusan) error {
	if err := config.DB.Create(jurusan).Error; err != nil {
		return err
	}
	return nil
}

func GetJurusans() (jurusans []model.Jurusan, err error) {
	if err = config.DB.Model(&model.Jurusan{}).Find(&jurusans).Error; err != nil {
		return
	}
	return
}

func GetJurusanByID(id uint) (jurusan model.Jurusan, err error) {
	if err = config.DB.First(&jurusan, id).Error; err != nil {
		return
	}
	return
}

func UpdateJurusan(jurusan *model.Jurusan) error {
	if err := config.DB.Save(jurusan).Error; err != nil {
		return err
	}
	return nil
}

func DeleteJurusan(id uint) error {
	if err := config.DB.Delete(&model.Jurusan{}, id).Error; err != nil {
		return err
	}
	return nil
}
