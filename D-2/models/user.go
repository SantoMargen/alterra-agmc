package models

import (
	"D-2/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FullName    string `json:"fullName"`
	Age         int    `json:"age"`
	Email       string `json:"email" validate:"required"`
	Username    string `json:"username"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPassword(u.Password)
	return
}
