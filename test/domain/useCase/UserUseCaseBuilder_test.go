package userUseCaseTest

import (
	userUseCase "dev-knowledge/domain/useCase"
	jwtservice "dev-knowledge/infrastructure/jwtService/test"
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

func (s *UserUseCaseBuilderShould) TestBuild_ValidParams_ReturnUseCase() {
	actualUserUseCase, err := userUseCase.NewBuilder().
		UserRepo(userRepoMock.GetUserRepository()).
		JwtService(jwtservice.NewJWTServiceMock()).
		Build()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualUserUseCase)
}
