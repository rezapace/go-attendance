package payload

import (
	"time"

	"presensee_project/model"
)

type CreateJadwalRequest struct {
	MatakuliahID uint      `json:"matakuliah_id" validate:"required"`
	RoomID       uint      `json:"room_id" validate:"required"`
	Sks          string    `json:"sks" validate:"required"`
	JamMulai     time.Time `json:"jam_mulai" validate:"required"`
	JamSelesai   time.Time `json:"jam_selesai" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	UserID       uint      `json:"user_id" form:"user_id" validate:"required"`
	DosenID      uint      `json:"dosen_id" form:"dosen_id" validate:"required"`
}

func (u *CreateJadwalRequest) ToEntity() *model.Jadwal {
	return &model.Jadwal{
		MatakuliahID: u.MatakuliahID,
		RoomID:       u.RoomID,
		Sks:          u.Sks,
		JamMulai:     u.JamMulai,
		JamSelesai:   u.JamSelesai,
		Name:         u.Name,
		Description:  u.Description,
		UserID:       u.UserID,
		DosenID:      u.DosenID,
	}
}

type UpdateJadwalRequest struct {
	ID           uint      `json:"id"`
	MatakuliahID uint      `json:"matakuliah_id"`
	RoomID       uint      `json:"room_id"`
	Sks          string    `json:"sks"`
	JamMulai     time.Time `json:"jam_mulai"`
	JamSelesai   time.Time `json:"jam_selesai"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	UserID       uint      `json:"user_id" form:"user_id"`
	DosenID      uint      `json:"dosen_id" form:"dosen_id"`
}

func (u *UpdateJadwalRequest) ToEntity() *model.Jadwal {
	return &model.Jadwal{
		MatakuliahID: u.MatakuliahID,
		RoomID:       u.RoomID,
		Sks:          u.Sks,
		JamMulai:     u.JamMulai,
		JamSelesai:   u.JamSelesai,
		Name:         u.Name,
		Description:  u.Description,
		UserID:       u.UserID,
		DosenID:      u.DosenID,
	}
}

type GetSingleJadwalResponse struct {
	ID           uint             `json:"id"`
	MatakuliahID uint             `json:"matakuliah_id"`
	RoomID       uint             `json:"room_id"`
	Sks          string           `json:"sks"`
	JamMulai     time.Time        `json:"jam_mulai"`
	JamSelesai   time.Time        `json:"jam_selesai"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	UserID       uint             `json:"user_id" form:"user_id"`
	DosenID      uint             `json:"dosen_id" form:"dosen_id"`
	Matakuliah   model.Matakuliah `json:"matakuliah"`
	Dosen        model.Dosen      `json:"dosen"`
	Room         model.Room       `json:"room"`
}

func NewGetSingleJadwalResponse(jadwal *model.Jadwal) *GetSingleJadwalResponse {
	return &GetSingleJadwalResponse{
		ID:           jadwal.ID,
		MatakuliahID: jadwal.MatakuliahID,
		RoomID:       jadwal.RoomID,
		Sks:          jadwal.Sks,
		JamMulai:     jadwal.JamSelesai,
		Name:         jadwal.Name,
		Description:  jadwal.Description,
		UserID:       jadwal.UserID,
		DosenID:      jadwal.DosenID,
	}
}

type GetPageJadwalResponse struct {
	ID           uint        `json:"id"`
	MatakuliahID uint        `json:"matakuliah_id"`
	RoomID       uint        `json:"room_id"`
	Sks          string      `json:"sks"`
	JamMulai     time.Time   `json:"jam_mulai"`
	JamSelesai   time.Time   `json:"jam_selesai"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	UserID       uint        `json:"user_id" form:"user_id"`
	DosenID      uint        `json:"dosen_id" form:"dosen_id"`
	Dosen        model.Dosen `json:"dosen"`
	Room         model.Room  `json:"room"`
}

func NewGetPageJadwalResponse(jadwal *model.Jadwal) *GetPageJadwalResponse {
	return &GetPageJadwalResponse{
		ID:           jadwal.ID,
		MatakuliahID: jadwal.MatakuliahID,
		RoomID:       jadwal.RoomID,
		Sks:          jadwal.Sks,
		JamMulai:     jadwal.JamMulai,
		JamSelesai:   jadwal.JamSelesai,
		Name:         jadwal.Name,
		Description:  jadwal.Description,
		UserID:       jadwal.UserID,
		DosenID:      jadwal.DosenID,
	}
}

type GetPageJadwalsResponse []GetPageJadwalResponse

func NewGetPageJadwalsResponse(jadwals *model.Jadwals) *GetPageJadwalsResponse {
	var briefJadwalsResponse GetPageJadwalsResponse
	for _, jadwal := range *jadwals {
		briefJadwalsResponse = append(briefJadwalsResponse, *NewGetPageJadwalResponse(&jadwal))
	}
	return &briefJadwalsResponse
}
