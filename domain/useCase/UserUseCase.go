package userUseCase

import (
	"context"
	"dev-knowledge/boundary/dto"
	"errors"
	"github.com/google/uuid"
)

type UserUseCase struct{}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{}
}

func (u UserUseCase) Create(ctx context.Context, createData *dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	responseDto, err := u.validateCreateUserData(createData)
	if err != nil {
		return responseDto, err
	}
	user := &dto.UserResponseDTO{
		UserID:   uuid.NewString(),
		Username: createData.Username,
		Email:    createData.Email,
	}

	return user, nil
}

func (u UserUseCase) validateCreateUserData(createData *dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	if createData.Username == "" {
		return nil, errors.New("username is empty")
	}
	if createData.Email == "" {
		return nil, errors.New("email is empty")
	}
	return nil, nil
}
