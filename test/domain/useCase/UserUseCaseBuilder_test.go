package userUseCaseTest

import (
	userUseCase "dev-knowledge/domain/useCase"
	commonTesting "dev-knowledge/infrastructure/testing"
	"dev-knowledge/test/adapters/repository/userRepoMock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserUseCaseBuilderShould struct {
	suite.Suite
}

func TestUserUseCaseBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserUseCaseBuilderShould))
}

func (s *UserUseCaseBuilderShould) TestBuild_WithoutParams_ReturnError() {
	expectedErrors := []error{
		userUseCase.ErrUserRepoIsRequired,
	}

	actualUserUseCase, err := userUseCase.NewBuilder().Build()
	assert.Nil(s.T(), actualUserUseCase)
	commonTesting.AssertErrors(s.T(), err, expectedErrors)
}

func (s *UserUseCaseBuilderShould) TestBuild_ValidParams_ReturnUseCase() {
	actualUserUseCase, err := userUseCase.NewBuilder().
		UserRepo(userRepoMock.GetUserRepository()).
		Build()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualUserUseCase)
}
