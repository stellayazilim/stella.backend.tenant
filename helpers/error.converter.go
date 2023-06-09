package helpers

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

func ListOfErrors(e error) []map[string]string {
	ve := e.(validator.ValidationErrors)
	var invalidFields []map[string]string

	for _, e := range ve {
		errors := map[string]string{}
		errors[ToCamel(e.Field())] = e.Tag()
		invalidFields = append(invalidFields, errors)
	}

	return invalidFields
}
func ToCamel(field string) string {

	runes := []rune(field)

	runes[0] = unicode.ToLower(runes[0])

	return string(runes)
}
