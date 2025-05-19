package verificationPrimitive

import (
	"dev-knowledge/infrastructure/errors"
	commonTime "dev-knowledge/infrastructure/tools/time"
	"time"
)

type Builder struct {
	verificationCode *VerificationCode
	errors           *errors.Errors
}

func NewRandomVerificationCode(lifeMinutes uint32) (*VerificationCode, error) {
	code, err := GenerateCode()
	if err != nil {
		return nil, err
	}

	return NewVerificationCode(code, lifeMinutes)
}

func NewVerificationCode(code string, lifeMinutes uint32) (*VerificationCode, error) {
	if lifeMinutes == 0 {
		return nil, ErrZeroLifeMinutes
	}

	createdAt := commonTime.Now()
	expireAt := createdAt.Add(time.Minute * time.Duration(lifeMinutes))

	verificationCode, err := NewBuilder().
		Code(code).
		CreatedAt(createdAt).
		ExpireAt(expireAt).
		Build()

	if err != nil {
		return nil, err
	}

	return verificationCode, nil
}

func NewBuilder() *Builder {
	return &Builder{
		verificationCode: &VerificationCode{},
		errors:           errors.NewErrors(),
	}
}

func (b *Builder) Code(code string) *Builder {
	b.verificationCode.code = code
	return b
}

func (b *Builder) ExpireAt(expireAt *commonTime.Time) *Builder {
	b.verificationCode.expireAt = expireAt
	return b
}

func (b *Builder) CreatedAt(createdAt *commonTime.Time) *Builder {
	b.verificationCode.createdAt = createdAt
	return b
}

func (b *Builder) Build() (*VerificationCode, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.verificationCode, nil
}

func (b *Builder) checkRequiredFields() {
	if b.verificationCode.code == "" {
		b.errors.AddError(ErrCodeRequired)
	}

	if b.verificationCode.createdAt == nil {
		b.errors.AddError(ErrCreatedAtRequired)
	}

	if b.verificationCode.expireAt == nil {
		b.errors.AddError(ErrExpireAtRequired)
	}

	expireAtBeforeCreated := b.verificationCode.createdAt != nil &&
		b.verificationCode.expireAt != nil &&
		!b.verificationCode.expireAt.After(b.verificationCode.createdAt)

	if expireAtBeforeCreated {
		b.errors.AddError(ErrExpireAtBeforeCreatedAt)
	}
}
