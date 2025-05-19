package email

import (
	emailPrimitive "dev-knowledge/common/domainPrimitive/primitive/email"
	emailPrimitiveStub "dev-knowledge/common/domainPrimitive/test/testDouble/stub/email"
	emailEntity "dev-knowledge/domain/entity/email"
	"dev-knowledge/infrastructure/errors"
	commonTime "dev-knowledge/infrastructure/tools/time"
	"dev-knowledge/test/domain/entity/entityStub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserEmailBuilderShould struct {
	suite.Suite
}

func TestUserEmailShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserEmailBuilderShould))
}

func (s *UserEmailBuilderShould) TestBuild_ParamsNotGiven_ReturnError() {
	actualUserEmail, err := emailEntity.NewBuilder().Build()

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), actualUserEmail)

	errs, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.True(s.T(), errs.Contains(emailEntity.ErrEmailIsRequired))
}

func (s *UserEmailBuilderShould) TestBuild_EmptyEmail_ReturnErrEmailRequired() {
	emptyEmail := emailPrimitive.Email("")

	actualUserEmail, err := emailEntity.NewBuilder().
		Email(emptyEmail).
		Build()

	assert.Nil(s.T(), actualUserEmail)
	assert.NotNil(s.T(), err)

	errs, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.True(s.T(), errs.Contains(emailEntity.ErrEmailIsRequired))
}

func (s *UserEmailBuilderShould) TestBuild_AllParams_ReturnUserEmailNoErr() {
	expectedID := emailEntity.NewEmailID()
	expectedEmail := emailPrimitiveStub.GetEmail()
	expectedCreatedAt := commonTime.Now()
	expectedVerificationCode := entityStub.GetVerificationCode()

	actualUserEmail, err := emailEntity.NewBuilder().
		ID(expectedID).
		Email(expectedEmail).
		VerificationCode(expectedVerificationCode).
		CreatedAt(expectedCreatedAt).
		Build()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualUserEmail)
	assert.EqualValues(s.T(), expectedEmail, actualUserEmail.Email())
	assert.Equal(s.T(), expectedCreatedAt.UnixNano(), actualUserEmail.CreatedAt().UnixNano())
	assert.Equal(s.T(), expectedID, actualUserEmail.ID())
	assert.Equal(s.T(), expectedVerificationCode.Code(), actualUserEmail.VerificationCode())
}

func (s *UserEmailBuilderShould) TestBuild_OnlyRequiredParams_ReturnUserEmailWithDefaultValues() {
	expectedEmail := emailPrimitiveStub.GetEmail()

	actualUserEmail, err := emailEntity.NewBuilder().
		Email(expectedEmail).
		Build()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualUserEmail)
	assert.EqualValues(s.T(), expectedEmail, actualUserEmail.Email())
	assert.NotEqual(s.T(), 0, actualUserEmail.CreatedAt().UnixNano())
	assert.NotEqualValues(s.T(), "", actualUserEmail.ID())
	assert.Empty(s.T(), actualUserEmail.VerificationCode())
}
