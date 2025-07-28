package userEntity

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	"dev-knowledge/domain/entity/user/agreement"
	"dev-knowledge/domain/entity/user/email"
	passwordEntity2 "dev-knowledge/domain/entity/user/password"
	"dev-knowledge/domain/entity/user/profile"
	"dev-knowledge/domain/entity/user/spec"
	commonTime "dev-knowledge/infrastructure/tools/time"
)

type User struct {
	id          *UserID
	profile     *profileEntity.Profile
	role        spec.UserRole
	password    *passwordEntity2.Password
	email       *emailEntity.UserEmails
	agreement   *agreementEntity.Agreement
	lastLoginAt *commonTime.Time
	createdAt   *commonTime.Time
	removed     bool
}

func CreateByRole(userRole spec.UserRole) (*User, error) {
	newAgreement, err := agreementEntity.NewAgreement()
	if err != nil {
		return nil, err
	}

	user, err := NewBuilder().
		Profile(profileEntity.NewEmptyProfile()).
		Role(userRole).
		Agreement(newAgreement).
		Build()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) AcceptTerms() {
	u.agreement.Accept()
}

func (u *User) SetFirstName(name string) error {
	err := u.profile.SetFirstName(name)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) SetLastName(name string) error {
	err := u.profile.SetLastName(name)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) SetPatronymic(name string) {
	u.profile.SetPatronymic(name)
}

func (u *User) SetPassword(plainTextPassword string, globalSalt passwordEntity2.Salt) error {
	password, err := passwordEntity2.NewPassword(plainTextPassword, globalSalt)
	if err != nil {
		return err
	}

	u.password = password
	return nil
}

func (u *User) SetLastLoginAt(lastLoginAt *commonTime.Time) {
	u.lastLoginAt = lastLoginAt
}

func (u *User) GenerateActivateEmailCode() (string, error) {
	if u.email == nil {
		return "", ErrUserEmailNotExists
	}

	return u.email.GenerateActivationCode()
}

func (u *User) ActivateEmail(verificationCode string) error {
	if u.email == nil || u.email.NonActivatedEmail() == "" {
		return ErrUserEmailNotExists
	}

	return u.email.Activate(verificationCode)
}

func (u *User) MatchPassword(plainTextPassword string, globalSalt passwordEntity2.Salt) error {
	if u.password == nil {
		return ErrPasswordIsRequired
	}

	return u.password.MatchPassword(plainTextPassword, globalSalt)
}
func (u *User) ID() *UserID {
	return u.id
}

func (u *User) CreatedAt() *commonTime.Time {
	return u.createdAt
}

func (u *User) Email() (e emailPrimitive.Email, exists bool) {
	if u.email == nil || u.email.ActivatedEmail() == "" {
		return "", false
	}

	return u.email.ActivatedEmail(), true
}

func (u *User) NonActivatedEmail() (e emailPrimitive.Email, exists bool) {
	if u.email == nil || u.email.NonActivatedEmail() == "" {
		return "", false
	}

	return u.email.NonActivatedEmail(), true
}

func (u *User) Password() (p passwordEntity2.Password, exists bool) {
	if u.password == nil {
		return passwordEntity2.Password{}, false
	}
	return *u.password, true
}

func (u *User) Profile() profileEntity.Profile {
	return *u.profile
}

func (u *User) Agreement() (a agreementEntity.Agreement, exist bool) {
	if u.agreement == nil {
		return agreementEntity.Agreement{}, false
	}

	return *u.agreement, true
}

func (u *User) LastLoginAt() *commonTime.Time {
	return u.lastLoginAt
}

func (u *User) Removed() bool {
	return u.removed
}

func (u *User) initEmail(email emailPrimitive.Email) error {
	newEmail, err := emailEntity.NewEmails(email)
	if err != nil {
		return err
	}

	u.email = newEmail
	return nil
}

func (u *User) SetEmail(email emailPrimitive.Email) error {
	if u.email != nil {
		return u.email.InitNewEmail(email)
	} else {
		return u.initEmail(email)
	}
}

func (u *User) Emails() (e *emailEntity.UserEmails, exists bool) {
	if u.email == nil {
		return nil, false
	}

	return u.email, true
}

func (u *User) Role() spec.UserRole {
	return u.role
}
