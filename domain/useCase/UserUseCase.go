package userUseCase

import (
	"context"
	"dev-knowledge/boundary/dto"
	repositoryInterface "dev-knowledge/boundary/repository"
	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo repositoryInterface.UserRepository
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
	return user, nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, userID string) (*dto.UserResponseDTO, error) {
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
