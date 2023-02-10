package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ReformatValidationError(e error) error {
	ve := e.(validator.ValidationErrors)
	errors := "Request validation error: "

	for _, e := range ve {
		subError := fmt.Sprintf("%s doesn't satisfy the %s condition. ", e.Field(), e.ActualTag())
		errors += subError
	}

	return fmt.Errorf(errors)
}
