package userUseCaseTest

import (
	"dev-knowledge/boundary/dto"
	userUseCase "dev-knowledge/domain/useCase"
	"dev-knowledge/infrastructure/errors"
	commonTesting "dev-knowledge/infrastructure/testing"
	"dev-knowledge/test/domain/useCase/stub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserGetMeShould struct {
	suite.Suite
	*userUseCaseTestCommon

	user *dto.UserResponseDTO
}

func TestUserGetMeShould(t *testing.T) {
	suite.Run(t, &UserGetMeShould{
		userUseCaseTestCommon: &userUseCaseTestCommon{},
	})
}

func (s *UserGetMeShould) SetupTest() {
	s.SetupUseCase()
	s.user = stub.GetRandomUserResponseDTO()
}

func (s *UserGetMeShould) TestGetMe_EmptyUserID_ReturnError() {
	expectedError := userUseCase.ErrUserIDIsRequired
	user, err := s.userUseCase.GetMe(s.ctx, "")

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), expectedError, err)
}

func (s *UserGetMeShould) TestGetMe_RepoError_ReturnError() {
	expectedUserID := commonTesting.RandomUUID()
	expectedError := errors.NewError("123", "123")
	s.userRepoMock.On("GetUserByID", s.ctx, expectedUserID).Return(nil, expectedError)

	user, err := s.userUseCase.GetMe(s.ctx, expectedUserID)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), expectedError, err)
}

func (s *UserGetMeShould) TestGetMe_Valid_ReturnUser() {
	expectedUserID := commonTesting.RandomUUID()
	s.userRepoMock.On("GetUserByID", s.ctx, expectedUserID).Return(s.user, nil)
	user, err := s.userUseCase.GetMe(s.ctx, expectedUserID)

	assert.Nil(s.T(), err)
	assert.EqualValues(s.T(), s.user, user)
}
