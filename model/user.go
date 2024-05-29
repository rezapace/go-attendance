package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Role      string    `json:"role" form:"role"`
	Dosen     Dosen     `json:"-" form:"dosen"`
	Mahasiswa Mahasiswa `json:"-" form:"mahasiswa"`
	Absen     Absen     `json:"-" form:"absen"`
}

type Users []User
