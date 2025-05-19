package userRepo

import "dev-knowledge/infrastructure/errors"

var (
	ErrLoggerIsRequired = errors.NewError("SYS", "Logger is required")
)
