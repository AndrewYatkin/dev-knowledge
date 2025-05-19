package emailEntity

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

const verificationCodeTtlMinutes uint32 = 60

type UserEmail struct {
	id               *EmailID
	email            emailPrimitive.Email
	verificationCode *verificationPrimitive.VerificationCode
	createdAt        *commonTime.Time
}

func (e *UserEmail) ID() *EmailID {
	return e.id
}

func (e *UserEmail) Email() emailPrimitive.Email {
	return e.email
}

func (e *UserEmail) VerificationCode() string {
	if e.verificationCode == nil {
		return ""
	}

	return e.verificationCode.Code()
}

func (e *UserEmail) VerificationCodeData() (vc verificationPrimitive.VerificationCode, exist bool) {
	if e.verificationCode == nil {
		return verificationPrimitive.VerificationCode{}, false
	}
	return *e.verificationCode, true
}

func (e *UserEmail) CreatedAt() *commonTime.Time {
	return e.createdAt
}

func (e *UserEmail) InitVerificationCode() error {
	verificationCode, err := verificationPrimitive.NewRandomVerificationCode(verificationCodeTtlMinutes)
	if err != nil {
		return err
	}

	e.verificationCode = verificationCode
	return nil
}

func (e *UserEmail) Verify(verificationCode string) error {
	if verificationCode == "" {
		return ErrCodeToVerifyIsEmpty
	}

	if e.verificationCode == nil {
		return ErrVerificationCodeNotExists
	}

	err := e.verificationCode.Verify(verificationCode)

	if err != nil {
		return err
	}

	e.verificationCode = nil
	return nil
}
