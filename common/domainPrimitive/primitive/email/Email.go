package emailPrimitive

import (
	"github.com/go-playground/validator"
	"strings"
)

type Email string

func EmailFrom(emailStr string) (Email, error) {
	trimmedEmailStr := strings.TrimSpace(emailStr)
	trimmedEmailStr = strings.ToLower(trimmedEmailStr)
	validate := validator.New()
	err := validate.Var(trimmedEmailStr, "required,email")
	if err != nil {
		return "", ErrWrongEmail
	}

	return Email(trimmedEmailStr), nil
}

func (e Email) String() string {
	return string(e)
}
