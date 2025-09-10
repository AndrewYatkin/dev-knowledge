package jwtservice

import (
	"context"
	commonMock "dev-knowledge/infrastructure/testing/mock"
)

type JWTServiceMock struct {
	*commonMock.BaseMock
}

func NewJWTServiceMock() *JWTServiceMock {
	return &JWTServiceMock{BaseMock: commonMock.NewBaseMock()}
}

func (m *JWTServiceMock) Verify(tokenString string) bool {
	result, err := m.ProcessMethod("Verify")
	if err != nil {
		return false
	}
	if res, ok := result.(bool); ok {
		return res
	}
	return true
}

func (m *JWTServiceMock) FillCtxWithParams(ctx context.Context, tokenString string) (context.Context, error) {
	return ctx, nil
}

func (m *JWTServiceMock) CreateUserToken(userID string, claims map[string]string) (string, error) {
	return "token", nil
}
