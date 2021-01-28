package requests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"log"
	"reflect"
	"regexp"
	"strings"
)

type RegisterRequest struct {
	Username    string `json:"username" validate:"required,min=6"`
	Password 	string `json:"password" validate:"required,regexPassword"`
}

type LoginRequest struct {
	Username    string `json:"username" validate:"required"`
	Password 	string `json:"password" validate:"required"`
}

func validation() (*validator.Validate, ut.Translator) {
	v := validator.New()
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} must be at least 6 characters.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())
		return t
	})

	_ = v.RegisterTranslation("regexPassword", trans, func(ut ut.Translator) error {
		return ut.Add("regexPassword", "{0} is not match.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("regexPassword", fe.Field())
		return t
	})

	_ = v.RegisterValidation("regexPassword", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`^[A-Za-z\d]{6,}$`)
		return re.MatchString(fl.Field().String())
	})

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v, trans
}

func RegisterValidation(request RegisterRequest) map[string]string {

	validation, trans := validation()
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

	validation, trans := validation()
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