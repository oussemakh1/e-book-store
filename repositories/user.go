package repositories

import (
	"rest-api/database"
	"rest-api/models"
)

func CreateUser(user models.User) (newUser models.User, err error) {
	
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return user, err
	}

	if err := db.Table("users").Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
	
}


func FindUsers() (users []models.User, err error) {
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return users, err
	}

	if err := db.Table("users").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}


func FindUser(id int) (user models.User, err error) {
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return user, err 
	}

	if err := db.Table("users").Where("id = ?", id).Find(&user).Error; err != nil {
			return user, nil
	}

	return user, nil
}


func UpdateUser(updateUser models.User, id int) (user models.User, err error) {
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return user, err 
	}

	if err := db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	if err := db.Model(&user).Updates(&updateUser).Error; err != nil {
		return user, err
	}

	return user, nil
}


func DeleteUser(id int) error {
	db, err := database.GetDB("DEV")
	defer db.Close()
	if err != nil {
		return err
	}

	var user models.User 
	if err := db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	if err := db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
