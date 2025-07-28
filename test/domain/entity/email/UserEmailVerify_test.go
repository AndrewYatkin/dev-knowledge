package email

import (
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	"dev-knowledge/domain/entity/user/email"
	"dev-knowledge/test/domain/entity/entityStub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserEmailVerifyShould struct {
	suite.Suite
}

func TestUserEmailVerifyShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserEmailVerifyShould))
}

func (s *UserEmailVerifyShould) TestVerify_VerifyCodeNotExists_ReturnError() {
	userEmail := entityStub.GetUserEmailWithoutVerify("email@email.com")

	err := userEmail.Verify("123")

	assert.Equal(s.T(), emailEntity.ErrVerificationCodeNotExists, err)
}

func (s *UserEmailVerifyShould) TestVerify_EmptyCode_ReturnError() {
	userEmail := entityStub.GetUserEmailWithoutVerify("email@email.com")

	err := userEmail.Verify("")

	assert.Equal(s.T(), emailEntity.ErrCodeToVerifyIsEmpty, err)
}

func (s *UserEmailVerifyShould) TestVerify_ValidCode_ReturnTrue() {
	userEmail := entityStub.GetUserEmailWithVerify("email@email.com", "123")
	err := userEmail.Verify("123")

	assert.Nil(s.T(), err)
	assert.Empty(s.T(), userEmail.VerificationCode())
}

func (s *UserEmailVerifyShould) TestGenerateAndVerify_ValidCode_ReturnTrue() {
	userEmail := entityStub.GetUserEmailWithoutVerify("email@email.com")

	err := userEmail.InitVerificationCode()
	assert.Nil(s.T(), err)

	err = userEmail.Verify(userEmail.VerificationCode())

	assert.Nil(s.T(), err)
	assert.Empty(s.T(), userEmail.VerificationCode())
}

func (s *UserEmailVerifyShould) TestGenerateAndVerify_invalidCode_ReturnFalseAndErr() {
	userEmail := entityStub.GetUserEmailWithoutVerify("email@email.com")

	err := userEmail.InitVerificationCode()
	assert.Nil(s.T(), err)

	err = userEmail.Verify("1234")

	assert.Equal(s.T(), verificationPrimitive.ErrMismatchedCode, err)
}
