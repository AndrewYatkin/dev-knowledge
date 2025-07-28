package emailEntity

import "dev-knowledge/infrastructure/errors"

var (
	ErrEmailIsRequired           = errors.NewError("3f5093eb-001", "Email is required")
	ErrCodeToVerifyIsEmpty       = errors.NewError("3f5093eb-002", "Code to verify is empty")
	ErrVerificationCodeNotExists = errors.NewError("3f5093eb-003", "Email verification code not exists")
	ErrNonActivatedEmailNotExist = errors.NewError("3f5093eb-004", "Non-activated email is not exist")
	ErrTimeoutHasNotExpired      = errors.NewError("45c739bc-005", "Timeout has not expired")
)
