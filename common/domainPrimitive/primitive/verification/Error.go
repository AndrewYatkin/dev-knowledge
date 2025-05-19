package verificationPrimitive

import "dev-knowledge/infrastructure/errors"

var (
	ErrCodeRequired            = errors.NewError("a3486fd2-001", "verification code is required")
	ErrCreatedAtRequired       = errors.NewError("a3486fd2-002", "verification createdAt is required")
	ErrExpireAtRequired        = errors.NewError("a3486fd2-003", "verification expireAt is required")
	ErrZeroLifeMinutes         = errors.NewError("a3486fd2-004", "life minutes for verification code can not be equals 0")
	ErrExpireAtBeforeCreatedAt = errors.NewError("a3486fd2-005", "expireAt before createdAt")
	ErrExpiredCode             = errors.NewError("a3486fd2-006", "verification code is expired")
	ErrMismatchedCode          = errors.NewError("a3486fd2-007", "mismatched verification code")
)
