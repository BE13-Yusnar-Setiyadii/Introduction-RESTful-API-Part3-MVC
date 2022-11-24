package controllers

import (
	"net/http"
	"strconv"
	"yusnar/rest/api/mvc/helper"
	"yusnar/rest/api/mvc/models"
	"yusnar/rest/api/mvc/repositories"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	result, errGet := repositories.GetUsers()
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error get all user"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all user", result))
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind user"))
	}
	errCreate := repositories.CreateUser(user)
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error create user"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success create user", user))
}

func GetUserByIdController(c echo.Context) error {
	user := models.User{}
	// var users []models.User
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error param user"))
	}
	iduint := uint(id)
	result, errGetById := repositories.GetUserById(user, iduint)
	if errGetById != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("not found, error get user by id"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get user by id", result))
}

func DeleteUserByIdController(c echo.Context) error {
	user := models.User{}
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error param user"))
	}
	iduint := uint(id)
	_, errDelete := repositories.DeleteUserById(user, iduint)
	if errDelete != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("not found, error delete user by id"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete user by id"))
}

func UpdateUserByIdController(c echo.Context) error {
	user := models.User{}
	id, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error param user"))
	}
	iduint := uint(id)
	_, errRead := repositories.UpdateUserById(user, iduint)
	if errRead != nil {
		return c.JSON(http.StatusFound, helper.FailedResponse("error id not found"))
	}
	c.Bind(&user)
	errUpdate := repositories.DataAfterUpdate(user)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update user by id"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update user by id", user))
}
