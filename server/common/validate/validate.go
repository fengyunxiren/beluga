package validate

import (
	"beluga/global"
	"errors"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() error {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("Init validator failed")
	}
	err := validate.RegisterValidation("password", PasswrodValidation)
	if err != nil {
		return err
	}
	return nil
}

func PasswrodValidation(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if len(password) < 8 || len(password) > 20 {
		return false
	}
	containDigit := strings.ContainsAny(password, global.DIGIT)
	containCapital := strings.ContainsAny(password, global.CAPITAL)
	containLowercase := strings.ContainsAny(password, global.LOWERCASE)
	containSpecial := strings.ContainsAny(password, global.SPECIAL)
	if containDigit && containCapital && containLowercase && containSpecial {
		return true
	} else {
		return false
	}
}
