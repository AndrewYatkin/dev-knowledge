package loggerFake

import (
	"context"
	"fmt"
)

type LogPublisherFake struct{}

func GetLogPublisher() *LogPublisherFake {
	return &LogPublisherFake{}
}

func (m *LogPublisherFake) LogWarn(ctx context.Context, messages ...string) {
	_ = fmt.Sprintf("logWarn messages: %s", messages)
}

func (m *LogPublisherFake) LogInfo(ctx context.Context, messages ...string) {
	_ = fmt.Sprintf("LogInfo messages: %s", messages)
}

func (m *LogPublisherFake) LogDebug(ctx context.Context, messages ...string) {
	_ = fmt.Sprintf("LogDebug messages: %s", messages)
}

func (m *LogPublisherFake) LogError(ctx context.Context, err ...error) {
	_ = fmt.Sprintf("LogError messages: %s", err)
}
