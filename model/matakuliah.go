package model

import (
	"gorm.io/gorm"
)

type Matakuliah struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Dosen string `json:"dosen" form:"dosen"`
}

type Matakuliahs []Matakuliah
