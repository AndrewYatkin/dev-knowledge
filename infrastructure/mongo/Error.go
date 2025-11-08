package mongo

import (
	"dev-knowledge/infrastructure/errors"
	"fmt"
)

var (
	ErrMongoDBIsRequired      = errors.NewError("SYS", "MongoDB is required")
	ErrLogPublisherIsRequired = errors.NewError("SYS", "Log publisher is required")

	DuplicateUniqueConstraintErrorCode = errors.ErrorCode("f69256a6-001")
	HandleOperationInMongoErrorCode    = errors.ErrorCode("f69256a6-002")
)

func ErrDuplicateUniqueConstraint(cause error) error {
	detailMsg := fmt.Sprintf("Duplicate unique constraint. Cause: %s", cause)
	return errors.NewError(DuplicateUniqueConstraintErrorCode, detailMsg)
}

func ErrHandleOperationInMongo(mongoDBName, collectionName, methodName string, cause error) error {
	detailMsg := fmt.Sprintf("Mongo handle operation error: %s.%s - %s. cause: %s",
		mongoDBName, collectionName, methodName, cause)
	return errors.NewError(HandleOperationInMongoErrorCode, detailMsg)
}
