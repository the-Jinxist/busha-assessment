package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ReformatValidationError(err error) error {
	ve := err.(validator.ValidationErrors)
	errors := "Request validation error. "

	for _, e := range ve {
		subError := fmt.Sprintf("%s doesn't follow the condition: %s", e.Field(), e.ActualTag())
		errors += fmt.Sprintf("\n %s", subError)
		// field := reflect.TypeOf(e.NameNamespace)

	}

	return fmt.Errorf(errors)
}
