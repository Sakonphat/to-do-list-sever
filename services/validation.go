package services

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"log"
	"reflect"
	"regexp"
	"strings"
)

func Validation() (*validator.Validate, ut.Translator) {
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

	_ = v.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0} must be maximum in 50 characters.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())
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