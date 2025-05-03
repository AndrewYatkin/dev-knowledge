package userUseCaseTest

import (
	userUseCase "dev-knowledge/domain/useCase"
	commonTesting "dev-knowledge/infrastructure/testing"
	"dev-knowledge/test/domain/useCase/stub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserCreateShould struct {
	suite.Suite
	*userUseCaseTestCommon
}

func TestUserCreateShould(t *testing.T) {
	suite.Run(t, &UserCreateShould{
		userUseCaseTestCommon: &userUseCaseTestCommon{},
	})
}

func (s *UserCreateShould) SetupTest() {
	s.SetupUseCase()
}

func (s *UserCreateShould) TestCreate_EmptyEmail_ReturnError() {
	incomingCreateUserDTO := stub.GetCreateUserDTOWithoutEmail()

	user, err := s.userUseCase.Create(s.ctx, incomingCreateUserDTO)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), userUseCase.ErrEmailIsRequired, err)
}

func (s *UserCreateShould) TestCreate_EmptyPassword_ReturnError() {
	incomingCreateUserDTO := stub.GetCrateUserDTOWithoutPassword()

	user, err := s.userUseCase.Create(s.ctx, incomingCreateUserDTO)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), userUseCase.ErrPasswordIsRequired, err)
}

func (s *UserCreateShould) TestCreate_InsertError_ReturnError() {
	incomingCreateUserDTO := stub.GetRandomCreateUserDTO()
	expectedError := commonTesting.RandomError()

	s.userRepoMock.On("Save", s.ctx, mock.AnythingOfType("*dto.UserResponseDTO")).Return(expectedError)

	user, err := s.userUseCase.Create(s.ctx, incomingCreateUserDTO)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), expectedError, err)
}

func (s *UserCreateShould) TestCreate_Valid_ReturnUser() {
	incomingCreateUserDTO := stub.GetRandomCreateUserDTO()

	s.userRepoMock.On("Save", s.ctx, mock.AnythingOfType("*dto.UserResponseDTO")).Return(nil)

	user, err := s.userUseCase.Create(s.ctx, incomingCreateUserDTO)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), user)
	assert.NotEmpty(s.T(), user.UserID)
	assert.Equal(s.T(), incomingCreateUserDTO.Email, user.Email)
	assert.Equal(s.T(), incomingCreateUserDTO.Username, user.Username)
}
