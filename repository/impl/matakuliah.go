package database

import (
	"presensee_project/config"
	"presensee_project/model"
	"time"
)

func CreateMatakuliah(matakuliah *model.Matakuliah) error {
	if err := config.DB.Create(matakuliah).Error; err != nil {
		return err
	}
	return nil
}

func GetMatakuliahs() (matakuliahs []model.Matakuliah, err error) {
	if err = config.DB.Model(&model.Matakuliah{}).Find(&matakuliahs).Error; err != nil {
		return
	}
	return
}

func GetMatakuliahByNameAndDate(name string, date time.Time) (model.Matakuliah, error) {
	var matakuliah model.Matakuliah
	if err := config.DB.Where("name = ? AND date = ?", name, date).First(&matakuliah).Error; err != nil {
		return model.Matakuliah{}, err
	}
	return matakuliah, nil
}

func GetMatakuliahByID(id uint) (matakuliah model.Matakuliah, err error) {
	if err = config.DB.First(&matakuliah, id).Error; err != nil {
		return
	}
	return
}

func UpdateMatakuliah(matakuliah *model.Matakuliah) error {
	if err := config.DB.Save(matakuliah).Error; err != nil {
		return err
	}
	return nil
}

func DeleteMatakuliah(id uint) error {
	if err := config.DB.Delete(&model.Matakuliah{}, id).Error; err != nil {
		return err
	}
	return nil
}
