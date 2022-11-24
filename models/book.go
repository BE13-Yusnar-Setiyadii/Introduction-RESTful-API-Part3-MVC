package models

import (
	"yusnar/rest/api/mvc/entities"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
	User        User
}

func BookCoreToModel(dataCore entities.BookCore) Book {
	return Book{
		Title:       dataCore.Title,
		Publisher:   dataCore.Publisher,
		Author:      dataCore.Author,
		PublishYear: dataCore.PublishYear,
		UserID:      dataCore.UserID,
	}
}

func ModelToBookCore(dataModel Book) entities.BookCore {
	return entities.BookCore{
		ID:          dataModel.ID,
		Title:       dataModel.Title,
		Publisher:   dataModel.Publisher,
		Author:      dataModel.Author,
		PublishYear: dataModel.PublishYear,
		UserID:      dataModel.UserID,
		User: entities.UserResponse{
			Name:        dataModel.User.Name,
			Email:       dataModel.User.Email,
			Telp_number: dataModel.User.Telp_number,
			Address:     dataModel.User.Address,
		},
	}
}

func ListModelToBookCore(dataModel []Book) []entities.BookCore {
	var result []entities.BookCore
	for _, v := range dataModel {
		result = append(result, ModelToBookCore(v))
	}
	return result
}
