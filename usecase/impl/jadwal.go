package usecase

import (
	"presensee_project/model"
	database "presensee_project/repository/impl"
)

func CreateJadwal(jadwal *model.Jadwal) error {
	err := database.CreateJadwal(jadwal)
	if err != nil {
		return err
	}
	return nil
}

func GetJadwal(id uint) (model.Jadwal, error) {
	jadwal, err := database.GetJadwalByID(id)
	if err != nil {
		return model.Jadwal{}, err
	}
	return jadwal, nil
}

func GetListJadwals() ([]model.Jadwal, error) {
	jadwals, err := database.GetJadwals()
	if err != nil {
		return nil, err
	}
	return jadwals, nil
}

func UpdateJadwal(jadwal *model.Jadwal) error {
	err := database.UpdateJadwal(jadwal)
	if err != nil {
		return err
	}
	return nil
}

func DeleteJadwal(id uint) error {
	err := database.DeleteJadwal(id)
	if err != nil {
		return err
	}
	return nil
}
