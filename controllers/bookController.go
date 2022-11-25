package controllers

import (
	"net/http"
	"strconv"
	"yusnar/rest/api/mvc/entities"
	"yusnar/rest/api/mvc/helper"
	"yusnar/rest/api/mvc/models"
	"yusnar/rest/api/mvc/repositories"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	result, errGet := repositories.GetBooks()
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error get all book"))
	}
	var dataResponse []entities.BookResponse
	for _, v := range result {
		dataResponse = append(dataResponse, entities.BookResponse{
			ID:          v.ID,
			Title:       v.Title,
			Publisher:   v.Publisher,
			Author:      v.Author,
			PublishYear: v.PublishYear,
			UserID:      v.UserID,
		})
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all book", dataResponse))
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	errBind := c.Bind(&book)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind book"))
	}
	result, errCreate := repositories.CreateBook(book)
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error create book"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success create user", result))
}
func GetBookByIdController(c echo.Context) error {
	book := models.Book{}
	// var users []models.User
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error param book"))
	}
	iduint := uint(id)
	result, errGetById := repositories.GetBookById(book, iduint)
	if errGetById != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("not found, error get book by id"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get book by id", result))
}

func DeleteBookByIdController(c echo.Context) error {
	book := models.Book{}
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error param book"))
	}
	iduint := uint(id)
	_, errDelete := repositories.DeleteBookById(book, iduint)
	if errDelete != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("not found, error delete book by id"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete book by id"))
}

func UpdateBookByIdController(c echo.Context) error {
	book := models.Book{}
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error param book"))
	}
	iduint := uint(id)
	result, errRead := repositories.UpdateBookById(book, iduint)
	if errRead != nil {
		return c.JSON(http.StatusFound, helper.FailedResponse("error id not found"))
	}
	c.Bind(&result)
	result2, errUpdate := repositories.DataBookAfterUpdate(result)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update book by id"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update book by id", result2))
}
