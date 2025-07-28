package userUseCase

import "dev-knowledge/infrastructure/errors"

var (
	ErrUserRepoIsRequired = errors.NewError("SYS", "UserRepository is required")
	ErrEmailIsRequired    = errors.NewError("e5885cb2-001", "Email is required")
	ErrPasswordIsRequired = errors.NewError("e5885cb2-002", "Password is required")
	ErrUserIDIsRequired   = errors.NewError("e5885cb2-003", "UserID is required")
)
