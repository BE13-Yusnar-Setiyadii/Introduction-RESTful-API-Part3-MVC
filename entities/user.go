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

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Telp_number string `json:"telp_number" form:"telp_number"`
	Address     string `json:"address" form:"address"`
}

type UserResponse struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Telp_number string `json:"phone"`
	Address     string `json:"address"`
}

func UserCoreToResponse(dataCore UserCore) UserResponse {
	return UserResponse{
		ID:          dataCore.ID,
		Name:        dataCore.Name,
		Email:       dataCore.Email,
		Telp_number: dataCore.Telp_number,
		Address:     dataCore.Address,
	}
}

func ListUserCoreToResponse(dataCore []UserCore) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, UserCoreToResponse(v))
	}
	return dataResponse
}
