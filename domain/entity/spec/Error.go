package spec

import (
	"dev-knowledge/infrastructure/errors"
	"fmt"
)

var (
	UnsupportedUserRoleErrorCode errors.ErrorCode = "80ebcff6-001"
)

func ErrUnsupportedUserRole(unsupportedUserRole string) *errors.Error {
	errMessage := fmt.Sprintf("Unsupported UserRole = '%s'", unsupportedUserRole)
	return errors.NewError(UnsupportedUserRoleErrorCode, errMessage)
}
