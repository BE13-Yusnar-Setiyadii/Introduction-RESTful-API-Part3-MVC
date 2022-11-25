package entities

import "time"

type BookCore struct {
	ID          uint
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookResponse struct {
	ID          uint   `json:"ID"`
	Title       string `json:"title"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	PublishYear string `json:"publish_year"`
	UserID      uint   `json:"user_id"`
}
