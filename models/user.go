package models

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
	"sever/config"
	"time"
)

type User struct {
	ID uint				`gorm:"column:id;primary_key;"`
	Uuid string			`gorm:"column:uuid;not null;"`
	Username string		`gorm:"column:username;not null;"`
	Password string		`gorm:"column:password;not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *User) TableName() string {
	return "users"
}

//fetch one user by username
func GetUser(user *User, username string) (err error) {
	if err = config.DB.Where("username = ?", username).First(user).Error; err != nil {
		return err
	}
	return nil
}

//fetch one user by uuid
func GetUserByUuid(user *User, uuid string) (err error) {
	if err = config.DB.Where("uuid = ?", uuid).First(user).Error; err != nil {
		return err
	}
	return nil
}

//insert a user
func CreateUser(user *User) (err error) {

	user.Uuid = uuid.NewV4().String()

	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}