package requests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"sever/services"
)

type RegisterRequest struct {
	Username    string `json:"username" validate:"required,min=6"`
	Password 	string `json:"password" validate:"required,regexPassword"`
}

type LoginRequest struct {
	Username    string `json:"username" validate:"required"`
	Password 	string `json:"password" validate:"required"`
}

func RegisterValidation(request RegisterRequest) map[string]string {

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

func LoginValidation(request LoginRequest) map[string]string {

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

func GetRegisterRequest(c *gin.Context) (RegisterRequest, error) {
	var request RegisterRequest

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	jsonErr := json.Unmarshal([]byte(reqBody), &request)

	if jsonErr != nil {
		return request, jsonErr
	}

	return request, nil
}

func GetLoginRequest(c *gin.Context) (LoginRequest, error) {
	var request LoginRequest

	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])

	jsonErr := json.Unmarshal([]byte(reqBody), &request)

	if jsonErr != nil {
		return request, jsonErr
	}

	return request, nil
}