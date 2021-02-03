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
			"message" : requestErr.Error(),
		})
		return
	}

	err := requests.RegisterValidation(request)
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
	if userInDbErr == nil {
		if user.Username == request.Username {
			err := make(map[string]string)
			err["username"] = "This username already exists in the system."
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code" : http.StatusUnprocessableEntity,
				"success" : false,
				"message" : err,
			})
			return
		}
	}

	hash, bcryptErr := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if bcryptErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : bcryptErr.Error(),
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
			"message" : creatErr.Error(),
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
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : http.StatusUnprocessableEntity,
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : errToken.Error(),
		})
		return
	}

	storeRedisErr := services.StoreRedis(user.Uuid, token)
	if storeRedisErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : storeRedisErr.Error(),
		})
		return
	}

	data := make(map[string]string)
	data["token"] = token.AccessToken
	data["token_type"] = "Bearer "

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"message" : "Successfully logged in.",
		"data" : data,
	})
	return
}

func Logout(c *gin.Context)  {

	accessToken, accessTokenErr := services.ParseJwtToken(c.Request)

	if accessTokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code" : http.StatusUnauthorized,
			"success" : false,
			"message" : "unauthorized",
		})
		return
	}

	deleted, delErr := services.DeleteJwtToken(accessToken)
	if delErr != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code" : http.StatusUnauthorized,
			"success" : false,
			"message" : "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"message" : "Successfully logged out.",
	})
	return
}