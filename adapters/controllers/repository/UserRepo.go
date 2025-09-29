package userRepo

import (
	"context"
	"dev-knowledge/boundary/dto"
	logInterface "dev-knowledge/infrastructure/logger/interface"
)

type UserRepo struct {
	logger logInterface.LogPublisher
}

func (u *UserRepo) Save(ctx context.Context, user *dto.UserResponseDTO) error {
	//TODO implement me
	u.logger.LogInfo(ctx, "userRepo implement me")

	return nil
}

func (u *UserRepo) GetUserByID(ctx context.Context, userID string) (*dto.UserResponseDTO, error) {
	//TODO implement me
	u.logger.LogInfo(ctx, "userRepo implement me")

	return &dto.UserResponseDTO{UserID: userID}, nil
}
