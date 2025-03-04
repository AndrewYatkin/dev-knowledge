package response

import (
	"dev-knowledge/infrastructure/errors"
	"fmt"
)

var (
	ErrWriteResponse = errors.NewError("3deacfe7-001", "An error occurred at write http response")

	UnmarshalRequestErrorCode errors.ErrorCode = "3deacfe7-011"
)

func ErrUnmarshalRequest(causeDescription string) error {
	errMsg := fmt.Sprintf("Malformed request. Cause - %s", causeDescription)
	return errors.NewError(UnmarshalRequestErrorCode, errMsg)
}
