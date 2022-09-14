package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName    string `json:"fullName"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
