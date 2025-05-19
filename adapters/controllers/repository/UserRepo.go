package userRepo

import (
	"context"
	"dev-knowledge/boundary/dto"
	loggerInterface "dev-knowledge/infrastructure/logger/interface"
)

type UserRepo struct {
	logger loggerInterface.Logger
}

func (u *UserRepo) Save(ctx context.Context, user *dto.UserResponseDTO) error {
	//TODO implement me
	u.logger.Info(ctx, "userRepo implement me")

	return nil
}

func (u *UserRepo) GetUserByID(ctx context.Context, userID string) (*dto.UserResponseDTO, error) {
	//TODO implement me
	u.logger.Info(ctx, "userRepo implement me")

	return &dto.UserResponseDTO{UserID: userID}, nil
}
