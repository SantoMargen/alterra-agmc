package database

import (
	"D-2/config"
	"D-2/models"
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
	if err := config.DB.Model(&models.User{}).Where("id=?", userId).First(&user).Error; err != nil {
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
	if err := config.DB.Model(&models.User{}).Where("id=?", userId).Updates(models.User{PhoneNumber: user.PhoneNumber}).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func DeletedUserById(userId int, user models.User) (models.User, error) {
	if err := config.DB.First(&user, userId).Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func LoginUser(user models.User) (models.User, error) {
	result := models.User{}
	if err := config.DB.Model(&models.User{}).Where("email=?", user.Email).First(&result).Error; err != nil || user.Email == "" || user.Password == "" {
		return models.User{}, err
	}
	return result, nil
}
