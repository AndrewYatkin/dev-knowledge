package jwtService

import "dev-knowledge/infrastructure/errors"

type JWTServiceBuider struct {
	service *JWTService
	errors  *errors.Errors
}

func NewBuilder() *JWTServiceBuider {
	return &JWTServiceBuider{
		service: &JWTService{},
		errors:  errors.NewErrors(),
	}
}

func (b *JWTServiceBuider) Secret(secret string) *JWTServiceBuider {
	b.service.secret = secret
	return b
}

func (b *JWTServiceBuider) ValidClaims(validClaims map[string]bool) *JWTServiceBuider {
	b.service.validClaims = validClaims
	return b
}

func (b *JWTServiceBuider) Build() (*JWTService, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.service, nil
}

func (b *JWTServiceBuider) checkRequiredFields() {
	if b.service.secret == "" {
		b.errors.AddError(errors.NewError("SYS", "JWTServiceBuilder: secret is required"))
	}
}

func (b *JWTServiceBuider) fillDefaultFields() {
	if b.service.validClaims == nil {
		b.service.validClaims = map[string]bool{
			string(RoleTokenKey):       true,
			string(ScopeTokenKey):      true,
			string(expirationTokenKey): true,
		}
	}
}
