package model

import (
	"gorm.io/gorm"
)

type Dosen struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	NIP    string `json:"nip" form:"nip"`
	Phone  string `json:"phone" form:"phone"`
	Image  string `json:"image" form:"image"`
	UserID uint   `json:"user_id" form:"user_id"`
}
