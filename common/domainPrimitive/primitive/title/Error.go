package titlePrimitive

import "dev-knowledge/infrastructure/errors"

var (
	ErrTitleIsEmpty = errors.NewError("8b3396e1-001", "Title cannot be empty")
	ErrTitleTooLong = errors.NewError("8b3396e1-002", "Title is too long")
)
