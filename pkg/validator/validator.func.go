package validator

import (
	v "github.com/go-playground/validator/v10"
)

func IsOmitStringType(fl v.FieldLevel) bool {
	value := fl.Field().String()

	if value == "" {
		return false
	}

	return len(value) > 1
}

func CheckNotEmpty(fl v.FieldLevel) bool {
	value := fl.Field().String()
	return value != ""
}