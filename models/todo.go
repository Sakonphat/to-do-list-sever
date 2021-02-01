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
func CreateTask(task *Task) (err error) {

	task.Uuid = uuid.NewV4().String()

	if err = config.DB.Create(task).Error; err != nil {
		return err
	}
	return nil
}

//fetch all Task
func GetAllTask(task *[]Task, userId uint) (err error) {
	if err := config.DB.Where("user_id = ?", userId).Order("id desc").Find(task).Error; err != nil {
		return err
	}
	return nil
}

//fetch one Task
func GetATaskByUuid(task *Task, uuid string) (err error) {
	if err := config.DB.Where("uuid = ?", uuid).First(task).Error; err != nil {
		return err
	}
	return nil
}

//update a Task
func UpdateATask(task *Task) (err error) {
	if err := config.DB.Save(task).Error; err != nil {
		return err
	}
	return nil
}

//update a Task by multiple key
func UpdateTaskByMany(task *Task, request map[string]interface{}) (err error) {
	if err := config.DB.Model(task).Updates(request).Error; err != nil {
		return err
	}
	return nil
}

//update complete Task
func UpdateCompletedTask(task *Task) (err error) {
	if err := config.DB.Model(task).Update("is_completed", true).Error; err != nil {
		return err
	}
	return nil
}

//update undo Task
func UpdateUndoTask(task *Task) (err error) {
	if err := config.DB.Model(task).Update("is_completed", false).Error; err != nil {
		return err
	}
	return nil
}

//delete a Task
func DeleteATask(task *Task) (err error) {
	if err := config.DB.Delete(task).Error; err != nil {
		return err
	}
	return nil
}

//delete a Task
func DeleteTasks(task *Task, userId uint) (err error) {
	if err := config.DB.Where("user_id = ?", userId).Delete(task).Error; err != nil {
		return err
	}
	return nil
}