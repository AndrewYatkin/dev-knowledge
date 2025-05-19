package verificationPrimitive

import (
	commonTime "dev-knowledge/infrastructure/tools/time"
)

type VerificationCode struct {
	code      string
	createdAt *commonTime.Time
	expireAt  *commonTime.Time
}

func (v *VerificationCode) Verify(code string) error {
	now := commonTime.Now()
	if v.expireAt.Before(now) {
		return ErrExpiredCode
	}

	if v.code != code {
		return ErrMismatchedCode
	}

	return nil
}

func (v *VerificationCode) Code() string {
	return v.code
}

func (v *VerificationCode) ExpireAt() *commonTime.Time {
	return v.expireAt
}

func (v *VerificationCode) CreatedAt() *commonTime.Time {
	return v.createdAt
}
