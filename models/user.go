package models

import (
	"sever/config"
)

type User struct {
	ID uint            `gorm:"column:id;primary_key;"`
	Username string    `gorm:"column:username;"`
	Password string    `gorm:"column:password;"`
}

//type User struct {
//	ID uint            `json:"id"`
//	Username string    `json:"username"`
//	Password string    `json:"password"`
//}

func (b *User) TableName() string {
	return "users"
}

//fetch one user
func GetUser(user *User, username string) (err error) {
	if err = config.DB.Where("username = ?", username).First(user).Error; err != nil {
		return err
	}
	return nil
}

//insert a user
func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}