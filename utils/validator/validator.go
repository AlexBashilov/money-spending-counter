package validator

import "github.com/go-playground/validator"

// InitValidator init and return validator
func InitValidator() *validator.Validate {
	validate := validator.New()

	return validate
}
