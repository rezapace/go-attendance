package usecase

import (
	"presensee_project/model"
	database "presensee_project/repository/impl"
)

func CreateJurusan(jurusan *model.Jurusan) error {
	err := database.CreateJurusan(jurusan)
	if err != nil {
		return err
	}
	return nil
}

func GetJurusan(id uint) (model.Jurusan, error) {
	jurusan, err := database.GetJurusanByID(id)
	if err != nil {
		return model.Jurusan{}, err
	}
	return jurusan, nil
}

func GetListJurusans() ([]model.Jurusan, error) {
	jurusans, err := database.GetJurusans()
	if err != nil {
		return nil, err
	}
	return jurusans, nil
}

func UpdateJurusan(jurusan *model.Jurusan) error {
	err := database.UpdateJurusan(jurusan)
	if err != nil {
		return err
	}
	return nil
}

func DeleteJurusan(id uint) error {
	err := database.DeleteJurusan(id)
	if err != nil {
		return err
	}
	return nil
}
