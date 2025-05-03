package userUseCaseTest

import (
	"context"
	userUseCase "dev-knowledge/domain/useCase"
	"dev-knowledge/test/adapters/repository/userRepoMock"
)

type userUseCaseTestCommon struct {
	userUseCase *userUseCase.UserUseCase

	userRepoMock *userRepoMock.UserRepositoryMock

	ctx context.Context
}

func (c *userUseCaseTestCommon) SetupUseCase() {
	c.ctx = context.Background()
	c.userRepoMock = userRepoMock.GetUserRepository()

	userUC, err := userUseCase.NewBuilder().
		UserRepo(c.userRepoMock).
		Build()
	if err != nil {
		panic(err)
	}
	c.userUseCase = userUC
}
