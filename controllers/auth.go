package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sever/controllers/requests"
	"sever/models"
	"sever/services"
)

//Register User
func Register(c *gin.Context) {

	request, requestErr := requests.GetRegisterRequest(c)
	if requestErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : "Can not get register request.",
			"errors" : requestErr,
		})
		return
	}

	err := requests.RegisterValidation(request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : "Register validation is error.",
			"errors" : err,
		})
		return
	}

	user := models.User{}
	userInDbErr := models.GetUser(&user, request.Username)
	if userInDbErr == nil {
		if user.Username == request.Username {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code" : http.StatusUnprocessableEntity,
				"success" : false,
				"message" : "This username already exists in the system.",
				"errors" : nil,
			})
			return
		}
	}

	hash, bcryptErr := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if bcryptErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : "Can not do something the password.",
			"errors" : bcryptErr,
		})
		return
	}

	user.Username = request.Username
	user.Password = string(hash)

	creatErr := models.CreateUser(&user)
	if creatErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : "Register is failed.",
			"errors" : nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"success" : true,
		"message" : "Register is success.",
	})
	return
}

func Login(c *gin.Context)  {

	request, requestErr := requests.GetLoginRequest(c)
	if requestErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : requestErr,
		})
		return
	}

	err := requests.LoginValidation(request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : err,
		})
		return
	}

	user := models.User{}
	userInDbErr := models.GetUser(&user, request.Username)
	if userInDbErr != nil {
		errMap := make(map[string]string)
		errMap["username"] = "This username is not exists in the system."
		c.JSON(http.StatusNotFound, gin.H{
			"code" : http.StatusNotFound,
			"success" : false,
			"message" : errMap,
		})
		return
	}

	errPwd := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if errPwd != nil {
		errMap := make(map[string]string)
		errMap["password"] = "Password is invalid."
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
			"success" : false,
			"message" : errMap,
		})
		return
	}

	token, errToken := services.GetJwtToken(user)
	if errToken != nil {
		errMap := make(map[string]string)
		errMap["token"] = errToken.Error()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : errMap,
		})
		return
	}

	data := make(map[string]string)
	data["token"] = token
	data["token_type"] = "Bearer "

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"message" : "Successfully logged in.",
		"data" : data,
	})
	return
}