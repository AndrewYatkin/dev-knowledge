package userEntity

import "dev-knowledge/infrastructure/errors"

var (
	ErrUserEmailNotExists  = errors.NewError("ac6310cc-002", "User email not exists")
	ErrAgreementIsRequired = errors.NewError("ac6310cc-004", "User agreement is required")
	ErrProfileIsRequired   = errors.NewError("ac6310cc-005", "User profile is required")
	ErrPasswordIsRequired  = errors.NewError("ac6310cc-006", "User password is required")
	ErrRoleIsRequired      = errors.NewError("ac6310cc-009", "User role is required")
)
