package requests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"sever/services"
)

type EditTaskRequest struct {
	Uuid    string 		`json:"uuid" validate:"required"`
	Title	string 		`json:"title" validate:"required,min=1,max=50"`
	Description string	`json:"description"`
}

type CreateTaskRequest struct {
	Title string 		`json:"title" validate:"required,min=1,max=50"`
	Description string	`json:"description"`
}

func ValidateCreateTaskRequest(request CreateTaskRequest) map[string]string {
	validation, trans := services.Validation()
	err := validation.Struct(request)

	if err != nil {
		errMap := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errMap[e.Field()] = e.Translate(trans)
		}

		return errMap
	}

	return nil
}

func GetCreateTaskRequest(c *gin.Context) (CreateTaskRequest, error) {
	var request CreateTaskRequest

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	jsonErr := json.Unmarshal([]byte(reqBody), &request)

	if jsonErr != nil {
		return request, jsonErr
	}

	return request, nil
}

func ValidateEditTaskRequest(request EditTaskRequest) map[string]string {
	validation, trans := services.Validation()
	err := validation.Struct(request)

	if err != nil {
		errMap := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errMap[e.Field()] = e.Translate(trans)
		}

		return errMap
	}

	return nil
}

func GetEditTaskRequest(c *gin.Context) (EditTaskRequest, error) {
	var request EditTaskRequest

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	jsonErr := json.Unmarshal([]byte(reqBody), &request)

	if jsonErr != nil {
		return request, jsonErr
	}

	return request, nil
}