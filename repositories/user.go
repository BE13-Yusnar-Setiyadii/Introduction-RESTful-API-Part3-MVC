package repositories

import (
	"errors"
	"yusnar/rest/api/mvc/config"
	"yusnar/rest/api/mvc/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func CreateUser(user models.User) error {
	tx := config.DB.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("create user failed")
	}
	return nil
}

func GetUserById(user models.User, iduint uint) (models.User, error) {
	// var users []models.User
	tx := config.DB.First(&user, iduint)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func DeleteUserById(user models.User, iduint uint) ([]models.User, error) {
	var users []models.User
	tx := config.DB.Delete(&user, iduint)
	if tx.Error != nil {
		return users, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users, errors.New("delete user failed")
	}
	return users, nil
}

func UpdateUserById(user models.User, iduint uint) ([]models.User, error) {
	var users []models.User
	tx := config.DB.First(&user, iduint)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil
}
func DataAfterUpdate(user models.User) error {
	tx := config.DB.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update user failed")
	}
	return nil
}
