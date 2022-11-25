package entities

import "time"

type UserCore struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	Telp_number string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserResponse struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Telp_number string `json:"phone"`
	Address     string `json:"address"`
}
