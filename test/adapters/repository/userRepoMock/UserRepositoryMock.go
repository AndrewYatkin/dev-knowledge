package userRepoMock

import (
	"context"
	"dev-knowledge/boundary/dto"
	commonMock "dev-knowledge/infrastructure/testing/mock"
)

type UserRepositoryMock struct {
	*commonMock.BaseMock
}

func GetUserRepository() *UserRepositoryMock {
	return &UserRepositoryMock{
		BaseMock: commonMock.NewBaseMock(),
	}
}

func (m *UserRepositoryMock) Save(ctx context.Context, user *dto.UserResponseDTO) error {
	_, err := m.ProcessMethod("Save", ctx, user)
	return err
}

func (m *UserRepositoryMock) GetUserByID(ctx context.Context, userID string) (*dto.UserResponseDTO, error) {
	result, err := m.ProcessMethod("GetUserByID", ctx, userID)
	if err != nil {
		return nil, err
	}
	if user, ok := result.(*dto.UserResponseDTO); ok {
		return user, nil
	}
	return nil, commonMock.ErrCannotCastResult
}
