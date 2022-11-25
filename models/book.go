package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title" form:"title"`
	Publisher   string `json:"publisher" form:"publisher"`
	Author      string `json:"author" form:"author"`
	PublishYear string `json:"publisher_year" form:"publisher_year"`
	UserID      uint   `json:"user_id" form:"user_id"`
	User        User
}
