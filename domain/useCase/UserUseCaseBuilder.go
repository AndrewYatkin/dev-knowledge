package userUseCase

import (
	repositoryInterface "dev-knowledge/boundary/repository"
	"dev-knowledge/infrastructure/errors"
	jwtServiceInterface "dev-knowledge/infrastructure/jwtService/interface"
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

func (b *Builder) JwtService(jwtService jwtServiceInterface.JWTService) *Builder {
	b.userUseCase.jwtService = jwtService
	return b
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
		b.errors.AddError(errors.NewError("SYS", "UserUseCase: UserRepository is required"))
	}
	if b.userUseCase.jwtService == nil {
		b.errors.AddError(errors.NewError("SYS", "UserUseCase: JwtService is required"))
	}
}
