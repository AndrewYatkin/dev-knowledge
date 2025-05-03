package repositoryInterface

import (
	"context"
	"dev-knowledge/boundary/dto"
)

type UserRepository interface {
	Save(ctx context.Context, user *dto.UserResponseDTO) error
	GetUserByID(ctx context.Context, userID string) (*dto.UserResponseDTO, error)
}
