package usecase

import (
	"presensee_project/model"
	database "presensee_project/repository/impl"
	"time"
)

func CreateMatakuliah(matakuliah *model.Matakuliah) error {
	err := database.CreateMatakuliah(matakuliah)
	if err != nil {
		return err
	}
	return nil
}

func GetMatakuliah(id uint) (model.Matakuliah, error) {
	matakuliah, err := database.GetMatakuliahByID(id)
	if err != nil {
		return model.Matakuliah{}, err
	}
	return matakuliah, nil
}

func GetMatakuliahByNameAndDate(name string, date time.Time) (model.Matakuliah, error) {
	matakuliah, err := database.GetMatakuliahByNameAndDate(name, date)
	if err != nil {
		return model.Matakuliah{}, err
	}
	return matakuliah, nil
}

func GetListMatakuliahs() ([]model.Matakuliah, error) {
	matakuliahs, err := database.GetMatakuliahs()
	if err != nil {
		return nil, err
	}
	return matakuliahs, nil
}

func UpdateMatakuliah(matakuliah *model.Matakuliah) error {
	err := database.UpdateMatakuliah(matakuliah)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMatakuliah(id uint) error {
	err := database.DeleteMatakuliah(id)
	if err != nil {
		return err
	}
	return nil
}
