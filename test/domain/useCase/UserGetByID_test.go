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

type UserGetByIDShould struct {
	suite.Suite
	*userUseCaseTestCommon

	user *dto.UserResponseDTO
}

func TestUserGetByIDShould(t *testing.T) {
	suite.Run(t, &UserGetByIDShould{
		userUseCaseTestCommon: &userUseCaseTestCommon{},
	})
}

func (s *UserGetByIDShould) SetupTest() {
	s.SetupUseCase()
	s.user = stub.GetRandomUserResponseDTO()
}

func (s *UserGetByIDShould) TestGetByID_EmptyUserID_ReturnError() {
	expectedError := userUseCase.ErrUserIDIsRequired
	user, err := s.userUseCase.GetUserByID(s.ctx, "")

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), expectedError, err)
}

func (s *UserGetByIDShould) TestGetByID_RepoError_ReturnError() {
	expectedUserID := commonTesting.RandomUUID()
	expectedError := errors.NewError("123", "123")
	s.userRepoMock.On("GetUserByID", s.ctx, expectedUserID).Return(nil, expectedError)

	user, err := s.userUseCase.GetUserByID(s.ctx, expectedUserID)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), expectedError, err)
}

func (s *UserGetByIDShould) TestGetByID_Valid_ReturnUser() {
	expectedUserID := commonTesting.RandomUUID()
	s.userRepoMock.On("GetUserByID", s.ctx, expectedUserID).Return(s.user, nil)
	user, err := s.userUseCase.GetUserByID(s.ctx, expectedUserID)

	assert.Nil(s.T(), err)
	assert.EqualValues(s.T(), s.user, user)
}
