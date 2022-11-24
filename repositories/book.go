package repositories

import (
	"errors"
	"yusnar/rest/api/mvc/config"
	"yusnar/rest/api/mvc/entities"
	"yusnar/rest/api/mvc/models"
)

func GetBooks() ([]entities.BookCore, error) {
	var books []models.Book
	tx := config.DB.Preload("User").Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = models.ListModelToBookCore(books)
	return dataCore, nil
}

func CreateBook(dataCore entities.BookCore) error {
	bookGorm := models.BookCoreToModel(dataCore)
	tx := config.DB.Create(&bookGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("create book failed")
	}
	return nil
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
