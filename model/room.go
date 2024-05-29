package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Location string `json:"location" form:"location"`
	Code     string `json:"code" form:"code"`
}
