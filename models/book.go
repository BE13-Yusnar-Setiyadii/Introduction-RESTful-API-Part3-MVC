package models

import (
	"yusnar/rest/api/mvc/entities"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string //`json:"title" form:"title"`
	Publisher   string //`json:"publisher" form:"publisher"`
	Author      string //`json:"author" form:"author"`
	PublishYear string //`json:"publisher_year" form:"publisher_year"`
	UserID      uint   //`json:"user_id" form:"user_id"`
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
		User: entities.UserCore{
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
