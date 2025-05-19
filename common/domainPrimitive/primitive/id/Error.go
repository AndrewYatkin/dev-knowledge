package idPrimitive

import (
	"dev-knowledge/infrastructure/errors"
	"fmt"
)

var (
	ErrEntityIDIsEmpty = errors.NewError("02e4e913-001", "Entity id is empty")

	CreateEntityIDErrorCode = errors.ErrorCode("02e4e913-002")
)

func ErrCreateEntityID(invalidID string) error {
	errMsg := fmt.Sprintf("Fail create entityID from string = %q", invalidID)
	return errors.NewError(CreateEntityIDErrorCode, errMsg)
}
