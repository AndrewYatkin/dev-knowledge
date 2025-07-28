package profileEntity

import "dev-knowledge/infrastructure/errors"

var (
	ErrFirstNameIsRequired = errors.NewError("bafd47d7-001", "Profile firstName is required")
	ErrLastNameIsRequired  = errors.NewError("bafd47d7-002", "Profile lastName is required")
)
