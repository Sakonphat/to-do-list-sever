package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"net/http"
	"sever/models"
	"sever/utils"
	"strings"
	"time"
)

type TokenDetails struct {
	AccessToken  string
	AccessUuid   string
	AtExpires    int64
}

func GetJwtToken(user models.User) (*TokenDetails, error)  {

	exp := time.Now().Add(time.Minute * 15).Unix()
	accessUuid := uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["uuid"] = user.Uuid
	atClaims["username"] = user.Username
	atClaims["access_uuid"] = accessUuid
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(utils.Env("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	tokenDetails := &TokenDetails{}
	tokenDetails.AccessToken = token
	tokenDetails.AccessUuid = accessUuid
	tokenDetails.AtExpires = exp

	return tokenDetails, nil
}

func ParseJwtToken(request *http.Request) (string, error) {

	token, err := verifyToken(request)
	if err != nil {
		return "", err
	}

	claims, claimsValid := token.Claims.(jwt.MapClaims)

	if claimsValid && token.Valid {

		accessUuid, accessUuidValid := claims["access_uuid"].(string)
		if !accessUuidValid {
			return "", err
		}

		return accessUuid, nil

	}

	return "", err
}

func GetUser(accessToken string) (*models.User, error) {

	uuidStr, err := client.Get(accessToken).Result()
	if err != nil {
		return nil, err
	}

	user := models.User{}

	userErr := models.GetUserByUuid(&user, uuidStr)
	if userErr != nil {
		return nil, userErr
	}

	return &user, nil
}

func DeleteJwtToken(accessToken string) (int64, error) {

	deleted, err := client.Del(accessToken).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}

func getToken(request *http.Request) string {

	bearerToken := request.Header.Get("Authorization")

	strArr := strings.Split(bearerToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func verifyToken(request *http.Request) (*jwt.Token, error) {

	tokenStr := getToken(request)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(utils.Env("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
