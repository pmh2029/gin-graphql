package validator

import (
	"github.com/gin-gonic/gin/binding"
	v "github.com/go-playground/validator/v10"
)

// New func: calls validator.New and add custom validators
func New() *v.Validate {
	validate, ok := binding.Validator.Engine().(*v.Validate)
	if !ok {
		return nil
	}

	_ = validate.RegisterValidation("omitstring", IsOmitStringType)
	_ = validate.RegisterValidation("notempty", CheckNotEmpty)
	return validate
}
