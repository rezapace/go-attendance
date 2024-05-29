package model

import (
	"time"

	"gorm.io/gorm"
)

type Jadwal struct {
	gorm.Model
	MatakuliahID uint      `json:"matakuliahID" form:"matakuliahID"`
	RoomID       uint      `json:"roomID" form:"roomID"`
	Sks          string    `json:"sks" form:"sks"`
	JamMulai     time.Time `json:"jam_mulai" form:"jama_mulai"`
	JamSelesai   time.Time `json:"jam_selesai" form:"jam_selesai"`
	Description  string    `json:"description" form:"description"`
	Name         string    `json:"name" form:"Name"`
	UserID       uint      `json:"user_id" form:"user_id" validate:"required"`
	DosenID      uint      `json:"dosen_id" form:"dosen_id" validate:"required"`
	Absen        Absen     `json:"-" form:"absen"`
}

type Jadwals []Jadwal
