package userUseCase

import (
	repositoryInterface "dev-knowledge/boundary/repository"
	"dev-knowledge/infrastructure/errors"
)

type Builder struct {
	userUseCase *UserUseCase
	errors      *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		userUseCase: &UserUseCase{},
		errors:      errors.NewErrors(),
	}
}

func (b *Builder) UserRepo(userRepo repositoryInterface.UserRepository) *Builder {
	b.userUseCase.userRepo = userRepo
	return b
}

func (b *Builder) Build() (*UserUseCase, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.userUseCase, nil
}

func (b *Builder) checkRequiredFields() {
	if b.userUseCase.userRepo == nil {
		b.errors.AddError(ErrUserRepoIsRequired)
	}
}
