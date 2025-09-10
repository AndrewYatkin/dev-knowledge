package jwtService

import "dev-knowledge/infrastructure/errors"

var (
	ErrJWTUnsupportedSigningMethod = errors.NewError("67806a4a-001", "unsupported signing method")
	ErrJWTWrongClaims              = errors.NewError("67806a4a-002", "wrong claims")
	ErrJWTMissingClaim             = errors.NewError("67806a4a-003", "missing claim")
	ErrJWTInvalidClaimType         = errors.NewError("67806a4a-004", "invalid claim type")
	ErrJWTInvalidClaims            = errors.NewError("67806a4a-005", "invalid claims provided")
	ErrJWTMissingUserID            = errors.NewError("67806a4a-006", "userID is required")
	ErrSignJWT                     = errors.NewError("67806a4a-007", "failed to sign token")
)
