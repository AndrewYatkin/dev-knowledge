package jwtservice_test

import (
	"context"
	jwtservice "dev-knowledge/infrastructure/jwtService/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type JWTServiceTestSuite struct {
	suite.Suite
	jwtService *jwtservice.JWTService
	secret     string
}

func TestJWTServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JWTServiceTestSuite))
}

func (s *JWTServiceTestSuite) SetupSuite() {
	s.secret = "test_secret"
	var err error
	s.jwtService, err = jwtservice.NewBuilder().
		Secret(s.secret).
		Build()
	assert.Nil(s.T(), err)
}

func (s *JWTServiceTestSuite) TestVerify() {
	validToken, err := s.jwtService.CreateUserToken("12345", map[string]string{"role": "admin"})
	assert.Nil(s.T(), err)
	assert.True(s.T(), s.jwtService.Verify(validToken))
}

func (s *JWTServiceTestSuite) TestWrongToken_VerifyFalse() {
	assert.False(s.T(), s.jwtService.Verify("invalid_token"))
}

func (s *JWTServiceTestSuite) TestFillCtxWithParams() {
	userID := "12345"
	role := "admin"
	validToken, err := s.jwtService.CreateUserToken(userID,
		map[string]string{
			jwtservice.RoleTokenKey: role},
	)
	assert.NoError(s.T(), err)

	ctx := context.Background()
	ctx, err = s.jwtService.FillCtxWithParams(ctx, validToken)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), userID, ctx.Value(jwtservice.UserIDKey))
	assert.Equal(s.T(), role, ctx.Value(jwtservice.UserRoleKey))
}

func (s *JWTServiceTestSuite) TestFillCtxWithParams_WithWrongToken_ReturnErr() {
	firstCtx := context.WithValue(context.Background(), "test", "test")
	ctx, err := s.jwtService.FillCtxWithParams(firstCtx, "invalid_token")
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ctx, firstCtx)
}

func (s *JWTServiceTestSuite) TestCreateUserToken() {
	userID := "12345"
	role := "admin"
	scope := "read:users"
	token, err := s.jwtService.CreateUserToken(userID,
		map[string]string{
			string(jwtservice.RoleTokenKey):  role,
			string(jwtservice.ScopeTokenKey): scope,
		})
	assert.Nil(s.T(), err)
	assert.NotEmpty(s.T(), token)
	assert.True(s.T(), s.jwtService.Verify(token))
}

func (s *JWTServiceTestSuite) TestCreateWithEmptyUserID_returnErr() {
	token, err := s.jwtService.CreateUserToken("", map[string]string{"role": "admin"})
	assert.Equal(s.T(), err, jwtservice.ErrJWTMissingUserID)
	assert.Empty(s.T(), token)
}

func (s *JWTServiceTestSuite) TestCreateWithWrongClaim_returnErr() {
	token, err := s.jwtService.CreateUserToken("12345", map[string]string{"invalid_claim": "value"})
	assert.ErrorIs(s.T(), err, jwtservice.ErrJWTInvalidClaims)
	assert.Empty(s.T(), token)
}
