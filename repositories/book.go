package repositories

import (
	"errors"
	"yusnar/rest/api/mvc/config"
	"yusnar/rest/api/mvc/entities"
	"yusnar/rest/api/mvc/models"
)

func GetBooks() ([]entities.BookCore, error) {
	var books []models.Book
	tx := config.DB.Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore []entities.BookCore
	for _, v := range books {
		dataCore = append(dataCore, entities.BookCore{
			ID:          v.ID,
			Title:       v.Title,
			Publisher:   v.Publisher,
			Author:      v.Author,
			PublishYear: v.PublishYear,
			UserID:      v.UserID,
		})

	}

	return dataCore, nil
}

func CreateBook(book models.Book) (models.Book, error) {
	tx := config.DB.Create(&book)
	if tx.Error != nil {
		return book, tx.Error
	}
	if tx.RowsAffected == 0 {
		return book, errors.New("create book failed")
	}
	return book, nil
}

func GetBookById(book models.Book, iduint uint) (models.Book, error) {
	tx := config.DB.First(&book, iduint)
	if tx.Error != nil {
		return book, tx.Error
	}
	return book, nil
}

func DeleteBookById(book models.Book, iduint uint) ([]models.Book, error) {
	var books []models.Book
	tx := config.DB.Unscoped().Delete(&book, iduint)
	if tx.Error != nil {
		return books, tx.Error
	}
	if tx.RowsAffected == 0 {
		return books, errors.New("delete book failed")
	}
	return books, nil
}

func UpdateBookById(book models.Book, iduint uint) (models.Book, error) {
	tx := config.DB.First(&book, iduint)
	if tx.Error != nil {
		return book, tx.Error
	}
	return book, nil
}
func DataBookAfterUpdate(book models.Book) (models.Book, error) {
	tx := config.DB.Save(&book)
	if tx.Error != nil {
		return book, tx.Error
	}
	if tx.RowsAffected == 0 {
		return book, errors.New("update book failed")
	}
	return book, nil
}
