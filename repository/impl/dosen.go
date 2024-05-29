package database

import (
	"presensee_project/config"
	"presensee_project/model"
)

func CreateDosen(dosen *model.Dosen) error {
	if err := config.DB.Create(dosen).Error; err != nil {
		return err
	}
	return nil
}

func GetDosens() (dosens []model.Dosen, err error) {
	if err = config.DB.Model(&model.Dosen{}).Find(&dosens).Error; err != nil {
		return
	}
	return
}

func GetDosenByID(id uint) (dosen model.Dosen, err error) {
	if err = config.DB.First(&dosen, id).Error; err != nil {
		return
	}
	return
}

func GetDosenByUserID(userID uint) (dosen model.Dosen, err error) {
	if err = config.DB.Where(model.Dosen{UserID: userID}).First(&dosen).Error; err != nil {
		return
	}
	return
}

func UpdateDosen(dosen *model.Dosen) error {
	if err := config.DB.Save(dosen).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDosen(id uint) error {
	if err := config.DB.Delete(&model.Dosen{}, id).Error; err != nil {
		return err
	}
	return nil
}
