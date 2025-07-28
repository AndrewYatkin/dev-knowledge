package passwordEntity

import "dev-knowledge/infrastructure/errors"

type Builder struct {
	hash Hash
	salt Salt

	errors *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) Hash(hash Hash) *Builder {
	b.hash = hash
	return b
}

func (b *Builder) Salt(salt Salt) *Builder {
	b.salt = salt
	return b
}

func (b *Builder) Build() (*Password, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.hash == "" {
		b.errors.AddError(ErrPasswordHashIsRequired)
	}

	if b.salt == "" {
		b.errors.AddError(ErrPasswordSaltIsRequired)
	}
}

func (b *Builder) createFromBuilder() *Password {
	return &Password{
		hash: b.hash,
		salt: b.salt,
	}
}
