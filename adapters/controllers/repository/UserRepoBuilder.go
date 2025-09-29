package userRepo

import (
	"dev-knowledge/infrastructure/errors"
	logInterface "dev-knowledge/infrastructure/logger/interface"
)

type Builder struct {
	userRepo *UserRepo
	errors   *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		userRepo: &UserRepo{},
		errors:   &errors.Errors{},
	}
}

func (b *Builder) Logger(logger logInterface.LogPublisher) *Builder {
	b.userRepo.logger = logger
	return b
}

func (b *Builder) Build() (*UserRepo, error) {
	b.chechRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.userRepo, nil
}

func (b *Builder) chechRequiredFields() {
	if b.userRepo.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
}
