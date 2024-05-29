package usecase

import (
	"presensee_project/model"
	database "presensee_project/repository/impl"
)

func CreateDosen(dosen *model.Dosen) error {
	err := database.CreateDosen(dosen)
	if err != nil {
		return err
	}
	return nil
}

func GetDosen(id uint) (model.Dosen, error) {
	dosen, err := database.GetDosenByID(id)
	if err != nil {
		return model.Dosen{}, err
	}
	return dosen, nil
}

func GetDosenByUserID(userID uint) (model.Dosen, error) {
	dosen, err := database.GetDosenByID(userID)
	if err != nil {
		return model.Dosen{}, err
	}
	return dosen, nil
}

func GetListDosens() ([]model.Dosen, error) {
	dosens, err := database.GetDosens()
	if err != nil {
		return nil, err
	}
	return dosens, nil
}

func UpdateDosen(dosen *model.Dosen) error {
	err := database.UpdateDosen(dosen)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDosen(id uint) error {
	err := database.DeleteDosen(id)
	if err != nil {
		return err
	}
	return nil
}
