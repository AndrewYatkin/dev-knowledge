package verification

import (
	verificationPrimitive "dev-knowledge/common/domainPrimitive/primitive/verification"
	commonTime "dev-knowledge/infrastructure/tools/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type VerifyVerificationCodeShould struct {
	suite.Suite
}

func TestVerifyVerificationCodeShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(VerifyVerificationCodeShould))
}

func (s *VerifyVerificationCodeShould) TestVerify_ValidCode_ReturnTrue() {
	expectedCodeValue := "123"
	verificationCode := GetVerificationCodeStub(expectedCodeValue)

	err := verificationCode.Verify(expectedCodeValue)
	assert.Nil(s.T(), err)

	err = verificationCode.Verify(expectedCodeValue)
	assert.Nil(s.T(), err)
}

func (s *VerifyVerificationCodeShould) TestVerify_ExpiredCode_ReturnFalseAndError() {
	expectedCodeValue := "123"
	expireAt := commonTime.Now().Add(time.Minute * time.Duration(-10))
	expiredVerificationCode := GetVerificationCodeStubByExpireAt(expectedCodeValue, expireAt)

	err := expiredVerificationCode.Verify(expectedCodeValue)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), verificationPrimitive.ErrExpiredCode, err)
}

func (s *VerifyVerificationCodeShould) TestVerify_NotValidCode_ReturnFalseAndError() {
	expiredVerificationCode := GetVerificationCodeStub("321")

	err := expiredVerificationCode.Verify("123")

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), verificationPrimitive.ErrMismatchedCode, err)
}
