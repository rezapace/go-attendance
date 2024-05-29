package model

import (
	"gorm.io/gorm"
)

type Jurusan struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Fakultas string `json:"fakultas" form:"fakultas" validate:"required"`
}
