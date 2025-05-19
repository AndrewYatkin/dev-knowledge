package verification

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type NewRandomVerificationCodeShould struct {
	suite.Suite
}

func TestNewRandomVerificationCodeShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(NewRandomVerificationCodeShould))
}

func (s *NewRandomVerificationCodeShould) TestCreate_ValidParams_ReturnCode() {
	lifeMinutes := uint32(10)

	actualVerificationCode, err := verificationEntity.NewRandomVerificationCode(lifeMinutes)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualVerificationCode)
	assert.NotNil(s.T(), actualVerificationCode.Code())
	assert.NotNil(s.T(), actualVerificationCode.CreatedAt())

	expectedExpireAt := actualVerificationCode.CreatedAt().Add(time.Minute * time.Duration(lifeMinutes))
	assert.Equal(s.T(), expectedExpireAt, actualVerificationCode.ExpireAt())
}

func (s *NewRandomVerificationCodeShould) TestCreate_ZeroLifeMinutes_ReturnError() {
	actualVerificationCode, err := verificationEntity.NewRandomVerificationCode(0)

	assert.Equal(s.T(), err, verificationEntity.ErrZeroLifeMinutes)
	assert.Nil(s.T(), actualVerificationCode)
}
