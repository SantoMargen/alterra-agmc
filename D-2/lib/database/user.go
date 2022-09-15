package database

import (
	"D-2/config"
	"D-2/middlewares"
	"D-2/models"
	"log"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil

}

func GetDetailUser(userId int) (models.User, error) {
	user := models.User{}
	if err := config.DB.First(&user, userId).Find(&user, userId).Error; err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func CreateNewUser(user models.User) (models.User, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UserUpdate(userId int, user models.User) (models.User, error) {
	config.DB.First(&user, userId)
	if err := config.DB.Model(&user).Updates(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func DeletedUserById(userId int, user models.User) (models.User, error) {
	if err := config.DB.First(&user, userId).Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func LoginUser(user models.User) (*string, error) {
	if err := config.DB.Where("email=? AND password=?", user.Email, user.Password).Error; err != nil {
		return nil, err
	}
	token, err := middlewares.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
