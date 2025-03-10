package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func ValidateStruct(s any) error {
	return validate.Struct(s)
}
