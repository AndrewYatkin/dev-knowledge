package emailEntity

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	"dev-knowledge/infrastructure/errors"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

type Builder struct {
	userEmail *UserEmail
	errors    *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		userEmail: &UserEmail{},
		errors:    errors.NewErrors(),
	}
}

func (b *Builder) ID(id *EmailID) *Builder {
	b.userEmail.id = id
	return b
}

func (b *Builder) Email(email emailPrimitive.Email) *Builder {
	b.userEmail.email = email
	return b
}

func (b *Builder) VerificationCode(code *verificationPrimitive.VerificationCode) *Builder {
	b.userEmail.verificationCode = code
	return b
}

func (b *Builder) CreatedAt(time *commonTime.Time) *Builder {
	b.userEmail.createdAt = time
	return b
}

func (b *Builder) Build() (*UserEmail, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()

	return b.userEmail, nil
}

func (b *Builder) checkRequiredFields() {
	if b.userEmail.email == "" {
		b.errors.AddError(ErrEmailIsRequired)
	}
}

func (b *Builder) fillDefaultFields() {
	if b.userEmail.id == nil {
		b.userEmail.id = NewEmailID()
	}

	if b.userEmail.createdAt == nil {
		b.userEmail.createdAt = commonTime.Now()
	}
}
