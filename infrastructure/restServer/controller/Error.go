package restServerController

import (
	"dev-knowledge/infrastructure/errors"
	"fmt"
)

var (
	ErrURLParamsIsEmpty             = errors.NewError("62106aaa-001", "URL Params is empty")
	ErrParseURLParams               = errors.NewError("62106aaa-002", "Can't parse URL Params")
	ParameterNotFoundByKeyErrorCode = errors.ErrorCode("62106aaa-003")
)

func ErrParameterNotFoundByKey(key string) error {
	errMsg := fmt.Sprintf("Parameter not found by key. Key = %q", key)
	return errors.NewError(ParameterNotFoundByKeyErrorCode, errMsg)
}
