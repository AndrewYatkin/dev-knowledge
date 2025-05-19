package emailEntity

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	commonTime "dev-knowledge/infrastructure/tools/time"
	"time"
)

const ActivationCodeGenerateTimeout = 60 * time.Second

type UserEmails struct {
	activatedEmail    *UserEmail
	notActivatedEmail *UserEmail
}

func NewEmails(email emailPrimitive.Email) (*UserEmails, error) {
	notActivatedEmail, err := NewBuilder().Email(email).Build()
	if err != nil {
		return nil, err
	}

	emails := &UserEmails{
		notActivatedEmail: notActivatedEmail,
	}

	return emails, nil
}

func (e *UserEmails) InitNewEmail(newEmail emailPrimitive.Email) error {
	notActivatedEmail, err := NewBuilder().Email(newEmail).Build()
	if err != nil {
		return err
	}

	e.notActivatedEmail = notActivatedEmail
	return nil
}

func (e *UserEmails) GenerateActivationCode() (string, error) {
	if e.notActivatedEmail == nil {
		return "", ErrNonActivatedEmailNotExist
	}

	err := e.canGenerateCode(e.notActivatedEmail)
	if err != nil {
		return "", err
	}

	err = e.notActivatedEmail.InitVerificationCode()
	if err != nil {
		return "", err
	}

	return e.notActivatedEmail.VerificationCode(), nil
}

func (e *UserEmails) Activate(verificationCode string) error {
	if e.notActivatedEmail == nil {
		return ErrNonActivatedEmailNotExist
	}

	err := e.notActivatedEmail.Verify(verificationCode)
	if err != nil {
		return err
	}

	e.activatedEmail = e.notActivatedEmail
	e.notActivatedEmail = nil
	return nil
}

func (e *UserEmails) HasActivated() bool {
	return e.activatedEmail != nil
}

func (e *UserEmails) ActivatedEmail() emailPrimitive.Email {
	if e.activatedEmail == nil {
		return ""
	}
	return e.activatedEmail.Email()
}

func (e *UserEmails) NonActivatedEmail() emailPrimitive.Email {
	if e.notActivatedEmail == nil {
		return ""
	}
	return e.notActivatedEmail.Email()
}

func (e *UserEmails) NonActivatedEmailData() (ue UserEmail, exists bool) {
	if e.notActivatedEmail == nil {
		return UserEmail{}, false
	}
	return *e.notActivatedEmail, true
}

func (e *UserEmails) ActivatedEmailData() (ue UserEmail, exists bool) {
	if e.activatedEmail == nil {
		return UserEmail{}, false
	}
	return *e.activatedEmail, true
}

func (e *UserEmails) canGenerateCode(email *UserEmail) error {
	if email.verificationCode == nil {
		return nil
	}

	nextNotifyTime := email.verificationCode.CreatedAt().Add(ActivationCodeGenerateTimeout)
	if commonTime.Now().Before(nextNotifyTime) {
		return ErrTimeoutHasNotExpired
	}

	return nil
}
