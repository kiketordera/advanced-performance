package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// MyCustomValidator check if the Atribute contains the word "hello"
func MyCustomValidator(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "hello")
}
