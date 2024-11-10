package services

import (
	"backend/api/src/config"
	"backend/api/src/models"
	"fmt"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	
	fmt.Println("all users : ", users)
	result := config.DB.Find(&users)



	fmt.Println("users : ", users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetDetailUser(userId int) (*models.User, error) {
	var user models.User
	// var tes *models.User

	fmt.Println("user detail 1 : ", user)
	findUser := config.DB.Where("id = ?", userId).First(&user)

	
	if findUser.Error != nil {
		return nil, findUser.Error
	}

	fmt.Println("user detail 2 : ", user)

	return &user, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	result := config.DB.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func UpdateUser(userId int, payload models.User) (*models.User, error) {

	var user models.User

	findUser := config.DB.Where("id = ?",userId).First(&user)

	if findUser.Error != nil {
		return nil, findUser.Error
	}

	body := models.User{Name: payload.Name, BirthDate: payload.BirthDate,Address: payload.Address,Phone: payload.Phone}


	updateUser := config.DB.Model(&user).Updates(body)

	if updateUser.Error != nil {
		return nil, updateUser.Error
	}

	fmt.Println("user update : ", user)

	return &user, nil
}

func DeleteUser(userId int) (*models.User, error) {

	var user models.User

	findUser := config.DB.Where("id = ?",userId).First(&user)

	if findUser.Error != nil {
		return nil, findUser.Error
	}
	
	result := config.DB.Delete(&user, userId)

	if result.Error != nil {
		fmt.Println("error : ", result.Error)
		return nil, result.Error
	}

	return &user, nil
}