package passwordEntity

import "dev-knowledge/infrastructure/errors"

var (
	ErrPasswordHashIsRequired      = errors.NewError("7cd41798-001", "Password hash is required")
	ErrPasswordSaltIsRequired      = errors.NewError("7cd41798-002", "Password salt is required")
	ErrPlainTextPasswordIsRequired = errors.NewError("7cd41798-003", "Plain text password is required")
	ErrGlobalSaltIsRequired        = errors.NewError("7cd41798-004", "Global salt is required")
	ErrShortSalt                   = errors.NewError("7cd41798-012", "Password salt is short")
	ErrHashIsEmpty                 = errors.NewError("7cd41798-014", "Password hash is empty")
	ErrLongSalt                    = errors.NewError("7cd41798-015", "Password salt is long")

	ErrMismatchedPassword                = errors.NewErrorWithLevel("7cd41798-005", "Mismatched password", errors.Levels.Info())
	ErrShortPassword                     = errors.NewErrorWithLevel("7cd41798-008", "Short password", errors.Levels.Info())
	ErrTooLongPassword                   = errors.NewErrorWithLevel("7cd41798-009", "Password is too long", errors.Levels.Info())
	ErrPasswordContainsProhibitedSymbols = errors.NewErrorWithLevel("7cd41798-010", "Password contains prohibited symbols", errors.Levels.Info())
	ErrPasswordNotContainRequiredSymbols = errors.NewErrorWithLevel("7cd41798-011", "Password doesn't contain required characters", errors.Levels.Info())

	ComparePasswordUnknownErrorCode = errors.ErrorCode("7cd41798-013")
)
