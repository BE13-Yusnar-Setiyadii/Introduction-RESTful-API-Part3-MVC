package models

import (
	"yusnar/rest/api/mvc/entities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string //`json:"name" form:"name"`
	Email       string //`json:"email" form:"email"`
	Password    string //`json:"password" form:"password"`
	Telp_number string `gorm:"type:varchar(15)"` //json:"telp_number" form:"telp_number"`
	Address     string //`json:"address" form:"address"`
	Books       []Book
}

func UserCoreToModel(dataCore entities.UserCore) User {
	userGorm := User{
		Name:        dataCore.Name,
		Email:       dataCore.Email,
		Password:    dataCore.Password,
		Telp_number: dataCore.Telp_number,
		Address:     dataCore.Address,
	}
	return userGorm
}

func ModelToUserCore(dataModel User) entities.UserCore {
	return entities.UserCore{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		Telp_number: dataModel.Telp_number,
		Address:     dataModel.Address,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}

func ListModelToUserCore(dataModel []User) []entities.UserCore {
	var dataCore []entities.UserCore
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToUserCore(v))
	}
	return dataCore
}
