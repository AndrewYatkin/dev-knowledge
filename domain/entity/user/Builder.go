package userEntity

import (
	agreementEntity "dev-knowledge/domain/entity/agreement"
	emailEntity "dev-knowledge/domain/entity/email"
	passwordEntity "dev-knowledge/domain/entity/password"
	profileEntity "dev-knowledge/domain/entity/profile"
	"dev-knowledge/domain/entity/spec"
	"dev-knowledge/infrastructure/errors"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

type Builder struct {
	id          *UserID
	profile     *profileEntity.Profile
	role        spec.UserRole
	password    *passwordEntity.Password
	email       *emailEntity.UserEmails
	agreement   *agreementEntity.Agreement
	lastLoginAt *commonTime.Time
	createdAt   *commonTime.Time
	removed     bool

	errors *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) ID(id *UserID) *Builder {
	b.id = id
	return b
}

func (b *Builder) Profile(profile *profileEntity.Profile) *Builder {
	b.profile = profile
	return b
}

func (b *Builder) Role(role spec.UserRole) *Builder {
	b.role = role
	return b
}

func (b *Builder) Password(password *passwordEntity.Password) *Builder {
	b.password = password
	return b
}

func (b *Builder) Email(email *emailEntity.UserEmails) *Builder {
	b.email = email
	return b
}

func (b *Builder) Agreement(agreement *agreementEntity.Agreement) *Builder {
	b.agreement = agreement
	return b
}

func (b *Builder) LastLoginAt(lastLoginAt *commonTime.Time) *Builder {
	b.lastLoginAt = lastLoginAt
	return b
}

func (b *Builder) CreatedAt(createdAt *commonTime.Time) *Builder {
	b.createdAt = createdAt
	return b
}

func (b *Builder) Removed(removed bool) *Builder {
	b.removed = removed
	return b
}

func (b *Builder) Build() (*User, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()

	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.agreement == nil {
		b.errors.AddError(ErrAgreementIsRequired)
	}
	if b.profile == nil {
		b.errors.AddError(ErrProfileIsRequired)
	}
	if b.role == "" {
		b.errors.AddError(ErrRoleIsRequired)
	}
}

func (b *Builder) fillDefaultFields() {
	if b.id == nil {
		b.id = NewUserID()
	}

	if b.createdAt == nil {
		b.createdAt = commonTime.Now()
	}
}

func (b *Builder) createFromBuilder() *User {
	return &User{
		id:          b.id,
		profile:     b.profile,
		role:        b.role,
		password:    b.password,
		email:       b.email,
		agreement:   b.agreement,
		lastLoginAt: b.lastLoginAt,
		createdAt:   b.createdAt,
		removed:     b.removed,
	}
}
