package services

import (
	"github.com/dgrijalva/jwt-go"
	"sever/models"
	"sever/utils"
	"time"
)

func GetJwtToken(user models.User) (string, error)  {

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(utils.Env("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
