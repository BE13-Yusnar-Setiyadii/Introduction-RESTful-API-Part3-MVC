package repositories

import (
	"errors"
	"yusnar/rest/api/mvc/config"
	"yusnar/rest/api/mvc/entities"
	"yusnar/rest/api/mvc/models"
)

func GetUsers() ([]entities.UserCore, error) {
	var users []models.User
	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = models.ListModelToUserCore(users)
	return dataCore, nil
}

func CreateUser(dataCore entities.UserCore) error {
	user := models.UserCoreToModel(dataCore)
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
	tx := config.DB.First(&user, iduint)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func DeleteUserById(user models.User, iduint uint) ([]models.User, error) {
	var users []models.User
	tx := config.DB.Unscoped().Delete(&user, iduint)
	if tx.Error != nil {
		return users, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users, errors.New("delete user failed")
	}
	return users, nil
}

func UpdateUserById(user models.User, iduint uint) (models.User, error) {
	tx := config.DB.First(&user, iduint)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}
func DataAfterUpdate(user models.User) (models.User, error) {
	tx := config.DB.Save(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, errors.New("update user failed")
	}
	return user, nil
}
