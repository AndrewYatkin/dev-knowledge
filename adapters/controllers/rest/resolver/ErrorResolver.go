package restResolver

import (
	userUseCase "dev-knowledge/domain/useCase"
	"dev-knowledge/infrastructure/errors"
	"dev-knowledge/infrastructure/restServer/response"
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
	case response.UnmarshalRequestErrorCode:
		return http.StatusBadRequest
	case userUseCase.ErrPermissionDenied.Code():
		return http.StatusForbidden
	default:
		return http.StatusUnprocessableEntity
	}
}
