package verification

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type NewVerificationCodeShould struct {
	suite.Suite
}

func TestNewVerificationCodeShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(NewVerificationCodeShould))
}

func (s *NewVerificationCodeShould) TestCreate_ValidParams_ReturnCode() {
	expectedCode := "123"
	lifeMinutes := uint32(10)

	actualVerificationCode, err := verificationEntity.NewVerificationCode(expectedCode, lifeMinutes)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualVerificationCode)
	assert.Equal(s.T(), expectedCode, actualVerificationCode.Code())
	assert.NotNil(s.T(), actualVerificationCode.CreatedAt())

	expectedExpireAt := actualVerificationCode.CreatedAt().Add(time.Minute * time.Duration(lifeMinutes))
	assert.Equal(s.T(), expectedExpireAt, actualVerificationCode.ExpireAt())
}

func (s *NewVerificationCodeShould) TestCreate_ZeroLifeMinutes_ReturnError() {
	actualVerificationCode, err := verificationEntity.NewVerificationCode("123", 0)

	assert.Equal(s.T(), err, verificationEntity.ErrZeroLifeMinutes)
	assert.Nil(s.T(), actualVerificationCode)
}

func (s *NewVerificationCodeShould) TestCreate_EmptyCode_ReturnError() {
	actualVerificationCode, err := verificationEntity.NewVerificationCode("", 10)

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), actualVerificationCode)
}
