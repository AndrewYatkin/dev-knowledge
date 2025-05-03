package usecaseInterface

import (
	"context"
	"dev-knowledge/boundary/dto"
)

type UserUseCaseInterface interface {
	Create(ctx context.Context, createData *dto.CreateUserDTO) (*dto.UserResponseDTO, error)
	GetUserByID(ctx context.Context, userIDStr string) (*dto.UserResponseDTO, error)
}
