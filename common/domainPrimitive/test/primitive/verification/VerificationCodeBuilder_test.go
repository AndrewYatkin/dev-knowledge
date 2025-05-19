package verification

import (
	"dev-knowledge/infrastructure/errors"
	commonTime "dev-knowledge/infrastructure/tools/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type VerificationCodeBuilderShould struct {
	suite.Suite
}

func TestVerificationCodeBuilderShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(VerificationCodeBuilderShould))
}

func (s *VerificationCodeBuilderShould) TestBuild_RequiredParams_ReturnCode() {
	expectedCode := "123"
	expectedCreatedAt := commonTime.FromUnixNano(99987654321)
	expectedExpireAt := commonTime.FromUnixNano(99997654321)

	actualVerificationCode, err := verificationEntity.NewBuilder().
		Code(expectedCode).
		CreatedAt(expectedCreatedAt).
		ExpireAt(expectedExpireAt).
		Build()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualVerificationCode)
	assert.Equal(s.T(), expectedCode, actualVerificationCode.Code())
	assert.Equal(s.T(), expectedCreatedAt, actualVerificationCode.CreatedAt())
	assert.Equal(s.T(), expectedExpireAt, actualVerificationCode.ExpireAt())
}

func (s *VerificationCodeBuilderShould) TestBuild_AllValidParams_ReturnCode() {
	expectedCode := "123"
	expectedCreatedAt := commonTime.FromUnixNano(99987654321)
	expectedExpireAt := commonTime.FromUnixNano(99997654321)

	actualVerificationCode, err := verificationEntity.NewBuilder().
		Code(expectedCode).
		CreatedAt(expectedCreatedAt).
		ExpireAt(expectedExpireAt).
		Build()

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), actualVerificationCode)
	assert.Equal(s.T(), expectedCode, actualVerificationCode.Code())
	assert.Equal(s.T(), expectedCreatedAt, actualVerificationCode.CreatedAt())
	assert.Equal(s.T(), expectedExpireAt, actualVerificationCode.ExpireAt())
}

func (s *VerificationCodeBuilderShould) TestBuild_ExpireAtBeforeCreatedAt_ReturnError() {
	expectedCode := "123"
	expectedCreatedAt := commonTime.FromUnixNano(99997654321)
	expectedExpireAt := commonTime.FromUnixNano(88887654321)

	actualVerificationCode, err := verificationEntity.NewBuilder().
		Code(expectedCode).
		CreatedAt(expectedCreatedAt).
		ExpireAt(expectedExpireAt).
		Build()

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), actualVerificationCode)

	errs, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.True(s.T(), errs.Contains(verificationEntity.ErrExpireAtBeforeCreatedAt))
}

func (s *VerificationCodeBuilderShould) TestBuild_CodeIsEmpty_ReturnError() {
	actualVerificationCode, err := verificationEntity.NewBuilder().
		CreatedAt(commonTime.FromUnixNano(99997654321)).
		ExpireAt(commonTime.FromUnixNano(88887654321)).
		Build()

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), actualVerificationCode)

	errs, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.True(s.T(), errs.Contains(verificationEntity.ErrCodeRequired))
}

func (s *VerificationCodeBuilderShould) TestBuild_CreatedAtIsNil_ReturnError() {
	actualVerificationCode, err := verificationEntity.NewBuilder().
		Code("123").
		ExpireAt(commonTime.FromUnixNano(88887654321)).
		Build()

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), actualVerificationCode)

	errs, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.True(s.T(), errs.Contains(verificationEntity.ErrCreatedAtRequired))
}

func (s *VerificationCodeBuilderShould) TestBuild_ExpireAtIsNil_ReturnError() {
	actualVerificationCode, err := verificationEntity.NewBuilder().
		Code("123").
		CreatedAt(commonTime.FromUnixNano(99997654321)).
		Build()

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), actualVerificationCode)

	errs, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.True(s.T(), errs.Contains(verificationEntity.ErrExpireAtRequired))
}

func (s *VerificationCodeBuilderShould) TestBuild_NoParams_ReturnError() {
	actualVerificationCode, err := verificationEntity.NewBuilder().Build()

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), actualVerificationCode)

	errs, ok := err.(*errors.Errors)
	assert.True(s.T(), ok)
	assert.True(s.T(), errs.Contains(verificationEntity.ErrCodeRequired))
	assert.True(s.T(), errs.Contains(verificationEntity.ErrCreatedAtRequired))
	assert.True(s.T(), errs.Contains(verificationEntity.ErrExpireAtRequired))
}
