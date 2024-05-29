package model

import (
	"time"

	"gorm.io/gorm"
)

type Absen struct {
	gorm.Model
	UserID       uint      `json:"user_id"`
	MahasiswaID  uint      `json:"mahasiswa_id"`
	JadwalID     uint      `json:"jadwal_id"`
	Matakuliah   string    `json:"matakuliah"`
	TimeAttemp   time.Time `json:"time_attemp" form:"time_attemp"`
	Status       string    `json:"status" form:"status"`
	Location     string    `json:"location" form:"location"`
	Image        string    `json:"image" form:"image"`
	IsKonfirmasi bool      `json:"is_konfirmasi" form:"is_konfirmasi"`
}

type Absens []Absen
