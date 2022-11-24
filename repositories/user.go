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
	var dataCore []entities.UserCore
	for _, v := range users {
		dataCore = append(dataCore, entities.UserCore{
			ID:          v.ID,
			Name:        v.Name,
			Email:       v.Email,
			Password:    v.Password,
			Telp_number: v.Telp_number,
			Address:     v.Address,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})

	}

	return dataCore, nil
}

func CreateUser(user models.User) (models.User, error) {
	tx := config.DB.Create(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, errors.New("create user failed")
	}
	return user, nil
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
