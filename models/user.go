package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Telp_number string `gorm:"type:varchar(15)" json:"telp_number" form:"telp_number"`
	Address     string `json:"address" form:"address"`
	Books       []Book
}
