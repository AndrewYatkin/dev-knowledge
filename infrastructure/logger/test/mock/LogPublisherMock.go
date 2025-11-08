package loggerMock

import (
	"context"
	commonTesting "dev-knowledge/infrastructure/testing/mock"
)

type LogPublisherMock struct {
	*commonTesting.BaseMock
}

func GetLogPublisherMock() *LogPublisherMock {
	return &LogPublisherMock{
		BaseMock: commonTesting.NewBaseMock(),
	}
}

func (m *LogPublisherMock) LogWarn(ctx context.Context, messages ...string) {
	_, err := m.ProcessMethod("LogWarn", ctx, messages)
	if err != nil {
		panic(err)
	}
}

func (m *LogPublisherMock) LogInfo(ctx context.Context, messages ...string) {
	_, err := m.ProcessMethod("LogInfo", ctx, messages)
	if err != nil {
		panic(err)
	}
}

func (m *LogPublisherMock) LogDebug(ctx context.Context, messages ...string) {
	_, err := m.ProcessMethod("LogDebug", ctx, messages)
	if err != nil {
		panic(err)
	}
}

func (m *LogPublisherMock) LogError(ctx context.Context, err ...error) {
	_, errs := m.ProcessMethod("LogError", ctx, err)
	if errs != nil {
		panic(errs)
	}
}
