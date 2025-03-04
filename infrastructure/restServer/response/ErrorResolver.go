package response

import (
	"dev-knowledge/infrastructure/errors"
	"net/http"
)

const (
	UnknownErrorCode = "UNKNOWN_CODE"
)

type ErrorResolver struct{}

func NewErrorResolver() *ErrorResolver {
	return &ErrorResolver{}
}

func (er *ErrorResolver) GetErrorCode(err error) string {
	switch errT := err.(type) {
	case *errors.Error:
		return string(errT.Code())
	default:
		return UnknownErrorCode
	}
}

func (er *ErrorResolver) GetErrorText(err error) string {
	switch errT := err.(type) {
	case *errors.Error:
		return errT.Message()
	default:
		return err.Error()
	}
}

func (er *ErrorResolver) GetHttpCode(err error) int {
	errs, ok := err.(*errors.Error)
	if !ok {
		return http.StatusInternalServerError
	}

	switch errs.Code() {
	default:
		return http.StatusUnprocessableEntity
	}
}
