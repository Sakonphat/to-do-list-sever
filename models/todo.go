package models

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
	"sever/config"
	"time"
)

type Task struct {
	ID uint				`gorm:"column:id;primary_key;"`
	Uuid string			`gorm:"column:uuid;not null;"`
	UserId uint			`gorm:"column:user_id;not null;"`
	Title string		`gorm:"column:title;not null;"`
	Description string	`gorm:"column:description;"`
	IsCompleted bool	`gorm:"column:is_completed;type:bool;default:false;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *Task) TableName() string {
	return "tasks"
}

//insert a Task
func CreateTask(todo *Task) (err error) {

	todo.Uuid = uuid.NewV4().String()

	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//fetch all Task
func GetAllTask(todo *[]Task, userId uint) (err error) {
	if err := config.DB.Where("user_id = ?", userId).Order("id desc").Find(todo).Error; err != nil {
		return err
	}
	return nil
}

//fetch one Task
func GetATaskByUuid(todo *Task, uuid string) (err error) {
	if err := config.DB.Where("uuid = ?", uuid).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//update a Task
func UpdateATaskByUuid(todo *Task, uuid string) (err error) {
	if err := config.DB.Where("uuid = ?", uuid).Updates(todo).Error; err != nil {
		return err
	}
	return nil
}

//delete a Task
func DeleteATaskByUuid(todo *Task, uuid string) (err error) {
	if err := config.DB.Where("id = ?", uuid).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}