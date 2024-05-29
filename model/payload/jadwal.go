package payload

import (
	"presensee_project/model"
	"time"
)

type CreateJadwalRequests struct {
	Hari         string    `json:"hari" validate:"required"`
	MatakuliahID uint      `json:"matakuliah_id" validate:"required"`
	RoomID       uint      `json:"room_id" validate:"required"`
	Jam          time.Time `json:"jam" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	UserID       uint      `json:"user_id" form:"user_id" validate:"required"`
	DosenID      uint      `json:"dosen_id" form:"dosen_id" validate:"required"`
}

type CreateJadwalResponse struct {
	JadwalID uint `json:"jadwal_id"`
}

type UpdateJadwalRequests struct {
	Hari         string `json:"hari" validate:"required"`
	MatakuliahID uint   `json:"matakuliah_id" validate:"required"`
	RoomID       uint   `json:"room_id" validate:"required"`
	Jam          string `json:"jam" validate:"required"`
	Name         string `json:"name" validate:"required"`
	UserID       uint   `json:"user_id" form:"user_id" validate:"required"`
}

type UpdateJadwalResponse struct {
	JadwalID uint `json:"jadwal_id"`
}

type GetJadwalResponse struct {
	JadwalID   uint             `json:"jadwal_id"`
	Hari       string           `json:"hari"`
	Matkul     string           `json:"matkul"`
	Ruangan    string           `json:"ruangan"`
	Jam        string           `json:"jam"`
	NamaDosen  string           `json:"nama_dosen"`
	Matakuliah model.Matakuliah `json:"matakuliah"`
	Room       model.Room       `json:"room"`
}

type GetJadwalsResponse struct {
	Jadwals []GetJadwalResponse `json:"jadwals"`
}
