package userUseCase

import (
	"context"
	"dev-knowledge/boundary/dto"
	repositoryInterface "dev-knowledge/boundary/repository"
	"dev-knowledge/domain/entity/user/spec"
	"dev-knowledge/infrastructure/jwtService"
	jwtServiceInterface "dev-knowledge/infrastructure/jwtService/interface"
	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo   repositoryInterface.UserRepository
	jwtService jwtServiceInterface.JWTService
}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (u UserUseCase) Create(ctx context.Context, createUserDTO *dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	err := u.validateCreateUserData(createUserDTO)
	if err != nil {
		return nil, err
	}
	user := &dto.UserResponseDTO{
		UserID:   uuid.NewString(),
		Username: createUserDTO.Username,
		Email:    createUserDTO.Email,
	}

	err = u.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := u.jwtService.CreateUserToken(
		user.UserID,
		map[string]string{jwtService.RoleTokenKey: spec.UserRoles.Admin().String()})
	if err != nil {
		return nil, err
	}
	user.Token = token
	return user, nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, userID string, executorRole string) (*dto.UserResponseDTO, error) {
	if !u.isAdmin(executorRole) {
		return nil, ErrPermissionDenied
	}
	if userID == "" {
		return nil, ErrUserIDIsRequired
	}

	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) GetMe(ctx context.Context, userID string) (*dto.UserResponseDTO, error) {
	if userID == "" {
		return nil, ErrUserIDIsRequired
	}

	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) validateCreateUserData(createData *dto.CreateUserDTO) error {
	if createData.Email == "" {
		return ErrEmailIsRequired
	}
	if createData.Password == "" {
		return ErrPasswordIsRequired
	}
	return nil
}

func (u *UserUseCase) isAdmin(executorRole string) bool {
	return executorRole == spec.UserRoles.Admin().String()
}
