package requests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"sever/services"
)

type TodoRequest struct {
	Uuid    string `json:"uuid" validate:"required"`
}

type CreateTaskRequest struct {
	Title string 		`json:"title" validate:"required,min=1,max=50"`
	Description string	`json:"description"`
}

func ValidateTaskRequest(request CreateTaskRequest) map[string]string {
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