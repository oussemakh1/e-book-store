package services

import (
	"rest-api/models"
	"rest-api/repositories"
)


func CreateUser(input models.CreateUserInput) (newUser models.User, err error) {
	user := models.NewUser(input)

	newUser, err = repositories.CreateUser(user)
	if err != nil {
		return user, err
	}
	return newUser, nil
}


func FindUsers() ([]models.User, error) {
	var users []models.User

	users, err := repositories.FindUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}


func FindUser(id int) (models.User, error) {
	var user models.User 
	user, err := repositories.FindUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}


func UpdateUser(input models.UpdateUserInput,id int) (newUser models.User, err error) {
	user := models.UpdateUser(input)

	newUser, err = repositories.UpdateUser(user, id)
	if err != nil {
		return user, err
	}
	return newUser, nil
}


func DeleteUser(id int) error {
	if err := repositories.DeleteUser(id); err != nil {
		return err
	}
	return nil
}