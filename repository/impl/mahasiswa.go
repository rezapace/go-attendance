package database

import (
	"presensee_project/config"
	"presensee_project/model"
)

func CreateMahasiswa(mahasiswa *model.Mahasiswa) error {
	if err := config.DB.Create(mahasiswa).Error; err != nil {
		return err
	}
	return nil
}

func GetMahasiswas() (mahasiswas []model.Mahasiswa, err error) {
	if err = config.DB.Model(&model.Mahasiswa{}).Find(&mahasiswas).Error; err != nil {
		return
	}
	return
}

func GetMahasiswaByID(id uint) (mahasiswa model.Mahasiswa, err error) {
	if err = config.DB.First(&mahasiswa, id).Error; err != nil {
		return
	}
	return
}

func GetMahasiswaByUserID(userID uint) (mahasiswa model.Mahasiswa, err error) {
	if err = config.DB.Where(model.Mahasiswa{UserID: userID}).First(&mahasiswa).Error; err != nil {
		return
	}
	return
}

func UpdateMahasiswa(mahasiswa *model.Mahasiswa) error {
	if err := config.DB.Save(mahasiswa).Error; err != nil {
		return err
	}
	return nil
}

func DeleteMahasiswa(id uint) error {
	if err := config.DB.Delete(&model.Mahasiswa{}, id).Error; err != nil {
		return err
	}
	return nil
}
