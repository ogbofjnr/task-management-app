package validator

import (
	"fmt"
	libValidator "github.com/go-playground/validator/v10"
	"strings"
)

type Validator struct {
	validator *libValidator.Validate
}

func NewValidator() Validator {
	baseValidator := libValidator.New()
	//baseValidator.RegisterValidation("is_rectangular", custom_rules.CheckRectangular)

	v := Validator{baseValidator}
	return v
}

func (v *Validator) Validate(s interface{}) error {
	return v.validator.Struct(s)
}

func (v *Validator) GetErrors(e error) []string {
	var errors []string
	for _, err := range e.(libValidator.ValidationErrors) {
		m := fmt.Sprintf("validation failed on field '%s', condition: %s", strings.ToLower(err.Field()), err.Tag())
		errors = append(errors, m)
	}
	return errors
}
